<template>
  <div>
    <ul type="circle">
      <li v-for="share in share_list" :key=share.id>
        {{share.ticker}}:{{share.amount}}株:配当 {{share.dividened}}$:総額{{share.total_cost}}$:平均取得金額{{share.mean_cost}}$
      </li>
    </ul>
    <div>
      配当総額 {{totalDividened}}$
    </div>
  </div>
</template>

<script>
import firebase from 'firebase';
export default {
  name :'ShareList',
  data: function(){
    return {
      share_list:function () {return [];},
    }
  },

  created(){
    this.updateList();
  },
  computed: {
    totalDividened: function () {
      let total = 0;
      for (let i in this.share_list) {
        total = total + Number(this.share_list[i].dividened);
      }
      return total;
    }
  },

  methods:{
    updateList(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/share'

        return fetch(url,{
          method:'GET',
          headers: {
              'Authorization': `Bearer ${idToken}`,
          },

        })
      }).then(res =>{
        if (res.ok) {
          return res.json();
        } else {
          return Promise.reject(new Error(`${res.status}: ${res.statusText}`));
        }

      }).then(json =>{
        //JSON.parseにしないとプロミスが解決されないで謎配列が入ってしまう
        this.share_list = JSON.parse(json.result);

      }).catch(err =>{
        alert(err);
      });
    },
  },
}
</script>
