<template>
  <div>
    <h1>Dashboards</h1>

    <DataTable :value="dashboards" :loading="loading" @onPage="handlePageChange">
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Name"></Column>
    </DataTable>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import { fetchDashboards } from '@/services/dashboars';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

export default {
  name: 'HomePage',
  components: {
    DataTable,
    Column
  },
  setup() {
    const dashboards = ref([]);
    let currentPage = 1;
    const loading = ref(false);

    const fetch = async () => {
      try {
        loading.value = true;
        const response = await fetchDashboards(currentPage);
        dashboards.value = response.results;
      } catch (error) {
        console.error('Error fetching activities:', error);
      } finally {
        loading.value = false;
      }
    };

    onMounted(() => {
      fetch();
    });

    return {
      dashboards,
      loading,
    };
  }
}
</script>
