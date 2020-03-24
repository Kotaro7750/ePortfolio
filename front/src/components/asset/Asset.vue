<template>
  <div class="container-fluid">
    <h3>資産</h3>
    <b-tabs>
      <b-tab >
        <template v-slot:title>
          <b-icon icon="list-ul"></b-icon>
        </template>
        <AssetList />
      </b-tab>
      <b-tab >
        <template v-slot:title>
          <b-icon icon="graph-up"></b-icon>
        </template>
        <AssetChart />
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import  AssetList  from "@/components/asset/AssetList.vue";
import  AssetChart  from "@/components/asset/AssetChart.vue";
export default {
  name:'Asset',
  components:{
    AssetList,
    AssetChart,
  },
}
</script>
