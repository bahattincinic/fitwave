<template>
  <div>
    <Menubar :model="items" />
  </div>
</template>

<script>
import Menubar from 'primevue/menubar';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/store/user';

export default {
  name: 'AppHeader',
  components: {
    Menubar,
  },
  setup() {
    const router = useRouter();
    const user = useUserStore();

    return { router, user };
  },
  computed: {
    items() {
      const accessToken = this.user.accessToken;
      const user = this.user.user;

      return [
        {
          label: 'Dashboard',
          icon: 'pi pi-fw pi-home',
          command: () => this.router.push('/'),
        },
        {
          label: 'Settings',
          icon: 'pi pi-fw pi-cog',
          command: () => this.router.push('/app/settings'),
        },
        {
          label: 'Data',
          icon: 'pi pi-fw pi-server',
          items: [
            {
              label: 'Activities',
              icon: 'pi pi-fw pi-calendar',
              command: () => this.router.push('/app/activities'),
            },
            {
              label: 'Gears',
              icon: 'pi pi-fw pi-sitemap',
              command: () => this.router.push('/app/gears'),
            },
            {
              label: 'Athletes',
              icon: 'pi pi-fw pi-user',
              command: () => this.router.push('/app/athletes'),
            },
          ],
        },
        {
          label: accessToken ? user.firstname : 'Anonymous',
          icon: 'pi pi-fw pi-user',
          items: [
            ...(accessToken
              ? [
                  {
                    label: 'Logout',
                    icon: 'pi pi-fw pi-sign-out',
                    command: () => this.logout(),
                  },
                ]
              : [
                  {
                    label: 'Login',
                    icon: 'pi pi-fw pi-sign-in',
                    command: () => this.router.push('/app/login'),
                  },
                ]),
          ],
        },
      ];
    },
  },
  methods: {
    logout() {
      const user = useUserStore();
      user.logout();
    },
  },
};
</script>
