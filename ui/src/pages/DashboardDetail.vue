<template>
  <div v-if="!loading" class="m-3">
    <div class="flex justify-content-between align-items-center">
      <div>
        <h1 class="pl-1">{{ dashboard.name }}</h1>
      </div>
      <div>
        <Button
          ref="toggleButton"
          type="button"
          icon="pi pi-cog"
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
  </div>
</template>

<script>
import {
  getDashboard,
  deleteDashboard,
  updateDashboard,
} from '@/services/dashboars';
import { fetchComponents } from '@/services/components';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import TieredMenu from 'primevue/tieredmenu';
import Dialog from 'primevue/dialog';
import { useHead } from '@unhead/vue';

export default {
  name: 'DashboardDetail',
  components: {
    Button,
    InputText,
    TieredMenu,
    Dialog,
  },
  data() {
    return {
      loading: false,
      dashboard: {},
      components: [],
      modal: {
        showUpdate: false,
        form: {},
      },
      menuItems: [
        {
          label: 'Update',
          icon: 'pi pi-file-edit',
          command: () => {
            this.openUpdateModal();
          },
        },
        {
          label: 'Delete',
          icon: 'pi pi-delete-left',
          command: () => {
            this.onDeleteDashboard();
          },
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
    openUpdateModal() {
      this.modal.showUpdate = true;
      this.modal.form.name = this.dashboard.name;
    },
    closeModal() {
      this.modal.showUpdate = false;
      this.modal.form = {};
    },
  },
};
</script>
