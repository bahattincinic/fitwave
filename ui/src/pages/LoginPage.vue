<template>
  <div class="container">
    <Card class="max-w-30rem">
      <template #title>
        <div class="text-center">Login</div>
      </template>
      <template #content>
        <form @submit.prevent="onLogin">
          <div class="field">
            <label for="username" class="pr-2">Username:</label>
            <InputText v-model="username" id="username" />
          </div>
          <div class="field">
            <label for="client_secret" class="pr-2">Password:</label>
            <Password v-model="password" id="password" />
          </div>
          <div class="text-center">
            <Button
              :disabled="loading || !username || !password"
              label="Login"
              type="submit"
              icon="pi pi-save"
            />
          </div>
        </form>
      </template>
    </Card>
  </div>
</template>

<script>
import { useHead } from '@unhead/vue';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import Button from 'primevue/button';
import Card from 'primevue/card';
import { useUserStore } from '@/store/user';
import { login } from '@/services/auth';
import { getUserConfig } from '@/services/config';

export default {
  name: 'LoginPage',
  components: {
    InputText,
    Password,
    Button,
    Card,
  },
  setup() {
    useHead({ title: 'Login' });

    return {
      user: useUserStore(),
    };
  },
  data() {
    return {
      username: '',
      password: '',
      loading: false,
    };
  },
  methods: {
    async onLogin() {
      this.loading = true;
      try {
        const resp = await login({
          username: this.username,
          password: this.password,
        });
        const config = await getUserConfig(resp.access_token);
        this.user.login(config, resp.access_token);
        this.$router.push('/');
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
  },
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 90vh;
}
</style>
