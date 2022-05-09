<template>
  <d2-container>
    <template slot="header">用户 / 友链 / 友链</template>
    <d2-crud
      ref="d2Crud"
      :options="options"
      :loading="loading"
      :columns="columns"
      :data="data"
      :pagination="pagination"
      @pagination-current-change="paginationCurrentChange"
      :form-options="formOptions"
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
          <el-button size="small" type="primary" style="margin-bottom: 5px" @click="dialogOptions.addVisible = true">
            <i class="el-icon-plus"></i> 新增
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-select style="width: 150px" size="small" v-model="pagination.category_id" clearable placeholder="请选择分类">
            <el-option
              v-for="item in categories"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input size="small" clearable placeholder="请输入关键字" v-model="pagination.key"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="paginationCurrentChange(1)">
            <i class="el-icon-search"></i> 搜索
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" @click="resetTable">
            <i class="el-icon-refresh"></i> 重置
          </el-button>
        </el-form-item>
      </el-form>
    </d2-crud>

    <!-- 添加友链弹窗 -->
    <el-dialog
      title="添加友链"
      :visible.sync="dialogOptions.addVisible"
      width="60%"
    >
      <el-form ref="addForm" :model="addForm" :rules="addRules" label-width="80px">
        <el-form-item label="网站名称" prop="name">
          <el-input size="small" v-model="addForm.name" type="text" autocomplete="off"
                    placeholder="请输入网站名称"></el-input>
        </el-form-item>
        <el-form-item label="网站地址" prop="url">
          <el-input size="small" v-model="addForm.url" type="text" autocomplete="off"
                    placeholder="请输入网站地址"></el-input>
        </el-form-item>
        <el-form-item label="分类" prop="category_id">
          <el-select size="small" v-model="addForm.category_id" clearable
                     @clear="addForm.category_id=null" placeholder="请选择分类">
            <el-option
              v-for="item in categories"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>&nbsp;
          <el-button size="small" type="primary" style="margin-bottom: 5px" @click="drawOptions.addVisible = true">
            <i class="el-icon-plus"></i> 新增
          </el-button>
        </el-form-item>
        <el-form-item label="网站描述" prop="desc">
          <el-input size="small" v-model="addForm.desc" type="text" autocomplete="off"
                    placeholder="请输入网站描述"></el-input>
        </el-form-item>
        <el-form-item label="网站图标" prop="icon">
          <el-input size="small" v-model="addForm.icon" type="text" autocomplete="off"
                    placeholder="请输入网站图标"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="btnOptions.addBtnLoading" @click="handleRowAdd">保存</el-button>
          <el-button @click="dialogOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 修改友链弹窗 -->
    <el-dialog
      title="修改友链"
      :visible.sync="dialogOptions.editVisible"
      width="60%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editRules" label-width="80px">
        <el-form-item label="网站名称" prop="name">
          <el-input size="small" v-model="editForm.name" type="text" autocomplete="off"
                    placeholder="请输入网站名称"></el-input>
        </el-form-item>
        <el-form-item label="网站地址" prop="url">
          <el-input size="small" v-model="editForm.url" type="text" autocomplete="off"
                    placeholder="请输入网站地址"></el-input>
        </el-form-item>
        <el-form-item label="分类" prop="category_id">
          <el-select size="small" v-model="editForm.category_id" clearable placeholder="请选择分类"
                     @clear="editForm.category_id=null">
            <el-option
              v-for="item in categories"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="网站描述" prop="desc">
          <el-input size="small" v-model="editForm.desc" type="text" autocomplete="off"
                    placeholder="请输入网站描述"></el-input>
        </el-form-item>
        <el-form-item label="网站图标" prop="icon">
          <el-input size="small" v-model="editForm.icon" type="text" autocomplete="off"
                    placeholder="请输入网站图标"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="btnOptions.editBtnLoading" @click="handleRowEdit">保存</el-button>
          <el-button @click="dialogOptions.editVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 添加分类抽屉 -->
    <el-drawer
      title="添加友链分类"
      :visible.sync="drawOptions.addVisible"
      direction="rtl"
    >
      <el-form ref="addCategoryForm" :model="addCategoryForm" :rules="addCategoryRules" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input style="width: 280px" size="small" type="text" v-model="addCategoryForm.name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="drawOptions.addBtnLoading"
                     @click="handleCategoryAdd">保存
          </el-button>
          <el-button @click="drawOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </d2-container>
</template>

<script>
import { createLink, deleteLink, getLinksByPage, multiDelLinks, updateLink } from '@api/aries/link'
import { addLinkCategory, getAllCategories } from '@api/aries/category'
import tableHandler from '@/components/aries/link/tableHandler'
import tag from '@/components/aries/link/tag'

export default {
  name: 'links',
  comments: {
    tableHandler,
    tag
  },
  data () {
    return {
      columns: [
        {
          title: '友链名称',
          key: 'name',
          width: '200'
        },
        {
          title: '友链地址',
          key: 'url',
          width: '250'
        },
        {
          title: '分类',
          component: {
            name: tag
          }
        },
        {
          title: '操作',
          component: {
            name: tableHandler,
            props: {
              openEditDialog: this.openEditDialog,
              handleRowRemove: this.handleRowRemove
            }
          }
        }
      ],
      data: [],
      categories: [],
      selection: [],
      addForm: {
        name: '',
        url: '',
        category_id: null,
        desc: '',
        icon: ''
      },
      editForm: {
        id: null,
        name: '',
        url: '',
        category_id: null,
        desc: '',
        icon: ''
      },
      addCategoryForm: {
        type: 1,
        name: ''
      },
      addRules: {
        name: [
          { required: true, trigger: 'blur', message: '请输入网站名称' },
          { max: 100, trigger: 'blur', message: '网站名称不能超过 100 个字符' }
        ],
        url: [
          { required: true, trigger: 'blur', message: '请输入网站地址' },
          { max: 255, trigger: 'blur', message: '网站地址不能超过 255 个字符' },
          { type: 'url', trigger: 'blur', message: '请输入正确的 URL' }
        ],
        desc: [
          { max: 255, trigger: 'blur', message: '网站描述不能超过 255 个字符' }
        ],
        icon: [
          { max: 255, trigger: 'blur', message: '网站图标不能超过 255 个字符' }
        ]
      },
      editRules: {
        name: [
          { required: true, trigger: 'blur', message: '请输入网站名称' },
          { max: 100, trigger: 'blur', message: '网站名称不能超过 100 个字符' }
        ],
        url: [
          { required: true, trigger: 'blur', message: '请输入网站地址' },
          { max: 255, trigger: 'blur', message: '网站地址不能超过 255 个字符' },
          { type: 'url', trigger: 'blur', message: '请输入正确的 URL' }
        ],
        desc: [
          { max: 255, trigger: 'blur', message: '网站描述不能超过 255 个字符' }
        ],
        icon: [
          { max: 255, trigger: 'blur', message: '网站图标不能超过 255 个字符' }
        ]
      },
      addCategoryRules: {
        name: [
          { required: true, trigger: 'blur', message: '请输入分类名称' }
        ]
      },
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
      dialogOptions: {
        addVisible: false,
        editVisible: false
      },
      btnOptions: {
        addBtnLoading: false,
        editBtnLoading: false
      },
      drawOptions: {
        addBtnLoading: false,
        addVisible: false
      },
      loading: false
    }
  },
  created () {
    this.fetchPageData()
    this.fetchLinkCategories()
  },
  methods: {
    // 获取分页数据
    fetchPageData () {
      this.loading = true
      setTimeout(() => {
        getLinksByPage({
          page: this.pagination.currentPage,
          size: this.pagination.pageSize,
          category_id: this.pagination.category_id,
          key: this.pagination.key
        })
          .then(res => {
            this.data = res.data.data
            this.pagination.total = res.data.total_num
          })
          .catch(() => {
          })
        this.loading = false
      }, 300)
    },
    // 获取友链分类
    fetchLinkCategories () {
      getAllCategories(1)
        .then(res => {
          this.categories = res.data
        })
        .catch(() => {
        })
    },
    // 跳页事件
    paginationCurrentChange (currentPage) {
      this.pagination.currentPage = currentPage
      this.fetchPageData()
    },
    // 重置表格数据
    resetTable () {
      this.pagination = {
        currentPage: 1, // 页码
        pageSize: 10, // 每页条数
        total: 0, // 总条数
        key: '' // 搜索关键词
      }
      this.fetchPageData()
    },
    // 重置表单信息
    resetForm (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].resetFields()
      }
    },
    // 打开友链编辑框
    openEditDialog (row) {
      this.dialogOptions.editVisible = true
      this.editForm = { ...row }
      this.editForm.category_id = this.editForm.category_id === 0 ? null : this.editForm.category_id
    },
    // 添加友链
    handleRowAdd () {
      this.$refs.addForm.validate(valid => {
        if (valid) {
          this.btnOptions.addBtnLoading = true
          setTimeout(() => {
            createLink(this.addForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.addVisible = false
                this.resetForm('addForm')
                this.fetchPageData()
              })
              .catch(() => {
              })
            this.btnOptions.addBtnLoading = false
          }, 300)
        }
      })
    },
    // 修改友链
    handleRowEdit () {
      // console.log(this.editForm)
      this.$refs.editForm.validate(valid => {
        if (valid) {
          this.btnOptions.editBtnLoading = true
          setTimeout(() => {
            updateLink(this.editForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.editVisible = false
                this.resetForm('editForm')
                this.fetchPageData()
              })
              .catch(() => {
              })
            this.btnOptions.editBtnLoading = false
          }, 300)
        }
      })
    },
    // 删除友链
    handleRowRemove (id) {
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(() => {
          setTimeout(() => {
            deleteLink(id)
              .then(res => {
                this.$message.success(res.msg)
                this.fetchPageData()
              })
              .catch(() => {
              })
          }, 300)
        })
        .catch(() => {
        })
    },
    // 全选
    handleSelectionChange (selection) {
      this.selection = selection
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
            multiDelLinks(ids)
              .then(res => {
                this.$message.success(res.msg)
                this.fetchPageData()
              })
              .catch(() => {
              })
          }, 300)
        }).catch(() => {
        })
      }
    },
    // 添加友链分类
    handleCategoryAdd () {
      this.drawOptions.addBtnLoading = true
      setTimeout(() => {
        addLinkCategory(this.addCategoryForm)
          .then(res => {
            this.$message.success(res.msg)
            this.drawOptions.addVisible = false
            this.resetForm('addCategoryForm')
            this.fetchLinkCategories()
          })
          .catch(() => {
          })
        this.drawOptions.addBtnLoading = false
      }, 300)
    }
  }
}
</script>

<style scoped>
</style>
