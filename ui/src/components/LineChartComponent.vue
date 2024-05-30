<template>
  <Chart type="line" :data="chartData" :options="chartOptions" />
</template>

<script>
import Chart from 'primevue/chart';

export default {
  name: 'LineChartComponent',
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
            backgroundColor: '#42A5F5',
            borderColor: '#42A5F5',
            data,
            fill: false,
          },
        ],
      };
    },
  },
};
</script>
