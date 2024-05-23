<template>
  <div>
    <Menubar :model="items" />
  </div>
</template>

<script>
import Menubar from 'primevue/menubar';
import { useRouter } from "vue-router";

export default {
  name: 'AppHeader',
  components: {
    Menubar
  },
  setup() {
    const router = useRouter();
    const items = [
      { label: 'Dashboard', icon: 'pi pi-fw pi-home', command: () => router.push('/') },
      { label: 'Settings', icon: 'pi pi-fw pi-cog', command: () => router.push('/settings') },
      { label: 'Activities', icon: 'pi pi-fw pi-calendar', command: () => router.push('/activities') },
      { label: 'Gears', icon: 'pi pi-fw pi-sitemap', command: () => router.push('/gears') },
      { label: 'Athletes', icon: 'pi pi-fw pi-user', command: () => router.push('/athletes') }
    ];

    return { items };
  }
}
</script>
