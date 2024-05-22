<template>
  <div>
    <h1>Settings Page</h1>
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
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';

export default {
  name: 'SettingsPage',
  components: {
    InputText,
    Button
  },
  setup() {
    const clientId = ref('');
    const clientSecret = ref('');

    onMounted(async () => {
      try {
        const config = await getUserConfig();
        clientId.value = config.client_id;
        clientSecret.value = config.client_secret;
      } catch (error) {
        console.error('Error fetching user config:', error);
      }
    });

    const saveSettings = async () => {
      try {
        await saveUserConfig({
          client_id: parseInt(clientId.value),
          client_secret: clientSecret.value
        });
        alert('Settings saved successfully!');
      } catch (error) {
        console.error('Error saving user config:', error);
        alert('Error saving settings.');
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
