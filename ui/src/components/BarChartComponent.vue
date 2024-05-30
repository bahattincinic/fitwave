<template>
  <Chart type="bar" :data="chartData" :options="chartOptions" />
</template>

<script>
import Chart from 'primevue/chart';

export default {
  name: 'BarChartComponent',
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
        scales: {
          x: {
            beginAtZero: true,
          },
          y: {
            beginAtZero: true,
          },
        },
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
            label: 'Data',
            data,
          },
        ],
      };
    },
  },
};
</script>
