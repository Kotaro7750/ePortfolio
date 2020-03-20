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
</style>
