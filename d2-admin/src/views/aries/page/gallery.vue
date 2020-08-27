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
          <el-button size="small" type="danger">
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
        <h4 style="width: 50%;line-height: 50px;margin: auto;color: #909399;text-align: center;font-weight: normal">暂无数据</h4>
      </el-col>
      <el-col style="margin: 0 3.32% 2% 0;" :span="4" v-for="item in data" :key="item.ID">
        <div class="image-container">
          <span class="demonstration">{{ item.file_name }}</span>
          <!--        <el-checkbox style="position: absolute;" v-show="isMultiVisible" v-model="item.checked"-->
          <!--                     @change="checked=>handleCheckBoxChange(checked,item)"></el-checkbox>-->
          <el-image class="attach-image" :src="item.url" lazy/>
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

    <el-dialog
      title="添加图库"
      :visible.sync="dialogOptions.addVisible"
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
import { getGalleriesByPage } from '@api/aries/gallery'

export default {
  name: 'gallery',
  data () {
    return {
      loading: false,
      isMultiVisible: false,
      pagination: {
        page: 1,
        size: 20,
        key: '',
        category_id: null,
        total_num: 0,
        total_pages: 0
      },
      data: [],
      addForm: {},
      editForm: {},
      addFormRules: {},
      editFormRules: {},
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
  },
  methods: {
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
            this.data = res.data.data
          })
          .catch(() => {
          })
        this.loading = false
      }, 300)
    },
    // 搜索
    search () {
      this.pagination.page = 1
      this.fetchPageData()
    },
    // 重置
    reset () {
      this.pagination.page = 1
      this.category_id = null
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
    // 取消批量操作
    handleMultiCancel () {
      this.isMultiVisible = false
    },
    // 打开添加弹窗
    handleOpenAddDialog () {
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
</style>
