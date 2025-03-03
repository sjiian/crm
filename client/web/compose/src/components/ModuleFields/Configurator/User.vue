<template>
  <div>
    <b-form-group>
      <b-form-checkbox v-model="f.options.presetWithAuthenticated">
        {{ $t('kind.user.presetWithCurrentUser') }}
      </b-form-checkbox>
    </b-form-group>

    <b-form-group
      v-if="f.options.roles"
      :label="$t('kind.user.roles.label')"
    >
      <vue-select
        v-model="f.options.roles"
        :options="roleOptions"
        :get-option-key="getOptionKey"
        :reduce="role => role.roleID"
        option-value="roleID"
        option-text="name"
        :close-on-select="false"
        append-to-body
        :placeholder="$t('kind.user.roles.placeholder')"
        :calculate-position="calculateDropdownPosition"
        multiple
        label="name"
        class="bg-white"
      />
    </b-form-group>

    <template v-if="f.isMulti">
      <b-form-group
        :label="$t('kind.select.optionType.label')"
      >
        <b-form-radio-group
          v-model="f.options.selectType"
          :options="selectOptions"
          stacked
          @change="updateIsUniqueMultiValue"
        />
      </b-form-group>

      <b-form-group v-if="shouldAllowDuplicates">
        <b-form-checkbox
          v-model="f.options.isUniqueMultiValue"
          :value="false"
          :unchecked-value="true"
        >
          {{ $t('kind.select.allow-duplicates') }}
        </b-form-checkbox>
      </b-form-group>
    </template>
  </div>
</template>

<script>
import base from './base'
import { VueSelect } from 'vue-select'

export default {
  i18nOptions: {
    namespaces: 'field',
  },

  components: {
    VueSelect,
  },

  extends: base,

  data () {
    return {
      selectOptions: [
        { text: this.$t('kind.select.optionType.default'), value: 'default', allowDuplicates: true },
        { text: this.$t('kind.select.optionType.multiple'), value: 'multiple' },
        { text: this.$t('kind.select.optionType.each'), value: 'each', allowDuplicates: true },
      ],
      roleOptions: [],
    }
  },

  computed: {
    shouldAllowDuplicates () {
      if (!this.f.isMulti) return false

      const { allowDuplicates } = this.selectOptions.find(({ value }) => value === this.f.options.selectType) || {}
      return !!allowDuplicates
    },
  },

  mounted () {
    this.$SystemAPI.roleList().then(({ set: roles = [] }) => {
      this.roleOptions = roles
    })
  },

  methods: {
    getOptionKey ({ roleID }) {
      return roleID
    },

    updateIsUniqueMultiValue (value) {
      const { allowDuplicates = false } = this.selectOptions.find(({ value: v }) => v === value) || {}
      if (!allowDuplicates) {
        this.f.options.isUniqueMultiValue = true
      }
    },
  },
}
</script>
