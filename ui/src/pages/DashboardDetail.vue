<template>
  <div v-if="loading" class="m-3">
    <Skeleton class="mb-2"></Skeleton>
    <Skeleton width="10rem" class="mb-2"></Skeleton>
    <Skeleton width="5rem" class="mb-2"></Skeleton>
    <Skeleton height="2rem" class="mb-2"></Skeleton>
    <Skeleton width="10rem" height="4rem"></Skeleton>
  </div>
  <div v-else class="m-3">
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

    <div v-if="!loading" class="mb-2 mt-2">
      <ComponentGrid
        :change-layout="changeLayout"
        :components="components"
        @refresh="refreshComponent"
        @edit="openEditComponentModal"
        @delete="deleteComponent"
      />
    </div>

    <DashboardModal
      :visible="modal.showUpdate"
      :loading="modal.loading"
      :row="dashboard"
      @save="onUpdateDashboard"
      @close="closeModal"
    />

    <ComponentModal
      :visible="modal.showComponent"
      :loading="modal.loading"
      :row="modal.form"
      @save="onSaveComponent"
      @close="closeModal"
      @set-loading="(v) => (this.modal.loading = v)"
    />
  </div>
</template>

<script>
import {
  runDashboard,
  getDashboard,
  deleteDashboard,
  updateDashboard,
} from '@/services/dashboars';
import { waitAsyncTask } from '@/services/user';
import {
  fetchComponents,
  createComponent,
  componentTypes,
  deleteComponent,
  updateComponent,
  runComponent,
} from '@/services/components';
import Button from 'primevue/button';
import TieredMenu from 'primevue/tieredmenu';
import Skeleton from 'primevue/skeleton';
import { useHead } from '@unhead/vue';
import ComponentGrid from '@/components/ComponentGrid';
import DashboardModal from '@/components/DashboardModal';
import ComponentModal from '@/components/ComponentModal';

export default {
  name: 'DashboardDetail',
  components: {
    Button,
    TieredMenu,
    ComponentGrid,
    Skeleton,
    DashboardModal,
    ComponentModal,
  },
  data() {
    return {
      componentTypes,
      loading: false,
      changeLayout: false,
      dashboard: {},
      components: [],
      modal: {
        showUpdate: false,
        showComponent: false,
        form: {},
        loading: false,
      },
      menuItems: [
        {
          label: 'Refresh',
          icon: 'pi pi-refresh',
          command: this.refreshDashboard,
        },
        {
          label: 'Change Layout',
          icon: 'pi pi-arrows-alt',
          command: () => {
            this.changeLayout = !this.changeLayout;
          },
        },
        {
          label: 'Update Dashboard',
          icon: 'pi pi-file-edit',
          command: () => {
            this.modal.showUpdate = true;
          },
        },
        {
          label: 'Create Component',
          icon: 'pi pi-file-edit',
          command: () => {
            this.modal.showComponent = true;
            this.modal.form = {};
          },
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

      if (this.components.length > 0) {
        await this.refreshDashboard();
      }
    },
    async refreshDashboard() {
      try {
        this.loading = true;

        const task = await waitAsyncTask(await runDashboard(this.dashboard.id));
        task.result.map((row) => {
          const component = this.components.find((comp) => comp.id === row.id);
          if (component) {
            component.results = row.results;
          }
        });
      } catch (error) {
        this.onError(error);
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
            this.onError(error);
          } finally {
            this.loading = false;
          }
        },
      });
    },
    async onUpdateDashboard(form) {
      try {
        this.modal.loading = true;

        await updateDashboard(this.dashboard.id, {
          name: form.name,
        });

        this.dashboard.name = form.name;
        this.$toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Dashboard has been updated successfully',
          life: 3000,
        });
      } catch (error) {
        this.onError(error);
      } finally {
        this.modal.loading = false;
        this.closeModal();
      }
    },
    async onSaveComponent(form) {
      this.modal.loading = true;

      const data = {
        name: form.name,
        type: form.type.code,
        query: form.query,
        configs: form.configs || null,
      };

      try {
        if (form.id) {
          await updateComponent(this.dashboard.id, form.id, data);
        } else {
          await createComponent(this.dashboard.id, data);
        }

        this.$toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Component has been saved successfully',
          life: 3000,
        });
      } catch (error) {
        this.onError(error);
      } finally {
        this.closeModal();
      }
      await this.fetch();
    },
    async refreshComponent(component) {
      const comp = this.components.find((comp) => comp.id === component.id);
      try {
        comp.loading = true;

        const task = await waitAsyncTask(
          await runComponent(this.dashboard.id, component.id)
        );
        const cmp = this.components.find((comp) => comp.id === component.id);
        if (cmp) {
          cmp.results = task.result;
        }
      } catch (error) {
        this.onError(error);
      } finally {
        comp.loading = false;
      }
    },
    async openEditComponentModal(component) {
      this.modal.showComponent = true;
      this.modal.form = component;
    },
    async deleteComponent(component) {
      this.$confirm.require({
        header: 'Confirmation',
        message: `Do you want to delete '${component.name}' Component`,
        icon: 'pi pi-info-circle',
        rejectClass: 'p-button-secondary p-button-outlined p-button-sm',
        acceptClass: 'p-button-danger p-button-sm',
        rejectLabel: 'Cancel',
        acceptLabel: 'Delete',
        accept: async () => {
          try {
            this.loading = true;
            await deleteComponent(this.dashboard.id, component.id);
          } catch (error) {
            this.onError(error);
          } finally {
            this.loading = false;
          }
          await this.fetch();
        },
      });
    },
    closeModal() {
      this.modal.showUpdate = false;
      this.modal.showComponent = false;
      this.modal.form = {};
      this.modal.loading = false;
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
