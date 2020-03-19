<script>
import { Doughnut ,mixins} from 'vue-chartjs'
import Chart from 'chart.js';

export default ({
  extends: Doughnut,
  mixins: [mixins.reactiveData],
  name:'DoughnutGraph',
  props:{
    labels:Array,
    dataset:Array,
    backgroundColor:Array,
    text:String,
  },
  watch: {
    dataset: function () {
      this.updateData();
    },
    labels: function () {
      this.updateData();
    },
    backgroundColor: function () {
      this.updateData();
    },
    text: function () {
      this.updateData();
    },
  },

  mounted() {
    this.updateData();
  },

  methods: {
    updateData() {
      let text = this.text;

      this.addPlugin({
        afterDraw(chart) {
          let ctx = chart.ctx

          let fontSize = (chart.height /10).toFixed(2);
          let fontStyle = 'normal';
          let fontFamily = "Helvetica Neue";
          ctx.fillStyle = '#000';
          ctx.font = Chart.helpers.fontString(fontSize, fontStyle, fontFamily);

          ctx.textAlign = 'center';
          ctx.textBaseline = 'middle';

          ctx.fillText(text, chart.width / 2 , chart.height / 2);
        }
      });

      this.chartData = {
        labels: this.labels, 
        datasets: [{
          data: this.dataset,
          backgroundColor: this.backgroundColor
          }
        ]
      };
    }
  }
})
</script>
