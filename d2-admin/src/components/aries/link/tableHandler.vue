<template>
  <div>
    <el-button size="small" @click="openEditDialog(scope.row)">编辑</el-button>
    <el-button size="small" type="danger" @click="handleRowRemove(scope.row.ID)">删除</el-button>
  </div>
</template>

<script>
export default {
  name: 'tableHandler',
  props: {
    scope: {
      default: null
    },
    openEditDialog: {
      type: Function
    },
    handleRowRemove: {
      type: Function
    }
  }
}
</script>

<style scoped>

</style>
