<template>
  <b-card title="月別予想配当" >
    <Loading v-if="isLoading"/>
    <BarGraph :labels="chart_label" :dataset="chart_data" :backgroundColor="chart_color" v-else/>
  </b-card>
</template>

<script>
import firebase from 'firebase/app';
import  Loading  from "@/components/Loading.vue";
import  BarGraph  from "@/components/BarGraph.vue";

export default {
  name :'DividendRatioForTicker',
  components:{
    Loading,
    BarGraph,
  },
  data: function(){
    return {
      isLoading:true,
      share_list:function () {return [];},
    }
  },

  created(){
    this.updateList();
  },

  computed: {
    chart_label: function () {
      return ["1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"];
    },
    
    chart_data: function () {
      let ret = [];

      for (let i in this.share_list) {
        let dividened_month = this.share_list[i].dividened_month;

        let monthly_dividened = [];
        for (let j = 1; j <= 12; ++j) {
          if (dividened_month.find(elem => elem === String(j))) {
            monthly_dividened.push(this.share_list[i].dividened / dividened_month.length);
          }else{
            monthly_dividened.push(0);
          }
        }

        ret.push({
          label:this.share_list[i].ticker,
          data:monthly_dividened,
          color:this.share_list[i].color,
        });
      }
      return ret;
    },
    
    chart_color: function () {
      let ret = [];
      for (let i in this.share_list) {
        ret.push(this.share_list[i].color);
      }
      return ret;
    },
  },

  methods:{
    updateList(){
      this.isLoading=true;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/share/ticker'

        return fetch(url,{
          method:'GET',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },

        })
      }).then(res =>{
        if (res.ok) {
          this.isLoading=false;
          return res.json();
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).then(json =>{
        //JSON.parseにしないとプロミスが解決されないで謎配列が入ってしまう
        this.share_list = JSON.parse(json.result);

      }).catch(err =>{
        alert(err);
      });
    },
  },
}
</script>
