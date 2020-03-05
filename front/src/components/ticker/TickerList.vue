<template>
  <div>
    <h4>一覧</h4>
    <ul type="circle">
      <li v-for="ticker in ticker_list" :key=ticker.id>
        {{ticker.ticker}}:{{ticker.devidened}}
        <button @click="updateTicker(ticker.id)">編集</button>
        <button @click="deleteTicker(ticker.id)">削除</button>
      </li>
    </ul>

  </div>
</template>

<script>

export default {
  name: 'TickerList',
  data: function(){
    return {
      ticker_list:function () {return [];}
    }
  },
  created(){
    let url = process.env.VUE_APP_API_URL + '/ticker'
    fetch(url,{
      method:'GET'
    }).then(res =>{
      return res.json();
    }).then(result =>{
      console.log(result.result);
      this.ticker_list = JSON.parse(result.result);
    }).catch(err =>{
      alert(err);
    });
  },
  methods:{
    deleteTicker(id){
      alert(id+"削除");
    },
    updateTicker(id){
      alert(id+"更新");
    }
  }
}
</script>
