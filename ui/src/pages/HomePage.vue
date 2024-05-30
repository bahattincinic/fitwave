<template>
  <div class="m-3">
    <div class="flex justify-content-between align-items-center">
      <div>
        <h1 class="pl-1">Dashboards</h1>
      </div>
      <div>
        <Button
          label="Create Dashboard"
          icon="pi pi-check"
          severity="warning"
          @click="openCreateModal"
        />
      </div>
    </div>

    <DataTable
      :value="dashboards"
      :loading="loading"
      selectionMode="single"
      @rowSelect="onRowSelect"
    >
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Name"></Column>
      <template #empty> No records found </template>
    </DataTable>

    <DashboardModal
      :visible="modalShow"
      :loading="loading"
      @save="onCreateDashboard"
      @close="modalShow = false"
    />
  </div>
</template>

<script>
import { useHead } from '@unhead/vue';
import { fetchDashboards, createDashboard } from '@/services/dashboars';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import DashboardModal from '@/components/DashboardModal';

export default {
  name: 'HomePage',
  components: {
    DataTable,
    Column,
    DashboardModal,
    Button,
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
      modalShow: false,
    };
  },
  methods: {
    async fetch() {
      try {
        this.loading = true;
        const response = await fetchDashboards();
        this.dashboards = response.results;
      } catch (error) {
        this.onError(error);
      } finally {
        this.loading = false;
      }
    },
    async onCreateDashboard(form) {
      try {
        await createDashboard({
          name: form.name,
        });
        this.$toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Dashboard has been created successfully',
          life: 3000,
        });
      } catch (error) {
        this.onError(error);
      } finally {
        this.loading = false;
        this.modalShow = false;
      }
      await this.fetch();
    },
    openCreateModal() {
      this.modalShow = true;
    },
    onRowSelect(event) {
      this.$router.push(`/dashboard/${event.data.id}`);
    },
    onError(err) {
      this.$toast.add({
        severity: 'error',
        summary: 'Error',
        detail: err.toString(),
        life: 3000,
      });
    },
  },
};
</script>
