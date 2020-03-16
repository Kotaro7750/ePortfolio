<template>
  <div>
    <h3>ティッカー</h3>

    <b-button-toolbar aria-label="Toolbar with button groups and input groups">

      <b-button-group class="mr-1">
        <b-button variant="success" v-b-modal.modal-add>
          <b-icon-plus></b-icon-plus>
        </b-button>
      </b-button-group>

      <b-input-group  prepend="目標利回り" >
        <b-form-input type="number" v-model="hopedYield"></b-form-input>
      </b-input-group>

    </b-button-toolbar>

    <TickerList ref="list" :yield="Number(hopedYield)"/>

    <b-modal centered id="modal-add" title="Add Ticker" @show="resetModal" @ok="addTicker">
      <SectorSelector :value="addedTicker.sector" @input="addedTicker.sector = $event"/>
      <b-form-group label="Ticker">
        <b-form-input type="text" v-model="addedTicker.ticker"></b-form-input>
      </b-form-group>
      <b-form-group label="Dividened">
        <b-form-input type="number" v-model="addedTicker.dividened"></b-form-input>
      </b-form-group>
    </b-modal>
  </div>
</template>

<script>
import  TickerList  from "@/components/ticker/TickerList.vue";
import  SectorSelector  from "@/components/SectorSelector.vue";
import firebase from 'firebase/app';

export default {
  name:'Ticker',
  components:{
    TickerList,
    SectorSelector,
  },
  data: function(){
    return {
      addedTicker:{
        ticker:"",
        dividened:0,
        sector:0,
      },
      hopedYield:3,
    }
  },

  methods:{
    resetModal() {
        this.addedTicker.ticker = '';
        this.addedTicker.dividened = 0;
        this.addedTicker.sector = 0;
      },
    addTicker(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/ticker'

        return fetch(url,{
          method:'POST',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({ticker:this.addedTicker.ticker, dividened:Number(this.addedTicker.dividened), sector_id:this.addedTicker.sector}),

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
