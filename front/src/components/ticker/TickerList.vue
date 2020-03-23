<template>
  <div class="container-fluid">
    <Loading v-if="isLoading"/>
    <div v-else class="table-responsive" style="overflow:hidden">
      <b-table responsive hover :items="ticker_list_table" :fields="fields" :sort-by.sync="sortBy" :sort-desc.sync="sortDesc" @row-clicked="tickerDetail">
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

    <b-modal id="modal-edit" centered title="ティッカー編集" @ok="updateTicker">
      <SectorSelector :selected_sector="edittingTicker.sector" :value="edittingTicker.sector" @input="edittingTicker.sector = $event"/>

      <b-form-group label="Dividened">
        <b-form-input type="number" v-model="edittingTicker.dividened"></b-form-input>
      </b-form-group>

      <b-form-group label="AboutURL">
        <b-form-input type="text" v-model="edittingTicker.about_url"></b-form-input>
      </b-form-group>

      <b-form-group label="Color">
        <b-form-input type="color" v-model="edittingTicker.color"></b-form-input>
      </b-form-group>
    </b-modal>

    <b-modal id="modal-delete" centered title="以下の履歴を削除していいですか？" @ok="deleteTicker">
      <p class="my-4">削除していいですか？</p>
    </b-modal>

  </div>
</template>

<script>
import  Loading  from "@/components/Loading.vue";
import  SectorSelector  from "@/components/SectorSelector.vue";
import firebase from 'firebase/app';

export default {
  name: 'TickerList',
  components:{
    Loading,
    SectorSelector,
  },
  props: {
    yield:Number,
  },

  data: function(){
    return {
      isLoading:true,
      ticker_list:function () {return [];},
      edittingTicker:{
        id:-1,
        ticker:"",
        dividened:0,
        sector:0,
        about_url:"",
        color:"",
      },

      deletedID:-1,
      fields: [
          { key: 'ticker', sortable: true },
          { key: 'sector', sortable: true },
          { key: 'dividened', sortable: true },
          { key: 'expected_price', sortable: false },
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
    ticker_list_table: function () {
      let ret = []
      for (let i in this.ticker_list) {
        ret.push({
          id:this.ticker_list[i].id,
          ticker:this.ticker_list[i].ticker,
          dividened:this.ticker_list[i].dividened,
          sector:this.ticker_list[i].sector,
          sector_id:this.ticker_list[i].sector_id,
          expected_price:(this.ticker_list[i].dividened*100/this.yield).toFixed(2),
          about_url:this.ticker_list[i].about_url,
          color:this.ticker_list[i].color,
        })
      }
      return ret
    },
  },

  methods:{
    tickerDetail(row){
      this.$router.push({name:'TickerDetail' ,params: {id:row.id,ticker:row.ticker,about_url:row.about_url}});
    },

    modalDelete(id) {
      this.deletedID = id;
      this.$bvModal.show('modal-delete');
    },

    modalEdit(item) {
      this.edittingTicker.id = item.id;
      this.edittingTicker.ticker = item.ticker;
      this.edittingTicker.dividened = item.dividened;
      this.edittingTicker.sector = item.sector_id;
      this.edittingTicker.about_url = item.about_url;
      this.edittingTicker.color = item.color;
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
        let url = process.env.VUE_APP_API_URL + '/ticker/' + String(this.edittingTicker.id);

        return fetch(url,{
          method:'PUT',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify(
          {
            id:this.edittingTicker.id,
            ticker:this.edittingTicker.ticker,
            dividened:Number(this.edittingTicker.dividened),
            sector_id:this.edittingTicker.sector,
            about_url:this.edittingTicker.about_url,
            color:this.edittingTicker.color,
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
      this.edittingTicker.id = id;
    },

    closeEditor(){
      this.edittingTicker.id = -1;
    },
  }
}
</script>
