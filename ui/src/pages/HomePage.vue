<template>
  <div class="m-3">
    <h1>Dashboards</h1>

    <DataTable :value="dashboards" :loading="loading">
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Name"></Column>
      <template #empty> No records found </template>
    </DataTable>
  </div>
</template>

<script>
import { useHead } from '@unhead/vue';
import { fetchDashboards } from '@/services/dashboars';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

export default {
  name: 'HomePage',
  components: {
    DataTable,
    Column,
  },
  setup() {
    useHead({ title: 'Dashboard' });
  },
  mounted() {
    this.fetch();
  },
  data() {
    return {
      loading: false,
      dashboards: [],
    };
  },
  methods: {
    async fetch() {
      try {
        this.loading = true;
        const response = await fetchDashboards();
        this.dashboards = response.results;
      } catch (error) {
        this.$toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>
