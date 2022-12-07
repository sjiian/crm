// eslint-disable-next-line
import { default as component } from './ReportEdit.vue'

const props = {
  report: {
    moduleID: '291598304537084005',
    filter: '',
    colorScheme: '#fffff',
    metrics: [
      {
        aggregate: 'AVG',
        field: 'count',
        moduleID: '291598304537084005',
      },
    ],
    dimensions: [{
      field: 'Rating',
      modifier: '(no grouping / buckets)',
      skipMissing: true,
    }],
  },
  modules: [
    {
      moduleID: '291598304537084005',
      handle: 'Case',
      name: 'Case',
      fields: [{
        kind: 'String',
        name: 'Sample',
        options: {
          multiLine: false,
          useRichTextEditor: false,
        },
      }],
      systemFields () {
        return Object.freeze([
          { isSystem: true, name: 'recordID', label: 'Record ID', kind: 'String' },
          { isSystem: true, name: 'revision', label: 'Revision', kind: 'Number' },
          { isSystem: true, name: 'ownedBy', label: 'Owned by', kind: 'User' },
          { isSystem: true, name: 'createdBy', label: 'Created by', kind: 'User' },
          { isSystem: true, name: 'createdAt', label: 'Created at', kind: 'DateTime' },
          { isSystem: true, name: 'updatedBy', label: 'Updated by', kind: 'User' },
          { isSystem: true, name: 'updatedAt', label: 'Updated at', kind: 'DateTime' },
          { isSystem: true, name: 'deletedBy', label: 'Deleted by', kind: 'User' },
          { isSystem: true, name: 'deletedAt', label: 'Deleted at', kind: 'DateTime' },
        ])
      }
    },
    {
      moduleID: '',
      handle: 'Party',
      name: 'Party',
      fields: [{
        kind: 'Record',
        name: 'AccountId',
        options: {
          multiLine: false,
          useRichTextEditor: false,
        },
      }],
      systemFields () {
        return Object.freeze([
          { isSystem: true, name: 'recordID', label: 'Record ID', kind: 'String' },
          { isSystem: true, name: 'revision', label: 'Revision', kind: 'Number' },
          { isSystem: true, name: 'ownedBy', label: 'Owned by', kind: 'User' },
          { isSystem: true, name: 'createdBy', label: 'Created by', kind: 'User' },
          { isSystem: true, name: 'createdAt', label: 'Created at', kind: 'DateTime' },
          { isSystem: true, name: 'updatedBy', label: 'Updated by', kind: 'User' },
          { isSystem: true, name: 'updatedAt', label: 'Updated at', kind: 'DateTime' },
          { isSystem: true, name: 'deletedBy', label: 'Deleted by', kind: 'User' },
          { isSystem: true, name: 'deletedAt', label: 'Deleted at', kind: 'DateTime' },
        ])
      }
    },
  ],
  supportedMetrics: -1,
  dimensionFieldKind: ['DateTime', 'Select', 'Number', 'Bool', 'String'],
  unSkippable: false,
}

export default {
  name: 'Edit',
  group: ['Chart', 'Report'],
  component,
  props,
  // add controls
  controls: [],

  scenarios: [
    {
      label: 'Full form',
      props,
    },
    {
      label: 'Empty form',
      props: {
        ...props,
        // add props
        modules: [],
      },
    },
  ],
}
