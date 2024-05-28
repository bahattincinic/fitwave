<template>
  <DataTable :value="rows" :tableStyle="style">
    <Column v-for="col of columns" :key="col" :field="col" :header="col"></Column>
  </DataTable>
</template>

<script>
import DataTable from "primevue/datatable";
import Column from "primevue/column";

export default {
  name: 'TableComponent',
  props: {
    rows: {
      type: Array,
      default: () => []
    },
    style: {
      type: String
    }
  },
  components: {
    DataTable,
    Column
  },
  computed: {
    columns() {
      if (this.rows.length === 0) {
        return {}
      }
      return Object.keys(this.rows[0]);
    }
  }
}
</script>
