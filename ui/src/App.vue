<template>
  <div id="app">
    <Toast />
    <ConfirmDialog />
    <AppHeader />
    <router-view />
  </div>
</template>

<script>
import Toast from 'primevue/toast';
import ConfirmDialog from 'primevue/confirmdialog';
import AppHeader from './components/AppHeader.vue';

export default {
  components: {
    AppHeader,
    Toast,
    ConfirmDialog,
  },
};
</script>
