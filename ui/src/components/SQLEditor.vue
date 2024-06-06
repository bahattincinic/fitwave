<template>
  <div>
    <MonacoEditor
      language="sql"
      :theme="theme"
      @mount="handleMount"
      v-model:value="codeValue"
      :options="options"
      :height="height"
      :width="width"
    />
  </div>
</template>

<script>
import { shallowRef } from 'vue';
import MonacoEditor from '@guolao/vue-monaco-editor';

export default {
  name: 'SQLEditor',
  components: { MonacoEditor },
  props: {
    code: {
      type: String,
      default: '',
    },
    height: {
      type: String,
      default: '300px',
    },
    width: {
      type: String,
      default: '700px',
    },
    theme: {
      type: String,
      default: 'vs-dark',
    },
  },
  emits: ['change'],
  setup() {
    return {
      editorRef: shallowRef(),
    };
  },
  data() {
    return {
      codeValue: this.code,
      options: {
        automaticLayout: true,
        formatOnType: true,
        formatOnPaste: true,
        autoIndent: true,
      },
    };
  },
  watch: {
    code(newValue) {
      this.codeValue = newValue;
    },
    codeValue(newValue) {
      this.$emit('change', newValue);
    },
  },
  methods: {
    handleMount(editor) {
      this.editorRef = editor;
    },
  },
};
</script>
