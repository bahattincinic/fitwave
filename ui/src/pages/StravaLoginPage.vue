<template>
  <div class="mt-4">
    <Card class="max-w-30rem">
      <template #title>Strava Login</template>
      <template #content>
        <div>
          <p class="mb-3">
            To log in, please click the "Login with Strava" button below. Once
            the login process is complete, Strava will redirect you back to our
            application.
          </p>
        </div>
        <Button
          :disabled="loading"
          severity="success"
          label="Login with Strava"
          icon="pi pi-user"
          @click="redirectToStrava()"
        />
      </template>
    </Card>
  </div>
</template>

<script>
import Button from 'primevue/button';
import {
  getStravaAccessToken,
  getStravaAuthorizationURL,
} from '@/services/auth';
import Card from 'primevue/card';
import { useHead } from '@unhead/vue';
import { useStravaStore } from '@/store/strava';
import { useUserStore } from '@/store/user';

export default {
  name: 'StravaLoginPage',
  components: {
    Button,
    Card,
  },
  setup() {
    useHead({ title: 'Strava Login' });

    const strava = useStravaStore();
    const user = useUserStore();

    return {
      strava,
      user,
    };
  },
  data() {
    return {
      loading: false,
      url: '',
    };
  },
  async mounted() {
    try {
      this.loading = true;

      await this.fetchURL();

      if (this.$route.query.code) {
        await this.login(this.$route.query.code);
      }
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
  methods: {
    redirectToStrava() {
      window.location.href = this.url;
    },
    async login(code) {
      const resp = await getStravaAccessToken(this.user.accessToken, code);
      this.strava.login(resp.access_token, resp.athlete);
      this.$router.push('/');
    },
    async fetchURL() {
      const resp = await getStravaAuthorizationURL(this.user.accessToken);
      this.url = resp.authorization_url;
    },
  },
};
</script>
