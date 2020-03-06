<template>
  <div >
    <input type="text"  v-model="ticker" v-if="!isEdit">
    <input type="text"  v-model="ticker" disabled="disabled" v-else>

    <input type="number"  v-model="dividened">
    <button @click="dispatch">{{message}}</button>
  </div>
</template>

<script>
export default {
  props: {
    message: String,
    isEdit: Boolean,
    edittedTicker: String,
  },
  data: function () {
      return {
          ticker:this.edittedTicker,
          dividened:''
        }
  },
  methods:{
    dispatch(){
      this.$emit("dispatch",{
        ticker:this.ticker,dividened:this.dividened
        });
    }
  }
}
</script>
