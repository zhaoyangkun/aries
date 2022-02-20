<template>
  <div>
    <el-tooltip content="编辑" placement="top-start" v-if="!scope.row.is_recycled">
      <el-button size="small" icon="el-icon-edit"
                 @click="openEditDialog(scope.row.ID)"></el-button>
    </el-tooltip>
    <el-tooltip content="向上" placement="top-start" v-if="!scope.row.is_recycled">
      <el-button size="small" icon="el-icon-top" @click="handleMoveUp(scope.row,scope.$index)"></el-button>
    </el-tooltip>
    <el-tooltip content="向下" placement="top-start" v-if="!scope.row.is_recycled">
      <el-button size="small" icon="el-icon-bottom" @click="handleMoveDown(scope.row,scope.$index)"></el-button>
    </el-tooltip>
    <el-tooltip content="加入回收站" placement="top-start" v-if="!scope.row.is_recycled">
      <el-button size="small" type="danger" icon="el-icon-close"
                 @click="handleRowRecycleOrRecover(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip content="恢复" placement="top-start" v-if="scope.row.is_recycled">
      <el-button size="small" type="success" icon="el-icon-check"
                 @click="handleRowRecycleOrRecover(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip content="彻底删除" placement="top-start" v-if="scope.row.is_recycled">
      <el-button size="small" type="danger" icon="el-icon-delete"
                 @click="handleRowRemove(scope.row.ID)"></el-button>
    </el-tooltip>
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
    },
    handleRowRecycleOrRecover: {
      type: Function
    },
    handleMoveUp: {
      type: Function
    },
    handleMoveDown: {
      type: Function
    }
  }
}
</script>

<style scoped>
</style>
