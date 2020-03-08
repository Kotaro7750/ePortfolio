<template>
  <div>
    <h3>株式購入</h3>
    <PurchaseList ref="list"/>
    <PurchaseEditor :message='"追加"' @dispatch="addPurchase"/>
  </div>
</template>

<script>
import  PurchaseEditor  from "@/components/purchase/PurchaseEditor.vue";
import  PurchaseList  from "@/components/purchase/PurchaseList.vue";
import firebase from 'firebase';

export default {
  name:'Purchase',
  components:{
    PurchaseEditor,
    PurchaseList,
  },
  methods:{
    addPurchase(input){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/purchase'

        return fetch(url,{
          method:'POST',
          headers: {
              "Content-Type": "application/json",
              'Authorization': `Bearer ${idToken}`,
          },
          body: JSON.stringify({ticker_id:input.ticker,date:input.date,cost:Number(input.cost),share:Number(input.share)}),

        })
      }).then(res =>{
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
