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

    <Dialog
      v-model:visible="modal.show"
      modal
      header="Create Dashboard"
      :style="{ width: '30rem' }"
    >
      <div class="flex align-items-center gap-3 mb-3">
        <label for="username" class="font-semibold w-6rem">Name</label>
        <InputText v-model="modal.form.name" id="name" />
      </div>
      <div class="flex justify-content-center gap-2">
        <Button
          :disabled="loading"
          type="button"
          label="Cancel"
          severity="secondary"
          @click="closeModal"
        />
        <Button
          :disabled="loading || !modal.form.name"
          type="button"
          label="Save"
          @click="onCreateDashboard"
        />
      </div>
    </Dialog>
  </div>
</template>

<script>
import { useHead } from '@unhead/vue';
import { fetchDashboards, createDashboard } from '@/services/dashboars';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Dialog from 'primevue/dialog';

export default {
  name: 'HomePage',
  components: {
    DataTable,
    Column,
    Button,
    Dialog,
    InputText,
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
      modal: {
        show: false,
        form: {
          name: '',
        },
      },
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
    async onCreateDashboard() {
      try {
        await createDashboard({
          name: this.modal.form.name,
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
        this.closeModal();
      }
      await this.fetch();
    },
    openCreateModal() {
      this.modal.show = true;
      this.modal.form.name = '';
    },
    closeModal() {
      this.modal.show = false;
      this.modal.form.name = '';
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
