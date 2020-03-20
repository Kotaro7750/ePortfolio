<template>
  <div class="container-fluid">
    <Loading v-if="isLoading"/>
    <b-container v-else>
      <b-row align-h="center">
        <b-card align="left" boarder-variant="dark">

          <template v-slot:header>
            <b-button  @click="getMD">
              <b-icon-arrow-clockwise></b-icon-arrow-clockwise>
            </b-button>

            <b-button variant="warning" @click="$emit('edit')">
              <b-icon-pencil></b-icon-pencil>
            </b-button>
          </template>

          <div class="md-content" v-html="showMD" v-if="isOK"></div>
          <div v-else>無効なURLです</div>
        </b-card>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import marked from "marked";
import firebase from 'firebase/app';
import  Loading  from "@/components/Loading.vue";

export default {
  name:'FinancialReportMD',
  components:{
    Loading,
  },

  props: {
    url:String,
  },
  data: function(){
    return {
      content:'',
      isLoading:true,
      isOK:false,
    }
  },

  created(){
    this.getMD();
  },

  computed:{
    showMD:function () {
      return marked(this.content);
    },
  },
  methods:{
    getMD(){
      this.isLoading = true;
      this.isOK = false;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/financial_report?url=' + encodeURIComponent(this.url)

        return fetch(url,{
          method:'GET',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },

        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          this.isLoading = false;
          this.isOK = true;
          return res.text();
        } else {
          this.isLoading = false;
          this.isOK = false;
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).then(text =>{
        this.content = text;

      }).catch(() =>{
        this.isLoading = false;
        this.isOK = false;
      });
    },
  },
}
</script>

<style >
  .md-content iframe{
    max-width: 100% !important;
  }

  .md-content h2{
    padding: 0.4em 0.5em;/*文字の上下 左右の余白*/
    background: #f4f4f4;/*背景色*/
    border-left: solid 5px #000000;/*左線*/
    border-bottom: solid 3px #d7d7d7;/*下線*/
  }

  .md-content h3{
    padding: 0.25em 0.5em;/*上下 左右の余白*/
    background: transparent;/*背景透明に*/
    border-left: solid 5px #000000;/*左線*/
  }

  .md-content table{
  border-collapse:collapse;
  margin:0 auto;
  }
  
  .md-content td{
    border-bottom:1px dashed #000;
  }
  
  .md-content th,tr:last-child td{
    border-bottom:2px solid #000000;
  }
  
  .md-content td,th{
    padding:10px;
  }
</style>
