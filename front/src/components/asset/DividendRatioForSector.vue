<template>
  <b-card title="セクター別配当割合" >
    <Loading v-if="isLoading"/>
    <DoughnutGraph :labels="chart_label" :dataset="chart_data" :backgroundColor="chart_color" v-else/>
  </b-card>
</template>

<script>
import firebase from 'firebase/app';
import  Loading  from "@/components/Loading.vue";
import  DoughnutGraph  from "@/components/DoughnutGraph.vue";

export default {
  name :'DividendRatioForSector',
  components:{
    Loading,
    DoughnutGraph,
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
      let ret = [];
      for (let i in this.share_list) {
        ret.push(this.share_list[i].sector);
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
    updateList(){
      this.isLoading=true;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/share/sector'

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
