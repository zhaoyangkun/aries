<template>
  <d2-container>
    <template slot="header">用户 / 页面 / 图库</template>
    <div class="operation-box">
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item v-show="!isMultiVisible">
          <el-button size="small" @click="isMultiVisible=true">
            <i class="el-icon-edit-outline"></i> 批量操作
          </el-button>
        </el-form-item>
        <el-form-item v-show="isMultiVisible">
          <el-button size="small" type="danger" @click="handleMultiDelGalleries">
            <i class="el-icon-delete"></i> 删除
          </el-button>
        </el-form-item>
        <el-form-item v-show="isMultiVisible">
          <el-button size="small" @click="handleMultiCancel">
            <i class="el-icon-close-tip"></i> 取消
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="handleOpenAddDialog">
            <i class="el-icon-plus"></i> 添加
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

    <!--图片列表-->
    <el-row v-loading="loading">
      <el-col :span="24" v-if="data.length === 0" style="margin: 0 auto 15px auto;height: 50px;background-color: white">
        <h4 style="width: 50%;line-height: 50px;margin: auto;color: #909399;text-align: center;font-weight: normal">
          暂无数据
        </h4>
      </el-col>
      <el-col style="margin: 0 3.32% 2% 0;" :span="4" v-for="item in data" :key="item.ID">
        <div class="image-container" :class="imgIsChecked(item.checked)">
          <span class="demonstration">{{ item.name }}</span>
          <el-checkbox style="position: absolute;" v-show="isMultiVisible" v-model="item.checked"
                       @change="checked=>handleCheckBoxChange(checked,item)"></el-checkbox>
          <el-image class="attach-image" :src="item.url" lazy @click="handleClickGallery(item)"/>
        </div>
      </el-col>

      <!--分页-->
      <el-col :span="24">
        <div class="page-box">
          <el-pagination
            background
            :page-sizes="[20, 50, 80, 110]"
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
      </el-col>
    </el-row>

    <!--添加图库-->
    <el-dialog
      title="添加图库"
      :visible.sync="dialogOptions.addVisible"
      :with-header="false"
      width="50%"
    >
      <el-form ref="addForm" :model="addForm" :rules="addFormRules" label-width="100px">
        <el-form-item label="图片预览">
          <h4 class="img-tip">单击打开附件</h4>
          <img alt="" class="img-show" :src="addImageSrc" @click="handleOpenDrawer('add')"/>
        </el-form-item>
        <el-form-item label="图片 URL" prop="url">
          <el-input placeholder="请输入图片 URL" v-model="addForm.url" clearable></el-input>
        </el-form-item>
        <el-form-item label="图片分类" prop="category_id">
          <el-select
            v-model="addForm.category_id"
            filterable
            default-first-option
            placeholder="请选择图片分类"
          >
            <el-option
              v-for="item in categories"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-tag
            :key="item.ID"
            v-for="item in categories"
            closable
            :disable-transitions="false"
            @click="handleShowEditInput(item.ID,item.name)"
            @close="handleDelCategory(item.ID,'addForm')"
          >
            {{ item.name }}
          </el-tag>
          <el-input
            class="input-new-tag"
            v-if="inputVisible"
            v-model="inputValue"
            ref="saveTagInput"
            size="small"
            @keyup.enter.native="handleInputConfirm"
            @blur="handleInputConfirm"
          >
          </el-input>
          <el-button v-else class="button-new-tag" size="small" @click="handleShowAddInput">
            <i class="el-icon-plus"></i> 新增分类
          </el-button>
        </el-form-item>
        <el-form-item label="图片名称" prop="name">
          <el-input placeholder="请输入图片名称" v-model="addForm.name" clearable></el-input>
        </el-form-item>
        <el-form-item label="拍摄地点" prop="location">
          <el-input placeholder="请输入拍摄地点" v-model="addForm.location" clearable></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.addBtnLoading" @click="handleAddGallery">保存</el-button>
          <el-button @click="dialogOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!--编辑图库-->
    <el-dialog
      title="编辑图库"
      :visible.sync="dialogOptions.editVisible"
      :with-header="false"
      width="50%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editFormRules" label-width="100px">
        <el-form-item label="图片预览">
          <h4 class="img-tip">单击打开附件</h4>
          <img alt="" class="img-show" :src="editImageSrc" @click="handleOpenDrawer('edit')"/>
        </el-form-item>
        <el-form-item label="图片 URL" prop="url">
          <el-input placeholder="请输入图片 URL" v-model="editForm.url" clearable></el-input>
        </el-form-item>
        <el-form-item label="图片分类" prop="category_id">
          <el-select
            v-model="editForm.category_id"
            filterable
            default-first-option
            placeholder="请选择图片分类"
          >
            <el-option
              v-for="item in categories"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-tag
            :key="item.ID"
            v-for="item in categories"
            closable
            :disable-transitions="false"
            @click="handleShowEditInput(item.ID,item.name)"
            @close="handleDelCategory(item.ID,'addForm')"
          >
            {{ item.name }}
          </el-tag>
          <el-input
            class="input-new-tag"
            v-if="inputVisible"
            v-model="inputValue"
            ref="saveTagInput"
            size="small"
            @keyup.enter.native="handleInputConfirm"
            @blur="handleInputConfirm"
          >
          </el-input>
          <el-button v-else class="button-new-tag" size="small" @click="handleShowAddInput">
            <i class="el-icon-plus"></i> 新增分类
          </el-button>
        </el-form-item>
        <el-form-item label="图片名称" prop="name">
          <el-input placeholder="请输入图片名称" v-model="editForm.name" clearable></el-input>
        </el-form-item>
        <el-form-item label="拍摄地点" prop="location">
          <el-input placeholder="请输入拍摄地点" v-model="editForm.location" clearable></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="dialogOptions.editBtnLoading" @click="handleEditGallery">保存</el-button>
          <el-button @click="dialogOptions.editVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!--附件抽屉-->
    <attachDrawer ref="attachDrawer" @changeClickedImg="changeClickedImg($event)"></attachDrawer>
  </d2-container>
</template>

<script>
import attachDrawer from '@/components/aries/common/attachDrawer'
import { createGallery, getGalleriesByPage, multiDelGalleries, updateGallery } from '@api/aries/gallery'
import { addGalleryCategory, deleteCategory, getAllCategories, updateGalleryCategory } from '@api/aries/category'

export default {
  name: 'gallery',
  components: {
    attachDrawer
  },
  data () {
    return {
      pagination: {
        page: 1,
        size: 20,
        key: '',
        category_id: null,
        total_num: 0,
        total_pages: 0
      },
      categoryIds: [],
      categories: [],
      selectGalleries: [],
      data: [],
      addImageSrc: '/image/none.jpg',
      editImageSrc: '/image/none.jpg',
      clickImg: '',
      mode: 'add',
      addForm: {
        category_id: null,
        url: '',
        name: '',
        desc: '',
        location: ''
      },
      editForm: {
        ID: null,
        category_id: null,
        url: '',
        name: '',
        desc: '',
        location: ''
      },
      addFormRules: {
        url: [
          { required: true, trigger: 'blur', message: '请输入图片 URL' },
          { max: 255, trigger: 'blur', message: '图片 URL 不能超过 255 个字符' }
        ],
        name: [
          { required: true, trigger: 'blur', message: '请输入图片名称' },
          { max: 255, trigger: 'blur', message: '图片名称不能超过 255 个字符' }
        ],
        desc: [
          { max: 255, trigger: 'blur', message: '图片描述不能超过 255 个字符' }
        ],
        location: [
          { max: 50, trigger: 'blur', message: '拍摄地点不能超过 50 个字符' }
        ]
      },
      editFormRules: {
        url: [
          { required: true, trigger: 'blur', message: '请输入图片 URL' },
          { max: 255, trigger: 'blur', message: '图片 URL 不能超过 255 个字符' }
        ],
        name: [
          { required: true, trigger: 'blur', message: '请输入图片名称' },
          { max: 255, trigger: 'blur', message: '图片名称不能超过 255 个字符' }
        ],
        desc: [
          { max: 255, trigger: 'blur', message: '图片描述不能超过 255 个字符' }
        ],
        location: [
          { max: 50, trigger: 'blur', message: '拍摄地点不能超过 50 个字符' }
        ]
      },
      inputVisible: false,
      inputValue: '',
      inputId: null,
      inputChange: false,
      loading: false,
      isMultiVisible: false,
      drawerVisible: false,
      dialogOptions: {
        addVisible: false,
        addBtnLoading: false,
        editVisible: false,
        editBtnLoading: false
      }
    }
  },
  created () {
    this.fetchPageData()
    this.fetchCategoryData()
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
        getGalleriesByPage({
          page: this.pagination.page,
          size: this.pagination.size,
          category_id: this.pagination.category_id,
          key: this.pagination.key
        })
          .then(res => {
            res.data.data.map(item => {
              this.$set(item, 'checked', false)
            })
            this.data = res.data.data
            this.pagination.total_num = res.data.total_num
            this.pagination.total_pages = res.data.total_pages
          })
          .catch(() => {
          })
        this.loading = false
      }, 300)
    },
    // 获取图片分类
    fetchCategoryData () {
      getAllCategories(2)
        .then(res => {
          this.categories = res.data
          if (this.categories.length > 0) {
            this.categories.forEach(category => {
              this.categoryIds.push(category.ID)
            })
          }
        })
        .catch(() => {
        })
    },
    // 搜索
    search () {
      this.pagination.page = 1
      this.fetchPageData()
    },
    // 重置
    reset () {
      this.pagination.page = 1
      this.pagination.category_id = null
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
    // 打开添加弹窗
    handleOpenAddDialog () {
      this.clearValidate('addForm')
      this.dialogOptions.addVisible = true
    },
    // 添加或修改分类
    handleInputConfirm () {
      if (this.inputValue) {
        if (this.inputChange) {
          updateGalleryCategory({
            id: this.inputId,
            type: 2,
            name: this.inputValue
          })
            .then(() => {
              this.fetchCategoryData()
            })
            .catch(() => {
            })
        } else {
          addGalleryCategory({
            type: 2,
            name: this.inputValue
          })
            .then(() => {
              this.fetchCategoryData()
            })
            .catch(() => {
            })
        }
      }
      this.inputVisible = false
      this.inputChange = false
      this.inputValue = ''
      this.inputId = null
    },
    // 显示添加文本框
    handleShowAddInput () {
      this.inputVisible = true
      this.inputChange = false
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
    },
    // 显示修改文本框
    handleShowEditInput (id, name) {
      this.inputVisible = true
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
      this.inputValue = name
      this.inputId = id
      this.inputChange = true
    },
    // 删除分类
    handleDelCategory (id, form) {
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(_ => {
          deleteCategory(id)
            .then(() => {
              this.fetchCategoryData()
              // 若删除项被选中，清空选中项
              if (this[form].category_id === id) {
                this[form].category_id = null
              }
            })
            .catch(() => {
            })
        })
        .catch(_ => {
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
        this.addForm.url = url
        this.addImageSrc = url
      } else {
        this.editForm.url = url
        this.editImageSrc = url
      }
      this.$refs.attachDrawer.closeDrawer()
    },
    // 添加图库
    handleAddGallery () {
      this.$refs.addForm.validate(valid => {
        if (valid) {
          this.dialogOptions.addBtnLoading = true
          setTimeout(_ => {
            createGallery(this.addForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.addVisible = false
                this.resetForm('addForm')
                this.addImageSrc = '/image/none.jpg'
                this.fetchPageData()
              })
              .catch(_ => {
              })
            this.dialogOptions.addBtnLoading = false
          }, 300)
        }
      })
    },
    // 打开修改图库弹窗
    handleClickGallery (item) {
      if (this.isMultiVisible) {
        if (item.checked) {
          const i = this.selectGalleries.indexOf(item.ID)
          if (i >= 0) {
            this.selectGalleries.splice(i, 1)
          }
        } else {
          this.selectGalleries.push(item.ID)
        }
        item.checked = !item.checked
      } else {
        this.editForm = { ...item }
        this.editImageSrc = this.editForm.url
        this.dialogOptions.editVisible = true
      }
    },
    // 修改图库
    handleEditGallery () {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          this.dialogOptions.editBtnLoading = true
          setTimeout(_ => {
            updateGallery(this.editForm)
              .then(res => {
                this.$message.success(res.msg)
                this.dialogOptions.editVisible = false
                this.fetchPageData()
              })
              .catch(_ => {
              })
            this.dialogOptions.editBtnLoading = false
          }, 300)
        }
      })
    },
    // 图片 class 绑定
    imgIsChecked (checked) {
      if (checked) {
        return 'img_checked'
      }
      return 'img_not_checked'
    },
    // 勾选图片
    handleCheckBoxChange (checked, item) {
      if (checked) {
        this.selectGalleries.push(item.ID)
      } else {
        const i = this.selectGalleries.indexOf(item.ID)
        if (i >= 0) {
          this.selectGalleries.splice(i, 1)
        }
      }
    },
    // 取消批量操作
    handleMultiCancel () {
      this.isMultiVisible = false
      this.selectGalleries = []
      const data = this.data
      data.map(item => {
        item.checked = false
        return item
      })
      this.data = data
    },
    // 批量删除图片
    handleMultiDelGalleries () {
      if (this.selectGalleries.length === 0) {
        this.$message.error('请选择要删除的图片')
        return
      }
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
        .then(() => {
          const ids = this.selectGalleries.join(',')
          multiDelGalleries(ids)
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

<style lang="scss" scoped>
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
