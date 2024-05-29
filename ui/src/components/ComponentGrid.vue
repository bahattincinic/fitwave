<template>
  <GridLayout
    v-if="layout.length > 0"
    :layout="layout"
    :col-num="2"
    :row-height="400"
    :is-draggable="true"
    :is-resizable="false"
    :auto-size="true"
    @layout-updated="onLayoutUpdated"
  >
    <GridItem
      v-for="item in layout"
      :key="item.i"
      :i="item.i"
      :x="item.x"
      :y="item.y"
      :w="item.w"
      :h="item.h"
    >
      <Panel :header="item.name" class="panel">
        <template #icons>
          <Button
            class="p-panel-header-icon p-link mr-2"
            @click="onToggle(item, $event)"
          >
            <span class="pi pi-cog"></span>
          </Button>
        </template>
        <div v-if="item.loading" class="m-3">
          <Skeleton class="mb-2"></Skeleton>
          <Skeleton width="10rem" class="mb-2"></Skeleton>
          <Skeleton width="5rem" class="mb-2"></Skeleton>
          <Skeleton height="2rem" class="mb-2"></Skeleton>
          <Skeleton width="10rem" height="4rem"></Skeleton>
        </div>
        <TableComponent v-else :rows="item.results" />
      </Panel>
    </GridItem>
  </GridLayout>

  <Menu ref="menu" :model="menuItems" popup />
</template>

<script>
import { GridLayout, GridItem } from 'vue-grid-layout-v3';
import Panel from 'primevue/panel';
import Menu from 'primevue/menu';
import Button from 'primevue/button';
import Skeleton from 'primevue/skeleton';
import TableComponent from '@/components/TableComponent';

export default {
  name: 'ComponentGrid',
  components: {
    GridLayout,
    GridItem,
    Panel,
    Button,
    Menu,
    TableComponent,
    Skeleton,
  },
  props: {
    components: {
      type: Array,
      required: true,
    },
  },
  watch: {
    components: {
      handler: function (newVal) {
        this.layout = this.calculateLayout(newVal);
      },
      deep: true,
    },
  },
  emits: ['refresh', 'edit', 'delete'],
  data() {
    return {
      layout: this.calculateLayout(this.components),
      selectedItem: null,
      menuItems: [
        {
          label: 'Refresh',
          icon: 'pi pi-refresh',
          command: () => {
            this.$emit('refresh', this.selectedItem.component);
          },
        },
        {
          label: 'Edit',
          icon: 'pi pi-file-edit',
          command: () => {
            this.$emit('edit', this.selectedItem.component);
          },
        },
        {
          label: 'Delete',
          icon: 'pi pi-delete-left',
          command: () => {
            this.$emit('delete', this.selectedItem.component);
          },
        },
      ],
    };
  },
  methods: {
    calculateLayout(components) {
      return components.map((component, index) => ({
        x: index % 2,
        y: Math.floor(index / 2),
        w: 1,
        h: 1,
        i: component.id,
        name: component.name,
        results: component.results,
        loading: component.loading || false,
        component,
      }));
    },
    onLayoutUpdated(newLayout) {
      this.layout = newLayout;
    },
    onToggle(item, event) {
      this.selectedItem = item;
      this.$refs.menu.toggle(event);
    },
  },
};
</script>

<style scoped>
.panel {
  max-height: 100%;
  overflow-y: auto;
}
</style>
