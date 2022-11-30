const buildVueConfig = require('./vue.config-builder')
const p = require('path')

module.exports = buildVueConfig({
  appFlavour: 'Admin Area',
  appName: 'admin',
  appLabel: 'Corteza Admin',
  theme: 'corteza-base',
  packageAlias: [
    { alias: 'corteza-webapp-admin', path: p.resolve('.') },
    { alias: 'corteza-js', path: p.resolve('../../../lib/js/dist') },
    // { alias: 'corteza-vue', path: p.resolve('../../../lib/vue/dist/index.d.ts') },
  ],
})
