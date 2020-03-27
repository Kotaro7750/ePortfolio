<template>
  <div class="container-fluid">
    <h1>{{$route.params.title}}</h1>
    <FinancialReportMD :url="$route.params.about_url"/>
  </div>
</template>

<script>
import  FinancialReportMD  from "@/components/ticker_detail/FinancialReportMD.vue";

export default {
  name:'MemoDetail',
  components:{
    FinancialReportMD,
  },
  data: function(){
    return {
    }
  },

  created(){
  },

  methods:{
  },
}
</script>
