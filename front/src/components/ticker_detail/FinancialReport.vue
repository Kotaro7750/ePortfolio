<template>
  <div>
    <Loading v-if="isLoading"/>
    <div v-else-if="noReport">
      <b-button variant="success" v-b-modal.modal-add>
        <b-icon-plus></b-icon-plus>
      </b-button>
    </div>
    <div v-else>
      <FinancialReportMD :url="targetURL" @edit="showEditModal"/>
    </div>

    
    <b-modal centered id="modal-add" title="Add FinancialReport" @ok="addFinancialReport">
      {{this.ticker}} {{this.year}}年第{{this.quarter}}四半期
      <b-form-group label="URL">
        <b-form-input type="text" v-model="targetURL"></b-form-input>
      </b-form-group>
    </b-modal>
    
    <b-modal centered id="modal-edit" title="Edit FinancialReport" @ok="editFinancialReport">
      {{this.ticker}} {{this.year}}年第{{this.quarter}}四半期
      <b-form-group label="URL">
        <b-form-input type="text" v-model="targetURL"></b-form-input>
      </b-form-group>
    </b-modal>
  </div>
</template>

<script>
import marked from "marked";
import firebase from 'firebase/app';
import  Loading  from "@/components/Loading.vue";
import  FinancialReportMD  from "@/components/ticker_detail/FinancialReportMD.vue";

export default {
  name:'FinancialReport',
  components:{
    Loading,
    FinancialReportMD,
  },

  props: {
    ticker_id:Number,
    ticker:String,
    year:Number,
    quarter:Number,
  },

  data: function(){
    return {
      content:'',
      noReport:false,
      isLoading:true,

      targetURL:'',
    }
  },

  created(){
    this.update();
  },

  watch:{
    year: function () {
      this.update();
    },
    quarter: function () {
      this.update();
    },
  },

  computed:{
    markdown:function () {
               return marked(this.content);
    },
  },

  methods:{
    showEditModal(){
      this.$bvModal.show('modal-edit');
    },

    update(){
      this.targetURL = '';
      this.isLoading = true;
      this.noReport = false;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/financial_report/' + Number(this.year) + '/' + this.quarter + '/' + this.ticker_id

        return fetch(url,{
          method:'GET',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },

        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          return res.json();
        } else {
          this.isLoading = true;
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).then(json =>{

        if (json.is_url_set) {
          this.noReport = false;
          this.isLoading = false;
          this.targetURL = json.url;
        }else{
          this.noReport = true;
          this.isLoading = false;
        }

      }).catch(err =>{
        alert(err);
      });

    },

    addFinancialReport(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/financial_report';

        return fetch(url,{
          method:'POST',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({ticker_id:this.ticker_id, year:this.year, quarter:this.quarter, url:this.targetURL})
        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          alert("success");
          this.update();
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }
      }).catch(err =>{
        alert(err);
        this.update();
      });

    },
    
    editFinancialReport(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/financial_report';

        return fetch(url,{
          method:'PUT',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({ticker_id:this.ticker_id, year:this.year, quarter:this.quarter, url:this.targetURL})
        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          alert("success");
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }
      }).catch(err =>{
        alert(err);
      });

      this.update();
    },
  }
}
</script>
