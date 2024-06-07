<template>
  <Dialog
    :visible="visible"
    @update:visible="
      (newVal) => {
        if (!newVal) onClose();
      }
    "
    maximizable
    :modal="true"
    :header="title"
    :style="{ width: '50rem' }"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
  >
    <div class="card flex justify-content-center">
      <Stepper>
        <StepperPanel header="Query">
          <template #content="{ nextCallback }">
            <div class="step-content">
              <div class="flex align-items-center gap-3 mb-3">
                <label for="name" class="font-semibold w-6rem">Name</label>
                <InputText v-model="form.name" id="name" class="query-name" />
              </div>
              <div class="flex align-items-center gap-3 mb-3">
                <SQLEditor
                  :code="form.query"
                  :dataSchema="dataSchema"
                  height="200px"
                  @change="(value) => (this.form.query = value)"
                />
              </div>

              <div class="flex justify-content-end">
                <Button
                  icon="pi pi-code"
                  class="mr-1"
                  aria-label="Format SQL"
                  severity="secondary"
                  @click="formatSQL"
                  v-tooltip.bottom="'Click to format SQL.'"
                />
                <Button
                  icon="pi pi-receipt"
                  class="mr-1"
                  aria-label="Data Schema"
                  :severity="detailSection === 'data' ? 'success' : 'secondary'"
                  @click="toggleDetailSection('data')"
                  v-tooltip.bottom="'Click to see Data Schema.'"
                />
                <Button
                  icon="pi pi-question"
                  aria-label="Dynamic Query Parameters"
                  :severity="
                    detailSection === 'dynamic' ? 'success' : 'secondary'
                  "
                  @click="toggleDetailSection('dynamic')"
                  v-tooltip.bottom="'Click to see Dynamic Query Parameters.'"
                />
              </div>

              <div
                class="flex align-items-center gap-3 mb-3"
                v-if="detailSection === 'dynamic'"
              >
                <DataTable :value="dynamicQueryOptions">
                  <Column sortable field="option" header="Option" />
                  <Column sortable field="value" header="Value" />
                  <Column sortable field="example" header="Example Usage" />
                </DataTable>
              </div>

              <div
                class="flex align-items-center gap-3 mb-3"
                v-if="detailSection === 'data'"
              >
                <DataTable :value="dataSchema">
                  <Column sortable field="table_name" header="Table Name" />
                  <Column sortable field="field_db_name" header="Field name" />
                  <Column sortable field="type" header="Type" />
                </DataTable>
              </div>

              <div class="flex pt-4 justify-content-between">
                <Button
                  :disabled="loading"
                  label="Close"
                  icon="pi pi-arrow-left"
                  iconPos="left"
                  @click="onClose"
                />
                <Button
                  :disabled="loading || !form.name || !form.query"
                  label="Next"
                  icon="pi pi-arrow-right"
                  iconPos="right"
                  @click="showNextStep(nextCallback)"
                />
              </div>
            </div>
          </template>
        </StepperPanel>
        <StepperPanel header="Format">
          <template #content="{ prevCallback }">
            <div class="step-content">
              <div class="flex align-items-center gap-3 mb-3">
                <label for="username" class="font-semibold w-6rem">Type</label>
                <Dropdown
                  v-model="form.type"
                  :options="componentTypes"
                  optionLabel="name"
                  placeholder="Select a Type"
                  checkmark
                  :highlightOnSelect="false"
                  class="w-full md:w-14rem"
                />
              </div>
              <div
                v-if="isChartComponent"
                class="flex align-items-center gap-3 mb-3"
              >
                <label for="username" class="font-semibold w-6rem"
                  >Chart X Axis</label
                >
                <Dropdown
                  v-model="form.x"
                  :options="chartChoices"
                  optionLabel="name"
                  placeholder="Select a X Axis"
                  checkmark
                  :highlightOnSelect="false"
                  class="w-full md:w-14rem"
                />
              </div>
              <div
                v-if="isChartComponent"
                class="flex align-items-center gap-3 mb-3"
              >
                <label for="username" class="font-semibold w-6rem"
                  >Chart Y AXis</label
                >
                <Dropdown
                  v-model="form.y"
                  :options="chartChoices"
                  optionLabel="name"
                  placeholder="Select a Y Axis"
                  checkmark
                  :highlightOnSelect="false"
                  class="w-full md:w-14rem"
                />
              </div>
              <div class="mb-3" v-if="queryResult">
                <TableComponent
                  v-if="form.type && form.type.code === componentTypeEnum.table"
                  :rows="queryResult"
                />
                <PieChartComponent
                  v-else-if="
                    form.type &&
                    form.x &&
                    form.y &&
                    form.type.code === componentTypeEnum.pieChart
                  "
                  :rows="queryResult"
                  :x="form.x.code"
                  :y="form.y.code"
                />
                <BarChartComponent
                  v-else-if="
                    form.type &&
                    form.x &&
                    form.y &&
                    form.type.code === componentTypeEnum.barChart
                  "
                  :rows="queryResult"
                  :x="form.x.code"
                  :y="form.y.code"
                />
                <LineChartComponent
                  v-else-if="
                    form.type &&
                    form.x &&
                    form.y &&
                    form.type.code === componentTypeEnum.lineChart
                  "
                  :rows="queryResult"
                  :x="form.x.code"
                  :y="form.y.code"
                />
              </div>
              <div class="flex pt-4 justify-content-between">
                <Button
                  :disabled="loading"
                  label="Back"
                  icon="pi pi-arrow-left"
                  iconPos="left"
                  @click="prevCallback"
                />
                <Button
                  :disabled="loading"
                  label="Close"
                  icon="pi pi-times"
                  severity="warning"
                  iconPos="left"
                  @click="onClose"
                />
                <Button
                  label="Save"
                  icon="pi pi-save"
                  iconPos="right"
                  @click="onSave"
                />
              </div>
            </div>
          </template>
        </StepperPanel>
      </Stepper>
    </div>
  </Dialog>
</template>

<script>
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Dropdown from 'primevue/dropdown';
import Stepper from 'primevue/stepper';
import StepperPanel from 'primevue/stepperpanel';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import {
  componentTypes,
  componentTypeEnum,
  dynamicQueryOptions,
} from '@/services/components';
import TableComponent from '@/components/TableComponent';
import PieChartComponent from '@/components/PieChartComponent';
import BarChartComponent from '@/components/BarChartComponent';
import LineChartComponent from '@/components/LineChartComponent';
import { runQuery, waitAsyncTask } from '@/services/user';
import SQLEditor from '@/components/SQLEditor';
import prettier from 'prettier/standalone';
import sqlPlugin from 'prettier-plugin-sql';

export default {
  name: 'ComponentModal',
  props: {
    row: {
      type: Object,
      default: () => {},
    },
    dataSchema: {
      type: Array,
      default: () => [],
    },
    loading: {
      type: Boolean,
      default: false,
    },
    visible: {
      type: Boolean,
      default: false,
    },
    accessToken: {
      type: String,
      required: true,
    },
  },
  components: {
    Dialog,
    InputText,
    SQLEditor,
    Button,
    TableComponent,
    PieChartComponent,
    LineChartComponent,
    BarChartComponent,
    Dropdown,
    Stepper,
    StepperPanel,
    DataTable,
    Column,
  },
  data() {
    return {
      componentTypes,
      componentTypeEnum,
      dynamicQueryOptions,
      queryResult: null,
      detailSection: '',
      form: this.getInitialForm(this.row),
    };
  },
  emits: ['close', 'save', 'set-loading'],
  watch: {
    row: {
      handler(newVal) {
        this.form = this.getInitialForm(newVal);
      },
      deep: true,
      immediate: true,
    },
    visible() {
      this.form = this.getInitialForm();
      this.queryResult = null;
    },
  },
  computed: {
    title() {
      return this.form.id ? 'Edit Component' : 'Create Component';
    },
    isChartComponent() {
      if (!this.form.type) {
        return false;
      }
      const options = [
        componentTypeEnum.barChart,
        componentTypeEnum.lineChart,
        componentTypeEnum.pieChart,
      ];
      return options.includes(this.form.type.code);
    },
    chartChoices() {
      if (!this.queryResult || this.queryResult.length === 0) {
        return [];
      }
      return Object.keys(this.queryResult[0]).map((k) => {
        return {
          name: k,
          code: k,
        };
      });
    },
  },
  methods: {
    onClose() {
      this.$emit('close');
    },
    getInitialForm(initial) {
      const data = initial || this.row || {};
      const configs = data.configs || {};
      return {
        name: data.name,
        id: data.id,
        type: componentTypes.find((t) => t.code === data.type),
        query: data.query,
        x: configs.x
          ? this.chartChoices.find((t) => t.code === configs.x)
          : null,
        y: configs.y
          ? this.chartChoices.find((t) => t.code === configs.y)
          : null,
      };
    },
    onSave() {
      const data = {
        name: this.form.name,
        id: this.form.id,
        type: this.form.type,
        query: this.form.query,
      };
      if (this.isChartComponent) {
        data.configs = {
          x: this.form.x.code,
          y: this.form.y.code,
        };
      }

      this.$emit('save', data);
    },
    async runPreviewQuery() {
      try {
        this.$emit('set-loading', true);

        const task = await waitAsyncTask(
          this.accessToken,
          await runQuery(this.accessToken, {
            query: this.form.query,
          })
        );
        this.queryResult = task.result;
        return true;
      } catch (error) {
        this.onError(error);
        return false;
      } finally {
        this.$emit('set-loading', false);
      }
    },
    async showNextStep(nextCallback) {
      const status = await this.runPreviewQuery();
      const hasResult = !!this.queryResult;

      if (status && !hasResult) {
        this.onError(new Error('No result found'));
      }

      if (status && hasResult) {
        const { x, y } = this.getInitialForm();
        this.form.x = x;
        this.form.y = y;
        nextCallback();
      }
    },
    toggleDetailSection(section) {
      this.detailSection = section === this.detailSection ? '' : section;
    },
    async formatSQL() {
      this.form.query = await prettier.format(this.form.query, {
        parser: 'sql',
        plugins: [sqlPlugin],
        keywordCase: 'upper',
      });
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
.query-name {
  width: 524px;
}
.query-input {
  width: 524px;
  height: 239px;
}
.step-content {
  min-width: 40rem;
}
</style>
