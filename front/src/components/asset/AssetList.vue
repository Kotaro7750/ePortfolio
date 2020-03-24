<template>
  <div class="container-fluid">
    <Loading v-if="isLoading"/>
    <div v-else class="table-responsive" style="overflow:hidden">
      <b-table responsive :items="share_list_table" :fields="fields" :sort-by.sync="sortBy" :sort-desc.sync="sortDesc" >
      </b-table>
    </div>
  </div>
  
</template>

<script>
import firebase from 'firebase/app';
import  Loading  from "@/components/Loading.vue";

export default {
  name :'AssetList',
  components:{
    Loading,
  },
  data: function(){
    return {
      isLoading:true,
      share_list:function () {return [];},
      fields: [
          { key: 'ticker', sortable: true },
          { key: 'sector', sortable: true },
          { key: 'amount', sortable: true },
          { key: 'total_cost', sortable: true },
          { key: 'mean_cost', sortable: true },
          { key: 'dividened', sortable: true },
          { key: 'dividened_yield', sortable: true },
        ],
      sortBy: 'ticker',
      sortDesc: true,
    }
  },

  created(){
    this.updateList();
  },

  computed: {
    totalDividened: function () {
      let total = 0;
      for (let i in this.share_list) {
        total = total + Number(this.share_list[i].dividened);
      }
      return total;
    },

    share_list_table: function () {
      let ret = [];
      let total_amount=0;
      let total_total_cost=0;
      let total_dividened=0;
      for (let i in this.share_list) {
        total_amount += this.share_list[i].amount;
        total_total_cost += this.share_list[i].total_cost;
        total_dividened += this.share_list[i].dividened;

        ret.push({
          ticker:this.share_list[i].ticker,
          sector:this.share_list[i].sector,
          amount:this.share_list[i].amount,
          total_cost:this.share_list[i].total_cost,
          mean_cost:this.share_list[i].mean_cost,
          dividened:this.share_list[i].dividened.toFixed(3),
          dividened_yield:((this.share_list[i].dividened / this.share_list[i].total_cost)*100).toFixed(4),
        });
      }

      
      ret.push({
        ticker:"ALL",
        amount:total_amount,
        total_cost:total_total_cost.toFixed(3),
        mean_cost:(total_total_cost/total_amount).toFixed(3),
        dividened:total_dividened.toFixed(3),
        dividened_yield:((total_dividened/total_total_cost)*100).toFixed(4),
        _rowVariant:'info',
      });
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
