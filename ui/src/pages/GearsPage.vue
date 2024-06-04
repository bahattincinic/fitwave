<template>
  <div class="m-3">
    <h1>Gears</h1>

    <div v-if="loading" class="m-3">
      <Skeleton class="mb-2"></Skeleton>
      <Skeleton width="10rem" class="mb-2"></Skeleton>
      <Skeleton width="5rem" class="mb-2"></Skeleton>
      <Skeleton height="2rem" class="mb-2"></Skeleton>
      <Skeleton width="10rem" height="4rem"></Skeleton>
    </div>
    <DataTable
      v-else
      :value="gears"
      selectionMode="single"
      @rowSelect="onRowSelect"
    >
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Name"></Column>
      <Column :field="formatDistance" header="Distance"></Column>
      <Column :field="athleteName" header="Athlete"></Column>
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
      <vue-json-pretty :data="modal.data" />
    </Dialog>
  </div>
</template>

<script>
import { fetchGears } from '@/services/gears';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Paginator from 'primevue/paginator';
import Dialog from 'primevue/dialog';
import Skeleton from 'primevue/skeleton';
import VueJsonPretty from 'vue-json-pretty';
import { useHead } from '@unhead/vue';
import { useUserStore } from '@/store/user';

export default {
  name: 'GearsPage',
  components: {
    DataTable,
    Column,
    Paginator,
    Dialog,
    VueJsonPretty,
    Skeleton,
  },
  setup() {
    useHead({ title: 'Gears' });

    return {
      user: useUserStore(),
    };
  },
  data() {
    return {
      gears: [],
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
  methods: {
    athleteName(row) {
      return `${row.athlete.firstname} ${row.athlete.lastname}`;
    },
    formatDistance(row) {
      return `${(row.distance / 1000).toFixed(1)} km`;
    },
    handlePageChange(event) {
      this.currentPage = event.page + 1;
      this.fetch();
    },
    async fetch() {
      try {
        this.loading = true;
        const response = await fetchGears(
          this.user.accessToken,
          this.currentPage
        );
        this.gears = response.results;
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
    onRowSelect(event) {
      this.modal.data = event.data;
      this.modal.show = true;
    },
  },
};
</script>
