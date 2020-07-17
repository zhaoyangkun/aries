<template>
  <div>
    <i class="el-icon-user" style="color: #2f74ff"
       v-if="scope.row.admin_user_id > 0"></i> {{ scope.row.nick_name }}
  </div>
</template>

<script>
export default {
  name: 'nickName',
  props: {
    scope: {
      default: null
    }
  }
}
</script>

<style scoped>

</style>
