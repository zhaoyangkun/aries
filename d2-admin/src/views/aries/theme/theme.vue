<template>
  <d2-container>
    <template slot="header">外观 / 主题</template>
    <el-row v-loading="loading">
      <el-col class="col-none-box" :span="24" v-if="theme_list.length === 0">
        <h4 class="none-tip">暂无数据</h4>
      </el-col>
      <el-col class="col-box" :span="5" v-for="item in theme_list" :key="item.ID">
        <div class="image-container" :class="themeIsEnabled(item.is_used)">
          <span class="demonstration">{{ item.theme_name }}</span>
          <el-image class="attach-image" :src="item.image" @click="showPreviewDiag(item.theme_name)" lazy/>
          <div class="theme-operation">
            <el-button class="theme-op-btn" style="width: 100%" v-if="!item.is_used"
                       @click="enableTheme(item.theme_name)">
              <i class="el-icon-check theme-op-btn-i"></i>启用
            </el-button>
            <el-button class="theme-op-btn" style="width: 100%;color: #409EFF;" v-else>
              <i class="el-icon-check theme-op-btn-i"></i>已启用
            </el-button>
            <!--            <el-button class="theme-op-btn">-->
            <!--              <i class="el-icon-setting theme-op-btn-i"></i>设置-->
            <!--            </el-button>-->
            <!--            <el-button class="theme-op-btn">-->
            <!--              <i class="el-icon-delete theme-op-btn-i"></i>删除-->
            <!--            </el-button>-->
          </div>
        </div>
      </el-col>
    </el-row>

    <!-- 主题详情弹窗 -->
    <el-dialog
      title="主题详情"
      :visible.sync="previewDialogVisible"
      :with-header="false"
      width="50%">
      <el-image style="width: 100%" :src="previewData.image" fit="cover"
                @click="openImgUrl(previewData.repo)"></el-image>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">主题名称：</h4></el-col>
        <el-col :span="19">
          <h4 class="pre-tip">{{ previewData.theme_name }}</h4>
        </el-col>
      </el-row>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">作者：</h4></el-col>
        <el-col :span="19">
          <h4 class="pre-tip">{{ previewData.author_name }}</h4>
        </el-col>
      </el-row>
      <el-row style="border-bottom: 1px solid #eaeefb">
        <el-col :span="5"><h4 class="pre-tip">仓库地址：</h4></el-col>
        <el-col :span="19">
          <h4 class="pre-tip">
            {{ previewData.repo }}
            <!--            <el-tooltip content="复制" placement="top" effect="light">-->
            <!--              <el-button-->
            <!--                size="mini"-->
            <!--                class="el-icon-document-copy"-->
            <!--                v-clipboard:copy="previewData.repo"-->
            <!--                v-clipboard:success="onCopySuccess"-->
            <!--                v-clipboard:error="onCopyError"-->
            <!--              ></el-button>-->
            <!--            </el-tooltip>-->
          </h4>
        </el-col>
      </el-row>
    </el-dialog>
  </d2-container>
</template>

<script>
import { enableTheme, getAllThemes, getThemeByName } from '@/api/aries/theme'

export default {
  name: 'theme',
  data () {
    return {
      theme_list: [],
      previewDialogVisible: false,
      loading: false,
      previewData: {
        theme_name: '',
        image: '',
        repo: ''
      }
    }
  },
  created () {
    this.getThemeData()
  },
  methods: {
    // class 绑定
    themeIsEnabled (isUsed) {
      if (isUsed) {
        return 'theme_enabled'
      }
      return 'theme_not_enabled'
    },
    // 获取主题数据
    getThemeData () {
      this.loading = true
      setTimeout(() => {
        getAllThemes()
          .then(res => {
            // console.log('res: ', res)
            this.theme_list = res.data
          })
          .catch(() => {
          })
        this.loading = false
      }, 300)
    },
    // 启用主题
    enableTheme (name) {
      enableTheme({
        theme_name: name
      })
        .then(res => {
          this.$message.success(res.msg)
          this.getThemeData()
        })
        .catch(() => {
        })
    },
    // 打开图片链接
    openImgUrl (url) {
      window.open(url)
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
    },
    // 显示主题详情窗口
    showPreviewDiag (name) {
      getThemeByName(name)
        .then(res => {
          this.previewData = res.data
          this.previewDialogVisible = true
        }).catch(() => {
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.col-none-box {
  margin: 0 auto 15px auto;
  height: 50px;
  background-color: white;
}

.col-box {
  margin: 0 4.16% 2% 0;
}

.image-container {
  height: 179px;
  padding: 0;
  margin: 0;
}

.attach-image {
  width: 100%;
  height: 115px;
  margin: 0;
  padding: 0;
  display: block;
  overflow: hidden;
}

.demonstration {
  padding-left: 5%;
  height: 30px;
  width: 95%;
  line-height: 30px;
  font-size: 15px;
  display: block;
  color: grey;
  background-color: white;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.none-tip {
  width: 50%;
  line-height: 50px;
  margin: auto;
  color: #909399;
  text-align: center;
  font-weight: normal;
}

.theme_not_enabled {
  border: 2px rgba(200, 200, 200, 0.18) solid;
}

.theme_enabled {
  border: #409EFF solid 2px;
}

.theme-operation {
  width: 96%;
  height: 30px;
  line-height: 30px;
  margin: 0;
  padding: 2px 2%;
  background-color: white;
}

.theme-op-btn {
  margin: 0;
  padding: 0;
  height: 30px;
  width: 33%;
  line-height: 30px;
  border: none;
  background-color: white;
}

.theme-op-btn-i {
  height: 30px;
  margin: 0;
  padding: 0;
  display: inherit;
}
</style>
