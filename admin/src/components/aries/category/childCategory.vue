<template>
  <div>
    <el-tag style="margin-right: 5px;margin-top: 3px" :key="text" :type="type" v-for="text in textList">
      {{ text }}
    </el-tag>
  </div>
</template>

<script>
export default {
  name: 'childCategory',
  props: {
    // 本行的所有数据，此字段不需要额外配置
    scope: {
      default: null
    },
    // 由父组件传入的值
    type: {
      type: String,
      default: 'default'
    }
  },
  mounted () {
    // console.log(this.scope.row)
  },
  computed: {
    // 获取信息列表
    textList () {
      const textList = []
      if (this.scope.row.children != null) {
        this.scope.row.children.forEach((val) => {
          textList.push(val.name)
        })
      }
      return textList
    }
  }
}
</script>

<style scoped>

</style>
