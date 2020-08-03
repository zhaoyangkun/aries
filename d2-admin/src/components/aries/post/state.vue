<template>
  <div>
    <el-tag :type="type" effect="plain">{{ text }}</el-tag>
  </div>
</template>

<script>
export default {
  name: 'state',
  props: {
    scope: {
      default: null
    }
  },
  computed: {
    type () {
      if (this.scope.row.is_recycled) {
        return 'danger'
      } else if (!this.scope.row.is_published) {
        return 'warning'
      } else if (this.scope.row.pwd !== '') {
        return 'info'
      } else {
        return 'success'
      }
    },
    text () {
      if (this.scope.row.is_recycled) {
        return '回收站'
      } else if (!this.scope.row.is_published) {
        return '草稿'
      } else if (this.scope.row.pwd !== '') {
        return '加密'
      } else {
        return '已发布'
      }
    }
  }
}
</script>

<style scoped>

</style>
