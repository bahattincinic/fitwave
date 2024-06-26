<template>
  <div class="m-3">
    <h1>Activities</h1>

    <div v-if="loading" class="m-3">
      <Skeleton class="mb-2"></Skeleton>
      <Skeleton width="10rem" class="mb-2"></Skeleton>
      <Skeleton width="5rem" class="mb-2"></Skeleton>
      <Skeleton height="2rem" class="mb-2"></Skeleton>
      <Skeleton width="10rem" height="4rem"></Skeleton>
    </div>
    <DataTable
      v-else
      :value="activities"
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
        <Message v-if="!strava.accessToken" severity="info">
          To download GPX file, you need to logged in to with Strava
        </Message>
        <Button
          :disabled="!strava.accessToken || loading || modal.loading"
          type="button"
          label="Download GPX"
          severity="secondary"
          @click="downloadGPXFile"
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
import Skeleton from 'primevue/skeleton';
import VueJsonPretty from 'vue-json-pretty';
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import Message from 'primevue/message';
import { useHead } from '@unhead/vue';
import { useStravaStore } from '@/store/strava';
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
    Skeleton,
    Message,
  },
  setup() {
    useHead({ title: 'Activities' });

    return {
      strava: useStravaStore(),
      user: useUserStore(),
    };
  },
  data() {
    return {
      activities: [],
      loading: false,
      currentPage: 1,
      count: 0,
      modal: {
        loading: false,
        show: false,
        data: {},
      },
    };
  },
  mounted() {
    this.fetch();
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
        const response = await fetchActivities(
          this.user.accessToken,
          this.currentPage
        );
        this.activities = response.results;
        this.count = response.count;
      } catch (error) {
        this.onError(error);
      } finally {
        this.loading = false;
      }
    },
    async downloadGPXFile() {
      try {
        this.modal.loading = true;

        const blob = await getActivityGPX(
          this.modal.data.id,
          this.strava.accessToken,
          this.user.accessToken
        );

        const fileName = `activity_${this.modal.data.id}.gpx`;
        const url = window.URL.createObjectURL(blob);

        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', fileName);
        document.body.appendChild(link);
        link.click();
        link.remove();
      } catch (error) {
        this.onError(error);
      } finally {
        this.modal.loading = false;
      }
    },
    onError(err) {
      this.$toast.add({
        severity: 'error',
        summary: 'Error',
        detail: err.toString(),
        life: 3000,
      });
    },
    onRowSelect(event) {
      this.modal.data = event.data;
      this.modal.show = true;
      this.modal.loading = false;
    },
  },
};
</script>
