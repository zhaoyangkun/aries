<template>
  <d2-container>
    <template slot="header">文章 / 文章</template>
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
          <el-select style="width: 150px" size="small" v-model="pagination.state" clearable
                     placeholder="请选择文章状态">
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
      :with-header="false"
      width="80%"
    >
      <el-form ref="addForm" :model="addForm" :rules="addRules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input size="small" type="text" v-model="addForm.title"></el-input>
        </el-form-item>
        <el-row>
          <el-col :span="8">
            <el-form-item label="分类" prop="category_id">
              <el-select size="small" v-model="addForm.category_id" clearable @clear="editForm.category_id=null"
                         placeholder="请选择分类">
                <el-option
                  v-for="item in categories"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID">
                </el-option>
              </el-select>
              &nbsp;<el-button size="mini" type="primary" @click="handleOpenAddDraw">新增</el-button>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="标签" prop="selectTagIds">
              <el-select size="small" filterable allow-create multiple :multiple-limit=3 v-model="addForm.selectTagIds"
                         clearable placeholder="请选择标签" @change="selectTrigger('addForm')">
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
        <el-form-item label="文章图片">
          <h4 class="img-tip">单击打开附件</h4>
          <img alt="" class="img-show" :src="addImageSrc" @click="handleOpenDrawer('add')"/>
        </el-form-item>
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
              <el-select size="small" v-model="editForm.category_id" clearable placeholder="请选择分类"
                         @clear="editForm.category_id=null">
                <el-option
                  v-for="item in categories"
                  :key="item.ID"
                  :label="item.name"
                  :value="item.ID">
                </el-option>
              </el-select>
              &nbsp;<el-button size="mini" type="primary" @click="handleOpenAddDraw">新增</el-button>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="标签" prop="tag_ids">
              <el-select size="small" filterable allow-create multiple :multiple-limit=3 v-model="editForm.selectTagIds"
                         clearable placeholder="请选择标签" @change="selectTrigger('editForm')">
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
        <el-form-item label="文章图片">
          <h4 class="img-tip">单击打开附件</h4>
          <img alt="" class="img-show" :src="editImageSrc" @click="handleOpenDrawer('edit')"/>
        </el-form-item>
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
                 :loading="dialogOptions.uploadLoading" @click="submitUpload">导 入
      </el-button>
      <el-button style="margin-top: 12px" size="small"
                 @click="dialogOptions.uploadVisible=false">取 消
      </el-button>
    </el-dialog>

    <!-- 添加分类抽屉 -->
    <el-drawer
      title="添加分类"
      :visible.sync="drawOptions.addVisible"
      direction="rtl"
    >
      <el-form ref="addCategoryForm" :model="addCategoryForm" :rules="addCategoryRules" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input style="width: 280px" size="small" type="text" v-model="addCategoryForm.name"></el-input>
        </el-form-item>
        <el-form-item label="url" prop="url">
          <el-input style="width: 280px" size="small" type="text" v-model="addCategoryForm.url"></el-input>
        </el-form-item>
        <el-form-item label="父级分类" prop="parent_id">
          <el-select size="small" v-model="addCategoryForm.parent_id" clearable placeholder="请选择父级分类">
            <el-option
              v-for="item in parentCategories"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="drawOptions.addBtnLoading"
                     @click="handleCategoryAdd">保存
          </el-button>
          <el-button @click="drawOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <!--附件抽屉-->
    <attachDrawer ref="attachDrawer" @changeClickedImg="changeClickedImg($event)"></attachDrawer>
  </d2-container>
</template>

<script>
import addEditor from '@/components/aries/post/addVditor'
import editEditor from '@/components/aries/post/editVditor'
import singleTag from '@/components/aries/post/singleTag'
import multiTag from '@/components/aries/post/multiTag'
import state from '@/components/aries/post/state'
import tableHandler from '@/components/aries/post/tableHandler'
import postTitle from '@/components/aries/post/postTitle'
import attachDrawer from '@/components/aries/common/attachDrawer'
import { addArticleCategory, getAllCategories, getAllParentCategories } from '@/api/aries/category'
import { addTag, getAllTags } from '@/api/aries/tag'
import {
  addPost,
  deletePost,
  getPostById,
  getPostsByPage,
  importPostFromFiles,
  movePostDown,
  movePostUp,
  recycleOrRecoverPost,
  updatePost
} from '@/api/aries/post'

export default {
  name: 'post',
  components: {
    /* eslint-disable vue/no-unused-components */
    addEditor,
    editEditor,
    singleTag,
    multiTag,
    tableHandler,
    postTitle,
    attachDrawer
  },
  data () {
    return {
      columns: [
        {
          title: '文章标题',
          width: '230',
          component: {
            name: postTitle
          }
        },
        {
          title: '状态',
          width: '80',
          component: {
            name: state
          }
        },
        {
          title: '分类',
          width: '80',
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
            name: tableHandler,
            props: {
              openEditDialog: this.openEditDialog,
              handleRowRemove: this.handleRowRemove,
              handleRowRecycleOrRecover: this.handleRowRecycleOrRecover,
              handleMoveUp: this.handleMoveUp,
              handleMoveDown: this.handleMoveDown
            }
          }
        }
      ],
      data: [],
      categories: [],
      parentCategories: [],
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
      addCategoryForm: {
        name: '',
        url: '',
        parent_id: null
      },
      addRules: {
        user_id: [
          { required: true, trigger: 'blur', message: '用户 ID 不能为空' }
        ],
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
      addCategoryRules: {
        name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
        url: [{ required: true, message: '请输入 URL', trigger: 'blur' }]
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
      drawOptions: {
        addVisible: false,
        addBtnLoading: false
      },
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false
      },
      addImageSrc: '/image/none.jpg',
      editImageSrc: '/image/none.jpg',
      loading: false,
      fileList: []
    }
  },
  created () {
    this.fetchPageData()
    this.fetchCategoryData()
    this.fetchParentCategories()
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
            this.data = res.data.data
            this.pagination.total = res.data.total_num
            this.loading = false
          })
          .catch(() => {
            this.loading = false
          })
      }, 300)
    },
    // 搜索
    search () {
      this.pagination.currentPage = 1
      this.fetchPageData()
    },
    // 获取分类数据
    fetchCategoryData () {
      getAllCategories(0)
        .then(res => {
          this.categories = res.data
        })
        .catch(() => {
        })
    },
    // 获取父级分类
    fetchParentCategories () {
      this.parentCategories = [{ value: 0, label: '请选择父级分类' }]
      getAllParentCategories(0)
        .then(res => {
          this.parentCategories = res.data
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
    // 清空表单验证
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
    openEditDialog (id) {
      // 显示弹窗
      this.dialogOptions.editVisible = true
      // 清空表单提示信息
      this.clearValidate('editForm')
      // 初始化表单数据
      getPostById(id)
        .then(res => {
          this.editForm = res.data
          this.editForm.category_id = this.editForm.category_id === 0 ? null : this.editForm.category_id
          const tagIds = []
          this.editForm.tag_list.forEach((tag) => {
            tagIds.push(tag.ID)
          })
          this.$set(this.editForm, 'selectTagIds', tagIds)
          this.editImageSrc = this.editForm.img
          this.$refs.editEditor.setContent(this.editForm.content)
        })
        .catch(() => {
        })
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
          }, 300)
        }
      })
    },
    // 修改文章事件
    handleRowEdit (isPublished) {
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
          }, 300)
        }
      })
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
              recycleOrRecoverPost(row.ID)
                .then(() => {
                  this.$message.success('恢复成功')
                  this.fetchPageData()
                })
                .catch(() => {
                })
            }, 300)
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
              recycleOrRecoverPost(row.ID)
                .then(() => {
                  this.$message.success('成功加入回收站')
                  this.fetchPageData()
                })
                .catch(() => {
                })
            }, 300)
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
          }, 300)
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
      if (this.fileList.length === 0) {
        this.$message.error('请选择要上传的文件')
        return
      }
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
      }, 300)
    },
    // 显示添加分类抽屉
    handleOpenAddDraw () {
      this.drawOptions.addVisible = true
      this.resetForm('addCategoryForm')
    },
    // 添加分类事件
    handleCategoryAdd () {
      this.$refs.addCategoryForm.validate((valid) => {
        if (valid) {
          this.drawOptions.addBtnLoading = true
          setTimeout(() => {
            addArticleCategory(this.addCategoryForm)
              .then(res => {
                this.$message.success(res.msg)
                this.fetchCategoryData()
                this.fetchParentCategories()
                this.drawOptions.addVisible = false
              })
              .catch(() => {
              })
            this.drawOptions.addBtnLoading = false
          }, 300)
        }
      })
    },
    // 标签改变事件
    selectTrigger (formName) {
      const nameList = []
      this.parentCategories.forEach(category => {
        nameList.push(category.name)
      })
      this[`${formName}`].selectTagIds.forEach((id, i) => {
        if (typeof (id) === 'string' && nameList.indexOf(id) === -1) {
          addTag({
            name: id
          })
            .then(res => {
              // 用 $set 动态修改标签列表，避免视图不更新
              this.$set(this.tags, this.tags.length, res.data)
              this[`${formName}`].selectTagIds[i] = res.data.ID
            })
            .catch(() => {
            })
        }
      })
    },
    // 向上移动文章事件
    handleMoveUp (row, index) {
      const form = {
        id: row.ID,
        order_id: row.order_id,
        is_top: row.is_top
      }
      movePostUp(form)
        .then(res => {
          this.$message.success(res.msg)
          this.fetchPageData()
          // const newData = [...this.data]
          // const temp = newData[index - 1]
          // newData[index - 1] = newData[index]
          // newData[index] = temp
          // this.data = newData
        })
        .catch(() => {
        })
    },
    // 向下移动文章事件
    handleMoveDown (row, index) {
      const form = {
        id: row.ID,
        order_id: row.order_id,
        is_top: row.is_top
      }
      movePostDown(form)
        .then(res => {
          this.$message.success(res.msg)
          this.fetchPageData()
          // const newData = [...this.data]
          // const temp = newData[index + 1]
          // newData[index + 1] = newData[index]
          // newData[index] = temp
          // this.data = newData
        })
        .catch(() => {
        })
    },
    // 打开弹窗
    handleOpenDrawer (mode) {
      this.$refs.attachDrawer.openDrawer()
      this.mode = mode
    },
    // 选中图片
    changeClickedImg (url) {
      if (this.mode === 'add') {
        this.addForm.img = url
        this.addImageSrc = url
      } else {
        this.editForm.img = url
        this.editImageSrc = url
      }
      this.$refs.attachDrawer.closeDrawer()
    }
  }
}
</script>

<style lang="scss" scoped>
:focus {
  outline: 0;
}

.operation-box {
  padding: 0;
}

.el-card__body {
  padding: 0;
}

.card-num {
  margin: 5px;
}

.no-tip {
  color: #99aabb;
}

.theme-pic {
  height: 160px;
  width: 320px;
}

.block {
  margin: 5px;
}

.demonstration {
  padding-left: 5%;
  height: 30px;
  width: 95%;
  line-height: 30px;
  font-size: 13px;
  display: block;
  color: grey;
  background-color: white;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.page-box {
  margin: auto;
}

.pre-tip {
  margin: 12px 0;
  line-height: 20px;
  font-weight: normal;
  color: grey;
}

.image-container {
  height: 126px;
  padding: 0;
  margin: 0;
}

.attach-image {
  width: 100%;
  height: 96px;
  margin: 0;
  padding: 0;
  display: inline-block;
  overflow: hidden;
}

.img_not_checked {
  border: 2px rgba(200, 200, 200, 0.18) solid;
}

.img_checked {
  border: #409EFF solid 2px;
}

.el-tag + .el-tag {
  margin-left: 10px;
}

.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}

.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}

.img-tip {
  height: 40px;
  line-height: 40px;
  margin: 0;
  padding: 0;
  color: #99aabb;
  font-weight: normal;
}

.img-show {
  width: 320px;
  height: 180px;
}
</style>
