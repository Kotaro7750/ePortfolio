<template>
  <div>
    <h4>一覧</h4>
    <Loading v-if="isLoading"/>
    <ul type="circle" v-else>
      <li v-for="ticker in ticker_list" :key=ticker.id>
        {{ticker.ticker}}:{{ticker.dividened}}
        <button @click="openEditor(ticker.id)" v-if="edittingID != ticker.id">編集</button>
        <button @click="closeEditor()" v-else>閉じる</button>
        <button @click="deleteTicker(ticker.id)">削除</button>

        <TickerEditor :edittedTicker="ticker.ticker" :isEdit="true" :message="`更新`" v-if="edittingID==ticker.id" @dispatch="updateTicker"/>
      </li>
    </ul>

  </div>
</template>

<script>
import  TickerEditor  from "@/components/ticker/TickerEditor.vue";
import  Loading  from "@/components/Loading.vue";
import firebase from 'firebase';

export default {
  name: 'TickerList',
  components:{
    TickerEditor,
    Loading,
  },

  data: function(){
    return {
      isLoading:true,
      ticker_list:function () {return [];},
      edittingID:-1,
    }
  },

  created(){
    this.updateList();
  },

  methods:{
    updateList(){
      this.isLoading = true;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/ticker'

        return fetch(url,{
          method:'GET',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },

        })
      }).then(res =>{
        if (res.ok) {
          this.isLoading = false;
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

    deleteTicker(id){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/ticker/' + String(id);

        return fetch(url,{
          method:'DELETE',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },
        })
      }).then(res =>{
        if (res.ok) {
          this.updateList();
          alert("success");
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).catch(err =>{
        alert(err);
      });

    },

    updateTicker(input){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/ticker/' + String(this.edittingID);

        return fetch(url,{
          method:'PUT',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({id:this.edittingID,ticker:input.ticker,dividened:Number(input.dividened)}),

        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          this.updateList();
          alert("success");
          this.closeEditor();
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).catch(err =>{
        alert(err);
      });
    },

    openEditor(id){
      this.edittingID = id;
    },

    closeEditor(){
      this.edittingID = -1;
    },
  }
}
</script>
