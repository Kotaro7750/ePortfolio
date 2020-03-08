<template>
  <div>
    <h4>一覧</h4>
    <ul type="circle">
      <li v-for="purchase_history in purchase_history_list" :key=purchase_history.id>
        {{purchase_history.ticker}}:{{purchase_history.date}}:{{purchase_history.share}}株:{{purchase_history.cost}}$
        <button @click="openEditor(purchase_history.id)" v-if="edittingID != purchase_history.id">編集</button>
        <button @click="closeEditor()" v-else>閉じる</button>
        <button @click="deletePurchase(purchase_history.id)">削除</button>

        <PurchaseEditor :message="`更新`" v-if="edittingID==purchase_history.id" @dispatch="updatePurchase"/>
      </li>
    </ul>

  </div>
</template>

<script>
import  PurchaseEditor  from "@/components/purchase/PurchaseEditor.vue";
import firebase from 'firebase';

export default {
  name: 'PurchaseList',
  components:{
    PurchaseEditor,
  },

  data: function(){
    return {
      purchase_history_list:function () {return [];},
      edittingID:-1,
    }
  },

  created(){
    this.updateList();
  },

  methods:{
    updateList(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase'

        return fetch(url,{
          method:'GET',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },

        })
      }).then(res =>{
        if (res.ok) {
          return res.json();
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).then(json =>{
        //JSON.parseにしないとプロミスが解決されないで謎配列が入ってしまう
        this.purchase_history_list = JSON.parse(json.result);

      }).catch(err =>{
        alert(err);
      });

    },

    deletePurchase(id){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase/' + String(id);

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

    updatePurchase(input){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase/' + String(this.edittingID);

        return fetch(url,{
          method:'PUT',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({id:this.edittingID, ticker_id:input.ticker,date:input.date, cost:Number(input.cost), share:Number(input.share)}),

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
