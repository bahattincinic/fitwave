<template>
  <div class="m-3">
    <h1>Settings</h1>
    <TabView>
      <TabPanel header="Config">
        <form @submit.prevent="saveSettings">
          <div class="field">
            <label for="client_id">Client ID:</label>
            <InputText v-model="clientId" id="client_id" />
          </div>
          <div class="field">
            <label for="client_secret">Client Secret:</label>
            <InputText v-model="clientSecret" id="client_secret" />
          </div>
          <Button
            :disabled="loading"
            label="Save"
            type="submit"
            icon="pi pi-save"
          />
        </form>
      </TabPanel>
      <TabPanel header="Sync Data">
        <Message v-if="!isSyncEligible" severity="error">
          You need to fill config from first to be able to sync your Strava
          data.
        </Message>
        <Message v-else-if="!accessToken" severity="info">
          You need to Login with Strava to be able to sync your data.
        </Message>

        <div v-if="isSyncEligible">
          <Button
            v-if="accessToken"
            :disabled="loading"
            severity="success"
            label="Sync Data"
            icon="pi pi-sync"
            @click="syncData()"
          />
          <Button
            v-else
            :disabled="loading"
            severity="success"
            label="Login with Strava"
            icon="pi pi-user"
            @click="$router.push('/login')"
          />
        </div>
      </TabPanel>
    </TabView>
  </div>
</template>

<script>
import { mapState } from 'pinia';
import {
  getUserConfig,
  saveUserConfig,
  triggerSync,
  waitAsyncTask,
} from '@/services/user';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import Message from 'primevue/message';
import { useUserStore } from '@/store/user';
import { useHead } from '@unhead/vue';

export default {
  name: 'SettingsPage',
  components: {
    InputText,
    Button,
    TabView,
    TabPanel,
    Message,
  },
  setup() {
    useHead({ title: 'Settings' });

    const userStore = useUserStore();

    return {
      userStore,
    };
  },
  data() {
    return {
      clientId: '',
      clientSecret: '',
      loading: false,
    };
  },
  computed: {
    ...mapState(useUserStore, ['user', 'accessToken']),
    isSyncEligible() {
      return !!this.clientId && !!this.clientSecret;
    },
  },
  mounted() {
    this.getConfig();
  },
  methods: {
    async getConfig() {
      this.loading = true;
      try {
        const config = await getUserConfig();
        this.clientId = config.client_id;
        this.clientSecret = config.client_secret;
      } catch (error) {
        this.onError(error);
      } finally {
        this.loading = false;
      }
    },
    async saveSettings() {
      try {
        this.loading = true;
        await saveUserConfig({
          client_id: parseInt(this.clientId),
          client_secret: this.clientSecret,
        });
        this.$toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Settings saved successfully!',
          life: 3000,
        });
      } catch (error) {
        this.onError(error);
      } finally {
        this.loading = false;
      }
    },
    async syncData() {
      try {
        this.loading = true;
        await waitAsyncTask(await triggerSync(this.accessToken));
        this.$toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Sync operation completed successfully',
          life: 3000,
        });
      } catch (error) {
        this.onError(error);
      } finally {
        this.loading = false;
      }
    },
    onError(err) {
      this.$toast.add({
        severity: 'error',
        summary: 'Error',
        detail: err.toString(),
        life: 3000,
      });
    },
  },
};
</script>

<style scoped>
.field {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
}

input {
  width: 100%;
}

button {
  margin-top: 1rem;
}
</style>
