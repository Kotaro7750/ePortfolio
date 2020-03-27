<template>
  <div class="container-fluid">
    <Loading v-if="isLoading"/>
    <div v-else class="table-responsive" style="overflow:hidden">
      <b-table responsive hover :items="memo_list_table" :fields="fields" :sort-by.sync="sortBy" :sort-desc.sync="sortDesc" @row-clicked="memoDetail">
        <template v-slot:cell(action)="row">
          <b-button size="sm" @click="modalEdit(row.item)" class="mr-2" variant="warning">
            <b-icon-pencil></b-icon-pencil>
          </b-button>
          <b-button size="sm" @click="modalDelete(row.item.id)" class="mr-2" variant="danger">
            <b-icon-trash></b-icon-trash>
          </b-button>
        </template>
      </b-table>
    </div>

    <b-modal id="modal-edit" centered title="メモ編集" @ok="updateMemo">
      <b-form-group label="Title">
        <b-form-input type="text" v-model="edittingMemo.title"></b-form-input>
      </b-form-group>

      <b-form-group label="URL">
        <b-form-input type="text" v-model="edittingMemo.about_url"></b-form-input>
      </b-form-group>
    </b-modal>

    <b-modal id="modal-delete" centered title="以下のメモを削除していいですか？" @ok="deleteMemo">
      <p class="my-4">削除していいですか？</p>
    </b-modal>

  </div>
</template>

<script>
import  Loading  from "@/components/Loading.vue";
import firebase from 'firebase/app';

export default {
  name: 'Memo',
  components:{
    Loading,
  },

  data: function(){
    return {
      isLoading:true,
      memo_list:function () {return [];},
      edittingMemo:{
        id:-1,
        title:"",
        about_url:"",
      },

      deletedID:-1,
      fields: [
          { key: 'title', sortable: true },
          { key: 'action',label:"", sortable: false }
        ],
      sortBy: 'date',
      sortDesc: true,
    }
  },

  created(){
    this.updateList();
  },

  computed: {
    memo_list_table: function () {
      let ret = []
      for (let i in this.memo_list) {
        ret.push({
          id:this.memo_list[i].id,
          title:this.memo_list[i].title,
          about_url:this.memo_list[i].about_url,
        })
      }
      return ret
    },
  },

  methods:{
    memoDetail(row){
      this.$router.push({name:'MemoDetail' ,params: {id:row.id,title:row.title,about_url:row.about_url}});
    },

    modalDelete(id) {
      this.deletedID = id;
      this.$bvModal.show('modal-delete');
    },

    modalEdit(item) {
      this.edittingMemo.id = item.id;
      this.edittingMemo.title = item.title;
      this.edittingMemo.about_url = item.about_url;
      this.$bvModal.show('modal-edit');
    },

    updateList(){
      this.isLoading = true;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/memo'

        return fetch(url,{
          method:'GET',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },

        })
      }).then(res =>{
        if (res.ok) {
          this.isLoading = false;
          return res.json();
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).then(json =>{
        //JSON.parseにしないとプロミスが解決されないで謎配列が入ってしまう
        this.memo_list = JSON.parse(json.result);

      }).catch(err =>{
        alert(err);
      });

    },

    deleteMemo(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/memo/' + String(this.deletedID);

        return fetch(url,{
          method:'DELETE',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },
        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          this.updateList();
          alert("success");
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).catch(err =>{
        alert(err);
      });

    },

    updateMemo(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/memo/' + String(this.edittingMemo.id);

        return fetch(url,{
          method:'PUT',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify(
          {
            id:this.edittingMemo.id,
            title:this.edittingMemo.title,
            about_url:this.edittingMemo.about_url,
          }),

        })
      }.bind(this)).then(res =>{
        if (res.ok) {
          this.updateList();
          alert("success");
          this.closeEditor();
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).catch(err =>{
        alert(err);
      });
    },

    openEditor(id){
      this.edittingMemo.id = id;
    },

    closeEditor(){
      this.edittingMemo.id = -1;
    },
  }
}
</script>
