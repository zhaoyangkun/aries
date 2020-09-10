<template>
  <d2-container>
    <template slot="header">用户 / 页面</template>
    <el-tabs :tab-position="tabPosition" @tab-click="handleTabClick" type="border-card">
      <el-tab-pane label="独立页面">
        <span slot="label"><i class="el-icon-link"></i> 独立页面</span>
        <el-table
          :data="independentPageData"
          style="width: 100%">
          <el-table-column
            prop="title"
            label="页面名称"
            width="150"
          >
          </el-table-column>
          <el-table-column
            label="页面地址"
            width="450"
          >
            <template slot-scope="scope">
              <el-link target="_blank" :href="scope.row.url" :underline="false">{{ scope.row.url }}</el-link>
            </template>
          </el-table-column>
          <el-table-column
            label="操作">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="handleManagePage(scope.row.title)">管理</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <el-tab-pane label="自定义页面">
        <span slot="label"><i class="fa fa-link"></i> 自定义页面</span>
        <div class="operation-box">
          <el-form :inline="true" class="demo-form-inline" slot="header">
            <el-form-item>
              <el-button size="small" type="danger" @click="handleMultiDelPages">
                <i class="el-icon-delete"></i> 批量删除
              </el-button>
            </el-form-item>
            <el-form-item>
              <el-button size="small" type="primary" @click="handleOpenAddDialog">
                <i class="el-icon-plus"></i> 新增
              </el-button>
            </el-form-item>
            <el-form-item>
              <el-input size="small" v-model="pagination.key" placeholder="请输入关键词"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button size="small" type="primary" @click="search"><i class="el-icon-search"></i> 搜索</el-button>
            </el-form-item>
            <el-form-item>
              <el-button size="small" @click="reset"><i class="el-icon-refresh"></i> 重置</el-button>
            </el-form-item>
          </el-form>
        </div>

        <el-table
          v-loading="loading"
          :data="customPageData"
          @selection-change="handleSelectChange"
          style="width: 100%"
        >
          <el-table-column
            type="selection"
            width="55">
          </el-table-column>
          <el-table-column
            prop="title"
            label="页面名称"
            width="150"
          >
          </el-table-column>
          <el-table-column
            label="页面地址"
            width="450"
          >
            <template slot-scope="scope">
              <el-link target="_blank" :href="`${blogVars.ContextPath}/custom/${scope.row.url}`" :underline="false">
                {{ blogVars.ContextPath + "/custom/" + scope.row.url }}
              </el-link>
            </template>
          </el-table-column>
          <el-table-column
            label="操作">
            <template slot-scope="scope">
              <el-button size="small" type="text" @click="handleOpenEditDialog(scope.row)">编辑</el-button>&nbsp;
              <el-popconfirm
                confirmButtonText='确定'
                cancelButtonText='取消'
                icon="el-icon-info"
                iconColor="red"
                title="确定要删除吗？"
                @onConfirm="handleDelPage(scope.row.ID)"
              >
                <el-button slot="reference" size="small" type="text">删除</el-button>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>
        <div class="page-box">
          <el-pagination
            background
            :page-sizes="[10, 30, 50]"
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="pagination.page"
            :page-size="pagination.size"
            :page-count="pagination.total_pages"
            :total="pagination.total_num"
            @size-change="handlePageSizeChange"
            @current-change="handleCurrentPageChange"
          >
          </el-pagination>
        </div>
      </el-tab-pane>
    </el-tabs>

    <el-dialog
      title="添加页面"
      :visible.sync="dialogOptions.addVisible"
      :with-header="false"
      width="60%"
    >
      <el-form ref="addForm" :model="addForm" :rules="addFormRules" label-width="100px">
        <el-form-item label="页面名称" prop="title">
          <el-input size="small" type="text" v-model="addForm.title" placeholder="请输入页面名称"
                    style="width: 50%"></el-input>
        </el-form-item>
        <el-form-item label="页面路径" prop="url">
          <el-input size="small" type="text" v-model="addForm.url" placeholder="页面路径不要以 / 开头"
                    style="width: 50%"></el-input>
          <span class="form-tip"> {{ blogVars.ContextPath + "/custom/" + addForm.url }}</span>
        </el-form-item>
        <el-form-item label="页面内容" prop="html">
          <addEditor ref="addEditor"></addEditor>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.addBtnLoading" @click="handleAddPage">保存</el-button>
          <el-button @click="dialogOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <el-dialog
      title="修改页面"
      :visible.sync="dialogOptions.editVisible"
      :with-header="false"
      width="60%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="100px">
        <el-form-item label="页面名称" prop="title">
          <el-input size="small" type="text" v-model="editForm.title" placeholder="请输入页面名称"
                    style="width: 50%"></el-input>
        </el-form-item>
        <el-form-item label="页面路径" prop="url">
          <el-input size="small" type="text" v-model="editForm.url" placeholder="页面路径不要以 / 开头"
                    style="width: 50%"></el-input>
          <span class="form-tip"> {{ blogVars.ContextPath + "/custom/" + editForm.url }}</span>
        </el-form-item>
        <el-form-item label="页面内容" prop="html">
          <editEditor :content="editForm.html" ref="editEditor"></editEditor>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.editBtnLoading" @click="handleEditPage">保存</el-button>
          <el-button @click="dialogOptions.editVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </d2-container>
</template>

<script>
import addEditor from '@/components/aries/post/addVditor'
import editEditor from '@/components/aries/post/editVditor'
import { getBlogVars } from '@api/aries/sys'
import { createPage, getPagesByPage, multiDelPages, updatePage } from '@api/aries/page'

export default {
  name: 'page',
  components: {
    addEditor,
    editEditor
  },
  data () {
    return {
      tabPosition: 'top',
      pagination: {
        page: 1,
        size: 10,
        key: '',
        total_num: 0,
        total_pages: 0
      },
      blogVars: {},
      selectList: [],
      independentPageData: [],
      customPageData: [],
      dialogOptions: {
        addVisible: false,
        editVisible: false,
        addBtnLoading: false,
        editBtnLoading: false
      },
      addForm: {
        title: '',
        url: '',
        html: '',
        md_html: ''
      },
      editForm: {
        ID: null,
        title: '',
        url: '',
        html: '',
        md_html: ''
      },
      addFormRules: {
        title: [
          { required: true, trigger: 'blur', message: '请输入页面名称' },
          { max: 100, trigger: 'blur', message: '页面名称不能超过 100 个字符' }
        ],
        url: [
          { required: true, trigger: 'blur', message: '请输入页面路径' },
          { max: 100, trigger: 'blur', message: '页面路径不能超过 100 个字符' }
        ],
        html: [
          { required: true, pattern: /^(?!\n$)/, trigger: 'blur', message: '请输入页面内容' },
          { max: 100000, trigger: 'blur', message: '页面内容不能超过 100000 个字符' }
        ],
        md_html: [
          { required: true, trigger: 'blur', message: '请输入 markdown 渲染的页面内容' },
          { max: 1000000, trigger: 'blur', message: 'markdown 渲染的页面内容不能超过 100000 个字符' }
        ]
      },
      editFormRules: {
        title: [
          { required: true, trigger: 'blur', message: '请输入页面名称' },
          { max: 100, trigger: 'blur', message: '页面名称不能超过 100 个字符' }
        ],
        url: [
          { required: true, trigger: 'blur', message: '请输入页面路径' },
          { max: 100, trigger: 'blur', message: '页面路径不能超过 100 个字符' }
        ],
        html: [
          { required: true, pattern: /^(?!\n$)/, trigger: 'blur', message: '请输入页面内容' },
          { max: 100000, trigger: 'blur', message: '页面内容不能超过 100000 个字符' }
        ],
        md_html: [
          { required: true, trigger: 'blur', message: '请输入 markdown 渲染的页面内容' },
          { max: 1000000, trigger: 'blur', message: 'markdown 渲染的页面内容不能超过 100000 个字符' }
        ]
      },
      loading: false
    }
  },
  created () {
    this.initIndependentPageData()
  },
  methods: {
    // 重置表单
    resetForm (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].resetFields()
      }
    },
    // 清空表单验证
    clearValidate (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].clearValidate()
      }
    },
    // 获取博客全局变量
    async fetchBlogVars () {
      await getBlogVars()
        .then(res => {
          this.blogVars = res.data
        })
        .catch(() => {
        })
    },
    // 获取独立页面数据
    async fetchIndependentPageData () {
      this.independentPageData = [
        {
          title: '日志页面',
          url: `${this.blogVars.ContextPath}/journals`
        },
        {
          title: '图库页面',
          url: `${this.blogVars.ContextPath}/galleries`
        }
      ]
    },
    // 初始化独立页面数据
    async initIndependentPageData () {
      await this.fetchBlogVars()
      await this.fetchIndependentPageData()
    },
    // 获取自定义页面数据
    fetchCustomPageData () {
      this.loading = true
      setTimeout(_ => {
        getPagesByPage({
          page: this.pagination.page,
          size: this.pagination.size,
          key: this.pagination.key
        })
          .then(res => {
            this.customPageData = res.data.data
            this.pagination.total_num = res.data.total_num
            this.pagination.total_pages = res.data.total_pages
          })
          .catch(_ => {
          })
        this.loading = false
      }, 300)
    },
    // 标签页切换
    handleTabClick (tab) {
      switch (tab.label) {
        case '独立页面':
          this.initIndependentPageData()
          break
        case '自定义页面':
          this.fetchCustomPageData()
          break
      }
    },
    // 搜索
    search () {
      this.pagination.page = 1
      this.fetchCustomPageData()
    },
    // 重置
    reset () {
      this.pagination.page = 1
      this.pagination.key = ''
      this.fetchCustomPageData()
    },
    // 页面管理
    handleManagePage (title) {
      if (title === '日志页面') {
        this.$router.replace('/page/journal')
      } else if (title === '图库页面') {
        this.$router.replace('/page/gallery')
      }
    },
    // 改变每页条数
    handlePageSizeChange (size) {
      this.pagination.size = size
      this.fetchCustomPageData()
    },
    // 改变当前页
    handleCurrentPageChange (page) {
      this.pagination.page = page
      this.fetchCustomPageData()
    },
    // 打开添加页面弹窗
    handleOpenAddDialog () {
      this.clearValidate('addForm')
      this.dialogOptions.addVisible = true
    },
    // 打开修改页面弹窗
    handleOpenEditDialog (row) {
      this.editForm = { ...row }
      this.dialogOptions.editVisible = true
    },
    // 添加页面
    handleAddPage () {
      this.dialogOptions.addBtnLoading = true
      setTimeout(_ => {
        this.addForm.html = this.$refs.addEditor.getContent()
        this.addForm.md_html = this.$refs.addEditor.getHTML()
        this.$refs.addForm.validate(valid => {
          if (valid) {
            createPage(this.addForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.addVisible = false
                this.resetForm('addForm')
                this.fetchCustomPageData()
              })
              .catch(_ => {
              })
          }
        }, 300)
        this.dialogOptions.addBtnLoading = false
      })
    },
    // 修改页面
    handleEditPage () {
      this.dialogOptions.editBtnLoading = true
      setTimeout(_ => {
        this.editForm.html = this.$refs.editEditor.getContent()
        this.editForm.md_html = this.$refs.editEditor.getHTML()
        this.$refs.editForm.validate(valid => {
          if (valid) {
            updatePage(this.editForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.editVisible = false
                this.fetchCustomPageData()
              })
              .catch(_ => {
              })
          }
        }, 300)
        this.dialogOptions.editBtnLoading = false
      })
    },
    // 删除页面
    handleDelPage (id) {
      const ids = [id].join(',')
      multiDelPages(ids)
        .then(res => {
          this.$message.success(res.msg)
          this.fetchCustomPageData()
        })
        .catch(_ => {
        })
    },
    // 选项发生变化
    handleSelectChange (selection) {
      selection.forEach(val => {
        this.selectList.push(val.ID)
      })
    },
    // 批量删除
    handleMultiDelPages () {
      if (this.selectList.length === 0) {
        this.$message.error('请选择要删除的页面')
        return
      }
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(_ => {
          const ids = this.selectList.join(',')
          multiDelPages(ids)
            .then(res => {
              this.$message.success(res.msg)
              this.fetchCustomPageData()
            })
            .catch(_ => {
            })
        })
        .catch(_ => {
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.page-box {
  margin-top: 10px;
}

.form-tip {
  color: darkgrey;
}
</style>
