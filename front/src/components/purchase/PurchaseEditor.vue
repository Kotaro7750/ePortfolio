<template>
  <div>
    <select v-model="selected_ticker">
      <option v-for="ticker in ticker_list" :key=ticker.id :value=ticker.id>{{ticker.ticker}}</option>
    </select>
    <input type="number" v-model="share">
    <input type="date" v-model="date">
    <input type="number" v-model="cost">
    <button @click="dispatch">{{message}}</button>
  </div>
</template>

<script>
import firebase from 'firebase';

export default {
  name:'PurchaseEditor',

  props: {
    message: String,
  },

  data: function () {
      return {
          ticker_list:function () {return [];},
          selected_ticker:'',
          cost:0,
          date:'',
          share:0,
        }
  },

  created(){
    this.updateList();
  },

  methods:{
    dispatch(){
      this.$emit("dispatch",{
        ticker:this.selected_ticker,share:this.share,cost:this.cost,date:this.date
        });
    },

    updateList(){
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
  }
}
</script>
