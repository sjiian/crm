import component from './CTranslatorButton.vue'
import { components } from '@cortezaproject/corteza-vue'
const { input, checkbox } = components.C3.controls

const props = {
  buttonClass: 'p-1 m-1',
  buttonVariant: 'dark',
  disabled: false,
  resource: 'compose:module/313892290634645612/313892291205333100',
}

export default {
  name: 'Button',
  group: ['Translator'],
  component,
  props,
  controls: [
    input('Class', 'buttonClass'),
    input('Variant', 'buttonVariant'),
    checkbox('Disabled', 'disabled'),
  ],
  scenarios: [
    {
      label: 'Full form',
      props: {
        ...props,
        buttonClass: 'p-1 m-1',
      },
    },
    {
      label: 'Empty form',
      props: {
        ...props,
        buttonClass: 'p-0 m-0',
        buttonVariant: 'light',
        disabled: true,
      },
    },
  ],
}
