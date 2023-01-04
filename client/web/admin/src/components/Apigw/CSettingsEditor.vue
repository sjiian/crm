<template>
  <b-card
    data-test-id="card-edit-authentication"
    class="shadow-sm"
    header-bg-variant="white"
    footer-bg-variant="white"
  >
    <b-form
      @submit.prevent="$emit('submit', settings)"
    >
      <b-form-group
        :label="$t('profiler.label')"
        label-size="lg"
        label-cols="2"
      >
        <b-form-checkbox
          v-model="settings['apigw.profiler-enabled']"
          :value="true"
          :unchecked-value="false"
          class="mt-3"
        >
          {{ $t('profiler.enabled') }}
        </b-form-checkbox>
        <b-form-checkbox
          v-model="settings['apigw.profiler-global']"
          :value="true"
          :unchecked-value="false"
          class="mt-3"
        >
          {{ $t('profiler.global') }}
        </b-form-checkbox>
        <b-form-checkbox
          v-model="settings['apigw.proxy-follow-redirects']"
          :value="true"
          :unchecked-value="false"
          class="mt-3"
        >
          {{ $t('proxy.follow') }}
        </b-form-checkbox>
      </b-form-group>
    </b-form>

    <template #header>
      <h3 class="m-0">
        {{ $t('title') }}
      </h3>
    </template>

    <template #footer>
      <c-submit-button
        class="float-right"
        :processing="processing"
        :success="success"
        @submit="$emit('submit', settings)"
      />
    </template>
  </b-card>
</template>

<script>
import CSubmitButton from 'corteza-webapp-admin/src/components/CSubmitButton'

export default {
  name: 'CSettingsEditor',

  i18nOptions: {
    namespaces: 'system.apigw',
    keyPrefix: 'editor.settings',
  },

  components: {
    CSubmitButton,
  },

  props: {
    settings: {
      type: Object,
      required: true,
    },

    processing: {
      type: Boolean,
      value: false,
    },

    success: {
      type: Boolean,
      value: false,
    },
  },

  data () {
    return {
      defaultMinPwd: 8,
      defaultMinUppCaseChrs: 0,
      defaultMinLowCaseChrs: 0,
    }
  },
}
</script>
