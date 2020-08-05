<template>
  <d2-container>
    <template slot="header">友链列表</template>
    <d2-crud
      ref="d2Crud"
      :options="options"
      :loading="loading"
      :columns="columns"
      :data="data"
      :pagination="pagination"
      @pagination-current-change="paginationCurrentChange"
      :form-options="formOptions"
    >
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item>
          <el-input size="small" clearable placeholder="请输入关键字" v-model="pagination.key"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary"><i class="el-icon-search"></i> 搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" @click="resetTable"><i class="el-icon-refresh"></i> 重置</el-button>
        </el-form-item>
      </el-form>
    </d2-crud>
  </d2-container>
</template>

<script>
export default {
  name: 'link',
  data () {
    return {
      columns: {},
      pagination: {
        currentPage: 1, // 页码
        pageSize: 10, // 每页条数
        total: 0, // 总条数
        key: '' // 搜索关键词
      },
      options: {
        border: true
      },
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false
      },
      loading: false
    }
  },
  methods: {
    paginationCurrentChange () {
    },
    resetTable () {
    }
  }
}
</script>

<style scoped>

</style>
