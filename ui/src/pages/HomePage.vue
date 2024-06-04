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

    <div v-if="loading" class="m-3">
      <Skeleton class="mb-2"></Skeleton>
      <Skeleton width="10rem" class="mb-2"></Skeleton>
      <Skeleton width="5rem" class="mb-2"></Skeleton>
      <Skeleton height="2rem" class="mb-2"></Skeleton>
      <Skeleton width="10rem" height="4rem"></Skeleton>
    </div>
    <DataTable
      v-else
      :value="dashboards"
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
import Skeleton from 'primevue/skeleton';
import DashboardModal from '@/components/DashboardModal';
import { useUserStore } from '@/store/user';

export default {
  name: 'HomePage',
  components: {
    DataTable,
    Column,
    DashboardModal,
    Button,
    Skeleton,
  },
  setup() {
    useHead({ title: 'Dashboard' });

    return {
      user: useUserStore(),
    };
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
        const response = await fetchDashboards(this.user.accessToken);
        this.dashboards = response.results;
      } catch (error) {
        this.onError(error);
      } finally {
        this.loading = false;
      }
    },
    async onCreateDashboard(form) {
      try {
        await createDashboard(this.user.accessToken, {
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
      this.$router.push(`/app/dashboard/${event.data.id}`);
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
