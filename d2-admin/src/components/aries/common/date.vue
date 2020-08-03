<template>
  <div>{{ formatDate }}</div>
</template>

<script>
import { dateFormat } from '@/plugin/time/date'

export default {
  name: 'date',
  props: {
    // 本行的所有数据，此字段不需要额外配置
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
  },
  methods: {
  }
}
</script>

<style scoped>

</style>
