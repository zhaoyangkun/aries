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
                ref="loginForm"
                label-position="top"
                :rules="rules"
                :model="loginForm"
                size="default">
                <el-form-item prop="username">
                  <el-input
                    type="text"
                    v-model="loginForm.username"
                    placeholder="用户名/邮箱">
                    <i slot="prepend" class="fa fa-user-circle-o"></i>
                  </el-input>
                </el-form-item>
                <el-form-item prop="pwd">
                  <el-input
                    type="password"
                    v-model="loginForm.pwd"
                    placeholder="密码">
                    <i slot="prepend" class="fa fa-key"></i>
                  </el-input>
                </el-form-item>
                <el-form-item prop="captcha_val">
                  <el-input
                    type="text"
                    v-model="loginForm.captcha_val"
                    placeholder="验证码">
                    <template slot="append">
                      <img @click="loadCaptcha" alt="captcha" class="login-captcha_val" :src="captcha.url">
                    </template>
                  </el-input>
                </el-form-item>
                  <el-button
                    size="default"
                    :loading="btnLoading"
                    @click="submit"
                    type="primary"
                    class="button-login">
                    登录
                  </el-button>
              </el-form>
            </el-card>
            <p
              class="page-login--options"
              flex="main:justify cross:center">
              <span @click="$router.push('/forgetPwd')"><d2-icon name="question-circle"/> 忘记密码</span>
              <span @click="toInitSetting"><d2-icon name="cogs"/> 初始化配置</span>
            </p>
          </div>
        </div>
        <div class="page-login--content-footer">
          <p class="page-login--content-footer-locales">
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs'
import { mapActions } from 'vuex'
import { createCaptcha } from '@/api/aries/auth'
import { getAllUsers } from '@/api/aries/user'

export default {
  name: 'login',
  data () {
    return {
      // logo路径
      logoPath: require('@/assets/img/logo@2x.png'),
      timeInterval: null,
      time: dayjs().format('HH:mm:ss'),
      btnLoading: false,
      captcha: {
        url: ''
      },
      // 表单
      loginForm: {
        username: '',
        pwd: '',
        captcha_id: '',
        captcha_val: ''
      },
      // 表单校验
      rules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          }
        ],
        pwd: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
          }
        ],
        captcha_val: [
          {
            required: true,
            message: '请输入验证码',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  created () {
    this.checkFirst()
  },
  mounted () {
    this.timeInterval = setInterval(() => {
      this.refreshTime()
    }, 1000)
    this.loadCaptcha()
  },
  beforeDestroy () {
    clearInterval(this.timeInterval)
  },
  methods: {
    ...mapActions('d2admin/account', [
      'login'
    ]),
    refreshTime () {
      this.time = dayjs().format('HH:mm:ss')
    },
    // 跳转到配置页面
    toInitSetting () {
      if (this.userList === 0) {
        this.$router.push('/register')
      } else {
        const h = this.$createElement
        this.$notify({
          title: '提示',
          message: h('i', { style: 'color: #1790fe' }, '已完成初始化配置，为了帐号安全，请先登录进入后台管理进行配置'),
          type: 'warning'
        })
      }
    },
    // 加载验证码
    loadCaptcha () {
      createCaptcha()
        .then(res => {
          const data = res.data
          this.loginForm.captcha_id = data.captcha_id
          this.captcha.url = data.captcha_url
        })
        .catch(() => {
        })
    },
    // 提交表单
    submit: function () {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.btnLoading = true
          setTimeout(() => {
            // 登录
            this.login({
              username: this.loginForm.username,
              pwd: this.loginForm.pwd,
              captchaId: this.loginForm.captcha_id,
              captchaVal: this.loginForm.captcha_val
            })
              .then(() => {
                this.$message.success('登录成功')
                // 重定向对象不存在则返回顶层路径
                this.$router.replace(this.$route.query.redirect || '/')
              })
              .catch(() => {
              })
            this.btnLoading = false
          }, 300)
        }
      })
    },
    // 判断是否是第一次运行，若为第一次，进入配置界面
    checkFirst () {
      getAllUsers()
        .then(res => {
          this.userList = res.data
          if (this.userList.length === 0) {
            this.$router.push('/register')
          }
        })
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

    .login-captcha_val {
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
