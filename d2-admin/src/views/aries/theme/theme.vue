<template>
  <d2-container>
    <template slot="header">外观 / 主题</template>
    <el-row v-loading="loading">
      <el-col class="col-none-box" :span="24" v-if="data.length === 0">
        <h4 class="none-tip">暂无数据</h4>
      </el-col>
      <el-col class="col-box" :span="5" v-for="item in data" :key="item.ID">
        <div class="image-container img_not_checked">
          <span class="demonstration">{{ item.theme_name }}</span>
          <el-image class="attach-image" :src="item.theme_image" lazy/>
          <div class="theme-operation">
            <el-button class="theme-op-btn">
              <i class="el-icon-check theme-op-btn-i"></i>启用
            </el-button>
            <el-button class="theme-op-btn">
              <i class="el-icon-setting theme-op-btn-i"></i>设置
            </el-button>
            <el-button class="theme-op-btn">
              <i class="el-icon-delete theme-op-btn-i"></i>删除
            </el-button>
          </div>
        </div>
      </el-col>
    </el-row>
  </d2-container>
</template>

<script>
export default {
  name: 'theme',
  data () {
    return {
      data: [
        {
          theme_name: 'xue',
          author: 'zyk',
          theme_image: 'https://i.loli.net/2020/09/03/Vl3aoG58uqNrsPL.png',
          is_used: false
        },
        {
          theme_name: 'xue',
          author: 'zyk',
          theme_image: 'https://i.loli.net/2020/09/03/Vl3aoG58uqNrsPL.png',
          is_used: false
        },
        {
          theme_name: 'xue',
          author: 'zyk',
          theme_image: 'https://i.loli.net/2020/09/03/Vl3aoG58uqNrsPL.png',
          is_used: false
        },
        {
          theme_name: 'xue',
          author: 'zyk',
          theme_image: 'https://i.loli.net/2020/09/03/Vl3aoG58uqNrsPL.png',
          is_used: false
        },
        {
          theme_name: 'xue',
          author: 'zyk',
          theme_image: 'https://i.loli.net/2020/09/03/Vl3aoG58uqNrsPL.png',
          is_used: false
        }
      ],
      loading: false
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

.img_not_checked {
  border: 2px rgba(200, 200, 200, 0.18) solid;
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
