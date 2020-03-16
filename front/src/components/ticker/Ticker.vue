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
      <b-form-group label="Ticker">
        <b-form-input type="text" v-model="addedTicker"></b-form-input>
      </b-form-group>
      <b-form-group label="Dividened">
        <b-form-input type="number" v-model="addedDividened"></b-form-input>
      </b-form-group>
    </b-modal>
  </div>
</template>

<script>
import  TickerList  from "@/components/ticker/TickerList.vue";
import firebase from 'firebase';

export default {
  name:'Ticker',
  components:{
    TickerList,
  },
  data: function(){
    return {
      addedTicker:"",
      addedDividened:0,
      hopedYield:3,
    }
  },

  methods:{
    resetModal() {
        this.addedTicker = '';
        this.addedDividened = 0;
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
          body: JSON.stringify({ticker:this.addedTicker,dividened:Number(this.addedDividened)}),

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
