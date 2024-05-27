<template>
  <div id="app">
    <Toast />
    <AppHeader />
    <router-view />
  </div>
</template>

<script>
import Toast from 'primevue/toast';
import AppHeader from './components/AppHeader.vue';

export default {
  components: {
    AppHeader,
    Toast,
  },
};
</script>
