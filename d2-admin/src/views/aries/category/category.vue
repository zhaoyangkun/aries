<template>
  <d2-container>
    <template slot="header">分类列表</template>
    <d2-crud
      ref="d2Crud"
      :loading="loading"
      :columns="columns"
      :data="data"
      :options="options"
      :pagination="pagination"
      @pagination-current-change="paginationCurrentChange"
      :form-options="formOptions"
      add-title="添加分类"
      :add-template="addTemplate"
      :add-rules="addRules"
      @row-add="handleRowAdd"
      edit-title="修改分类"
      :edit-template="editTemplate"
      :edit-rules="editRules"
      @row-edit="handleRowEdit"
      @dialog-cancel="handleDialogCancel"
      :rowHandle="rowHandle"
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
          <el-input size="small" placeholder="请输入名称" v-model="pagination.key"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="fetchData"><i class="el-icon-search"></i> 搜索</el-button>
        </el-form-item>
      </el-form>
    </d2-crud>
  </d2-container>
</template>

<script>
import {
  getCategoriesByPage, getAllParentCategories, addCategory,
  updateCategory, deleteCategory, multiDelCategory
} from '@/api/aries/category'

export default {
  name: 'category',
  data () {
    return {
      columns: [
        {
          title: '分类名称',
          key: 'name',
          width: '180'
        },
        {
          title: '访问 URL',
          key: 'url',
          width: '180'
        },
        {
          title: '子级分类',
          key: 'childrenStr'
        }
      ],
      data: [],
      pagination: {
        currentPage: 1, // 页码
        pageSize: 10, // 每页条数
        total: 0, // 总条数
        key: '' // 搜索关键词
      },
      parentCategories: [], // 父级分类列表
      selection: [], // 选中条目
      options: {
        border: true
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
          title: '分类名称',
          value: ''
        },
        url: {
          title: 'URL',
          value: ''
        },
        parent_id: {
          title: '父级分类',
          value: 0,
          component: {
            name: 'el-select',
            options: []
          }
        }
      },
      addRules: {
        name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
        url: [{ required: true, message: '请输入 URL', trigger: 'blur' }]
      },
      editTemplate: {
        name: {
          title: '分类名称',
          value: ''
        },
        url: {
          title: 'URL',
          value: ''
        },
        parent_id: {
          title: '父级分类',
          value: 0,
          component: {
            name: 'el-select',
            options: []
          }
        }
      },
      editRules: {
        name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
        url: [{ required: true, message: '请输入 URL', trigger: 'blur' }]
      },
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false
      },
      loading: false
    }
  },
  mounted () {
    this.fetchData()
    this.fetchParentCategories()
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
      getCategoriesByPage({
        page: this.pagination.currentPage,
        size: this.pagination.pageSize,
        category_type: 0,
        key: this.pagination.key
      })
        .then(res => {
          const pageData = res.data.data
          pageData.forEach(function (val) {
            val.childrenStr = ''
            if (val.children !== null) {
              val.children.forEach(function (child) {
                val.childrenStr += child.name + ' '
              })
            }
          })
          this.data = pageData
          this.pagination.total = res.data.total_num
          this.loading = false
        })
        .catch(() => {
          this.loading = false
        })
    },
    // 获取父级分类
    fetchParentCategories () {
      this.parentCategories = [{ value: 0, label: '请选择父级分类' }]
      getAllParentCategories(0)
        .then(res => {
          const _this = this
          res.data.forEach(function (val) {
            _this.parentCategories.push({
              value: val.ID,
              label: val.name
            })
          })
          this.addTemplate.parent_id.component.options = this.parentCategories
          this.editTemplate.parent_id.component.options = this.parentCategories
        })
        .catch(() => {
        })
    },
    // 添加数据弹窗
    addRow () {
      this.$refs.d2Crud.showDialog({
        mode: 'add'
      })
    },
    // 添加数据
    handleRowAdd (row, done) {
      this.formOptions.saveLoading = true
      setTimeout(() => {
        addCategory(row)
          .then(res => {
            this.$message({
              message: res.msg,
              type: 'success'
            })
            done()
            this.fetchData()
            this.fetchParentCategories()
          })
          .catch(() => {
          })
        this.formOptions.saveLoading = false
      }, 300)
    },
    handleRowEdit (row, done) {
      this.formOptions.saveLoading = true
      const data = row.row
      setTimeout(() => {
        updateCategory(data.ID, data)
          .then(res => {
            this.$message({
              message: res.msg,
              type: 'success'
            })
            done()
            this.fetchData()
            this.fetchParentCategories()
          })
          .catch(() => {
          })
        this.formOptions.saveLoading = false
      }, 300)
    },
    // 取消弹窗
    handleDialogCancel (done) {
      this.$message({
        message: '取消保存',
        type: 'warning'
      })
      done()
    },
    // 全选
    handleSelectionChange (selection) {
      this.selection = selection
    },
    // 删除
    handleRowRemove ({ index, row }, done) {
      setTimeout(() => {
        deleteCategory(row.ID)
          .then(res => {
            this.$message({
              message: res.msg,
              type: 'success'
            })
            done()
            this.fetchData()
            this.fetchParentCategories()
          })
          .catch(() => {})
      }, 300)
    },
    // 批量删除
    handleRowListRemove () {
      if (this.selection.length === 0) {
        this.$message({
          message: '请勾选要删除的条目',
          type: 'error'
        })
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
            multiDelCategory(ids)
              .then(res => {
                this.$message({
                  message: res.msg,
                  type: 'success'
                })
                this.fetchData()
                this.fetchParentCategories()
              })
              .catch(() => {
              })
          }, 300)
        }).catch(() => {
        })
      }
    }
  }
}
</script>
