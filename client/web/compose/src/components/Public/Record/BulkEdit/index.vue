<template>
  <div>
    <b-button
      v-if="!openOnSelect"
      :title="$t('recordList.bulkRecord.title')"
      variant="outline-light"
      class="text-primary border-0"
      size="sm"
      @click="showModal = true"
    >
      <font-awesome-icon
        :icon="['far', 'edit']"
      />
    </b-button>

    <b-modal
      :visible="showModal"
      :title="modalTitle || $t('recordList.bulkRecord.title')"
      body-class="p-0"
      footer-class="flex-column align-items-stretch"
      centered
      no-fade
      @hide="onModalHide"
    >
      <b-card
        v-if="fields.length"
        class="pt-0"
      >
        <div
          v-for="(field, index) in fields"
          :key="field.fieldID"
          class="position-relative"
        >
          <field-editor
            :namespace="namespace"
            :module="module"
            :field="getField(field)"
            :errors="fieldErrors(field)"
            :record="record"
          />

          <c-input-confirm
            class="position-absolute"
            :tooltip="$t('recordList.bulkRecord.field.remove')"
            style="top: -2px; right: -4.5px; z-index: 2;"
            @confirmed="fields.splice(index, 1)"
          />
        </div>
      </b-card>

      <template #modal-footer>
        <vue-select
          v-model="selectedField"
          :placeholder="getFieldSelectorPlaceholder"
          :get-option-label="getFieldLabel"
          :get-option-key="getOptionKey"
          :options="moduleFields"
          append-to-body
          :calculate-position="calculateDropdownPosition"
          :selectable="option => !fields.includes(option.name)"
          :reduce="f => f.name"
          class="bg-white position-relative"
          @input="addField"
        />

        <hr class="my-3">

        <div
          class="d-flex justify-content-between align-items-center"
        >
          <b-button
            variant="light"
            :disabled="processing"
            @click="onReset"
          >
            {{ $t('recordList.bulkRecord.reset') }}
          </b-button>

          <div>
            <b-button
              variant="link"
              rounded
              class="text-primary"
              @click="onModalHide"
            >
              {{ $t('general.label.cancel') }}
            </b-button>
            <b-button
              variant="primary"
              :disabled="!fields.length || processing"
              @click="handleBulkUpdateSelectedRecords(query)"
            >
              {{ $t('general.label.save') }}
            </b-button>
          </div>
        </div>
      </template>
    </b-modal>
  </div>
</template>

<script>
import FieldEditor from 'corteza-webapp-compose/src/components/ModuleFields/Editor'
import { VueSelect } from 'vue-select'
import { compose } from '@cortezaproject/corteza-js'
import record from 'corteza-webapp-compose/src/mixins/record.js'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'BulkEdit',

  components: {
    VueSelect,
    FieldEditor,
  },

  mixins: [
    record,
  ],

  props: {
    namespace: {
      type: compose.Namespace,
      required: true,
    },

    module: {
      type: compose.Module,
      required: true,
    },

    selectedFields: {
      type: Array,
      default: () => ([]),
    },

    initialRecord: {
      type: Object,
      default: () => ({}),
    },

    openOnSelect: {
      type: Boolean,
      default: false,
    },

    modalTitle: {
      type: String,
      default: '',
    },

    query: {
      type: String,
      default: '',
    },
  },

  data () {
    return {
      showModal: false,
      selectedField: undefined,
      fields: [],
    }
  },

  computed: {
    moduleFields () {
      return [
        ...[...this.module.fields].sort((a, b) =>
          (a.label || a.name).localeCompare(b.label || b.name),
        ),
        ...this.module.systemFields().filter(({ name }) => name === 'ownedBy'),
      ].filter((field) => this.isFieldEditable(field))
    },

    getFieldSelectorPlaceholder () {
      return this.$t(`recordList.bulkRecord.field.add${this.fields.length ? 'Another' : ''}`)
    },
  },

  watch: {
    query: {
      handler (query) {
        if (!this.openOnSelect || !query.length) return

        this.record = new compose.Record(this.module, this.initialRecord)
        this.showModal = true
      },
    },

    selectedFields: {
      handler (fields = []) {
        if (!fields.length) return

        fields.forEach(f => {
          if (this.fields.includes(f)) return
          this.fields.push(f)
        })
      },
    },
  },

  created () {
    this.record = new compose.Record(this.module, {})
  },

  methods: {
    onModalHide () {
      this.showModal = false

      if (this.openOnSelect) {
        this.fields = []
        this.record = new compose.Record(this.module, {})
      }
      this.$emit('close')
    },

    getFieldLabel ({ kind, label, name }) {
      return label || name || kind
    },

    addField (field) {
      if (!field) return

      this.fields.push(field)
      this.selectedField = null
    },

    onReset () {
      this.record = new compose.Record(this.module, this.initialRecord)
      this.fields = []
    },

    getField (fieldName) {
      const field = this.moduleFields.find(
        ({ name }) => name === fieldName,
      )

      return field || {}
    },

    isFieldEditable (field) {
      if (!field) return false

      const { canCreateOwnedRecord } = this.module || {}
      const { createdAt, canManageOwnerOnRecord } = this.record || {}
      const { name, canUpdateRecordValue, isSystem, expressions = {} } = field || {}

      if (!canUpdateRecordValue) return false

      if (isSystem) {
        // Make ownedBy field editable if correct permissions
        if (name === 'ownedBy') {
          // If not created we check module permissions, otherwise the canManageOwnerOnRecord
          return createdAt ? canManageOwnerOnRecord : canCreateOwnedRecord
        }

        return false
      }

      return !expressions.value
    },

    getOptionKey ({ fieldID }) {
      return fieldID
    },
  },
}
</script>

<style lang="scss">
.position-initial {
  position: initial;
}
</style>
