<template>
  <d2-container>
    <template slot="header">评论列表</template>
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
          <el-select style="width: 150px" size="small" v-model="pagination.state" clearable placeholder="请选择评论状态">
            <el-option
              v-for="item in stateList"
              :key="item.value"
              :label="item.name"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select style="width: 150px" size="small" v-model="pagination.type" clearable placeholder="请选择评论类型">
            <el-option
              v-for="item in typeList"
              :key="item.value"
              :label="item.name"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input size="small" clearable placeholder="请输入关键字" v-model="pagination.key"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="fetchPageData"><i class="el-icon-search"></i> 搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" @click="resetTable"><i class="el-icon-refresh"></i> 重置</el-button>
        </el-form-item>
      </el-form>
    </d2-crud>

    <!-- 回复弹窗 -->
    <el-dialog
      title="回复"
      :visible.sync="dialogOptions.replyVisible"
      width="50%"
    >
      <el-form ref="replyForm" :model="replyForm" label-width="80px">
        <el-form-item label="内容" prop="content">
          <editVditor ref="editEditor"></editVditor>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.replyBtnLoading"
                     @click="handleReply">回复
          </el-button>
          <el-button @click="dialogOptions.replyVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </d2-container>
</template>

<script>
import { addComment, deleteComment, getCommentsByPage, updateComment } from '@/api/aries/comment'
import nickName from '@/components/aries/comment/nickName'
import date from '@/components/aries/common/date'
import state from '@/components/aries/comment/state'
import tableHandler from '@/components/aries/comment/tableHandler'
import editVditor from '@/components/aries/post/editVditor'
import { getAllUsers } from '@/api/aries/user'

export default {
  name: 'comment',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    date,
    // eslint-disable-next-line vue/no-unused-components
    state,
    // eslint-disable-next-line vue/no-unused-components
    nickName,
    // eslint-disable-next-line vue/no-unused-components
    tableHandler,
    editVditor
  },
  data () {
    return {
      loading: false,
      options: {
        border: true
      },
      columns: [
        {
          title: '呢称',
          width: '120',
          component: {
            name: nickName
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
          title: '评论内容',
          key: 'content',
          width: '200'
        },
        {
          title: '所属文章',
          key: 'article.title',
          width: '180'
        },
        {
          title: '发布时间',
          width: '138',
          component: {
            name: date,
            props: {
              fmt: 'yyyy-MM-dd hh:mm'
            }
          }
        },
        {
          title: '操作',
          width: '200',
          component: {
            name: tableHandler,
            props: {
              openReplyDialog: this.openReplyDialog,
              handleRowRemove: this.handleRowRemove,
              handleRowRecycleOrRecover: this.handleRowRecycleOrRecover
            }
          }
        }
      ],
      data: [],
      pagination: {
        currentPage: 1, // 页码
        pageSize: 10, // 每页条数
        total: 0, // 总条数
        state: null, // 状态
        type: null, // 类型
        key: '' // 关键词
      },
      formOptions: {
        labelWidth: '80px',
        labelPosition: 'left',
        saveLoading: false
      },
      dialogOptions: {
        replyVisible: false,
        replyBtnLoading: false
      },
      typeList: [
        { value: 1, name: '文章评论' },
        { value: 2, name: '页面评论' }
      ],
      stateList: [
        { value: 1, name: '回收站' },
        { value: 2, name: '待审核' },
        { value: 3, name: '已发布' }
      ],
      currRow: {},
      replyForm: {
        article_id: null,
        admin_user_id: null,
        root_comment_id: null,
        parent_comment_id: null,
        email: '',
        nick_name: '',
        url: '',
        content: '',
        md_content: ''
      }
    }
  },
  created () {
    this.fetchPageData()
    this.fetchUserData()
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
        getCommentsByPage({
          page: this.pagination.currentPage,
          size: this.pagination.pageSize,
          state: this.pagination.state,
          type: this.pagination.type,
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
    // 获取用户数据
    fetchUserData () {
      getAllUsers()
        .then(res => {
          this.replyForm.admin_user_id = res.data[0].ID
          this.replyForm.email = res.data[0].email
          this.replyForm.nick_name = res.data[0].nickname
          this.replyForm.url = res.data[0].site_url
        })
        .catch(() => {
        })
    },
    // 重置表格
    resetTable () {
      this.pagination = {
        currentPage: 1, // 页码
        pageSize: 10, // 每页条数
        total: 0, // 总条数
        state: null, // 状态
        key: '' // 关键词
      }
      this.fetchPageData()
    },
    // 重置表单信息
    resetForm (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].resetFields()
      }
    },
    // 打开回复弹窗
    openReplyDialog (row) {
      this.resetForm('replyForm')
      this.dialogOptions.replyVisible = true
      this.currRow = row
    },
    // 回复事件
    handleReply () {
      if (this.replyForm.root_comment_id === 0) {
        this.replyForm.root_comment_id = this.currRow.ID
      } else {
        this.replyForm.root_comment_id = this.currRow.root_comment_id
      }
      this.replyForm.article_id = this.currRow.article_id
      this.replyForm.parent_comment_id = this.currRow.ID
      this.replyForm.content = this.$refs.editEditor.getContent()
      this.replyForm.md_content = this.$refs.editEditor.getHTML()
      this.dialogOptions.replyBtnLoading = true
      console.log('replyForm: ', this.replyForm)
      setTimeout(() => {
        addComment(this.replyForm)
          .then(() => {
            this.$message.success('回复成功')
            this.dialogOptions.replyVisible = false
            this.fetchPageData()
          })
          .catch(() => {
          })
        this.dialogOptions.replyBtnLoading = false
      }, 300)
    },
    // 将评论加入回收站或者恢复
    handleRowRecycleOrRecover (row) {
      const data = row
      if (data.is_recycled) {
        this.$confirm('确定要恢复吗?', '恢复', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'success'
        })
          .then(() => {
            data.is_recycled = !data.is_recycled
            setTimeout(() => {
              updateComment(data)
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
            data.is_recycled = !data.is_recycled
            setTimeout(() => {
              updateComment(data)
                .then(() => {
                  this.$message.success('加入回收站成功')
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
    // 删除评论
    handleRowRemove (id) {
      this.$confirm('确定要彻底删除吗?一旦彻底删除，将无法恢复', '彻底删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(() => {
          setTimeout(() => {
            deleteComment(id)
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
    }
  }
}
</script>

<style scoped>

</style>
