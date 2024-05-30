<template>
  <Chart type="pie" :data="chartData" :options="chartOptions" />
</template>

<script>
import Chart from 'primevue/chart';

export default {
  name: 'PieChartComponent',
  props: {
    rows: {
      type: Array,
      default: () => [],
    },
    x: {
      type: String,
      required: true,
    },
    y: {
      type: String,
      required: true,
    },
  },
  components: {
    Chart,
  },
  data() {
    return {
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false,
      },
    };
  },
  computed: {
    chartData() {
      if (this.rows.length === 0) {
        return {};
      }

      const labels = this.rows.map((row) => row[this.x]);
      const data = this.rows.map((row) => row[this.y]);

      return {
        labels,
        datasets: [
          {
            data,
          },
        ],
      };
    },
  },
};
</script>
