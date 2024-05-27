<template>
  <div class="m-3">
    <Toast />
    <h1>Dashboards</h1>

    <DataTable :value="dashboards" :loading="loading">
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Name"></Column>
      <template #empty> No records found </template>
    </DataTable>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { useHead } from '@unhead/vue';
import { fetchDashboards } from '@/services/dashboars';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

export default {
  name: 'HomePage',
  components: {
    DataTable,
    Column,
    Toast,
  },
  setup() {
    useHead({ title: 'Dashboard' });

    const dashboards = ref([]);
    const loading = ref(false);
    const toast = useToast();

    const fetch = async () => {
      try {
        loading.value = true;
        const response = await fetchDashboards();
        dashboards.value = response.results;
      } catch (error) {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
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
  },
};
</script>
