<template>
  <div class="m-3">
    <Toast />
    <h1>Activities</h1>

    <DataTable :value="activities" :loading="loading">
      <Column field="id" header="ID"></Column>
      <Column field="type" header="Type"></Column>
      <Column field="name" header="Name"></Column>
      <Column :field="athleteName" header="Athlete"></Column>
      <Column field="gear.name" header="Gear"></Column>
      <template #empty>
        No records found
      </template>
    </DataTable>

    <Paginator :rows="20" :totalRecords="count" @page="handlePageChange"></Paginator>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { fetchActivities } from '@/services/activities';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Paginator from 'primevue/paginator';

export default {
  name: 'ActivitiesPage',
  components: {
    DataTable,
    Column,
    Paginator,
    Toast
  },
  setup() {
    const activities = ref([]);
    const count = ref(0);
    let currentPage = 1;
    const loading = ref(false);
    const toast = useToast();

    const fetch = async () => {
      try {
        loading.value = true;
        const response = await fetchActivities(currentPage);
        activities.value = response.results;
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
      activities,
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
