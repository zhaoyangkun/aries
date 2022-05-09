<template>
  <div class="page-login">
    <div class="page-login--layer page-login--layer-area">
      <ul class="circles">
        <li v-for="n in 10" :key="n"></li>
      </ul>
    </div>
    <div
      class="page-login--layer page-login--layer-time"
      flex="main:center cross:center">
      {{ time }}
    </div>
    <div class="page-login--layer">
      <div
        class="page-login--content"
        flex="dir:top main:justify cross:stretch box:justify">
        <div class="page-login--content-header">
          <p class="page-login--content-header-motto">
          </p>
        </div>
        <div
          class="page-login--content-main"
          flex="dir:top main:center cross:center">
          <!-- logo -->
          <!--          <img alt="logo" class="page-login&#45;&#45;logo" :src="logoPath">-->
          <!-- form -->
          <div class="page-login--form">
            <el-card style="border-radius: 8px;margin-bottom: 5px">
              <h2 class="tip">Aries</h2>
              <el-form
                ref="regForm"
                label-position="top"
                :rules="rules"
                :model="regForm"
                size="default">
                <el-form-item prop="username">
                  <el-input
                    type="text"
                    v-model="regForm.username"
                    placeholder="用户名">
                    <i slot="prepend" class="fa fa-user-circle-o"></i>
                  </el-input>
                </el-form-item>
                <el-form-item prop="email">
                  <el-input
                    type="email"
                    v-model="regForm.email"
                    placeholder="邮箱">
                    <i slot="prepend" class="fa fa-envelope-o"></i>
                  </el-input>
                </el-form-item>
                <el-form-item prop="pwd">
                  <el-input
                    type="password"
                    v-model="regForm.pwd"
                    placeholder="密码">
                    <i slot="prepend" class="fa fa-key"></i>
                  </el-input>
                </el-form-item>
                <el-form-item prop="second_pwd">
                  <el-input
                    type="password"
                    v-model="regForm.second_pwd"
                    placeholder="确认密码">
                    <i slot="prepend" class="fa fa-key"></i>
                  </el-input>
                </el-form-item>
                <el-form-item prop="site_name">
                  <el-input
                    type="text"
                    v-model="regForm.site_name"
                    placeholder="网站名称">
                    <i slot="prepend" class="fa fa-location-arrow"></i>
                  </el-input>
                </el-form-item>
                <el-form-item prop="theme_name">
                  <i slot="prepend" class="fa fa-github" aria-hidden="true"></i>
                  <el-select v-model="regForm.theme_name" clearable placeholder="选择主题">
                    <el-option
                      v-for="item in themeList"
                      :key="item.ID"
                      :label="item.theme_name"
                      :value="item.theme_name">
                    </el-option>
                  </el-select>
                </el-form-item>
                <el-form-item prop="site_url">
                  <el-input
                    type="text"
                    v-model="regForm.site_url"
                    placeholder="网站地址">
                    <i slot="prepend" class="fa fa-chain"></i>
                  </el-input>
                </el-form-item>
                <el-button
                  size="default" :loading="btnLoading" @click="submit" type="primary"
                  class="button-login">
                  保存配置
                </el-button>
              </el-form>
            </el-card>
            <p
              class="page-login--options"
              flex="main:justify cross:center">
              <span @click="toLogin"><d2-icon name="user"/> 登录</span>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs'
import { authRegister } from '@/api/aries/auth'
import { getAllUsers } from '@/api/aries/user'
import { getAllThemes } from '@/api/aries/theme'

export default {
  name: 'register',
  data () {
    // 自定义校验函数
    const validatePwd = (rule, value, callback) => {
      if (this.regForm.pwd !== value) {
        callback(new Error('两次密码不一致！'))
      } else {
        callback()
      }
    }
    return {
      timeInterval: null,
      time: dayjs().format('HH:mm:ss'),
      labelPosition: 'left',
      btnLoading: false,
      userList: [],
      themeList: [],
      // 表单
      regForm: {
        username: '',
        email: '',
        pwd: '',
        second_pwd: '',
        site_name: '',
        site_url: '',
        theme_name: ''
      },
      // 表单校验
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 30, message: '用户名长度需在 3~30 之间', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '邮箱格式错误', trigger: 'blur' }
        ],
        pwd: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码长度需在 6~20 之间', trigger: 'blur' }
        ],
        second_pwd: [
          { required: true, message: '请输入确认密码', trigger: 'blur' },
          { min: 6, max: 20, message: '确认密码长度需在 6~20 之间', trigger: 'blur' },
          { validator: validatePwd, trigger: 'blur' }
        ],
        site_name: [
          { required: true, message: '请输入网站名称', trigger: 'blur' },
          { max: 50, message: '网站名称长度不能超过 50', trigger: 'blur' }
        ],
        site_url: [
          { required: true, message: '请输入网站地址', trigger: 'blur' },
          { max: 255, message: '网站地址长度不能超过 255', trigger: 'blur' },
          { type: 'url', message: '请输入正确的网站地址', trigger: 'blur' }
        ],
        theme_name: [
          { required: true, message: '请选择主题', trigger: 'blur' }
        ]
      }
    }
  },
  created () {
    this.checkFirst()
    this.getThemes()
  },
  mounted () {
    this.timeInterval = setInterval(() => {
      this.refreshTime()
    }, 1000)
  },
  beforeDestroy () {
    clearInterval(this.timeInterval)
  },
  methods: {
    refreshTime () {
      this.time = dayjs().format('HH:mm:ss')
    },
    // 获取主题
    getThemes () {
      getAllThemes()
        .then(res => {
          this.themeList = res.data
        })
    },
    // 提交表单
    submit () {
      this.$refs.regForm.validate((valid) => {
        if (valid) {
          this.btnLoading = true
          setTimeout(() => {
            // 注册
            authRegister(this.$data.regForm)
              .then(res => {
                this.$message.success(res.msg)
                this.toLogin()
              })
              .catch(() => {
              })
            this.btnLoading = false
          }, 300)
        }
      })
    },
    // 校验是否为第一次注册
    checkFirst () {
      getAllUsers()
        .then(res => {
          this.userList = res.data
          if (this.userList.length > 0) {
            this.$router.push('/login')
          } else {
            const h = this.$createElement
            this.$notify({
              title: '提示',
              message: h(
                'i',
                { style: 'color: #1790fe' },
                'Aries 博客初始化成功，请先配置博客参数，再登录'
              ),
              type: 'success'
            })
          }
        })
    },
    // 跳转到登录页面
    toLogin () {
      this.$router.push('/login')
    }
  }
}
</script>

<style lang="scss">
.page-login {
  @extend %unable-select;
  $backgroundColor: #F0F2F5;
  // ---
  background-color: $backgroundColor;
  height: 100%;
  position: relative;
  // 层
  .page-login--layer {
    @extend %full;
    overflow: auto;
  }

  .page-login--layer-area {
    overflow: hidden;
  }

  // 时间
  .page-login--layer-time {
    font-size: 24em;
    font-weight: bold;
    color: rgba(0, 0, 0, 0.03);
    overflow: hidden;
  }

  // 登陆页面控件的容器
  .page-login--content {
    height: 100%;
    min-height: 500px;
  }

  // header
  .page-login--content-header {
    padding: 1em 0;

    .page-login--content-header-motto {
      margin: 0px;
      padding: 0px;
      color: $color-text-normal;
      text-align: center;
      font-size: 12px;
    }
  }

  // main
  .page-login--logo {
    width: 240px;
    margin-bottom: 2em;
    margin-top: -2em;
  }

  // 登录表单
  .page-login--form {
    width: 350px;
    // 卡片
    .el-card {
      margin-bottom: 15px;
    }

    // 登录按钮
    .button-login {
      width: 100%;
    }

    // 输入框左边的图表区域缩窄
    .el-input-group__prepend {
      padding: 0px 14px;
    }

    .login-code {
      height: 40px - 2px;
      display: block;
      margin: 0px -20px;
      border-top-right-radius: 2px;
      border-bottom-right-radius: 2px;
    }

    // 登陆选项
    .page-login--options {
      margin: 0px;
      padding: 0px;
      font-size: 14px;
      color: $color-primary;
      margin-bottom: 15px;
      font-weight: bold;
    }

    .page-login--quick {
      width: 100%;
    }
  }

  // 快速选择用户面板
  .page-login--quick-user {
    @extend %flex-center-col;
    padding: 10px 0px;
    border-radius: 4px;

    &:hover {
      background-color: $color-bg;

      i {
        color: $color-text-normal;
      }

      span {
        color: $color-text-normal;
      }
    }

    i {
      font-size: 36px;
      color: $color-text-sub;
    }

    span {
      font-size: 12px;
      margin-top: 10px;
      color: $color-text-sub;
    }
  }

  // footer
  .page-login--content-footer {
    padding: 1em 0;

    .page-login--content-footer-locales {
      padding: 0px;
      margin: 0px;
      margin-bottom: 15px;
      font-size: 12px;
      line-height: 12px;
      text-align: center;
      color: $color-text-normal;

      a {
        color: $color-text-normal;
        margin: 0 .5em;

        &:hover {
          color: $color-text-main;
        }
      }
    }

    .page-login--content-footer-copyright {
      padding: 0px;
      margin: 0px;
      margin-bottom: 10px;
      font-size: 12px;
      line-height: 12px;
      text-align: center;
      color: $color-text-normal;

      a {
        color: $color-text-normal;
      }
    }

    .page-login--content-footer-options {
      padding: 0px;
      margin: 0px;
      font-size: 12px;
      line-height: 12px;
      text-align: center;

      a {
        color: $color-text-normal;
        margin: 0 1em;
      }
    }
  }

  // 背景
  .circles {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    margin: 0px;
    padding: 0px;

    li {
      position: absolute;
      display: block;
      list-style: none;
      width: 20px;
      height: 20px;
      background: #FFF;
      animation: animate 25s linear infinite;
      bottom: -200px;
      @keyframes animate {
        0% {
          transform: translateY(0) rotate(0deg);
          opacity: 1;
          border-radius: 0;
        }
        100% {
          transform: translateY(-1000px) rotate(720deg);
          opacity: 0;
          border-radius: 50%;
        }
      }

      &:nth-child(1) {
        left: 15%;
        width: 80px;
        height: 80px;
        animation-delay: 0s;
      }

      &:nth-child(2) {
        left: 5%;
        width: 20px;
        height: 20px;
        animation-delay: 2s;
        animation-duration: 12s;
      }

      &:nth-child(3) {
        left: 70%;
        width: 20px;
        height: 20px;
        animation-delay: 4s;
      }

      &:nth-child(4) {
        left: 40%;
        width: 60px;
        height: 60px;
        animation-delay: 0s;
        animation-duration: 18s;
      }

      &:nth-child(5) {
        left: 65%;
        width: 20px;
        height: 20px;
        animation-delay: 0s;
      }

      &:nth-child(6) {
        left: 75%;
        width: 150px;
        height: 150px;
        animation-delay: 3s;
      }

      &:nth-child(7) {
        left: 35%;
        width: 200px;
        height: 200px;
        animation-delay: 7s;
      }

      &:nth-child(8) {
        left: 50%;
        width: 25px;
        height: 25px;
        animation-delay: 15s;
        animation-duration: 45s;
      }

      &:nth-child(9) {
        left: 20%;
        width: 15px;
        height: 15px;
        animation-delay: 2s;
        animation-duration: 35s;
      }

      &:nth-child(10) {
        left: 85%;
        width: 150px;
        height: 150px;
        animation-delay: 0s;
        animation-duration: 11s;
      }
    }
  }

  .tip {
    text-align: center;
    margin: 15px 0;
    color: #1790fe;
    background-image: linear-gradient(-60deg, #29bdd9, #1790fe);
    -webkit-text-fill-color: transparent;
    -webkit-background-clip: text;
  }
}
</style>
