<template>
  <d2-container>
    <template slot="header">系统 / 附件</template>
    <div class="operation-box">
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item v-if="!isMultiVisible">
          <el-button size="small" @click="isMultiVisible=true">
            <i class="el-icon-edit-outline"></i> 批量操作
          </el-button>
        </el-form-item>
        <el-form-item v-if="isMultiVisible">
          <el-button size="small" type="danger">
            <i class="el-icon-delete"></i> 删除
          </el-button>
        </el-form-item>
        <el-form-item v-if="isMultiVisible">
          <el-button size="small" @click="isMultiVisible=false">
            <i class="el-icon-close-tip"></i> 取消
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="handleOpenUploadDialog">
            <i class="el-icon-plus"></i> 上传
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-select style="width: 150px" size="small" v-model="pagination.storage_name" clearable
                     placeholder="请选择图床类型">
            <el-option
              v-for="item in picBedTypes"
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
    </div>

    <el-row v-loading="loading">
      <el-col :span="24" v-if="data.length === 0" style="background-color: white;margin: 0 auto 15px auto">
        <h4 style="width: 50%;margin: 10px auto;color: #909399;text-align: center">暂无数据</h4>
      </el-col>
      <el-col :span="4" v-for="item in data" :key="item.ID"
              style="height: 120px;margin: 0 3.32% 2% 0;padding: 0;background-color: white;">
        <span class="demonstration">{{ item.file_name }}</span>
        <el-image @click="handleOpenImageDialog(item)" style="width: 100%;height: 90px;left: 0;top: 0" :src="item.url"
                  lazy></el-image>
      </el-col>

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
        :on-remove="handleRemove"
        :before-upload="handleBeforeUpload"
        :file-list="fileList"
        list-type="picture">
        <el-button size="small" type="primary">点击上传</el-button>
        <div slot="tip" class="el-upload__tip">只能上传 jpeg/jpg/png/gif/bmp 文件，且不超过 5 MB</div>
      </el-upload>
    </el-dialog>

    <!-- 图片预览弹窗 -->
    <el-dialog
      title="图片预览"
      :visible.sync="previewDialogVisible"
      :with-header="false"
      width="50%">
      <el-image style="width: 100%" :src="previewData.url" fit="cover" @click="openImgUrl(previewData.url)"></el-image>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">图床类型：</h4></el-col>
        <el-col :span="19">
          <h4 class="pre-tip" v-if="previewData.storage_type==='sm.ms'">sm.ms</h4>
          <h4 class="pre-tip" v-if="previewData.storage_type==='imgbb'">imgbb</h4>
          <h4 class="pre-tip" v-if="previewData.storage_type==='cos'">腾讯云 COS</h4>
        </el-col>
      </el-row>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">图片名称：</h4></el-col>
        <el-col :span="19"><h4 class="pre-tip">{{ previewData.file_name }}</h4></el-col>
      </el-row>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">图片大小：</h4></el-col>
        <el-col :span="19"><h4 class="pre-tip">{{ previewData.size }} KB</h4></el-col>
      </el-row>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">图片地址：</h4></el-col>
        <el-col :span="19">
          <h4 class="pre-tip">
            {{ previewData.url }}
            <el-tooltip content="复制" placement="top" effect="light">
              <el-button
                size="mini"
                class="el-icon-document-copy"
                v-clipboard:copy="previewData.url"
                v-clipboard:success="onCopySuccess"
                v-clipboard:error="onCopyError"
              ></el-button>
            </el-tooltip>
          </h4>
        </el-col>
      </el-row>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">Markdown 链接：</h4></el-col>
        <el-col :span="19">
          <h4 class="pre-tip">
            ![image]({{ previewData.url }})
            <el-tooltip content="复制" placement="top" effect="light">
              <el-button
                size="mini"
                class="el-icon-document-copy"
                v-clipboard:copy="`![image](${previewData.url})`"
                v-clipboard:success="onCopySuccess"
                v-clipboard:error="onCopyError"
              ></el-button>
            </el-tooltip>
          </h4>
        </el-col>
      </el-row>
    </el-dialog>

  </d2-container>
</template>

<script>

import { getImagesByPage, uploadImgToAttachment } from '@api/aries/picture'

export default {
  name: 'attachment',
  created () {
    this.fetchPageData()
  },
  data () {
    return {
      picBedTypes: [
        { value: 'sm.ms', name: 'sm.ms' },
        { value: 'imgbb', name: 'imgbb' },
        { value: 'cos', name: '腾讯云' }
      ],
      pagination: {
        page: 1,
        size: 20,
        key: '',
        storage_name: null,
        total_num: 0,
        total_pages: 0
      },
      data: [],
      fileList: [],
      previewData: {
        url: '',
        file_name: '',
        size: 0,
        storage_type: '',
        createAt: null
      },
      imageTypes: ['jpeg', 'jpg', 'png', 'gif', 'bmp'],
      limit: 5,
      progressPercent: 0,
      uploadDialogVisible: false,
      previewDialogVisible: false,
      isMultiVisible: false,
      loading: false
    }
  },
  methods: {
    // 获取分页数据
    fetchPageData () {
      this.loading = true
      setTimeout(() => {
        getImagesByPage({
          page: this.pagination.page,
          size: this.pagination.size,
          key: this.pagination.key,
          storage_name: this.pagination.storage_name
        })
          .then(res => {
            this.data = res.data.data
            this.pagination.total_num = res.data.total_num
            this.pagination.total_pages = res.data.total_pages
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
      this.pagination.size = 18
      this.pagination.key = ''
      this.pagination.storage_name = null
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
    // 打开上传弹窗
    handleOpenUploadDialog () {
      this.uploadDialogVisible = true
    },
    handleCloseUploadDialog () {
      this.fetchPageData()
      this.fileList = []
    },
    // 图片删除
    handleRemove (file, fileList) {
      console.log(file, fileList)
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
    uploadImg: function (file) {
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
    // 图片预览
    handleOpenImageDialog (file) {
      this.previewData = file
      this.previewDialogVisible = true
    },
    // 获取文件类型
    getFileType (fileName) {
      return fileName.substring(fileName.lastIndexOf('.') + 1)
    },
    // 打开图片链接
    openImgUrl (url) {
      window.open(url)
    },
    // 复制内容到剪贴板
    doCopy (val) {
      this.$copyText(val).then(e => {
      }, e => {
      })
    },
    // 复制成功
    onCopySuccess () {
      const h = this.$createElement
      this.$notify({
        title: ' 提示',
        type: 'success',
        message: h('i', { style: 'color: teal' }, '已复制到剪贴板')
      })
    },
    // 复制失败
    onCopyError () {
      const h = this.$createElement
      this.$notify({
        title: '提示',
        type: 'error',
        message: h('i', { style: 'color: teal' }, '复制失败')
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
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.page-box {
  margin: auto;
}

.pre-tip {
  margin: 12px 0;
  height: 20px;
  line-height: 20px;
  font-weight: normal;
  color: grey;
}
</style>
