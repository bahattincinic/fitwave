<template>
  <div>
    <Toast />
    <h1>Settings</h1>
    <form @submit.prevent="saveSettings">
      <div class="field">
        <label for="client_id">Client ID:</label>
        <InputText v-model="clientId" id="client_id" />
      </div>
      <div class="field">
        <label for="client_secret">Client Secret:</label>
        <InputText v-model="clientSecret" id="client_secret" />
      </div>
      <Button label="Save" type="submit" icon="pi pi-save" />
    </form>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { getUserConfig, saveUserConfig } from '@/services/user';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';

export default {
  name: 'SettingsPage',
  components: {
    InputText,
    Button,
    Toast
  },
  setup() {
    const clientId = ref('');
    const clientSecret = ref('');
    const toast = useToast();

    onMounted(async () => {
      try {
        const config = await getUserConfig();
        clientId.value = config.client_id;
        clientSecret.value = config.client_secret;
      } catch (error) {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      }
    });

    const saveSettings = async () => {
      try {
        await saveUserConfig({
          client_id: parseInt(clientId.value),
          client_secret: clientSecret.value
        });
        toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Settings saved successfully!',
          life: 3000,
        });
      } catch (error) {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error.toString(),
          life: 3000,
        });
      }
    };

    return {
      clientId,
      clientSecret,
      saveSettings
    };
  }
}
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
