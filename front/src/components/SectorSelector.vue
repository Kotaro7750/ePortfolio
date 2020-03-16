<template>
  <b-form-group label="Sector">
    <b-form-select :value="selected_sector" @input="$emit('input',$event)" :options="list"/>
  </b-form-group>
</template>

<script>
import firebase from 'firebase/app';

export default {
  name: 'SectorSelector',
  props: {
    selected_sector:Number,
  },
  data: function () {
      return {
          sector_list:function () {return [];},
        }
  },

  created(){
    this.updateList();
  },
  computed:{
    list(){
      let ret = [];
      for (let i in this.sector_list) {
        ret.push({
          value:this.sector_list[i].id,
          text:this.sector_list[i].sector,
        });
      }
        return ret;
    },
  },
  
  methods: {
    updateList(){
      firebase.auth().currentUser.getIdToken(true).then(function(idToken) {
        let url = process.env.VUE_APP_API_URL + '/sector'

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
        this.sector_list = JSON.parse(json.result);

      }).catch(err =>{
        alert(err);
      });

    },
  }
}
</script>
