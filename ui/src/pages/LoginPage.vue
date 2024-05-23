<template>
  <div>
    <Toast />

    <Button
      :disabled="loading"
      severity="success"
      label="Login with Strava"
      icon="pi pi-user"
      @click="redirectToStrava()"
    />
  </div>
</template>

<script>
import Toast from 'primevue/toast';
import Button from 'primevue/button';
import { onMounted, ref } from 'vue';
import { useToast } from 'primevue/usetoast';
import { getAccessToken, getAuthorizationURL } from '@/services/auth';
import {useRoute, useRouter} from 'vue-router';
import { useUserStore } from "@/store/user";
import Cookies from "js-cookie";

export default {
  name: 'LoginPage',
  components: {
    Toast,
    Button
  },
  setup() {
    const loading = ref(false);
    const toast = useToast();
    const user = useUserStore();
    const url = ref('');
    const { query } = useRoute();
    const router = useRouter()

    const login = async (code) => {
      const resp = await getAccessToken(code);
      user.setAccessToken(resp.access_token);
      user.setUser(resp.athlete);
      Cookies.set('accessToken', resp.access_token);
      await router.push('/');
    }

    const fetchURL = async () => {
      const resp = await getAuthorizationURL();
      url.value = resp.authorization_url;
    }

    onMounted(async () => {
      try {
        await fetchURL();
        if (query.code) {
          await login(query.code);
        }
      } catch (error) {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      }
    })

    return {
      loading,
      toast,
      url,
      user
    }
  },
  methods: {
    redirectToStrava() {
      window.location.href = this.url;
    }
  }
};
</script>
