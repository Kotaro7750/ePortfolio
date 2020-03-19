<template>
  <div class="container-fluid">
    <h3>株式購入</h3>
    <b-container fluid>
        <b-col cols="1">
          <b-button variant="success" v-b-modal.modal-add>
            <b-icon-plus></b-icon-plus>
          </b-button>
        </b-col>
    </b-container>
    <PurchaseList ref="list"/>
    
    <b-modal id="modal-add" centered title="購入" @ok="addPurchase">
      <TickerSelector :value="addedPurchase.ticker" @input="addedPurchase.ticker = $event"/>
      <b-form-group label="Share">
        <b-form-input type="number" v-model="addedPurchase.share"></b-form-input>
      </b-form-group>

      <b-form-group label="Date">
        <b-form-input type="date" v-model="addedPurchase.date"></b-form-input>
      </b-form-group>

      <b-form-group label="Cost">
        <b-form-input type="number" v-model="addedPurchase.cost"></b-form-input>
      </b-form-group>
    </b-modal>
  </div>
</template>

<script>
import  PurchaseList  from "@/components/purchase/PurchaseList.vue";
import  TickerSelector  from "@/components/ticker/TickerSelector.vue";
import firebase from 'firebase/app';

export default {
  name:'Purchase',
  components:{
    PurchaseList,
    TickerSelector,
  },

  data: function () {
      return {
          addedPurchase:{
            ticker:-1,
            share:1,
            date:"",
            cost:0,
          }
        }
  },

  methods:{
    addPurchase(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase'

        return fetch(url,{
          method:'POST',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({ticker_id:this.addedPurchase.ticker,date:this.addedPurchase.date,cost:Number(this.addedPurchase.cost),share:Number(this.addedPurchase.share)}),

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
    },
  }
}
</script>
