<template>
  <div class="container-fluid">
    <h1>{{$route.params.ticker}}</h1>
    <b-tabs>
      <b-tab active title="概要">
        <FinancialReportMD :url="$route.params.about_url" />
      </b-tab>
      <b-tab title="決算">
        <b-button-toolbar key-nav>

        <b-button-group class="mx-1">
          <b-button @click="year=year-1">&laquo;</b-button>
          <b-button @click="stepQuater(false)">&lsaquo;</b-button>
        </b-button-group>

        <b-button-group class="mx-1 w-25">
          <b-form-input type="number" v-model="year"></b-form-input>
        </b-button-group>

        <b-button-group class="mx-1">
          <b-button @click="stepQuater(true)">&rsaquo;</b-button>
          <b-button @click="year=year+1">&raquo;</b-button>
        </b-button-group>

        <b-button-group class="mx-1">
          <b-button @click="quarter=1" :variant="activeVariant(1)">Q1</b-button>
          <b-button @click="quarter=2" :variant="activeVariant(2)">Q2</b-button>
          <b-button @click="quarter=3" :variant="activeVariant(3)">Q3</b-button>
          <b-button @click="quarter=4" :variant="activeVariant(4)">Q4</b-button>
        </b-button-group>

        </b-button-toolbar>

        <FinancialReport :ticker_id="Number($route.params.id)" :ticker="$route.params.ticker" :year="Number(year)" :quarter="Number(quarter)"/>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import  FinancialReport  from "@/components/ticker_detail/FinancialReport.vue";
import  FinancialReportMD  from "@/components/ticker_detail/FinancialReportMD.vue";

export default {
  name:'TickerDetail',
  components:{
    FinancialReport,
    FinancialReportMD,
  },
  data: function(){
    return {
      content:'',
      year:0,
      quarter:0,
    }
  },

  created(){
    let today = new Date();
    this.year = today.getFullYear();
    this.quarter = Math.floor(today.getMonth()/4) + 1;
  },

  methods:{
    activeVariant(quarter) {
      if (quarter == this.quarter) {
        return "success";
      }else{
        return "";
      }
    },

    stepQuater(isNext){
      if (isNext) {
        if (this.quarter == 4) {
          this.quarter = 1;
          this.year = this.year + 1;
        }else{
          this.quarter = this.quarter + 1;
        }
      }else {
        if (this.quarter == 1) {
          this.quarter = 4;
          this.year = this.year - 1;
        }else{
          this.quarter = this.quarter - 1;
        }
      }
    },
  },
}
</script>
