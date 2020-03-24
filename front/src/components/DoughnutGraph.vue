<script>
import { Doughnut ,mixins} from 'vue-chartjs'

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
