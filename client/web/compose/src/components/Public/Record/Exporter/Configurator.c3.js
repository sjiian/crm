// eslint-disable-next-line
import { default as component } from './Configurator.vue'
import { compose } from '@cortezaproject/corteza-js'
import { components } from '@cortezaproject/corteza-vue'
const { checkbox, select } = components.C3.controls

const props = {
  allowJSON: true,
  allowCSV: true,
  module: new compose.Module(),
  preselectedFields: [],
  query: undefined,
  filterRangeType: 'all',
  filterRangeBy: 'created_at',
  dateRange: 'lastMonth',
  startDate: null,
  endDate: null,
  systemFields: ['ownedBy', 'createdAt', 'createdBy', 'updatedAt', 'updatedBy'],
  disabledTypes: ['User', 'Record', 'File'],
  processingCount: false,
  selectedFields: [],
  processing: false,
}

export default {
  name: 'Configurator',
  group: ['Exporter'],
  component,
  props,
  controls: [
    checkbox('Allow JSON', 'allowJSON'),
    checkbox('Allow CSV', 'allowCSV'),
    checkbox('Processing', 'processing'),
    select('Filter range by', 'filterRangeBy', [{ value: 'created_at', text: 'created_at' }, { value: 'deleted_at', text: 'deleted_at' }, undefined]),
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
        allowJSON: false,
        allowCSV: false,
        module: undefined,
        recordCount: 0,
        selection: [],
        filterRangeBy: '',
        filterRangeType: '',
        systemFields: [],
        disabledTypes: [],
        processing: false,
      },
    },
  ],
}
