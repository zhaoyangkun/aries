<template>
  <div>
    <el-tag :type="type" v-if="text !== ''">{{ text }}</el-tag>
  </div>
</template>

<script>
export default {
  name: 'singleTag',
  props: {
    // 本行的所有数据，此字段不需要额外配置
    scope: {
      default: null
    },
    // 由父组件传入的值
    type: {
      type: String,
      default: 'success'
    }
  },
  mounted () {
  },
  computed: {
    text () {
      return this.scope.row.category.name
    }
  }
}
</script>

<style scoped>

</style>
