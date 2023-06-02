package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cortezaproject/corteza/server/pkg/plugin/automation"
	"github.com/hashicorp/go-hclog"
	hcp "github.com/hashicorp/go-plugin"
	"go.uber.org/zap"
)

type (
	plugin struct {
		log *zap.Logger
		ar  automation.AutomationRegistry

		rawPlugins map[string]*cp
		clients    map[string]*hcp.Client
		pluginMap  map[string]hcp.Plugin
	}

	CortezaPluginAutomationFunction interface {
		automation.AutomationFunction
		automation.AutomationRegistry
	}
)

var (
	pl *plugin
)

const PLUGIN_PATH = "./plugin"

func Service() *plugin {
	return pl
}

func Setup(l *zap.Logger, ar automation.AutomationRegistry) {
	if pl != nil {
		return
	}

	pl = New(l, ar)
}

func New(l *zap.Logger, ar automation.AutomationRegistry) *plugin {
	return &plugin{
		log: l.Named("plugin"),
		ar:  ar,

		rawPlugins: make(map[string]*cp),
		clients:    make(map[string]*hcp.Client),
		pluginMap:  make(map[string]hcp.Plugin),
	}
}

func (p *plugin) register(path string) {
	var (
		logger = hclog.New(&hclog.LoggerOptions{
			Name:   "plugin",
			Output: os.Stdout,
			Level:  hclog.Debug,
		})

		config = &hcp.ClientConfig{
			HandshakeConfig:  handshakeConfig,
			Plugins:          p.pluginMap,
			Cmd:              exec.Command("sh", "-c", path),
			AllowedProtocols: allowedProtocols,
			Logger:           logger,
		}
	)

	slug := makeSlug(path)

	p.clients[slug] = hcp.NewClient(config)

	p.log.Info("registered new GRPC client", zap.String("slug", slug))
}

func (p *plugin) unregister() {
	for _, c := range p.clients {
		c.Kill()
	}
}

// finds the registered client and prepares it in the internal registry
func (p *plugin) dispense(slug string) (err error) {
	if _, is := p.clients[slug]; !is {
		err = fmt.Errorf("could not find grpc client %s", slug)
		return
	}

	c, err := p.clients[slug].Client()

	if err != nil {
		p.log.Error("client register error", zap.Error(err))
		return
	}

	raw, err := c.Dispense(slug)

	if err != nil {
		p.log.Error("client dispense error", zap.Error(err))
		return
	}

	cc := &cp{}

	switch s := raw.(type) {
	case automation.AutomationFunction:
		cc.af = s
	}

	p.rawPlugins[slug] = cc

	return
}

func (p *plugin) findPlugins(path string) (ss []string, err error) {
	ss, err = filepath.Glob(filepath.Join(PLUGIN_PATH, "*"))
	return
}

func (p *plugin) registerPlugins(ctx context.Context) {

	filePaths, err := p.findPlugins(PLUGIN_PATH)

	if err != nil {
		p.log.Error("invalid file glob pattern", zap.Error(err))
		return
	}

	for _, filePath := range filePaths {

		slug := makeSlug(filePath)

		// todo - this should be handled via Manager
		p.pluginMap[slug] = &automation.AutomationFunctionPlugin{}

		p.register(filePath)

		if err := p.dispense(slug); err != nil {
			p.log.Error("could not dispense", zap.String("slug", slug))
		}
	}
}

func (p *plugin) registerAutomations() {
	for _, raw := range p.rawPlugins {
		p.ar.AddFunctions(raw.automationFunction())
	}
}

func (p *plugin) Register(ctx context.Context) {
	p.registerPlugins(ctx)
	p.registerAutomations()
}

func makeSlug(path string) string {
	return filepath.Base(path)
}
