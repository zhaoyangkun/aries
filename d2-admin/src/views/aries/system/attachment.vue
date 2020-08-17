<template>
  <d2-container>
    <template slot="header">系统 / 附件</template>
    <div class="operation-box">
      <el-form :inline="true" class="demo-form-inline" slot="header">
        <el-form-item>
          <el-button size="small" type="danger">
            <i class="el-icon-delete"></i> 批量删除
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="openUploadDialog">
            <i class="el-icon-plus"></i> 上传
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-row>
      <el-card style="padding-bottom: 10px">
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">abcdefghijklmnopqrstuvwxyz.jpg</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="block">
            <span class="demonstration">默认</span>
            <el-image :src="src" lazy></el-image>
          </div>
        </el-col>
      </el-card>
      <el-col :span="24">
        <div class="page-box">
          <el-pagination
            :page-sizes="[18, 36, 54, 72, 90]"
            :page-size="100"
            layout="total, sizes, prev, pager, next, jumper"
            :total="50">
          </el-pagination>
        </div>
      </el-col>
    </el-row>

    <el-dialog
      title="上传"
      :visible.sync="uploadDialogVisible"
      :with-header="false"
      width="50%"
    >
      <el-upload
        class="upload-demo"
        action=""
        :http-request="uploadImg"
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :file-list="fileList"
        list-type="picture">
        <el-button size="small" type="primary">点击上传</el-button>
        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
      </el-upload>
    </el-dialog>
  </d2-container>
</template>

<script>

import { uploadImgToPicBed } from '@api/aries/sys'

export default {
  name: 'attachment',
  data () {
    return {
      src: '/image/xue.jpg',
      fileList: [
        // {
        //   name: 'food.jpeg',
        //   url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        // },
        // {
        //   name: 'food2.jpeg',
        //   url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        // }
      ],
      uploadDialogVisible: false
    }
  },
  methods: {
    openUploadDialog () {
      this.uploadDialogVisible = true
    },
    handleRemove (file, fileList) {
      console.log(file, fileList)
    },
    handlePreview (file) {
      console.log(file)
    },
    uploadImg (file) {
      console.log('file: ', file)
      const formData = new FormData()
      formData.append('file[]', file.file)
      uploadImgToPicBed(formData)
        .then(res => {
          console.log(res)
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
  height: 28px;
  width: 180px;
  line-height: 28px;
  display: block;
  color: grey;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.page-box {
  margin: 5px auto;
}
</style>
