<template>
  <div>
    <!--选择附件抽屉-->
    <el-drawer
      title="选择附件"
      :visible.sync="drawerVisible"
      direction="rtl"
      style="overflow: auto"
    >
      <el-row style="padding: 0 8px">
        <el-row>
          <el-col :span="3" class="drawer-form-item">
            <el-tooltip content="上传" placement="top-start">
              <el-button size="small" type="success" @click="uploadDialogVisible=true"><i class="el-icon-upload2"></i>
              </el-button>
            </el-tooltip>
          </el-col>
          <el-col :span="12" class="drawer-form-item">
            <el-input size="small" v-model="pagination.key" placeholder="请输入关键词"></el-input>
          </el-col>
          <el-col :span="3" class="drawer-form-item">
            <el-tooltip content="搜索" placement="top-start">
              <el-button size="small" type="primary" @click="handleSearch"><i class="el-icon-search"></i></el-button>
            </el-tooltip>
          </el-col>
          <el-col :span="3" class="drawer-form-item">
            <el-tooltip content="重置" placement="top-start">
              <el-button size="small" @click="handleReset"><i class="el-icon-refresh"></i></el-button>
            </el-tooltip>
          </el-col>
        </el-row>
        <el-divider class="divider"></el-divider>

        <!--图片列表-->
        <el-row v-loading="loading">
          <el-col :span="24" v-if="data.length===0">
            <h4 class="empty-tip">暂无数据</h4>
          </el-col>
          <el-col :span="12" v-for="item in data" :key="item.ID">
            <div class="image-container">
              <el-image class="attach-image" :src="item.url" lazy @click="changeClickedImg(item.url)"></el-image>
            </div>
          </el-col>
        </el-row>
        <el-divider class="divider"></el-divider>

        <!--分页-->
        <el-col :span="24">
          <div class="page-box">
            <el-pagination
              background
              layout="prev, pager, next"
              :current-page="pagination.page"
              :page-size="pagination.size"
              :page-count="pagination.total_pages"
              @current-change="handleCurrentPageChange"
            >
            </el-pagination>
          </div>
        </el-col>
      </el-row>
    </el-drawer>

    <!--上传图片弹窗-->
    <el-dialog
      title="上传"
      :visible.sync="uploadDialogVisible"
      :with-header="false"
      @closed="handleCloseUploadDialog"
      width="50%"
    >
      <el-upload
        class="upload-demo"
        action=""
        accept=".jpeg,.jpg,.png,.gif,.bmp"
        :limit="limit"
        :multiple="true"
        :http-request="uploadImg"
        :before-upload="handleBeforeUpload"
        :on-exceed="handleOnExceed"
        :file-list="fileList"
        list-type="picture">
        <el-button size="small" type="primary">点击上传</el-button>
        <div slot="tip" class="el-upload__tip">只能上传 jpeg/jpg/png/gif/bmp 文件，且不超过 5 MB，一次最多上传 10 个。</div>
      </el-upload>
    </el-dialog>
  </div>
</template>

<script>
import { getImagesByPage, uploadImgToAttachment } from '@api/aries/picture'

export default {
  name: 'attachDrawer',
  data () {
    return {
      pagination: {
        page: 1,
        size: 10,
        key: '',
        total_pages: 1
      },
      data: [],
      fileList: [],
      imageTypes: ['jpeg', 'jpg', 'png', 'gif', 'bmp'],
      limit: 10,
      drawerVisible: false,
      uploadDialogVisible: false,
      loading: false
    }
  },
  created () {
    this.fetchPageData()
  },
  methods: {
    // 显示抽屉
    openDrawer () {
      this.drawerVisible = true
    },
    // 隐藏抽屉
    closeDrawer () {
      this.drawerVisible = false
    },
    // 获取分页数据
    fetchPageData () {
      this.loading = true
      setTimeout(() => {
        getImagesByPage({
          page: this.pagination.page,
          size: this.pagination.size,
          key: this.pagination.key
        })
          .then(res => {
            this.data = res.data.data
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
    // 改变当前页
    handleCurrentPageChange (page) {
      this.pagination.page = page
      this.fetchPageData()
    },
    // 关闭上传弹窗
    handleCloseUploadDialog () {
      this.fetchPageData()
      this.fileList = []
    },
    // 获取文件类型
    getFileType (fileName) {
      return fileName.substring(fileName.lastIndexOf('.') + 1)
    },
    // 上传前
    handleBeforeUpload (file) {
      const fileType = this.getFileType(file.name)
      const fileSize = file.size / 1024 / 1024
      if (this.imageTypes.indexOf(fileType) === -1) {
        this.$message.error('仅支持上传 jpeg, jpg, png, gif 和 bmp 格式的图片')
        return false
      }
      if (fileSize > 5) {
        this.$message.error('上传文件大小不能超过 5 MB')
        return false
      }
    },
    // 校验文件个数
    handleOnExceed (files, fileList) {
      if (fileList.length > this.limit) {
        this.$message.error(`最多允许上传 ${this.limit} 个文件`)
        return false
      }
    },
    // 上传图片
    uploadImg (file) {
      const formData = new FormData()
      formData.append('file[]', file.file)
      const config = {
        onUploadProgress: progressEvent => {
          const percent = progressEvent.loaded / progressEvent.total * 100 | 0
          file.onProgress({ percent: percent })
        }
      }
      uploadImgToAttachment(formData, config)
        .then(() => {
          file.onSuccess()
        })
        .catch(() => {
        })
    },
    // 改变父组件的值
    changeClickedImg (url) {
      this.$emit('changeClickedImg', url)
    }
  }
}
</script>

<style lang="scss">
.el-drawer__body {
  overflow: auto;
}

.el-drawer__container ::-webkit-scrollbar {
  display: none;
}
</style>

<style lang="scss" scoped>
.drawer-form-item {
  margin-right: 2%;
}

.image-container {
  height: 105px;
  padding: 0;
  margin: 0 1% 1% 0;
}

.attach-image {
  width: 100%;
  height: 105px;
  margin: 0;
  padding: 0;
  display: inline-block;
  overflow: hidden;
}

.page-box {
  margin-bottom: 10px;
}

.empty-tip {
  color: #99aabb;
  font-weight: normal;
  text-align: center;
}

.divider {
  margin: 12px 0;
}
</style>
