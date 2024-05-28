<template>
  <div v-if="!loading" class="m-3">
    <div class="flex justify-content-between align-items-center">
      <div>
        <h1 class="pl-1">{{ dashboard.name }}</h1>
      </div>
      <div>
        <Button
          label="Delete"
          icon="pi pi-check"
          severity="danger"
          @click="onDeleteDashboard($event)"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { getDashboard, deleteDashboard } from '@/services/dashboars';
import { fetchComponents } from '@/services/components';
import Button from 'primevue/button';
import { useHead } from '@unhead/vue';

export default {
  name: 'DashboardDetail',
  components: {
    Button,
  },
  data() {
    return {
      loading: false,
      dashboard: {},
      components: [],
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
    onDeleteDashboard(event) {
      this.$confirm.require({
        target: event.currentTarget,
        message: 'Do you want to delete this record?',
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
  },
};
</script>
