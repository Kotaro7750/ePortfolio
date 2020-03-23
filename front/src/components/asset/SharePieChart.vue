<template>
  <div class="container-fluid">
    <Loading v-if="isLoading"/>
    <b-card-group deck v-else> 
      <b-card title="ティッカー別配当割合">
        <DoughnutGraph :labels="chart_label" :dataset="chart_data" :backgroundColor="chart_color" />
      </b-card>
      <b-card title="配当履歴">
        <div class="dividend_history">
          <iframe width="600" height="371" seamless frameborder="0" scrolling="yes" :src="dividened_history_url"></iframe>
        </div>
      </b-card>
    </b-card-group > 
  </div>
  
</template>

<script>
import firebase from 'firebase/app';
import  Loading  from "@/components/Loading.vue";
import  DoughnutGraph  from "@/components/asset/DoughnutGraph.vue";

export default {
  name :'SharePieChart',
  components:{
    Loading,
    DoughnutGraph,
  },
  data: function(){
    return {
      isLoading:true,
      share_list:function () {return [];},
      dividened_history_url:process.env.VUE_APP_DIVIDENED_URL,
    }
  },

  created(){
    this.updateList();
  },

  computed: {
    chart_label: function () {
      let ret = [];
      for (let i in this.share_list) {
        ret.push(this.share_list[i].ticker);
      }
      return ret;
    },
    
    chart_data: function () {
      let ret = [];
      for (let i in this.share_list) {
        ret.push(this.share_list[i].dividened);
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
    hoge(){
      return "hoggehoe";
    },
    updateList(){
      this.isLoading=true;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/share'

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

<style scoped>
  .dividend_history /deep/ iframe{
    max-width: 100% !important;
  }
</style>
