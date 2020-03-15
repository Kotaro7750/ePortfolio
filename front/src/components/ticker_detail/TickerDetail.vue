<template>
  <div>
    <h1>{{$route.params.ticker}}</h1>
    <b-tabs>
      <b-tab title="概要">
      </b-tab>
      <b-tab active title="決算">
        <b-input-group>
          <b-form-input type="number" v-model="year"></b-form-input>
          <b-button-group>
            <b-button @click="quarter=1" :variant="activeVariant(1)">Q1</b-button>
            <b-button @click="quarter=2" :variant="activeVariant(2)">Q2</b-button>
            <b-button @click="quarter=3" :variant="activeVariant(3)">Q3</b-button>
            <b-button @click="quarter=4" :variant="activeVariant(4)">Q4</b-button>
          </b-button-group>
        </b-input-group>

        <FinancialReport :ticker_id="Number($route.params.id)" :ticker="$route.params.ticker" :year="Number(year)" :quarter="Number(quarter)"/>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import  FinancialReport  from "@/components/ticker_detail/FinancialReport.vue";

export default {
  name:'TickerDetail',
  components:{
    FinancialReport,
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
  },
}
</script>
