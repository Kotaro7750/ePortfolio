<template>
  <div>
    <h3>ティッカー</h3>
    <TickerList/>
    <TickerEditor :message='"追加"' @dispatch="addTicker"/>
  </div>
</template>

<script>
import  TickerList  from "@/components/ticker/TickerList.vue";
import  TickerEditor  from "@/components/ticker/TickerEditor.vue";

export default {
  name:'Ticker',
  components:{
    TickerList,
    TickerEditor,
  },

  methods:{
    addTicker(input){
      alert(input.ticker + ' ' + input.dividened);
    }
  },
}
</script>
