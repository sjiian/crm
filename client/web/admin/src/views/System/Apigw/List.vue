<template>
  <b-container
    class="py-3"
  >
    <span
      class="text-nowrap"
    >
      <c-settings-editor
        :settings="apigwSettings"
        @submit="onSettingsSubmit"
      />
    </span>
    <c-content-header
      :title="$t('title')"
    >
      <span
        class="text-nowrap"
      >
        <b-button
          v-if="canCreate"
          data-test-id="button-add"
          variant="primary"
          :to="{ name: 'system.apigw.new' }"
        >
          {{ $t('new') }}
        </b-button>
        <b-button
          v-if="$Settings.get('apigw.profilerEnabled', false)"
          data-test-id="button-profiler"
          class="ml-2"
          variant="info"
          :to="{ name: 'system.apigw.profiler' }"
        >
          {{ $t('profiler') }}
        </b-button>
        <c-permissions-button
          v-if="canGrant"
          data-test-id="button-permissions"
          resource="corteza::system:apigw-route/*"
          button-variant="light"
          class="ml-2"
        >
          <font-awesome-icon :icon="['fas', 'lock']" />
          {{ $t('permissions') }}
        </c-permissions-button>
      </span>
      <b-dropdown
        v-if="false"
        variant="link"
        right
        menu-class="shadow-sm"
        :text="$t('export')"
      >
        <b-dropdown-item-button variant="link">
          {{ $t('yaml') }}
        </b-dropdown-item-button>
      </b-dropdown>
    </c-content-header>

    <c-resource-list
      :primary-key="primaryKey"
      :filter="filter"
      :sorting="sorting"
      :pagination="pagination"
      :fields="fields"
      :items="items"
      :row-class="genericRowClass"
      :translations="{
        searchPlaceholder: $t('filterForm.query.placeholder'),
        notFound: $t('admin:general.notFound'),
        noItems: $t('admin:general.resource-list.no-items'),
        loading: $t('loading'),
        showingPagination: 'admin:general.pagination.showing',
        singlePluralPagination: 'admin:general.pagination.single',
        prevPagination: $t('admin:general.pagination.prev'),
        nextPagination: $t('admin:general.pagination.next'),
      }"
      @search="filterList"
    >
      <template #header>
        <c-resource-list-status-filter
          v-model="filter.deleted"
          data-test-id="filter-deleted-routes"
          :label="$t('filterForm.deleted.label')"
          :excluded-label="$t('filterForm.excluded.label')"
          :inclusive-label="$t('filterForm.inclusive.label')"
          :exclusive-label="$t('filterForm.exclusive.label')"
          @change="filterList"
        />
      </template>

      <template #actions="{ item }">
        <b-button
          size="sm"
          variant="link"
          :to="{ name: editRoute, params: { [primaryKey]: item[primaryKey] } }"
        >
          <font-awesome-icon
            :icon="['fas', 'pen']"
          />
        </b-button>
      </template>
    </c-resource-list>
  </b-container>
</template>

<script>
import * as moment from 'moment'
import listHelpers from 'corteza-webapp-admin/src/mixins/listHelpers'
import editorHelpers from 'corteza-webapp-admin/src/mixins/editorHelpers'
import { mapGetters } from 'vuex'
import { components } from '@cortezaproject/corteza-vue'
import CSettingsEditor from 'corteza-webapp-admin/src/components/Apigw/CSettingsEditor'
const { CResourceList } = components

export default {
  components: {
    CResourceList,
    CSettingsEditor,
  },

  mixins: [
    listHelpers,
    editorHelpers,
  ],

  i18nOptions: {
    namespaces: [ 'system.apigw' ],
    keyPrefix: 'list',
  },

  data () {
    return {
      id: 'routes',

      settings: [],

      primaryKey: 'routeID',
      editRoute: 'system.apigw.edit',

      filter: {
        query: '',
        deleted: 0,
      },

      sorting: {
        sortBy: 'createdAt',
        sortDesc: true,
      },

      state: {
        processing: false,
        success: false,
      },

      fields: [
        {
          key: 'endpoint',
          sortable: true,
        },
        {
          key: 'method',
          sortable: false,
        },
        {
          key: 'enabled',
          formatter: (v) => v ? 'Yes' : 'No',
        },
        {
          key: 'createdAt',
          sortable: true,
          formatter: (v) => moment(v).fromNow(),
        },
        {
          key: 'actions',
          label: '',
          tdClass: 'text-right',
        },
      ].map(c => ({
        ...c,
        // Generate column label translation key
        label: this.$t(`columns.${c.key}`),
      })),
    }
  },

  computed: {
    ...mapGetters({
      can: 'rbac/can',
    }),

    canCreate () {
      return this.can('system/', 'apigw-route.create')
    },

    canGrant () {
      return this.can('system/', 'grant')
    },

    apigwSettings () {
      if (this.settings.length > 0) {
        return this.settings.reduce((map, obj) => {
          const { name, value } = obj
          const split = name.split('.')

          if (split[0] === 'apigw') {
            map[name] = value
          }

          return map
        }, {})
      }
      return {}
    },
  },

  created () {
    this.fetchSettings()
  },

  methods: {
    onSettingsSubmit (settings) {
      this.state.processing = true

      const values = Object.entries(settings).map(([name, value]) => {
        return { name, value }
      })

      this.$SystemAPI.settingsUpdate({ values })
        .then(() => {
          this.animateSuccess('state')
          this.toastSuccess(this.$t('notification:gateway.settings.success'))
        })
        .catch(this.toastErrorHandler(this.$t('notification:gateway.settings.error')))
        .finally(() => {
          this.state.processing = false
        })
    },

    fetchSettings () {
      this.incLoader()

      this.$SystemAPI.settingsList()
        .then(settings => {
          this.settings = settings
        })
        .catch(e => {
          this.toastErrorHandler(this.$t('notification:settings.system.fetch.error'))(e)
          this.$router.push({ name: 'dashboard' })
        })
        .finally(() => {
          this.decLoader()
        })
    },

    items () {
      return this.procListResults(this.$SystemAPI.apigwRouteList(this.encodeListParams()))
    },
  },
}
</script>
