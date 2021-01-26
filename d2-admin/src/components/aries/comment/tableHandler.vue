<template>
  <div>
    <el-tooltip content="回复" placement="top-start" v-if="scope.row.is_checked">
      <el-button size="small" class="el-icon-chat-line-round"
                 @click="openReplyDialog(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip content="通过审核" placement="top-start" v-if="!scope.row.is_checked && !scope.row.is_recycled">
      <el-button size="small" class="el-icon-success" @click="handleChecked(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip content="加入回收站" placement="top-start" v-if="!scope.row.is_recycled">
      <el-button size="small" class="el-icon-close" type="danger"
                 @click="handleRowRecycleOrRecover(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip content="恢复" placement="top-start" v-if="scope.row.is_recycled">
      <el-button size="small" class="el-icon-check" type="success"
                 @click="handleRowRecycleOrRecover(scope.row)"></el-button>
    </el-tooltip>
    <el-tooltip content="彻底删除" placement="top-start" v-if="scope.row.is_recycled">
      <el-button size="small" class="el-icon-delete" type="danger"
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
    openReplyDialog: {
      type: Function
    },
    handleRowRemove: {
      type: Function
    },
    handleRowRecycleOrRecover: {
      type: Function
    },
    handleChecked: {
      type: Function
    }
  }
}
</script>

<style scoped>
</style>
