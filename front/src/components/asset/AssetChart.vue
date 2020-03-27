<template>
  <div class="container-fluid">
    <b-card-group deck > 
      <DividendRatioForTicker />
      <DividendRatioForSector />
    </b-card-group > 
    <b-card-group deck > 
      <ExpectedDividenedForMonth />
      <b-card title="配当履歴">
        <div class="dividend_history">
          <iframe width="600" height="371" seamless frameborder="0" scrolling="yes" :src="dividened_history_url"></iframe>
        </div>
      </b-card>
    </b-card-group > 
  </div>
  
</template>

<script>
import DividendRatioForTicker  from "@/components/asset/DividendRatioForTicker.vue";
import DividendRatioForSector  from "@/components/asset/DividendRatioForSector.vue";
import ExpectedDividenedForMonth  from "@/components/asset/ExpectedDividenedForMonth.vue";

export default {
  name :'AssetChart',
  components:{
    DividendRatioForTicker,
    DividendRatioForSector,
    ExpectedDividenedForMonth,
  },
  data: function(){
    return {
      dividened_history_url:process.env.VUE_APP_DIVIDENED_URL,
    }
  },
}
</script>

<style scoped>
  .dividend_history /deep/ iframe{
    max-width: 100% !important;
  }
</style>
