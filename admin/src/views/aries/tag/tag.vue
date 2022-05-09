<template>
  <d2-container>
    <template slot="header">文章 / 标签</template>
    <d2-crud
      ref="d2Crud"
      :loading="loading"
      :columns="columns"
      :data="data"
      :options="options"
      :pagination="pagination"
      @pagination-current-change="paginationCurrentChange"
      :form-options="formOptions"
      :row-handle="rowHandle"
      add-title="添加标签"
      :add-template="addTemplate"
      :add-rules="addRules"
      @row-add="handleRowAdd"
      @dialog-cancel="handleDialogCancel"
      edit-title="修改标签"
      :edit-template="editTemplate"
      :edit-rules="editRules"
      @row-edit="handleRowEdit"
      @row-remove="handleRowRemove"
      selection-row
      @selection-change="handleSelectionChange"
    >
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item>
          <el-button size="small" type="danger" style="margin-bottom: 5px" @click="handleRowListRemove">
            <i class="el-icon-delete"></i> 批量删除
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" style="margin-bottom: 5px" @click="addRow">
            <i class="el-icon-plus"></i> 新增
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-input size="small" placeholder="请输入标签名称" v-model="pagination.key"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="search"><i class="el-icon-search"></i> 搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" @click="reset"><i class="el-icon-refresh"></i> 重置</el-button>
        </el-form-item>
      </el-form>
    </d2-crud>
  </d2-container>
</template>

<script>
import { addTag, deleteTag, getTagsByPage, multiDelTags, updateTag } from '@/api/aries/tag'

export default {
  name: 'tag',
  data () {
    return {
      columns: [
        {
          title: '标签名称',
          key: 'name',
          width: '180'
        }
      ],
      data: [],
      selection: [],
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
      rowHandle: {
        remove: {
          icon: 'el-icon-delete',
          size: 'small',
          fixed: 'right',
          confirm: true
        },
        edit: {
          icon: 'el-icon-edit',
          text: '编辑',
          size: 'small'
        }
      },
      addTemplate: {
        name: {
          title: '标签名称',
          value: ''
        }
      },
      editTemplate: {
        name: {
          title: '标签名称',
          value: ''
        }
      },
      addRules: {
        name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }]
      },
      editRules: {
        name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }]
      },
      loading: false
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    // 分页
    paginationCurrentChange (currentPage) {
      this.pagination.currentPage = currentPage
      this.fetchData()
    },
    // 获取分页数据
    fetchData () {
      this.loading = true
      setTimeout(() => {
        getTagsByPage({
          page: this.pagination.currentPage,
          size: this.pagination.pageSize,
          key: this.pagination.key
        })
          .then(res => {
            const data = res.data
            this.data = data.data
            this.pagination.total = data.total_num
          })
          .catch(() => {
          })
        this.loading = false
      }, 300)
    },
    // 搜索
    search () {
      this.pagination.currentPage = 1
      this.fetchData()
    },
    // 重置
    reset () {
      this.pagination = {
        currentPage: 1, // 页码
        pageSize: 10, // 每页条数
        total: 0, // 总条数
        key: '' // 搜索关键词
      }
      this.fetchData()
    },
    // 添加数据弹窗
    addRow () {
      this.$refs.d2Crud.showDialog({
        mode: 'add'
      })
    },
    // 添加事件
    handleRowAdd (row, done) {
      this.formOptions.saveLoading = true
      setTimeout(() => {
        addTag(row)
          .then(res => {
            this.$message.success(res.msg)
            done()
            this.fetchData()
            this.fetchParentCategories()
          })
          .catch(() => {
          })
        this.formOptions.saveLoading = false
      }, 300)
    },
    // 修改事件
    handleRowEdit (row, done) {
      this.formOptions.saveLoading = true
      const data = { ...row.row }
      setTimeout(() => {
        updateTag(data)
          .then(res => {
            this.$message.success(res.msg)
            done()
            this.fetchData()
            this.fetchParentCategories()
          })
          .catch(() => {
          })
        this.formOptions.saveLoading = false
      }, 300)
    },
    // 全选
    handleSelectionChange (selection) {
      this.selection = selection
    },
    // 删除
    handleRowRemove (row) {
      setTimeout(() => {
        deleteTag(row.row.ID)
          .then(res => {
            this.$message.success(res.msg)
            this.fetchData()
            this.fetchParentCategories()
          })
          .catch(() => {
          })
      }, 300)
    },
    // 批量删除
    handleRowListRemove () {
      if (this.selection.length === 0) {
        this.$message.error('请勾选要删除的条目')
      } else {
        this.$confirm('确定要删除吗?', '删除', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          let ids = []
          this.selection.forEach(function (val) {
            ids.push(val.ID)
          })
          ids = ids.join() // 将数组转换成字符串：[1, 2, 3] --> '1, 2, 3'
          setTimeout(() => {
            multiDelTags(ids)
              .then(res => {
                this.$message.success(res.msg)
                this.fetchData()
                this.fetchParentCategories()
              })
              .catch(() => {
              })
          }, 300)
        }).catch(() => {
        })
      }
    },
    // 取消弹窗
    handleDialogCancel (done) {
      this.$message.warning('取消保存')
      done()
    }
  }
}
</script>
