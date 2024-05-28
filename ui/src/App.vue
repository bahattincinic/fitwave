<template>
  <div id="app">
    <Toast />
    <ConfirmPopup />
    <AppHeader />
    <router-view />
  </div>
</template>

<script>
import Toast from 'primevue/toast';
import ConfirmPopup from 'primevue/confirmpopup';
import AppHeader from './components/AppHeader.vue';

export default {
  components: {
    AppHeader,
    Toast,
    ConfirmPopup,
  },
};
</script>
