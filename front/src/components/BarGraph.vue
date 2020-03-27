<script>
import { Bar ,mixins} from 'vue-chartjs'

export default ({
  extends: Bar,
  mixins: [mixins.reactiveData],
  name:'BarGraph',

  props:{
    labels:Array,
    dataset:Array,
    backgroundColor:Array,
    text:String,
  },

  data: function () {
      return {
        options:{
          scales:{
              xAxes:[{stacked:true,}],
              yAxes:[{
                  id: 'y-axis',
                  stacked: true,
                  type: 'linear',
                  display: true,
                  position: 'left',
                  ticks: {
                      beginAtZero: true,
                  },
              }]
          },
        }
      }
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

  computed: {
    datasetParser(){
      let ret = []
      for (let i in this.dataset) {
        ret.push({
          label:this.dataset[i].label,
          data: this.dataset[i].data,
          backgroundColor: this.dataset[i].color,
                yAxisID: 'y-axis'
        })
      }
      return ret;
    }
  },

  methods: {
    updateData() {
      this.renderChart({
        labels: this.labels, 
        datasets: this.datasetParser
      },this.options);
    }
  }
})
</script>
