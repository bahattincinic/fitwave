<template>
  <div class="m-3">
    <Toast />
    <h1>Athletes</h1>

    <DataTable :value="athletes" :loading="loading">
      <Column field="id" header="ID"></Column>
      <Column field="firstname" header="First Name"></Column>
      <Column field="lastname" header="Last Name"></Column>
      <Column field="country" header="Country"></Column>
      <Column field="city" header="City"></Column>
      <template #empty> No records found </template>
    </DataTable>

    <Paginator
      :rows="20"
      :totalRecords="count"
      @page="handlePageChange"
    ></Paginator>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { fetchAthletes } from '@/services/athletes';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Paginator from 'primevue/paginator';
import { useHead } from '@unhead/vue';

export default {
  name: 'AthletesPage',
  components: {
    DataTable,
    Column,
    Paginator,
    Toast,
  },
  setup() {
    useHead({ title: 'Athletes' });

    const athletes = ref([]);
    const count = ref(0);
    let currentPage = 1;
    const loading = ref(false);
    const toast = useToast();

    const fetch = async () => {
      try {
        loading.value = true;
        const response = await fetchAthletes(currentPage);
        athletes.value = response.results;
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
      athletes,
      count,
      loading,
      handlePageChange,
    };
  },
};
</script>
