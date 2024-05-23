<template>
  <div>
    <Toast />
    <h1>Gears</h1>

    <DataTable :value="gears" :loading="loading">
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Name"></Column>
      <Column field="distance" header="Distance"></Column>
      <Column :field="athleteName" header="Athlete"></Column>
    </DataTable>

    <Paginator :rows="20" :totalRecords="count" @page="handlePageChange"></Paginator>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { fetchGears } from '@/services/gears';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Paginator from 'primevue/paginator';

export default {
  name: 'GearsPage',
  components: {
    DataTable,
    Column,
    Paginator,
    Toast
  },
  setup() {
    const gears = ref([]);
    const count = ref(0);
    let currentPage = 1;
    const loading = ref(false);
    const toast = useToast();

    const fetch = async () => {
      try {
        loading.value = true;
        const response = await fetchGears(currentPage);
        gears.value = response.results;
        count.value = response.count;
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

    const handlePageChange = (event) => {
      currentPage = event.page + 1;
      fetch();
    };

    onMounted(() => {
      fetch();
    });

    return {
      gears,
      count,
      loading,
      handlePageChange
    };
  },
  methods: {
    athleteName(row) {
      return `${row.athlete.firstname} ${row.athlete.lastname}`;
    }
  }
}
</script>
