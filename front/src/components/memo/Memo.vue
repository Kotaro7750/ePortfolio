<template>
  <div class="container-fluid">
    <h3>メモ</h3>

    <b-button-toolbar aria-label="Toolbar with button groups and input groups">

      <b-button-group class="mr-1">
        <b-button variant="success" v-b-modal.modal-add>
          <b-icon-plus></b-icon-plus>
        </b-button>
      </b-button-group>

    </b-button-toolbar>

    <MemoList ref="list" />

    <b-modal centered id="modal-add" title="Add Memo" @show="resetModal" @ok="addMemo">
      <b-form-group label="Title">
        <b-form-input type="text" v-model="addedMemo.title"></b-form-input>
      </b-form-group>

      <b-form-group label="URL">
        <b-form-input type="text" v-model="addedMemo.about_url"></b-form-input>
      </b-form-group>
    </b-modal>
  </div>
</template>

<script>
import  MemoList  from "@/components/memo/MemoList.vue";
import firebase from 'firebase/app';

export default {
  name:'Memo',
  components:{
    MemoList,
  },
  data: function(){
    return {
      addedMemo:{
        title:"",
        about_url:"",
      },
      hopedYield:3,
    }
  },

  methods:{
    resetModal() {
        this.addedMemo.title = '';
        this.addedMemo.about_url = '';
      },
    addMemo(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/memo'

        return fetch(url,{
          method:'POST',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify(
          {
            title:this.addedMemo.title,
            about_url:this.addedMemo.about_url,
          }),

        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          this.$refs.list.updateList();
          alert("success");
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }
      }).catch(err =>{
        alert(err);
      });
    }
  },
}
</script>
