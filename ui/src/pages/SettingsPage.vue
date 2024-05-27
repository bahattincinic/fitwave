<template>
  <div class="m-3">
    <Toast />
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
import { ref, onMounted } from 'vue';
import { mapState } from 'pinia';
import {
  getUserConfig,
  saveUserConfig,
  triggerSync,
  getTaskDetail,
} from '@/services/user';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
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
    Toast,
    TabView,
    TabPanel,
    Message,
  },
  setup() {
    useHead({ title: 'Settings' });

    const clientId = ref('');
    const clientSecret = ref('');
    const toast = useToast();
    const userStore = useUserStore();
    const loading = ref(false);

    onMounted(async () => {
      try {
        const config = await getUserConfig();
        clientId.value = config.client_id;
        clientSecret.value = config.client_secret;
        loading.value = true;
      } catch (error) {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      } finally {
        loading.value = false;
      }
    });

    return {
      clientId,
      clientSecret,
      toast,
      loading,
      userStore,
    };
  },
  computed: {
    ...mapState(useUserStore, ['user', 'accessToken']),
    isSyncEligible() {
      return !!this.clientId && !!this.clientSecret;
    },
  },
  methods: {
    async saveSettings() {
      try {
        this.loading = true;
        await saveUserConfig({
          client_id: parseInt(this.clientId),
          client_secret: this.clientSecret,
        });
        this.toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Settings saved successfully!',
          life: 3000,
        });
      } catch (error) {
        this.toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      } finally {
        this.loading = false;
      }
    },
    async delay(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
    async syncData() {
      try {
        this.loading = true;
        const task = await triggerSync(this.accessToken);

        let taskStatus = task.status;
        while (taskStatus !== 'success') {
          await this.delay(5000);
          const taskDetail = await getTaskDetail(task.id);
          taskStatus = taskDetail.status;
        }

        this.toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Sync operation completed successfully',
          life: 3000,
        });
      } catch (error) {
        this.toast.add({
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
