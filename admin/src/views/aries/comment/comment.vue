<template>
  <d2-container>
    <template slot="header">用户 / 评论</template>
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
      <el-form ref="replyForm" :model="replyForm" :rules="replyFormRules" label-width="80px">
        <el-form-item label="内容" prop="content">
          <editVditor ref="editEditor"></editVditor>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.replyBtnLoading"
                     @click="handleReply">通过审核并回复
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
import { getSysSettingItem } from '@api/aries/sys'

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
          key: 'article_title',
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
              handleRowRecycleOrRecover: this.handleRowRecycleOrRecover,
              handleChecked: this.handleChecked
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
        user_img: '',
        content: '',
        md_content: '',
        type: 0,
        is_recycled: false,
        is_checked: true
      },
      replyFormRules: {
        content: [
          { required: true, pattern: /^(?!\n$)/, trigger: 'blur', message: '请输入回复内容' },
          { min: 6, max: 1200, trigger: 'blur', message: '内容字数要在 6 ~ 1200 之间' }
        ]
      }
    }
  },
  created () {
    this.fetchPageData()
    this.fetchUserData()
    this.fetchSiteSetting()
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
          this.replyForm.user_img = res.data[0].user_img
        })
        .catch(() => {
        })
    },
    // 获取网站设置
    fetchSiteSetting () {
      getSysSettingItem('网站设置')
        .then(res => {
          this.replyForm.url = res.data.site_url
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
      this.currRow = { ...row }
    },
    // 回复事件
    handleReply () {
      const OS = this.getOS()
      const browser = this.getBrowser()
      this.replyForm.root_comment_id = this.currRow.ID
      if (this.currRow.root_comment_id > 0) {
        this.replyForm.root_comment_id = this.currRow.root_comment_id
      }
      this.replyForm.device = `${browser.browser} ${browser.version} in ${OS}`
      this.replyForm.article_id = this.currRow.article_id
      this.replyForm.parent_comment_id = this.currRow.ID
      this.replyForm.type = this.currRow.type
      this.replyForm.content = this.$refs.editEditor.getContent()
      this.replyForm.md_content = this.$refs.editEditor.getHTML()
      this.$refs.replyForm.validate((valid) => {
        if (valid) {
          if (this.replyForm.nick_name === '') {
            this.$message.error('请先设置呢称')
            return
          }
          if (this.replyForm.user_img === '') {
            this.$message.error('请先设置用户头像')
            return
          }
          this.dialogOptions.replyBtnLoading = true
          setTimeout(() => {
            addComment(this.replyForm)
              .then(() => {
                this.$message.success('回复成功')
                this.dialogOptions.replyVisible = false
                this.fetchPageData()
              })
              .catch(() => {
              })
          }, 300)
          this.dialogOptions.replyBtnLoading = false
        }
      })
    },
    // 将评论加入回收站或者恢复
    handleRowRecycleOrRecover (row) {
      const data = { ...row }
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
    // 通过审核
    handleChecked (row) {
      const data = { ...row }
      setTimeout(() => {
        data.is_checked = true
        updateComment(data)
          .then(() => {
            this.$message.success('成功通过审核')
            this.fetchPageData()
          })
          .catch(() => {
          })
      }, 300)
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
    },
    getOS () {
      const sUserAgent = navigator.userAgent

      const isIphone = sUserAgent.indexOf('iPhone') > -1
      if (isIphone) return 'iPhone'
      const isIpod = sUserAgent.indexOf('iPod') > -1
      if (isIpod) return 'iPod'
      const isIpad = sUserAgent.indexOf('iPad') > -1
      if (isIpad) return 'iPad'
      const isAndroid = sUserAgent.indexOf('Android') > -1
      if (isAndroid) return 'Android'

      const isWin =
        navigator.platform === 'Win32' || navigator.platform === 'Windows'
      const isMac =
        navigator.platform === 'Mac68K' ||
        navigator.platform === 'MacPPC' ||
        navigator.platform === 'Macintosh' ||
        navigator.platform === 'MacIntel'
      if (isMac) return 'Mac'
      const isUnix = navigator.platform === 'X11' && !isWin && !isMac
      if (isUnix) return 'Unix'
      const isLinux = String(navigator.platform).indexOf('Linux') > -1
      if (isLinux) return 'Linux'
      if (isWin) {
        const isWin2K =
          sUserAgent.indexOf('Windows NT 5.0') > -1 ||
          sUserAgent.indexOf('Windows 2000') > -1
        if (isWin2K) return 'Win2000'
        const isWinXP =
          sUserAgent.indexOf('Windows NT 5.1') > -1 ||
          sUserAgent.indexOf('Windows XP') > -1
        if (isWinXP) return 'WinXP'
        const isWin2003 =
          sUserAgent.indexOf('Windows NT 5.2') > -1 ||
          sUserAgent.indexOf('Windows 2003') > -1
        if (isWin2003) return 'Win2003'
        const isWinVista =
          sUserAgent.indexOf('Windows NT 6.0') > -1 ||
          sUserAgent.indexOf('Windows Vista') > -1
        if (isWinVista) return 'WinVista'
        const isWin7 =
          sUserAgent.indexOf('Windows NT 6.1') > -1 ||
          sUserAgent.indexOf('Windows 7') > -1
        if (isWin7) return 'Win7'
        const isWin10 =
          sUserAgent.indexOf('Windows NT 10') > -1 ||
          sUserAgent.indexOf('Windows 10') > -1
        if (isWin10) return 'Win10'
      }
      return 'unknown OS'
    },
    getBrowser () {
      const ua = navigator.userAgent.toLowerCase()
      let s

      s = ua.match(/edg\/([\d.]+)/)
      if (s) return { browser: 'Edge', version: s[1] }
      s = ua.match(/rv:([\d.]+)\) like gecko/)
      if (s) return { browser: 'IE', version: s[1] }
      s = ua.match(/msie ([\d.]+)/)
      if (s) return { browser: 'IE', version: s[1] }
      s = ua.match(/firefox\/([\d.]+)/)
      if (s) return { browser: 'Firefox', version: s[1] }
      s = ua.match(/chrome\/([\d.]+)/)
      if (s) return { browser: 'Chrome', version: s[1] }
      s = ua.match(/opera.([\d.]+)/)
      if (s) return { browser: 'Opera', version: s[1] }
      s = ua.match(/version\/([\d.]+).*safari/)
      if (s) return { browser: 'Safari', version: s[1] }

      return { browser: 'Unknown Browser', version: '' }
    }
  }
}
</script>

<style scoped>
</style>
