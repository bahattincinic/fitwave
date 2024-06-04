<template>
  <div id="app">
    <Toast />
    <ConfirmDialog />
    <AppHeader v-if="showHeader" />
    <router-view />
  </div>
</template>

<script>
import Toast from 'primevue/toast';
import ConfirmDialog from 'primevue/confirmdialog';
import AppHeader from './components/AppHeader.vue';
import { useUserStore } from '@/store/user';

export default {
  components: {
    AppHeader,
    Toast,
    ConfirmDialog,
  },
  setup() {
    return {
      user: useUserStore(),
    };
  },
  computed: {
    showHeader() {
      return !['Login', 'Setup'].includes(this.$route.name);
    },
  },
};
</script>
