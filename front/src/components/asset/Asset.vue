<template>
  <div>
    <h3>資産</h3>
    <b-tabs>
      <b-tab title="一覧">
        <ShareList />
      </b-tab>
      <b-tab >
        <template v-slot:title>
          <b-icon icon="pie-chart"></b-icon>
        </template>
        <SharePieChart />
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import  ShareList  from "@/components/asset/ShareList.vue";
import  SharePieChart  from "@/components/asset/SharePieChart.vue";
export default {
  name:'Asset',
  components:{
    ShareList,
    SharePieChart,
  },
}
</script>
