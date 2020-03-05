<template>
  <div>
    <h4>一覧</h4>
    <ul type="circle">
      <li v-for="ticker in ticker_list" :key=ticker.id>
        {{ticker.ticker}}:{{ticker.devidens}}
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
      ticker_list:
        [
          {"id":0,"ticker":"CSCO","devidens":1.44},
          {"id":1,"ticker":"MMM","devidens":5.88}
        ],
    }
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
