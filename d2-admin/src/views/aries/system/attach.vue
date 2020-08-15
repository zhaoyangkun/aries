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
          <el-button size="small" type="primary">
            <i class="el-icon-plus"></i> 上传
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-row>
      <el-col :span="6">
        <el-card>
          <span class="demonstration">图片</span>
          <el-image :src="src" lazy></el-image>
        </el-card>
      </el-col>
    </el-row>
  </d2-container>
</template>

<script>
export default {
  name: 'attach',
  data () {
    return {
      src: 'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg'
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
</style>
