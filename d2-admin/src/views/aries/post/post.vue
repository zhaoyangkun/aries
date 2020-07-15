<template>
  <d2-container>
    <template slot="header">文章列表</template>
    <d2-crud
      ref="d2Crud"
      :loading="loading"
      :columns="columns"
      :data="data"
      :options="options"
      :pagination="pagination"
      @pagination-current-change="paginationCurrentChange"
      :form-options="formOptions"
    >
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item>
          <el-button size="small" type="primary" @click="openAddDialog">
            <i class="el-icon-plus"></i> 写文章
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="success" @click="openUploadDialog">
            <i class="el-icon-bottom"></i> 文件导入文章
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
          <el-select style="width: 150px" size="small" v-model="pagination.state" clearable placeholder="请选择文章状态">
            <el-option
              v-for="item in states"
              :key="item.value"
              :label="item.name"
              :value="item.value">
            </el-option>
          </el-select>
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
    </d2-crud>

    <!-- 添加文章弹窗 -->
    <el-dialog
      title="写文章"
      :visible.sync="dialogOptions.addVisible"
      :fullscreen="dialogOptions.addFullScreen"
      width="80%"
    >
      <el-form ref="addForm" :model="addForm" :rules="addRules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input size="small" type="text" v-model="addForm.title"></el-input>
        </el-form-item>
        <el-row>
          <el-col :span="8">
            <el-form-item label="分类" prop="category_id">
              <el-select size="small" v-model="addForm.category_id" clearable placeholder="请选择分类">
                <el-option
                  v-for="item in categories"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="标签" prop="selectTagIds">
              <el-select size="small" multiple :multiple-limit=3
                         v-model="addForm.selectTagIds" clearable
                         placeholder="请选择标签">
                <el-option
                  v-for="item in tags"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="URL" prop="url">
              <el-input size="small" type="text" v-model="addForm.url"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="密码" prop="pwd">
              <el-input size="small" type="password" v-model="addForm.pwd"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="关键词" prop="keywords">
              <el-input size="small" type="text" v-model="addForm.keywords"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="3">
            <el-form-item label="置顶" prop="is_top">
              <el-switch v-model="addForm.is_top"></el-switch>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item label="允许评论" prop="is_allow_commented">
              <el-switch v-model="addForm.is_allow_commented"></el-switch>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="内容" prop="content">
          <addEditor ref="addEditor"></addEditor>
        </el-form-item>
        <el-form-item label="摘要" prop="summary">
          <el-input :rows="5" type="textarea" v-model="addForm.summary"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="success" :loading="dialogOptions.addDraftBtnLoading"
                     @click="handleRowAdd(false)">保存为草稿
          </el-button>
          <el-button type="primary" :loading="dialogOptions.addBtnLoading"
                     @click="handleRowAdd(true)">发布
          </el-button>
          <el-button @click="dialogOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 修改文章弹窗 -->
    <el-dialog
      title="修改文章"
      :visible.sync="dialogOptions.editVisible"
      :fullscreen="dialogOptions.editFullScreen"
      width="80%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editRules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input size="small" type="text" v-model="editForm.title"></el-input>
        </el-form-item>
        <el-row>
          <el-col :span="8">
            <el-form-item label="分类" prop="category_id">
              <el-select size="small" v-model="editForm.category_id" clearable placeholder="请选择分类">
                <el-option
                  v-for="item in categories"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="标签" prop="tag_ids">
              <el-select size="small" multiple :multiple-limit=3
                         v-model="editForm.selectTagIds" clearable
                         placeholder="请选择标签">
                <el-option
                  v-for="item in tags"
                  :key="item.name"
                  :label="item.name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="URL" prop="url">
              <el-input size="small" type="text" v-model="editForm.url"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="密码" prop="pwd">
              <el-input size="small" type="password" v-model="editForm.pwd"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="关键词" prop="keywords">
              <el-input size="small" type="text" v-model="editForm.keywords"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="3">
            <el-form-item label="置顶" prop="is_top">
              <el-switch v-model="editForm.is_top"></el-switch>
            </el-form-item>
          </el-col>
          <el-col :span="3">
            <el-form-item label="允许评论" prop="is_allow_commented">
              <el-switch v-model="editForm.is_allow_commented"></el-switch>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="内容" prop="content">
          <editEditor :content="editForm.content" ref="editEditor"></editEditor>
        </el-form-item>
        <el-form-item label="摘要" prop="summary">
          <el-input :rows="5" type="textarea" v-model="editForm.summary"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="success" :loading="dialogOptions.editDraftBtnLoading"
                     @click="handleRowEdit(false)">保存为草稿
          </el-button>
          <el-button type="primary" :loading="dialogOptions.editBtnLoading"
                     @click="handleRowEdit(true)">发布
          </el-button>
          <el-button @click="dialogOptions.editVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 导入文章弹窗 -->
    <el-dialog
      title="导入文章"
      :visible.sync="dialogOptions.uploadVisible"
      width="50%">
      <el-upload
        ref="upload"
        class="upload-demo"
        :multiple="true"
        :on-change="handleFileChange"
        :on-exceed="handleExceed"
        :on-remove="handleFileRemove"
        :auto-upload="false"
        :file-list="fileList"
        :limit="10"
        accept=".md"
        action=""
      >
        <el-button size="small" type="primary">选择文件</el-button>
        <div slot="tip" class="el-upload__tip">只能上传 md 格式文件，且不超过 2 MB</div>
      </el-upload>
      <el-button style="margin-top: 12px" size="small" type="success"
                 :loading="dialogOptions.uploadLoading" @click="submitUpload">导 入</el-button>
      <el-button style="margin-top: 12px" size="small"
                 @click="dialogOptions.uploadVisible=false">取 消</el-button>
    </el-dialog>

  </d2-container>
</template>

<script>
// eslint-disable-next-line import/no-duplicates
import addEditor from '@/components/aries/post/vditor'
// eslint-disable-next-line import/no-duplicates
import editEditor from '@/components/aries/post/vditor'
import singleTag from '@/components/aries/post/singleTag'
import multiTag from '@/components/aries/post/multiTag'
import state from '@/components/aries/post/state'
import tableHandle from '@/components/aries/post/tableHandle'
import { getAllCategories } from '@/api/aries/category'
import { getAllTags } from '@/api/aries/tag'
import { addPost, deletePost, getPostsByPage, importPostFromFiles, updatePost } from '@/api/aries/post'

export default {
  name: 'post',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    addEditor,
    // eslint-disable-next-line vue/no-unused-components
    editEditor,
    // eslint-disable-next-line vue/no-unused-components
    singleTag,
    // eslint-disable-next-line vue/no-unused-components
    multiTag,
    // eslint-disable-next-line vue/no-unused-components
    tableHandle
  },
  data () {
    return {
      columns: [
        {
          title: '文章标题',
          width: '260',
          key: 'title'
        },
        {
          title: '状态',
          width: '100',
          component: {
            name: state
          }
        },
        {
          title: '分类',
          width: '100',
          component: {
            name: singleTag,
            props: {
              type: 'success'
            }
          }
        },
        {
          title: '标签',
          width: '180',
          component: {
            name: multiTag
          }
        },
        {
          title: '评论数',
          width: '70',
          key: 'comment_count'
        },
        {
          title: '浏览数',
          width: '70',
          key: 'visit_count'
        },
        {
          title: '操作',
          component: {
            name: tableHandle,
            props: {
              openEditDialog: this.openEditDialog,
              handleRowRemove: this.handleRowRemove,
              handleRowRecycleOrRecover: this.handleRowRecycleOrRecover
            }
          }
        }
      ],
      data: [],
      categories: [],
      tags: [],
      states: [
        { name: '已发布', value: 1 },
        { name: '回收', value: 2 },
        { name: '草稿', value: 3 },
        { name: '加密', value: 4 }
      ],
      pagination: {
        currentPage: 1, // 页码
        pageSize: 10, // 每页条数
        total: 0, // 总条数
        key: '', // 搜索关键词
        state: null,
        category_id: null
      },
      addForm: {
        user_id: null,
        category_id: null,
        order_id: 1,
        selectTagIds: [],
        tag_ids: '',
        is_top: false,
        is_recycled: false,
        is_published: true,
        is_allow_commented: true,
        pwd: '',
        url: '',
        title: '',
        summary: '',
        img: '',
        content: '',
        md_content: '',
        keywords: ''
      },
      editForm: {
        user_id: null,
        category_id: null,
        order_id: 1,
        selectTagIds: [],
        tag_ids: '',
        is_top: false,
        is_recycled: false,
        is_published: true,
        is_allow_commented: true,
        pwd: '',
        url: '',
        title: '',
        summary: '',
        img: '',
        content: '',
        md_content: '',
        keywords: ''
      },
      addRules: {
        user_id: [
          { required: true, trigger: 'blur', message: '用户 ID 不能为空' }
        ],
        // category_id: [
        //   { required: true, trigger: 'blur', message: '请选择分类' }
        // ],
        title: [
          { required: true, trigger: 'blur', message: '请输入文章标题' },
          { max: 255, trigger: 'blur', message: 'URL 长度不能超过 255 位' }
        ],
        pwd: [
          { max: 64, trigger: 'blur', message: '密码长度不能超过 64 位' }
        ],
        url: [
          { max: 255, trigger: 'blur', message: 'URL 长度不能超过 255 位' }
        ],
        content: [
          { required: true, pattern: /^(?!\n$)/, trigger: 'blur', message: '请输入文章内容' },
          { max: 100000, trigger: 'blur', message: '内容字数不能超过 100000' }
        ],
        summary: [
          { max: 255, trigger: 'blur', message: '摘要字数不能超过 255' }
        ],
        md_content: [
          { required: true, trigger: 'blur', message: '请输入文章内容' }
        ]
      },
      editRules: {
        user_id: [
          { required: true, trigger: 'blur', message: '用户ID不能为空' }
        ],
        // category_id: [
        //   { required: true, trigger: 'blur', message: '请选择分类' }
        // ],
        title: [
          { required: true, trigger: 'blur', message: '请输入文章标题' },
          { max: 255, trigger: 'blur', message: 'URL 长度不能超过 255 位' }
        ],
        pwd: [
          { max: 64, trigger: 'blur', message: '密码长度不能超过 64 位' }
        ],
        url: [
          { max: 255, trigger: 'blur', message: 'URL 长度不能超过 255 位' }
        ],
        content: [
          { required: true, pattern: /^(?!\n$)/, trigger: 'blur', message: '请输入文章内容' },
          { max: 100000, trigger: 'blur', message: '内容字数不能超过 100000' }
        ],
        summary: [
          { max: 255, trigger: 'blur', message: '摘要字数不能超过 255' }
        ],
        md_content: [
          { required: true, trigger: 'blur', message: '请输入文章内容' }
        ]
      },
      options: {
        border: true
      },
      dialogOptions: {
        addBtnLoading: false,
        addDraftBtnLoading: false,
        addVisible: false,
        addFullScreen: true,
        editBtnLoading: false,
        editDraftBtnLoading: false,
        editVisible: false,
        editFullScreen: true,
        uploadVisible: false,
        uploadLoading: false
      },
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false
      },
      loading: false,
      fileList: []
    }
  },
  created () {
    this.fetchPageData()
    this.fetchCategoryData()
    this.fetchTagData()
  },
  methods: {
    // 分页
    paginationCurrentChange (currentPage) {
      this.pagination.currentPage = currentPage
      this.fetchPageData()
    },
    // 获取分页数据
    fetchPageData () {
      this.loading = true
      setTimeout(() => {
        getPostsByPage({
          page: this.pagination.currentPage,
          size: this.pagination.pageSize,
          category_id: this.pagination.category_id,
          state: this.pagination.state,
          key: this.pagination.key
        })
          .then(res => {
            const pageData = res.data.data
            pageData.forEach((val) => {
              let tagStr = ''
              val.tag_list.forEach((tag, index) => {
                tagStr += `${tag.name},`
              })
              val.tagStr = tagStr
            })
            this.data = pageData
            this.pagination.total = res.data.total_num
            this.loading = false
          })
          .catch(() => {
            this.loading = false
          })
      }, 500)
    },
    // 搜索
    search () {
      this.pagination.currentPage = 1
      this.fetchPageData()
    },
    // 获取分类数据
    fetchCategoryData () {
      getAllCategories(3)
        .then(res => {
          this.categories = res.data
        })
        .catch(() => {
        })
    },
    // 获取标签数据
    fetchTagData () {
      getAllTags()
        .then(res => {
          this.tags = res.data
        })
        .catch(() => {
        })
    },
    // 重置
    reset () {
      this.pagination = {
        currentPage: 1,
        pageSize: 10,
        total: 0,
        key: '',
        state: '',
        category_id: ''
      }
      this.fetchPageData()
    },
    // 重置表单信息
    resetForm (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].resetFields()
      }
    },
    clearValidate (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].clearValidate()
      }
    },
    // 打开添加文章弹窗
    openAddDialog () {
      // 显示弹窗
      this.dialogOptions.addVisible = true
      // 清空表单数据
      this.resetForm('addForm')
    },
    // 打开编辑文章弹窗
    openEditDialog (row) {
      // 显示弹窗
      this.dialogOptions.editVisible = true
      // 清空表单提示信息
      this.clearValidate('editForm')
      // 初始化表单数据
      const tagIds = []
      row.tag_list.forEach((tag) => {
        tagIds.push(tag.ID)
      })
      this.editForm = row
      this.$set(this.editForm, 'selectTagIds', tagIds)
    },
    // 打开导入文章弹窗
    openUploadDialog () {
      // 显示弹窗
      this.dialogOptions.uploadVisible = true
      // 清空上传文件列表
      this.fileList = []
    },
    // 添加文章事件
    handleRowAdd (isPublished) {
      setTimeout(() => {
        // 获取编辑器组件中的文本内容
        this.addForm.content = this.$refs.addEditor.getContent()
        this.addForm.md_content = this.$refs.addEditor.getHTML()
        // 获取 user_id
        this.addForm.user_id = Number(localStorage.getItem('uuid'))
        // 将数组转换成字符串
        this.addForm.tag_ids = this.addForm.selectTagIds.join()
        this.addForm.is_published = isPublished
        // 校验表单
        this.$refs.addForm.validate((valid) => {
          if (valid) {
            if (isPublished) {
              this.dialogOptions.addBtnLoading = true
            } else {
              this.dialogOptions.addDraftBtnLoading = true
            }
            setTimeout(() => {
              addPost(this.addForm)
                .then(res => {
                  this.$message.success(res.msg)
                  this.dialogOptions.addVisible = false
                  this.fetchPageData()
                })
                .catch(() => {
                })
              if (isPublished) {
                this.dialogOptions.addBtnLoading = false
              } else {
                this.dialogOptions.addDraftBtnLoading = false
              }
            }, 500)
          }
        })
      }, 500)
    },
    // 修改文章事件
    handleRowEdit (isPublished) {
      setTimeout(() => {
        // 获取编辑器组件中的文本内容
        this.editForm.content = this.$refs.editEditor.getContent()
        this.editForm.md_content = this.$refs.editEditor.getHTML()
        // 将数组转换成字符串
        this.editForm.tag_ids = this.editForm.selectTagIds.join()
        this.editForm.is_published = isPublished
        // 校验表单
        this.$refs.editForm.validate((valid) => {
          if (valid) {
            if (isPublished) {
              this.dialogOptions.editBtnLoading = true
            } else {
              this.dialogOptions.editDraftBtnLoading = true
            }
            setTimeout(() => {
              updatePost(this.editForm)
                .then(res => {
                  this.$message.success(res.msg)
                  this.dialogOptions.editVisible = false
                  this.fetchPageData()
                })
                .catch(() => {
                })
              if (isPublished) {
                this.dialogOptions.editBtnLoading = false
              } else {
                this.dialogOptions.editDraftBtnLoading = false
              }
            }, 500)
          }
        })
      }, 500)
    },
    // 将文章加入回收站或者恢复
    handleRowRecycleOrRecover (row) {
      let tagIds = []
      row.tag_list.forEach((tag) => {
        tagIds.push(tag.ID)
      })
      tagIds = tagIds.join()
      row.tag_ids = tagIds
      if (row.is_recycled) {
        this.$confirm('确定要恢复吗?', '恢复', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'success'
        })
          .then(() => {
            row.is_recycled = !row.is_recycled
            setTimeout(() => {
              updatePost(row)
                .then(res => {
                  this.$message.success('恢复成功')
                  this.fetchPageData()
                })
                .catch(() => {
                })
            }, 500)
          })
          .catch(() => {
          })
      } else {
        this.$confirm('确定要加入回收站吗?', '加入回收站', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
          .then(() => {
            row.is_recycled = !row.is_recycled
            setTimeout(() => {
              updatePost(row)
                .then(res => {
                  this.$message.success('成功加入回收站')
                  this.fetchPageData()
                })
                .catch(() => {
                })
            }, 500)
          })
          .catch(() => {
          })
      }
    },
    // 删除文章
    handleRowRemove (id) {
      this.$confirm('确定要彻底删除吗?一旦彻底删除，将无法恢复', '彻底删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(() => {
          setTimeout(() => {
            deletePost(id)
              .then(res => {
                this.$message.success(res.msg)
                this.fetchPageData()
              })
              .catch(() => {
              })
          }, 500)
        })
        .catch(() => {
        })
    },
    // 获取文件类型
    getFileType (fileName) {
      return fileName.substring(fileName.lastIndexOf('.') + 1)
    },
    handleExceed (files, fileList) {
      this.$message.warning(
        `当前限制选择 10 个文件，共选择了 ${files.length + fileList.length} 个文件`
      )
    },
    // 文件变动事件
    handleFileChange (file, fileList) {
      this.fileList = fileList
    },
    // 文件删除事件
    handleFileRemove (file, fileList) {
      this.fileList = fileList
    },
    // 上传文件事件
    submitUpload () {
      const formData = new FormData()
      for (let i = 0; i < this.fileList.length; i++) {
        const file = this.fileList[i]
        // 校验文件格式
        if (this.getFileType(file.name) !== 'md') {
          this.$message.error('只支持导入 md 格式的文件')
          return
        }
        // 校验文件大小
        if (file.size > 2 * 1024 * 1024) {
          this.$message.error('文件大小不能超过 2 MB')
          return
        }
        formData.append('file[]', file.raw)
      }
      this.dialogOptions.uploadLoading = true
      setTimeout(() => {
        importPostFromFiles(formData)
          .then(res => {
            this.$message.success(res.msg)
            this.dialogOptions.uploadVisible = false
            this.fetchPageData()
          })
          .catch(() => {
          })
        this.dialogOptions.uploadLoading = false
      }, 500)
    }
  }
}
</script>

<style scoped>
</style>
