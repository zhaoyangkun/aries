<template>
  <div>
    <el-tooltip effect="dark" content="编辑" placement="top-start">
      <el-button size="small" type="primary" icon="el-icon-edit"
                 @click="openEditDialog(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip effect="dark" content="加入回收站" placement="top-start" v-if="!scope.row.is_recycled">
      <el-button size="small" type="danger" icon="el-icon-close"
                 @click="handleRowRecycleOrRecover(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip effect="dark" content="恢复" placement="top-start" v-if="scope.row.is_recycled">
      <el-button size="small" type="success" icon="el-icon-check"
                 @click="handleRowRecycleOrRecover(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip effect="dark" content="彻底删除" placement="top-start" v-if="scope.row.is_recycled">
      <el-button size="small" type="danger" icon="el-icon-delete"
                 @click="handleRowRemove(scope.row.ID)"></el-button>
    </el-tooltip>
  </div>
</template>

<script>
export default {
  name: 'tableHandle',
  props: {
    scope: {
      default: null
    },
    openEditDialog: {
      type: Function
    },
    handleRowRemove: {
      type: Function
    },
    handleRowRecycleOrRecover: {
      type: Function
    }
  }
}
</script>

<style scoped>

</style>
