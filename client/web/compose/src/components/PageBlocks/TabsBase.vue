<template>
  <wrap
    v-bind="$props"
    :scrollable-body="false"
    v-on="$listeners"
  >
    <div
      v-if="!options.tabs.length"
      class="d-flex h-100 align-items-center justify-content-center"
    >
      <p class="mb-0">
        {{ $t('tabs.noTabs') }}
      </p>
    </div>

    <b-tabs
      v-else
      card
      :nav-class="navClass"
      :nav-wrapper-class="navWrapperClass"
      :content-class="contentClass"
      v-bind="{
        align: block.options.style.alignment,
        justified: block.options.style.justify === 'justify',
        pills: block.options.style.appearance === 'pills',
        tabs: block.options.style.appearance === 'tabs',
        small: block.options.style.appearance === 'small',
        vertical: block.options.style.orientation === 'vertical',
        end: block.options.style.position === 'end'
      }"
      no-fade
      class="h-100"
      :class="{ 'd-flex flex-column': block.options.style.orientation !== 'vertical' }"
    >
      <b-tab
        v-for="(tab, index) in tabbedBlocks"
        :key="`${getTabTitle(tab, index)}-${index}`"
        class="h-100"
        :title-item-class="getTitleItemClass(index)"
        :title-link-class="getTitleItemClass(index)"
        no-body
        :lazy="isTabLazy(tab)"
      >
        <template #title>
          <div v-if="editable" class="position-absolute">
            <b-button
              :variant="block.options.style.appearance === 'pills' ? 'secondary' : 'primary'"
              size="sm"
              class="tab-menu-item-tool mr-2"
              pill
              style="left: 5px;"
              @click="onTabMenuEditItemClick(tab)"
            >
              <font-awesome-icon
                :icon="['far', 'edit']"
                style="height: 9px;"
              />
            </b-button>

            <b-button
              variant="danger"
              size="sm"
              class="tab-menu-item-tool"
              pill
              style="left: 28px;"
              @click="onTabMenuDeleteItemClick(index)"
            >
              <font-awesome-icon
                :icon="['far', 'trash-alt']"
                style="height: 9px;"
              />
            </b-button>
          </div>

          <span>
            {{ getTabTitle(tab, index) }}
          </span>
        </template>

        <page-block-tab
          v-if="tab.block"
          v-bind="{ ...$attrs, ...$props, page, block: tab.block, blockIndex: index, boundingRect: { xywh: block.xywh} }"
          :record="record"
          :module="module"
          :magnified="magnified"
        />

        <div
          v-else
          class="d-flex h-100 align-items-center justify-content-center"
        >
          <p class="mb-0">
            {{ $t('tabs.noBlock') }}
          </p>
        </div>
      </b-tab>
    </b-tabs>
  </wrap>
</template>

<script>
import base from './base'
import { compose } from '@cortezaproject/corteza-js'
import { fetchID } from 'corteza-webapp-compose/src/lib/block'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'TabBase',

  components: {
    PageBlockTab: () => import('corteza-webapp-compose/src/components/PageBlocks'),
  },
  extends: base,

  computed: {
    tabbedBlocks () {
      return this.block.options.tabs.map(({ blockID, title }) => {
        const unparsedBlock = this.blocks.find(b => fetchID(b) === blockID)

        if (!unparsedBlock) {
          return { title }
        }

        let block = JSON.parse(JSON.stringify(unparsedBlock))

        // Blocks should display as Plain, to avoid card shadow/border
        block.style.wrap.kind = 'Plain'
        block.style.border.enabled = false
        block = compose.PageBlockMaker(block)
        return {
          block,
          title,
        }
      })
    },

    contentClass () {
      return `overflow-hidden mh-100 ${this.block.options.style.orientation === 'vertical' ? 'd-block' : 'flex-fill'}`
    },

    navClass () {
      const { orientation } = this.block.options.style
      const style = orientation === 'vertical' ? 'px-3' : 'px-2'
      return `bg-white ${style}`
    },

    navWrapperClass () {
      const { orientation, position } = this.block.options.style
      let border = 'border-bottom'
      let style = 'bg-white mh-100'

      if (orientation === 'vertical') {
        border = position === 'end' ? 'border-left' : 'border-right'
        style = `${style} overflow-auto`
      } else if (position === 'end') {
        border = 'border-top'
      }

      return `${border} ${style}`
    },
  },

  methods: {
    onTabMenuEditItemClick (tab) {
      const blockIndex = this.blocks.findIndex(block => fetchID(block) === tab.block.blockID)
      if (blockIndex > -1) {
        this.$emit('edit-tab-item-block', blockIndex)
      }
    },

    onTabMenuDeleteItemClick (tabIndex) {
      this.block.options.tabs.splice(tabIndex, 1)
    },

    getTitleItemClass (index) {
      const { justify, alignment } = this.block.options.style
      return `order-${index} text-truncate text-${alignment} ${justify !== 'none' ? 'flex-fill' : ''} position-relative`
    },

    getTabTitle ({ title, block = {} }, tabIndex) {
      const { title: blockTitle, kind } = block
      return title || blockTitle || kind || `${this.$t('tabs.tab')} ${tabIndex + 1}`
    },

    isTabLazy ({ block = {} }) {
      const { kind } = block
      return ['Calendar', 'Metric', 'Geometry'].includes(kind)
    },
  },
}
</script>

<style scoped>
.tab-menu-item-tool {
  height: 20px;
  width: 20px;
  padding: 2px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}
</style>
