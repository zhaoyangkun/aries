<template>
  <div>
    <div v-if="scope.row.is_top">
      <i class="el-icon-magic-stick" style="color:red;"></i> {{ scope.row.title }}
    </div>
    <div v-else>
      {{ scope.row.title }}
    </div>
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
