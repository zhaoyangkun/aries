<template>
  <d2-container>
    <template slot="header">用户 / 页面 / 日志</template>
    <div class="operation-box">
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item>
          <el-button size="small" type="danger" @click="handleMultiDelJournals">
            <i class="el-icon-delete"></i> 批量删除
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="dialogOptions.addVisible=true">
            <i class="el-icon-plus"></i> 写日志
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-input size="small" clearable placeholder="请输入关键字" v-model="pagination.key"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="handleSearch">
            <i class="el-icon-search"></i> 搜索
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" @click="handleReset">
            <i class="el-icon-refresh"></i> 重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table
      v-loading="loading"
      @selection-change="handleSelectChange"
      :data="tableData"
      style="width: 100%;margin-bottom: 20px;"
      row-key="ID"
    >
      <el-table-column
        type="selection"
        width="55"
      >
      </el-table-column>
      <el-table-column
        prop="content"
        label="内容"
      >
      </el-table-column>
      <el-table-column
        prop="is_secret"
        label="状态"
        width="100"
      >
        <template slot-scope="scope">
          <el-tag type="info" effect="plain" v-if="scope.row.is_secret">私密</el-tag>
          <el-tag type="success" effect="plain" v-else>公开</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        prop="CreatedAt"
        label="创建时间"
        width="160"
      >
        <template slot-scope="scope">
          <date :scope="scope" fmt="yyyy-MM-dd hh:mm"></date>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        width="160">
        <template slot-scope="scope">
          <el-button size="small" type="text" @click="handleOpenEditDialog(scope.row)">编辑</el-button>&nbsp;
          <el-popconfirm
            confirmButtonText='确定'
            cancelButtonText='取消'
            icon="el-icon-info"
            iconColor="red"
            title="确定要删除吗？"
            @onConfirm="handleDelJournal(scope.row.ID)"
          >
            <el-button slot="reference" size="small" type="text">删除</el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

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

    <el-dialog
      title="写日志"
      :visible.sync="dialogOptions.addVisible"
      :with-header="false"
      width="50%"
    >
      <el-form ref="addForm" :model="addForm" :rules="addFormRules" label-width="80px">
        <el-form-item label="日志内容" prop="content">
          <el-input size="small" type="textarea" :rows="8" v-model="addForm.content"></el-input>
        </el-form-item>
        <el-form-item label="是否私密" prop="is_secret">
          <el-switch
            v-model="addForm.is_secret"
            active-text="私密"
            inactive-text="公开"
            active-color="#EBEEF5"
            inactive-color="#409EFF"
          >
          </el-switch>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.addBtnLoading" @click="handleAddJournal">保存</el-button>
          <el-button @click="dialogOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <el-dialog
      title="编辑日志"
      :visible.sync="dialogOptions.editVisible"
      :with-header="false"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="80px">
        <el-form-item label="日志内容" prop="content">
          <el-input size="small" type="textarea" :rows="8" v-model="editForm.content"></el-input>
        </el-form-item>
        <el-form-item label="是否私密" prop="is_secret">
          <el-switch
            v-model="editForm.is_secret"
            active-text="私密"
            inactive-text="公开"
            active-color="#EBEEF5"
            inactive-color="#409EFF"
          >
          </el-switch>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.editBtnLoading" @click="handleEditJournal">保存</el-button>
          <el-button @click="dialogOptions.editVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </d2-container>
</template>

<script>
import date from '@/components/aries/common/date'
import { createJournal, getJournalsByPage, multiDelJournals, updateJournal } from '@api/aries/journal'

export default {
  name: 'journal',
  components: {
    date
  },
  data () {
    return {
      pagination: {
        page: 1,
        size: 10,
        key: '',
        total_num: 0,
        total_pages: 0
      },
      selectList: [],
      tableData: [],
      addForm: {
        is_secret: false,
        content: ''
      },
      addFormRules: {
        content: [
          { required: true, trigger: 'blur', message: '请输入日志内容' },
          { max: 255, trigger: 'blur', message: '日志内容不能超过 255 个字符' }
        ]
      },
      editForm: {
        ID: null,
        is_secret: false,
        content: ''
      },
      editFormRules: {
        content: [
          { required: true, trigger: 'blur', message: '请输入日志内容' },
          { max: 255, trigger: 'blur', message: '日志内容不能超过 255 个字符' }
        ]
      },
      dialogOptions: {
        addVisible: false,
        addBtnLoading: false,
        editVisible: false,
        editBtnLoading: false
      },
      loading: false
    }
  },
  created () {
    this.fetchPageData()
  },
  methods: {
    // 重置表单
    resetForm (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].resetFields()
      }
    },
    // 清空表单校验
    clearValidate (formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].clearValidate()
      }
    },
    // 获取分页数据
    fetchPageData () {
      this.loading = true
      setTimeout(() => {
        getJournalsByPage({
          page: this.pagination.page,
          size: this.pagination.size,
          key: this.pagination.key
        })
          .then(res => {
            this.tableData = res.data.data
            this.pagination.total_num = res.data.total_num
            this.pagination.total_pages = res.data.total_pages
          })
          .catch(() => {
          })
        this.loading = false
      }, 300)
    },
    // 搜索
    handleSearch () {
      this.pagination.page = 1
      this.fetchPageData()
    },
    // 重置
    handleReset () {
      this.pagination.page = 1
      this.pagination.key = ''
      this.fetchPageData()
    },
    // 改变每页条数
    handlePageSizeChange (size) {
      this.pagination.size = size
      this.fetchPageData()
    },
    // 改变当前页
    handleCurrentPageChange (page) {
      this.pagination.page = page
      this.fetchPageData()
    },
    // 添加日志
    handleAddJournal () {
      this.$refs.addForm.validate(valid => {
        if (valid) {
          this.dialogOptions.addBtnLoading = true
          setTimeout(() => {
            createJournal(this.addForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.addVisible = false
                this.resetForm('addForm')
                this.fetchPageData()
              })
              .catch(() => {
              })
            this.dialogOptions.addBtnLoading = false
          }, 300)
        }
      })
    },
    // 打开修改日志弹窗
    handleOpenEditDialog (data) {
      this.editForm = { ...data }
      this.dialogOptions.editVisible = true
    },
    // 修改日志
    handleEditJournal () {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          this.dialogOptions.editBtnLoading = true
          setTimeout(() => {
            updateJournal(this.editForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.editVisible = false
                this.resetForm('editForm')
                this.fetchPageData()
              })
              .catch(() => {
              })
            this.dialogOptions.editBtnLoading = false
          }, 300)
        }
      })
    },
    // 删除日志
    handleDelJournal (id) {
      multiDelJournals([id].join(','))
        .then(res => {
          this.$message.success(res.msg)
          this.fetchPageData()
        })
        .catch(() => {
        })
    },
    // 选项发生变化
    handleSelectChange (selection) {
      selection.forEach(val => {
        this.selectList.push(val.ID)
      })
    },
    // 批量删除日志
    handleMultiDelJournals () {
      if (this.selectList.length === 0) {
        this.$message.error('请勾选要删除的日志')
        return
      }
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(() => {
          const ids = this.selectList.join(',')
          multiDelJournals(ids)
            .then(res => {
              this.$message.success(res.msg)
              this.fetchPageData()
            })
            .catch(() => {
            })
        })
        .catch(() => {
        })
    }
  }
}
</script>

<style scoped>
</style>
