<template>
  <div>
    <MonacoEditor
      language="sql"
      :theme="theme"
      @mount="handleMount"
      @beforeMount="OnBeforeMount"
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

let autocompleteSet = false;

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
    dataSchema: {
      type: Array,
      default: () => [],
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
    OnBeforeMount(monaco) {
      if (!autocompleteSet) {
        monaco.languages.register({ id: 'sql' });
        monaco.languages.setMonarchTokensProvider('sql', {
          ignoreCase: true,
          tokenizer: {
            root: [
              [/\bselect\b/i, 'keyword'],
              [/\bfrom\b/i, 'keyword'],
              [/\bwhere\b/i, 'keyword'],
              [/\border by\b/i, 'keyword'],
              [/\bgroup by\b/i, 'keyword'],
            ],
          },
        });

        monaco.languages.registerCompletionItemProvider('sql', {
          provideCompletionItems: (model, position) => {
            const suggestions = [];
            const uniqueTables = new Set();

            const textUntilPosition = model.getValueInRange({
              startLineNumber: position.lineNumber,
              startColumn: 1,
              endLineNumber: position.lineNumber,
              endColumn: position.column,
            });

            const wordMatch = textUntilPosition.match(/(\w+)$/);
            const wordStart = wordMatch
              ? position.column - wordMatch[0].length
              : position.column;

            this.dataSchema.forEach((field) => {
              const range = {
                startLineNumber: position.lineNumber,
                startColumn: wordStart,
                endLineNumber: position.lineNumber,
                endColumn: position.column,
              };

              if (!uniqueTables.has(field.table_name)) {
                uniqueTables.add(field.table_name);
                suggestions.push({
                  label: field.table_name,
                  kind: monaco.languages.CompletionItemKind.Class,
                  insertText: field.table_name,
                  range,
                });
              }

              suggestions.push({
                label: field.field_db_name,
                kind: monaco.languages.CompletionItemKind.Field,
                insertText: field.field_db_name,
                detail: `${field.table_name}.${field.field_db_name} (${field.type})`,
                range,
              });
            });

            return { suggestions };
          },
        });

        autocompleteSet = true;
      }
    },
  },
};
</script>
