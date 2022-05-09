<template>
  <d2-container>
    <template slot="header">外观 / 菜单</template>
    <div class="operation-box">
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item>
          <el-button size="small" type="danger" @click="handleNavsMultiDel">
            <i class="el-icon-delete"></i> 批量删除
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="openAddDialog">
            <i class="el-icon-plus"></i> 新增
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table
      v-loading="loading"
      :data="tableData"
      @selection-change="handleSelectChange"
      style="width: 100%;margin-bottom: 20px;"
      row-key="ID"
      :tree-props="{children: 'child_navs', hasChildren: 'hasChildren'}">
      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column
        prop="name"
        label="菜单名称"
        width="180">
      </el-table-column>
      <el-table-column
        prop="url"
        label="URL"
        width="180">
      </el-table-column>
      <el-table-column
        label="操作">
        <template slot-scope="scope">
          <el-tooltip content="编辑" placement="top-start">
            <el-button size="small" icon="el-icon-edit" @click="openEditDialog(scope.row)"></el-button>
          </el-tooltip>
          <el-tooltip content="向上移动" placement="top-start">
            <el-button size="small" icon="el-icon-top"
                       @click="handleNavMoveUp(scope.row,scope.$index)"></el-button>
          </el-tooltip>
          <el-tooltip content="向下移动" placement="top-start">
            <el-button size="small" icon="el-icon-bottom"
                       @click="handleNavMoveDown(scope.row,scope.$index)"></el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top-start">
            <el-button size="small" type="danger" icon="el-icon-delete"
                       @click="handleNavDelete(scope.row.ID)"></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加菜单弹窗 -->
    <el-dialog
      title="添加菜单"
      :visible.sync="dialogOptions.addVisible"
      :with-header="false"
      width="50%"
    >
      <el-form ref="addForm" :model="addForm" :rules="addRules" label-width="80px">
        <el-form-item label="菜单名称" prop="name">
          <el-input size="small" type="text" v-model="addForm.name"></el-input>
        </el-form-item>
        <el-form-item label="父级菜单" prop="parent_nav_id">
          <el-select size="small" v-model="addForm.parent_nav_id" clearable @clear="addForm.parent_nav_id=null"
                     placeholder="请选择分类">
            <el-option
              v-for="item in tableData"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="打开方式" prop="open_type">
          <el-radio v-model="addForm.open_type" :label="curWindow">当前窗口</el-radio>
          <el-radio v-model="addForm.open_type" :label="newWindow">新窗口</el-radio>
        </el-form-item>
        <el-form-item label="URL" prop="url">
          <el-input size="small" type="text" v-model="addForm.url"></el-input>
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-input size="small" type="text" v-model="addForm.icon"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.addBtnLoading" @click="handleNavAdd">保存</el-button>
          <el-button @click="dialogOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 修改菜单弹窗 -->
    <el-dialog
      title="修改菜单"
      :visible.sync="dialogOptions.editVisible"
      :with-header="false"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editRules" label-width="80px">
        <el-form-item label="菜单名称" prop="name">
          <el-input size="small" type="text" v-model="editForm.name"></el-input>
        </el-form-item>
        <el-form-item label="父级菜单" prop="parent_nav_id">
          <el-select size="small" v-model="editForm.parent_nav_id" clearable @clear="editForm.parent_nav_id=null"
                     placeholder="请选择分类">
            <el-option
              v-for="item in tableData"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="打开方式" prop="open_type">
          <el-radio v-model="editForm.open_type" :label="curWindow">当前窗口</el-radio>
          <el-radio v-model="editForm.open_type" :label="newWindow">新窗口</el-radio>
        </el-form-item>
        <el-form-item label="URL" prop="url">
          <el-input size="small" type="text" v-model="editForm.url"></el-input>
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-input size="small" type="text" v-model="editForm.icon"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.editBtnLoading" @click="handleNavEdit">保存</el-button>
          <el-button @click="dialogOptions.editVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </d2-container>
</template>

<script>
import { addNav, deleteNav, getNavs, moveNavDown, moveNavUp, multiDelNavs, updateNav } from '@api/aries/nav'

export default {
  name: 'navs',
  data () {
    return {
      tableData: [],
      selectList: [],
      addForm: {
        name: '',
        parent_nav_id: null,
        open_type: 0,
        url: '',
        icon: ''
      },
      editForm: {
        ID: null,
        name: '',
        parent_nav_id: null,
        open_type: 0,
        url: '',
        icon: ''
      },
      addRules: {
        name: [
          { required: true, trigger: 'blur', message: '请输入菜单名称' },
          { max: 100, trigger: 'blur', message: '菜单名称不能超过 100 个字符' }
        ],
        url: [
          { required: true, trigger: 'blur', message: '请输入 URL' },
          { max: 255, trigger: 'blur', message: '菜单名称不能超过 255 个字符' }
        ],
        icon: [
          { max: 255, trigger: 'blur', message: '图标不能超过 255 个字符' }
        ]
      },
      editRules: {
        name: [
          { required: true, trigger: 'blur', message: '请输入菜单名称' },
          { max: 100, trigger: 'blur', message: '菜单名称不能超过 100 个字符' }
        ],
        url: [
          { required: true, trigger: 'blur', message: '请输入 URL' },
          { max: 255, trigger: 'blur', message: '菜单名称不能超过 255 个字符' }
        ],
        icon: [
          { max: 255, trigger: 'blur', message: '图标不能超过 255 个字符' }
        ]
      },
      dialogOptions: {
        addVisible: false,
        addBtnLoading: false,
        editVisible: false,
        editBtnLoading: false
      },
      curWindow: 0,
      newWindow: 1,
      loading: false
    }
  },
  created () {
    this.fetchTableData()
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
    // 获取菜单数据
    fetchTableData () {
      this.loading = true
      setTimeout(() => {
        getNavs()
          .then(res => {
            this.tableData = res.data
          })
          .catch(() => {
          })
        this.loading = false
      }, 300)
    },
    // 打开添加菜单弹窗
    openAddDialog () {
      this.resetForm('addForm')
      this.dialogOptions.addVisible = true
    },
    // 打开修改菜单弹窗
    openEditDialog (row) {
      if (row.parent_nav_id === 0) {
        row.parent_nav_id = null
      }
      this.editForm = { ...row }
      this.dialogOptions.editVisible = true
    },
    // 添加菜单
    handleNavAdd () {
      this.$refs.addForm.validate(valid => {
        if (valid) {
          this.dialogOptions.addBtnLoading = true
          setTimeout(() => {
            addNav(this.addForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.addVisible = false
                this.fetchTableData()
              })
              .catch(() => {
              })
            this.dialogOptions.addBtnLoading = false
          }, 300)
        }
      })
    },
    // 修改菜单
    handleNavEdit () {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          if (this.editForm.parent_nav_id === this.editForm.ID) {
            this.$message.error('不能将自身设置为父级菜单')
            return
          }
          this.dialogOptions.editBtnLoading = true
          setTimeout(() => {
            updateNav(this.editForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.editVisible = false
                this.fetchTableData()
              })
              .catch(() => {
              })
            this.dialogOptions.editBtnLoading = false
          }, 300)
        }
      })
    },
    // 向上移动菜单
    handleNavMoveUp (row, index) {
      const navType = row.parent_nav_id === 0 ? 'parent' : 'child'
      moveNavUp(navType, row.order_id)
        .then(res => {
          this.$message.success(res.msg)
          this.fetchTableData()
        })
        .catch(() => {
        })
    },
    // 向下移动菜单
    handleNavMoveDown (row, index) {
      const navType = row.parent_nav_id === 0 ? 'parent' : 'child'
      moveNavDown(navType, row.order_id)
        .then(res => {
          this.$message.success(res.msg)
          this.fetchTableData()
        })
        .catch(() => {
        })
    },
    // 删除菜单
    handleNavDelete (id) {
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(() => {
          deleteNav(id)
            .then(res => {
              this.$message.success(res.msg)
              this.fetchTableData()
            })
            .catch(() => {
            })
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
    // 批量删除菜单
    handleNavsMultiDel () {
      if (this.selectList.length === 0) {
        this.$message.error('请勾选要删除的条目')
        return
      }
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(() => {
          const ids = this.selectList.join(',')
          multiDelNavs(ids)
            .then(res => {
              this.$message.success(res.msg)
              this.fetchTableData()
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

<style lang="scss" scoped>
.operation-box {
  padding: 0;
}
</style>
