// eslint-disable-next-line
import { default as component } from './NamespaceItem.vue'
import { components } from '@cortezaproject/corteza-vue'
const { checkbox, input } = components.C3.controls

const props = {
  namespace: {
    meta: {
      subtitle: 'Subtitle',
      description: 'Lorem ipsum dolor',
    },
    name: 'CRM',
    enabled: true,
    canUpdateNamespace: true,
  },
}

export default {
  name: 'Item',
  group: ['Namespaces'],
  component,
  props,
  controls: [
    // add more controls
    input('Subtitle', 'namespace.meta.subtitle'),
    input('Description', 'namespace.meta.description'),
    input('Name', 'namespace.name'),
    checkbox('Enable namespace', 'namespace.enabled'),
  ],

  scenarios: [
    {
      label: 'Full form',
      props,
    },
    {
      label: 'Empty form',
      props: {
        ...props,
        namespace: {
          meta: {
            subtitle: '',
            description: '',
          },
          name: '',
          enabled: false,
          canUpdateNamespace: false,
        },
      },
    },
  ],
}
