<template>
  <b-form-select v-model="selected_ticker" :options="list"/>
</template>

<script>
import firebase from 'firebase';

export default {
  name: 'TickerSelector',
  data: function () {
      return {
          ticker_list:function () {return [];},
          selected_ticker:'',
        }
  },

  created(){
    this.updateList();
  },
  computed:{
    list(){
      let ret = [];
      for (let i in this.ticker_list) {
        ret.push({
          value:this.ticker_list[i].id,
          text:this.ticker_list[i].ticker,
        });
      }
        return ret;
    },
  },
  
  methods: {
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
