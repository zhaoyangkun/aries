<template>
  <div>{{ formatDate }}</div>
</template>

<script>
import { dateFormat } from '@/plugin/time/date'

export default {
  name: 'date',
  props: {
    // 本行数据
    scope: {
      default: null
    },
    fmt: {
      type: String,
      default: 'yyyy-MM-dd hh:mm:ss'
    }
  },
  computed: {
    formatDate () {
      return dateFormat(this.fmt, new Date(this.scope.row.CreatedAt))
    }
  }
}
</script>

<style scoped>
</style>
