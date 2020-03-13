<template>
  <div>
    <h4>一覧</h4>
    <Loading v-if="isLoading"/>
    <b-table v-else responsive :items="purchase_history_table" :fields="fields" :sort-by.sync="sortBy" :sort-desc.sync="sortDesc">
      <template v-slot:cell(action)="row">
        <b-button size="sm" @click="modalEdit(row.item)" class="mr-2" variant="warning">編集</b-button>
        <b-button size="sm" @click="modalDelete(row.item.id)" class="mr-2" variant="danger">削除</b-button>
      </template>
    </b-table>

    <b-modal id="modal-edit" centered title="履歴編集" @ok="updatePurchase">
      <b-form-group label="Share">
        <b-form-input type="number" v-model="edittingShare"></b-form-input>
      </b-form-group>

      <b-form-group label="Date">
        <b-form-input type="date" v-model="edittingDate"></b-form-input>
      </b-form-group>

      <b-form-group label="Cost">
        <b-form-input type="number" v-model="edittingCost"></b-form-input>
      </b-form-group>
    </b-modal>

    <b-modal id="modal-delete" centered title="以下の履歴を削除していいですか？" @ok="deletePurchase">
      <p class="my-4">削除していいですか？</p>
    </b-modal>
  </div>
</template>

<script>
import firebase from 'firebase';
import  Loading  from "@/components/Loading.vue";

export default {
  name: 'PurchaseList',
  components:{
    Loading,
  },

  data: function(){
    return {
      isLoading:true,
      purchase_history_list:function () {return [];},

      edittingID:-1,
      edittingDate:'',
      edittingShare:0,
      edittingCost:0,

      deletedID:-1,
      fields: [
          { key: 'date', sortable: true },
          { key: 'ticker', sortable: true },
          { key: 'share', sortable: true },
          { key: 'cost', sortable: true },
          { key: 'action', sortable: false }
        ],
      sortBy: 'date',
      sortDesc: true,
    }
  },
  computed: {
    purchase_history_table: function () {
      let ret = []
      for (let i in this.purchase_history_list) {
        ret.push({
          id:this.purchase_history_list[i].id,
          date:this.purchase_history_list[i].date.split('T')[0],
          ticker:this.purchase_history_list[i].ticker,
          ticker_id:this.purchase_history_list[i].ticker_id,
          share:this.purchase_history_list[i].share,
          cost:this.purchase_history_list[i].cost,
        })
      }
      return ret
    },
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
      console.log(item);
      this.edittingID = item.id;
      this.edittingDate = item.date;
      this.edittingShare = item.share;
      this.edittingTicker = item.ticker_id;
      this.edittingCost = item.cost;
      this.$bvModal.show('modal-edit');
    },

    updateList(){
      this.isLoading = true;
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase'

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
        this.purchase_history_list = JSON.parse(json.result);

      }).catch(err =>{
        alert(err);
      });

    },

    deletePurchase(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase/' + String(this.deletedID);

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

    updatePurchase(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase/' + String(this.edittingID);

        return fetch(url,{
          method:'PUT',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({
            id:this.edittingID,
            ticker_id:this.edittingTicker,
            date:this.edittingDate,
            cost:Number(this.edittingCost),
            share:Number(this.edittingShare)
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
      this.edittingID = id;
    },

    closeEditor(){
      this.edittingID = -1;
    },
  }
}
</script>
