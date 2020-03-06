<template>
  <div>
    <h4>一覧</h4>
    <ul type="circle">
      <li v-for="(ticker,index) in ticker_list" :key=ticker.id>
        {{ticker.ticker}}:{{ticker.devidened}}
        <button @click="openEditor(index)" v-if="edittingIndex != index">編集</button>
        <button @click="closeEditor()" v-else>閉じる</button>
        <button @click="deleteTicker(ticker.id)">削除</button>

        <TickerEditor :edittedTicker="ticker.ticker" :isEdit="true" :message="`更新`" v-if="edittingIndex==index" @dispatch="updateTicker"/>
      </li>
    </ul>

  </div>
</template>

<script>
import  TickerEditor  from "@/components/ticker/TickerEditor.vue";

export default {
  name: 'TickerList',
  components:{
    TickerEditor,
  },

  data: function(){
    return {
      ticker_list:function () {return [];},
      edittingIndex:-1,
    }
  },

  created(){
    let url = process.env.VUE_APP_API_URL + '/ticker'

    fetch(url,{
      method:'GET'

    }).then(res =>{
      if (res.ok) {
        return res.json();
      } else {
        return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
      }

    }).then(json =>{
      //JSON.parseにしないとプロミスが解決されないで謎配列が入ってしまう
      this.ticker_list = JSON.parse(json.result);

    }).catch(err =>{
      alert(err);
    });
  },

  methods:{
    deleteTicker(id){
      alert(id+"削除");
    },
    updateTicker(input){
      alert(input.ticker + ' ' + input.dividened);
      this.closeEditor();
    },
    openEditor(index){
      this.edittingIndex = index;
    },
    closeEditor(){
      this.edittingIndex = -1;
    },
  }
}
</script>
