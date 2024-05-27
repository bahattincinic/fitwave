<template>
  <div class="m-3">
    <h1>Activities</h1>

    <DataTable
      :value="activities"
      :loading="loading"
      selectionMode="single"
      @rowSelect="onRowSelect"
    >
      <Column field="id" header="ID"></Column>
      <Column field="type" header="Type"></Column>
      <Column field="name" header="Name"></Column>
      <Column :field="athleteName" header="Athlete"></Column>
      <Column field="gear.name" header="Gear"></Column>
      <template #empty> No records found </template>
    </DataTable>

    <Paginator
      :rows="20"
      :totalRecords="count"
      @page="handlePageChange"
    ></Paginator>

    <Dialog
      v-model:visible="modal.show"
      modal
      header="Gear Detail"
      :style="{ width: '50rem' }"
    >
      <div class="mb-3">
        <Button
          :disabled="!accessToken || loading"
          type="button"
          label="Download GPX"
          severity="secondary"
          @click="downloadGPXFile"
          v-tooltip.right="
            'To download GPX file, you need to logged in to with Strava'
          "
        />
      </div>
      <div>
        <vue-json-pretty :data="modal.data" />
      </div>
    </Dialog>
  </div>
</template>

<script>
import { fetchActivities, getActivityGPX } from '@/services/activities';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Paginator from 'primevue/paginator';
import VueJsonPretty from 'vue-json-pretty';
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { useHead } from '@unhead/vue';
import { mapState } from 'pinia';
import { useUserStore } from '@/store/user';

export default {
  name: 'ActivitiesPage',
  components: {
    DataTable,
    Column,
    Paginator,
    Dialog,
    Button,
    VueJsonPretty,
  },
  setup() {
    useHead({ title: 'Activities' });
  },
  data() {
    return {
      activities: [],
      loading: false,
      currentPage: 1,
      count: 0,
      modal: {
        show: false,
        data: {},
      },
    };
  },
  mounted() {
    this.fetch();
  },
  computed: {
    ...mapState(useUserStore, ['accessToken']),
  },
  methods: {
    athleteName(row) {
      return `${row.athlete.firstname} ${row.athlete.lastname}`;
    },
    handlePageChange(event) {
      this.currentPage = event.page + 1;
      this.fetch();
    },
    async fetch() {
      try {
        this.loading = true;
        const response = await fetchActivities(this.currentPage);
        this.activities = response.results;
        this.count = response.count;
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
    async downloadGPXFile() {
      try {
        this.loading = true;
        const blob = await getActivityGPX(this.modal.data.id, this.accessToken);
        const url = window.URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        const fileName = `activity_${this.modal.data.id}.gpx`;
        link.setAttribute('download', fileName);
        document.body.appendChild(link);
        link.click();
        link.remove();
      } finally {
        this.loading = false;
      }
    },
    onRowSelect(event) {
      this.modal.data = event.data;
      this.modal.show = true;
    },
  },
};
</script>
