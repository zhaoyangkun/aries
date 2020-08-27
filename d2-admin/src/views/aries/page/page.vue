<template>
  <d2-container>
    <template slot="header">用户 / 页面</template>
    <el-tabs :tab-position="tabPosition" @tab-click="handleTabClick" type="border-card">
      <el-tab-pane label="独立页面">
        <span slot="label"><i class="el-icon-link"></i> 独立页面</span>
        <el-table
          :data="independentPageData"
          style="width: 100%">
          <el-table-column
            prop="title"
            label="页面名称"
            width="150"
          >
          </el-table-column>
          <el-table-column
            label="页面地址"
            width="450"
          >
            <template slot-scope="scope">
              <el-link target="_blank" :href="scope.row.url" :underline="false">{{ scope.row.url }}</el-link>
            </template>
          </el-table-column>
          <el-table-column
            label="操作">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="handleManagePage(scope.row.title)">管理</el-button>
            </template>
          </el-table-column>
        </el-table>

      </el-tab-pane>
      <el-tab-pane label="自定义页面">
        <span slot="label"><i class="fa fa-link"></i> 自定义页面</span>
      </el-tab-pane>
    </el-tabs>
  </d2-container>
</template>

<script>
import { getBlogVars } from '@api/aries/sys'

export default {
  name: 'page',
  data () {
    return {
      tabPosition: 'top',
      blogVars: {},
      independentPageData: [],
      customPageData: []
    }
  },
  created () {
    this.initIndependentPageData()
  },
  methods: {
    // 获取博客全局变量
    async fetchBlogVars () {
      await getBlogVars()
        .then(res => {
          this.blogVars = res.data
        })
        .catch(() => {
        })
    },
    // 获取独立页面数据
    async fetchIndependentPageData () {
      this.independentPageData = [
        {
          title: '日志页面',
          url: `${this.blogVars.ContextPath}/journals`
        },
        {
          title: '图库页面',
          url: `${this.blogVars.ContextPath}/galleries`
        }
      ]
    },
    // 初始化独立页面数据
    async initIndependentPageData () {
      await this.fetchBlogVars()
      await this.fetchIndependentPageData()
    },
    // 标签页切换
    handleTabClick (tab) {
      console.log(tab)
    },
    // 页面管理
    handleManagePage (title) {
      console.log('title:　', title)
      if (title === '日志页面') {
        this.$router.replace('/page/journal')
      } else if (title === '图库页面') {
        this.$router.replace('/page/gallery')
      }
    }
  }
}
</script>

<style scoped>

</style>
