<template>
  <div>
    <h1>Athletes</h1>

    <DataTable :value="athletes" :loading="loading" @onPage="handlePageChange">
      <Column field="id" header="ID"></Column>
      <Column field="firstname" header="First Name"></Column>
      <Column field="lastname" header="Last Name"></Column>
      <Column field="country" header="Country"></Column>
      <Column field="city" header="City"></Column>
    </DataTable>

    <Paginator :rows="20" :totalRecords="count" @page="handlePageChange"></Paginator>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import { fetchAthletes } from '@/services/athletes';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Paginator from 'primevue/paginator';

export default {
  name: 'AthletesPage',
  components: {
    DataTable,
    Column,
    Paginator
  },
  setup() {
    const athletes = ref([]);
    const count = ref(0);
    let currentPage = 1;
    const loading = ref(false);

    const fetch = async () => {
      try {
        loading.value = true;
        const response = await fetchAthletes(currentPage);
        athletes.value = response.results;
        count.value = response.count;
      } catch (error) {
        console.error('Error fetching athletes:', error);
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
      handlePageChange
    };
  }
}
</script>
