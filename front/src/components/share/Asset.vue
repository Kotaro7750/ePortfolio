<template>
  <div>
    <h3>資産</h3>
    <ShareList />
  </div>
</template>

<script>
import  ShareList  from "@/components/share/ShareList.vue";
export default {
  name:'Asset',
  components:{
    ShareList,
  },
}
</script>
