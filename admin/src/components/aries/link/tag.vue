<template>
  <el-tag v-if="scope.row.category.name">{{ scope.row.category.name }}</el-tag>
</template>

<script>
export default {
  name: 'tag',
  props: {
    scope: {
      default: null
    }
  }
}
</script>

<style scoped>
</style>
