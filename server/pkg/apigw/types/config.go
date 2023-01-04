package types

import "time"

type (
	Config struct {
		Enabled bool

		ProfilerGlobal  bool
		ProfilerEnabled bool

		ProxyFollowRedirects bool
		ProxyOutboundTimeout time.Duration
	}
)
