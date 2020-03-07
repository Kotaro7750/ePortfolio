<template>
  <div>
    <h3>ティッカー</h3>
    <TickerList ref="list"/>
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
      let url = process.env.VUE_APP_API_URL + '/ticker'

      fetch(url,{
        method:'POST',
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ticker:input.ticker,dividened:Number(input.dividened)}),

      }).then(res =>{
        if (res.ok) {
          this.$refs.list.updateList();
          alert("success");
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }
      }).catch(err =>{
        alert(err);
      });
    }
  },
}
</script>
