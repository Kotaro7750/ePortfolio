<template>
  <div >
    <input type="text"  v-model="ticker">
    <input type="number"  v-model="divedened">
    <button @click="dispatch">{{message}}</button>
  </div>
</template>

<script>
export default {
  props: {
    message: String,
  },
  data: function () {
      return {
          ticker:'',
          divedened:''
        }
  },
  methods:{
    dispatch(){
      alert(this.ticker + ' and ' + this.divedened);
    }
  }
}
</script>
