// eslint-disable-next-line
import { default as component } from './Pagination.vue'
import { components } from '@cortezaproject/corteza-vue'
const { checkbox } = components.C3.controls

const props = {
  hasPrevPage: true,
  hasNextPage: true,
}

export default {
  name: 'Pagination',
  group: ['ModuleFields'],
  component,
  props,
  controls: [
    checkbox('Enable previous page', 'hasPrevPage'),
    checkbox('Enable next page', 'hasNextPage'),
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
        hasPrevPage: false,
        hasNextPage: false,
      },
    },
  ],
}
