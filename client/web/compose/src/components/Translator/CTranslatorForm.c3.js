import component from './CTranslatorForm.vue'
import { components } from '@cortezaproject/corteza-vue'
const { select } = components.C3.controls

const props = {
  // passed but not used
  // disabled: false,

  languages: [
    { localizedName: 'English', name: 'English', tag: 'en' },
    { localizedName: 'German', name: 'German', tag: 'de' },
    { localizedName: 'Italian', name: 'italiano', tag: 'it' },
  ],

  primaryResource: 'compose:page/291598304267993189/291598304588464229',

  translations: [
    { key: 'title', lang: 'en', message: 'Title', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'title', lang: 'de', message: 'Title', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'title', lang: 'it', message: 'Titolo', resource: 'compose:page/291598304267993189/291598304588464229' },

    { key: 'description', lang: 'en', message: 'Description', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'description', lang: 'de', message: 'Bezeichnung', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'description', lang: 'it', message: 'Descrizione', resource: 'compose:page/291598304267993189/291598304588464229' },

    { key: 'name', lang: 'en', message: 'Name', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'name', lang: 'de', message: 'Name', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'name', lang: 'it', message: 'Nome', resource: 'compose:page/291598304267993189/291598304588464229' },

    { key: 'contact', lang: 'en', message: 'Contact', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'contact', lang: 'de', message: 'Contakt', resource: 'compose:page/291598304267993189/291598304588464229' },
    { key: 'contact', lang: 'it', message: 'Contatto', resource: 'compose:page/291598304267993189/291598304588464229' },
  ],

  // Passed to modal; won't be shown
  // titles: {
  //   'compose:page/291598304267993189/291598304588464229': "Page"
  // },

  highlightKey: 'title',
}

const translations = [
  { value: '', text: 'None' },
  { value: 'title', text: 'Title' },
  { value: 'description', text: 'Description' },
  { value: 'name', text: 'Name' },
  { value: 'contact', text: 'Contact' },
]

export default {
  name: 'Form',
  group: ['Translator'],
  component,
  props,
  controls: [
    select('Highlighted key', 'highlightKey', translations),
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
        // nested values don;t get overwritten 
        // languages: [],
        // nested values don;t get overwritten
        // translations: [],
        highlightKey: '',
      },
    },
  ],
}
