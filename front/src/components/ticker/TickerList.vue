<template>
  <div>
    <h4>一覧</h4>
    <Loading v-if="isLoading"/>
    <b-table v-else responsive :items="ticker_list" :fields="fields" :sort-by.sync="sortBy" :sort-desc.sync="sortDesc">
      <template v-slot:cell(action)="row">
        <b-button size="sm" @click="modalEdit(row.item)" class="mr-2" variant="warning">
          <b-icon-pencil></b-icon-pencil>
        </b-button>
        <b-button size="sm" @click="modalDelete(row.item.id)" class="mr-2" variant="danger">
          <b-icon-trash></b-icon-trash>
        </b-button>
      </template>
    </b-table>

    <b-modal id="modal-edit" centered title="ティッカー編集" @ok="updateTicker">
      <b-form-group label="Dividened">
        <b-form-input type="number" v-model="edittingDividened"></b-form-input>
      </b-form-group>
    </b-modal>

    <b-modal id="modal-delete" centered title="以下の履歴を削除していいですか？" @ok="deleteTicker">
      <p class="my-4">削除していいですか？</p>
    </b-modal>

  </div>
</template>

<script>
import  Loading  from "@/components/Loading.vue";
import firebase from 'firebase';

export default {
  name: 'TickerList',
  components:{
    Loading,
  },

  data: function(){
    return {
      isLoading:true,
      ticker_list:function () {return [];},
      edittingID:-1,
      edittingTicker:"",
      edittingDividened:0,

      deletedID:-1,
      fields: [
          { key: 'ticker', sortable: true },
          { key: 'dividened', sortable: true },
          { key: 'action', sortable: false }
        ],
      sortBy: 'date',
      sortDesc: true,
    }
  },

  created(){
    this.updateList();
  },

  methods:{
    modalDelete(id) {
      this.deletedID = id;
      this.$bvModal.show('modal-delete');
    },

    modalEdit(item) {
      this.edittingID = item.id;
      this.edittingTicker = item.ticker;
      this.edittingDividened = item.dividened;
      this.$bvModal.show('modal-edit');
    },

    updateList(){
      this.isLoading = true;
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
          this.isLoading = false;
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

    deleteTicker(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/ticker/' + String(this.deletedID);

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

    updateTicker(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/ticker/' + String(this.edittingID);

        return fetch(url,{
          method:'PUT',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({id:this.edittingID,ticker:this.edittingTicker,dividened:Number(this.edittingDividened)}),

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
      this.edittingID = id;
    },

    closeEditor(){
      this.edittingID = -1;
    },
  }
}
</script>
