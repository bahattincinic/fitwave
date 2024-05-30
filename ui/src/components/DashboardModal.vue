<template>
  <Dialog
    :visible="visible"
    @update:visible="
      (newVal) => {
        if (!newVal) onClose();
      }
    "
    :modal="true"
    :header="title"
    :style="{ width: '30rem' }"
  >
    <div class="flex align-items-center gap-3 mb-3">
      <label for="username" class="font-semibold w-6rem">Name</label>
      <InputText v-model="form.name" id="name" />
    </div>
    <div class="flex justify-content-center gap-2">
      <Button
        :disabled="loading"
        type="button"
        label="Cancel"
        severity="secondary"
        @click="onClose"
      />
      <Button
        :disabled="loading || !form.name"
        type="button"
        :label="form.id ? 'Save' : 'Create'"
        @click="onSave"
      />
    </div>
  </Dialog>
</template>

<script>
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';

export default {
  name: 'DashboardModal',
  props: {
    row: {
      type: Object,
      default: () => {},
    },
    loading: {
      type: Boolean,
      default: false,
    },
    visible: {
      type: Boolean,
      default: false,
    },
  },
  components: {
    Dialog,
    InputText,
    Button,
  },
  data() {
    return {
      form: this.getInitialForm(this.row),
    };
  },
  emits: ['close', 'save'],
  watch: {
    row: {
      handler(newVal) {
        this.form = this.getInitialForm(newVal);
      },
      deep: true,
      immediate: true,
    },
    visible() {
      this.form = this.getInitialForm(this.row);
    },
  },
  computed: {
    title() {
      return this.form.id ? 'Edit Dashboard' : 'Create Dashboard';
    },
  },
  methods: {
    getInitialForm(initial) {
      const data = initial || this.row || {};
      return {
        id: data.id,
        name: data.name,
      };
    },
    onClose() {
      this.$emit('close');
    },
    onSave() {
      this.$emit('save', this.form);
    },
  },
};
</script>
