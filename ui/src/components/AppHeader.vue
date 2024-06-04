<template>
  <div>
    <Menubar :model="items" />
  </div>
</template>

<script>
import Menubar from 'primevue/menubar';
import { useStravaStore } from '@/store/strava';
import { useUserStore } from '@/store/user';

export default {
  name: 'AppHeader',
  components: {
    Menubar,
  },
  setup() {
    return {
      strava: useStravaStore(),
      user: useUserStore(),
    };
  },
  computed: {
    items() {
      const accessToken = this.strava.accessToken;
      const user = this.strava.user;

      return [
        {
          label: 'Dashboard',
          icon: 'pi pi-fw pi-home',
          command: () => this.$router.push('/'),
        },
        {
          label: 'Settings',
          icon: 'pi pi-fw pi-cog',
          command: () => this.$router.push('/app/settings'),
        },
        {
          label: 'Data',
          icon: 'pi pi-fw pi-server',
          items: [
            {
              label: 'Activities',
              icon: 'pi pi-fw pi-calendar',
              command: () => this.$router.push('/app/activities'),
            },
            {
              label: 'Gears',
              icon: 'pi pi-fw pi-sitemap',
              command: () => this.$router.push('/app/gears'),
            },
            {
              label: 'Athletes',
              icon: 'pi pi-fw pi-user',
              command: () => this.$router.push('/app/athletes'),
            },
          ],
        },
        {
          label: accessToken ? user.firstname : 'User',
          icon: 'pi pi-fw pi-user',
          items: [
            ...(this.user.loginNeeded()
              ? [
                  {
                    label: 'Logout App',
                    icon: 'pi pi-fw pi-sign-out',
                    command: () => {
                      this.user.logout();
                      this.$router.push('/app/login');
                    },
                  },
                ]
              : []),
            ...(accessToken
              ? [
                  {
                    label: 'Logout Strava',
                    icon: 'pi pi-fw pi-sign-out',
                    command: () => this.strava.logout(),
                  },
                ]
              : [
                  {
                    label: 'Login Strava',
                    icon: 'pi pi-fw pi-sign-in',
                    command: () => this.$router.push('/app/strava-login'),
                  },
                ]),
          ],
        },
      ];
    },
  },
};
</script>
