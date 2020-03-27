<template>
  <div>
    <b-form-group :state="state" :label="label" label-for="months">
      <b-form-tags
        input-id="months"
        :input-attrs="{ 'aria-describedby': 'tags-validation-help' }"

        :value="months"
        @input="$emit('input',$event)"

        :state="state"
        :tag-validator="tagValidator"
        separator=" "
      ></b-form-tags>
    </b-form-group>
  </div>
</template>

<script>
  export default {
    name:"MonthSelector",
    props: {
      label:String,
      months:Array,
    },
    data() {
      return {
        dirty: false
      }
    },
    computed: {
      state() {
        // Overall component validation state
        return this.dirty ? this.months.length > 0  : null
      }
    },
    watch: {
      months() {
        this.dirty = true
      }
    },
    methods: {
      tagValidator(month) {
        return Number.isInteger(Number(month)) && 1 <= Number(month) && Number(month) <= 12
      }
    }
  }
</script>
