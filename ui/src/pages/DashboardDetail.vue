<template>
  <div class="m-3">
    <div class="flex justify-content-between align-items-center">
      <div>
        <h1 class="pl-1">{{ dashboard.name }}</h1>
      </div>
      <div>
        <Button
          ref="toggleButton"
          type="button"
          icon="pi pi-cog"
          :disabled="loading"
          @click="$refs.menu.toggle($event)"
          aria-haspopup="true"
          aria-controls="overlay_menu"
        />
        <TieredMenu ref="menu" id="overlay_menu" :model="menuItems" popup />
      </div>
    </div>

    <Dialog
      v-model:visible="modal.showUpdate"
      modal
      header="Update Dashboard"
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
          @click="onUpdateDashboard"
        />
      </div>
    </Dialog>
    <Dialog
      v-model:visible="modal.showCreateComponent"
      maximizable modal
      header="Create Component"
      :style="{ width: '50rem' }"
      :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
    >
      <div class="flex align-items-center gap-3 mb-3">
        <label for="username" class="font-semibold w-6rem">Name</label>
        <InputText v-model="modal.form.name" id="name" />
      </div>
      <div class="flex align-items-center gap-3 mb-3">
        <label for="username" class="font-semibold w-6rem">Type</label>
        <Dropdown
          v-model="modal.form.type"
          :options="componentTypes"
          optionLabel="name"
          placeholder="Select a Type"
          checkmark
          :highlightOnSelect="false"
          class="w-full md:w-14rem"
        />
      </div>
      <div class="flex align-items-center gap-3 mb-3">
        <label for="username" class="font-semibold w-6rem">Query</label>
        <Textarea v-model="modal.form.query" id="query" rows="5" cols="30" />
      </div>

      <div v-if="modal.form.result" class="mb-4">
        <TableComponent :rows="modal.form.result" />
      </div>

      <div class="flex justify-content-end gap-2">
        <Button
          :disabled="loading"
          type="button"
          label="Cancel"
          severity="secondary"
          @click="closeModal"
        />
        <Button
          :disabled="loading || !modal.form.query"
          type="button"
          label="Preview"
          severity="warning"
          @click="onPreviewQuery"
        />
        <Button
          :disabled="loading || !modal.form.name || !modal.form.type || !modal.form.query"
          type="button"
          label="Create"
          @click="onCreateComponent"
        />
      </div>
    </Dialog>
  </div>
</template>

<script>
import {
  getDashboard,
  deleteDashboard,
  updateDashboard,
} from '@/services/dashboars';
import {getTaskDetail, runQuery} from '@/services/user';
import { fetchComponents, createComponent, componentTypes } from '@/services/components';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Textarea from 'primevue/textarea';
import Dropdown from 'primevue/dropdown';
import TieredMenu from 'primevue/tieredmenu';
import Dialog from 'primevue/dialog';
import { useHead } from '@unhead/vue';
import TableComponent from "@/components/TableComponent";

export default {
  name: 'DashboardDetail',
  components: {
    Button,
    InputText,
    TieredMenu,
    Dialog,
    Textarea,
    Dropdown,
    TableComponent
  },
  data() {
    return {
      componentTypes,
      loading: false,
      dashboard: {},
      components: [],
      modal: {
        showUpdate: false,
        showCreateComponent : false,
        form: {},
      },
      menuItems: [
        {
          label: 'Update Dashboard',
          icon: 'pi pi-file-edit',
          command: this.openUpdateModal,
        },
        {
          label: 'Create Component',
          icon: 'pi pi-file-edit',
          command: this.openCreateComponentModal,
        },
        {
          label: 'Delete',
          icon: 'pi pi-delete-left',
          command: this.onDeleteDashboard,
        },
      ],
    };
  },
  mounted() {
    this.fetch();
  },
  methods: {
    async fetch() {
      const dashId = this.$route.params.id;
      try {
        this.loading = true;
        this.dashboard = await getDashboard(dashId);
        useHead({ title: this.dashboard.name });

        const resp = await fetchComponents(dashId);
        this.components = resp.results;
      } catch (error) {
        this.$router.push('/');
      } finally {
        this.loading = false;
      }
    },
    onDeleteDashboard() {
      this.$confirm.require({
        header: 'Confirmation',
        message: 'Do you want to delete this dashboard?',
        icon: 'pi pi-info-circle',
        rejectClass: 'p-button-secondary p-button-outlined p-button-sm',
        acceptClass: 'p-button-danger p-button-sm',
        rejectLabel: 'Cancel',
        acceptLabel: 'Delete',
        accept: async () => {
          try {
            this.loading = true;
            await deleteDashboard(this.dashboard.id);
            this.$router.push('/');
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
      });
    },
    async onUpdateDashboard() {
      try {
        await updateDashboard(this.dashboard.id, {
          name: this.modal.form.name,
        });
        this.$toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Dashboard has been updated successfully',
          life: 3000,
        });
      } catch (error) {
        this.$toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      } finally {
        this.loading = false;
        this.closeModal();
      }
      await this.fetch();
    },
    async onCreateComponent() {
      try {
        await createComponent(this.dashboard.id, {
          name: this.modal.form.name,
          type: this.modal.form.type.code,
          query: this.modal.form.query,
        });
        this.$toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Component has been created successfully',
          life: 3000,
        });
      } catch (error) {
        this.$toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      } finally {
        this.loading = false;
        this.closeModal();
      }
      await this.fetch();
    },
    async onPreviewQuery() {
      try {
        this.loading = true;
        const task = await this.waitTask(await runQuery({
          query: this.modal.form.query
        }));
        this.modal.form.result = task.result;
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
    async waitTask(task) {
      const delay = async (ms) => {
        return new Promise((resolve) => setTimeout(resolve, ms));
      };

      let taskStatus = task.status;
      while (taskStatus !== 'success') {
        await delay(1000);
        task = await getTaskDetail(task.id);
        taskStatus = task.status;
      }

      return task;
    },
    openUpdateModal() {
      this.modal.showUpdate = true;
      this.modal.form.name = this.dashboard.name;
    },
    openCreateComponentModal() {
      this.modal.showCreateComponent = true;
      this.modal.form = {
        name: '',
        query: '',
        type: '',
      }
    },
    closeModal() {
      this.modal.showUpdate = false;
      this.modal.showCreateComponent = false;
      this.modal.form = {};
    },
  },
};
</script>
