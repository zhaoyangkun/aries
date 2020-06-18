<template>
  <d2-container>
    <template slot="header">基础表格</template>
    <d2-crud
      ref="d2Crud"
      :columns="columns"
      :data="data"/>
  </d2-container>
</template>

<script>
export default {
  name: 'article',
  data () {
    return {
      columns: [
        {
          title: '分类名称',
          key: 'name',
          width: '180'
        },
        {
          title: '所属类型',
          key: 'type',
          width: '180'
        },
        {
          title: '访问 URL',
          key: 'url',
          width: '180'
        },
        {
          title: '子级分类',
          key: 'children'
        }
      ],
      data: [
        {
          name: '后端',
          type: '文章',
          url: '/backend',
          children: 'Java,Go,Python'
        },
        {
          name: '后端',
          type: '文章',
          url: '/backend',
          children: 'Java,Go,Python'
        },
        {
          name: '后端',
          type: '文章',
          url: '/backend',
          children: 'Java,Go,Python'
        },
        {
          name: '后端',
          type: '文章',
          url: '/backend',
          children: 'Java,Go,Python'
        },
        {
          name: '后端',
          type: '文章',
          url: '/backend',
          children: 'Java,Go,Python'
        }
      ]
    }
  },
  mounted () {
    console.log(this.$refs.d2Crud.d2CrudData)
  }
}
</script>
