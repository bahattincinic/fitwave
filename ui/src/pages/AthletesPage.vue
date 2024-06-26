<template>
  <div class="m-3">
    <h1>Athletes</h1>

    <div v-if="loading" class="m-3">
      <Skeleton class="mb-2"></Skeleton>
      <Skeleton width="10rem" class="mb-2"></Skeleton>
      <Skeleton width="5rem" class="mb-2"></Skeleton>
      <Skeleton height="2rem" class="mb-2"></Skeleton>
      <Skeleton width="10rem" height="4rem"></Skeleton>
    </div>
    <DataTable
      v-else
      :value="athletes"
      selectionMode="single"
      @rowSelect="onRowSelect"
    >
      <Column field="id" header="ID"></Column>
      <Column sortable field="firstname" header="First Name"></Column>
      <Column sortable field="lastname" header="Last Name"></Column>
      <Column sortable field="country" header="Country"></Column>
      <Column sortable field="city" header="City"></Column>
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
import { fetchAthletes } from '@/services/athletes';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Paginator from 'primevue/paginator';
import VueJsonPretty from 'vue-json-pretty';
import Dialog from 'primevue/dialog';
import { useHead } from '@unhead/vue';
import Skeleton from 'primevue/skeleton';
import { useUserStore } from '@/store/user';

export default {
  name: 'AthletesPage',
  components: {
    DataTable,
    Column,
    Paginator,
    Dialog,
    VueJsonPretty,
    Skeleton,
  },
  setup() {
    useHead({ title: 'Athletes' });

    return {
      user: useUserStore(),
    };
  },
  data() {
    return {
      athletes: [],
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
    handlePageChange(event) {
      this.currentPage = event.page + 1;
      this.fetch();
    },
    async fetch() {
      try {
        this.loading = true;
        const response = await fetchAthletes(
          this.user.accessToken,
          this.currentPage
        );
        this.athletes = response.results;
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
