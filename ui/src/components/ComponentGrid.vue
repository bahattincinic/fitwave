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
        <TableComponent :rows="item.results" />
      </Panel>
    </GridItem>
  </GridLayout>
</template>

<script>
import { GridLayout, GridItem } from 'vue-grid-layout-v3';
import Panel from 'primevue/panel';
import TableComponent from '@/components/TableComponent';

export default {
  name: 'ComponentGrid',
  components: {
    GridLayout,
    GridItem,
    Panel,
    TableComponent,
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
  data() {
    return {
      layout: this.calculateLayout(this.components),
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
      }));
    },
    onLayoutUpdated(newLayout) {
      this.layout = newLayout;
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
