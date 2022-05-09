<template>
  <d2-container>
    <el-row>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章</span>
            <el-button size="medium" icon="el-icon-notebook-2" style="float: right; padding: 3px 0"
                       type="text" @click="$router.push('/post')"></el-button>
          </div>
          <div class="text item"><h2 style="color: #2f74ff" class="card-num">{{ indexData.article_count }}</h2></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>评论</span>
            <el-button size="medium" icon="el-icon-s-comment" style="float: right; padding: 3px 0"
                       type="text" @click="$router.push('/comment')"></el-button>
          </div>
          <div class="text item"><h2 style="color: #4dc820" class="card-num">{{ indexData.comment_count }}</h2></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>访问量</span>
            <el-button size="medium" icon="el-icon-view" style="float: right; padding: 3px 0" type="text"></el-button>
          </div>
          <div class="text item"><h2 style="color: slategrey" class="card-num">0</h2></div>
        </el-card>
      </el-col>
      <el-col :span="24">
        <el-tabs class="box-card" type="border-card">
          <el-tab-pane label="最近文章">
            <el-timeline>
              <p class="no-tip" v-if="indexData.latest_articles.length === 0">暂无文章</p>
              <el-timeline-item :key="item.ID" v-for="item in indexData.latest_articles"
                                :timestamp="formatDate(item.CreatedAt)" placement="top">
                {{ item.title }}
              </el-timeline-item>
            </el-timeline>
          </el-tab-pane>
          <el-tab-pane label="最近评论">
            <el-timeline>
              <p class="no-tip" v-if="indexData.latest_comments.length === 0">暂无评论</p>
              <el-timeline-item :key="item.ID" v-for="item in indexData.latest_comments"
                                :timestamp="formatDate(item.CreatedAt)" placement="top">
                {{ item.content }}
              </el-timeline-item>
            </el-timeline>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
  </d2-container>
</template>

<script>
import { getAdminIndexData } from '@api/aries/sys'
import { dateFormat } from '@/plugin/time/date'

export default {
  name: 'index',
  data () {
    return {
      indexData: {
        article_count: 0,
        comment_count: 0,
        latest_articles: [],
        latest_comments: []
      },
      fmt: 'yyyy-MM-dd',
      currentDate: new Date()
    }
  },
  created () {
    this.fetchAdminIndexData()
  },
  methods: {
    fetchAdminIndexData () {
      getAdminIndexData()
        .then(res => {
          this.indexData = res.data
        })
        .catch(() => {
        })
    },
    formatDate (time) {
      return dateFormat(this.fmt, new Date(time))
    }
  }
}
</script>

<style scoped lang="scss">
.box-card {
  margin: 10px;
}

.card-num {
  margin: 5px;
}

.no-tip {
  color: #99aabb;
}
</style>
