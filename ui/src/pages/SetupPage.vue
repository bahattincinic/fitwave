<template>
  <div class="container">
    <Card class="w-3">
      <template #title>
        <div class="text-center">FitWave Setup</div>
      </template>
      <template #content>
        <form @submit.prevent="onSave">
          <div class="field">
            <label for="client_id" class="pr-2">Client ID:</label>
            <InputText v-model="clientId" id="client_id" />
          </div>
          <div class="field">
            <label for="client_secret" class="pr-2">Client Secret:</label>
            <InputText v-model="clientSecret" id="client_secret" />
          </div>
          <div class="field">
            <label for="login_type" class="pr-2">Login Type:</label>
            <Dropdown
              v-model="loginType"
              :options="loginTypes"
              optionLabel="name"
              placeholder="Select a Login Type"
              checkmark
              :highlightOnSelect="false"
              class="w-full md:w-14rem"
            />
          </div>
          <div
            v-if="loginType && loginType.code === loginTypeEnum.protected"
            class="field"
          >
            <label for="username" class="pr-2">Username:</label>
            <InputText v-model="username" id="username" />
          </div>
          <div
            v-if="loginType && loginType.code === loginTypeEnum.protected"
            class="field"
          >
            <label for="client_secret" class="pr-2">Password:</label>
            <Password v-model="password" id="password" />
          </div>

          <div class="text-center">
            <Button
              :disabled="loading || !clientId || !clientSecret || !loginType"
              label="Complete Setup"
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
import { useUserStore } from '@/store/user';
import { loginTypeEnum, loginTypes } from '@/services/auth';
import { completeSetup } from '@/services/config';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Dropdown from 'primevue/dropdown';

export default {
  name: 'SetupPage',
  components: {
    InputText,
    Password,
    Button,
    Card,
    Dropdown,
  },
  setup() {
    useHead({ title: 'Setup' });

    return {
      user: useUserStore(),
    };
  },
  data() {
    return {
      loginTypes,
      loginTypeEnum,
      loading: false,
      clientId: '',
      clientSecret: '',
      loginType: '',
      username: '',
      password: '',
    };
  },
  methods: {
    async onSave() {
      this.loading = true;
      const loginType = this.loginType && this.loginType.code;

      if (
        loginType === loginTypeEnum.protected &&
        (!this.username || !this.password)
      ) {
        this.$toast.add({
          severity: 'error',
          summary: 'Error',
          detail: 'username and password are required for protected login',
          life: 3000,
        });
        return;
      }

      try {
        const resp = await completeSetup({
          client_id: parseInt(this.clientId),
          client_secret: this.clientSecret,
          login_type: this.loginType.code,
          ...(loginType === loginTypeEnum.protected
            ? {
                login_username: this.username,
                login_password: this.password,
              }
            : {}),
        });

        this.user.setConfig(resp);
        this.user.setSetupCompleted(this.loginType, true);

        if (loginType === loginTypeEnum.anonymous) {
          this.$router.push('/');
        } else {
          this.$router.push('/app/login');
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
