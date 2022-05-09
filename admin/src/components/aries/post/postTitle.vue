<template>
  <div>
    <i v-if="scope.row.is_top" class="el-icon-magic-stick" style="color:red;"></i>
    <span v-if="scope.row.is_recycled" style="color: darkgrey">
      {{ scope.row.title }}
    </span>
    <span v-else>
      {{ scope.row.title }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'postTitle',
  props: {
    // 本行的所有数据，此字段不需要额外配置
    scope: {
      default: null
    }
  }
}
</script>

<style scoped>

</style>
