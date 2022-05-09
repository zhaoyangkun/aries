<template>
  <d2-container>
    <template slot="header">用户 / 用户信息</template>
    <el-tabs :tab-position="tabPosition" type="border-card">
      <el-tab-pane label="个人信息" style="width: 500px">
        <el-form :model="userForm" ref="userForm" :rules="userRules" label-width="100px">
          <el-form-item label="头像" prop="user_img">
            <el-input size="small" type="text" autocomplete="off" placeholder="头像链接"
                      v-model="userForm.user_img"></el-input>
          </el-form-item>
          <el-form-item label="用户名" prop="username">
            <el-input size="small" type="text" autocomplete="off" placeholder="用户名" readonly
                      v-model="userForm.username"></el-input>
          </el-form-item>
          <el-form-item label="昵称" prop="nickname">
            <el-input size="small" type="text" autocomplete="off" placeholder="昵称"
                      v-model="userForm.nickname"></el-input>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input size="small" type="text" autocomplete="off" placeholder="邮箱"
                      v-model="userForm.email"></el-input>
          </el-form-item>
          <el-form-item label="个性签名" prop="signature">
            <el-input size="small" type="textarea" :rows="3" autocomplete="off" placeholder="个性签名"
                      v-model="userForm.signature"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.userInfoLoading" @click="handleUpdateUserInfo">保存
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="密码" style="width: 500px">
        <el-form :model="pwdForm" ref="pwdForm" :rules="pwdRules" label-width="100px">
          <el-form-item label="旧密码" prop="old_pwd">
            <el-input size="small" type="password" autocomplete="off" placeholder="旧密码" show-password
                      v-model="pwdForm.old_pwd"></el-input>
          </el-form-item>
          <el-form-item label="新密码" prop="new_pwd">
            <el-input size="small" type="password" autocomplete="off" placeholder="新密码" show-password
                      v-model="pwdForm.new_pwd"></el-input>
          </el-form-item>
          <el-form-item label="确认新密码" prop="confirm_new_pwd">
            <el-input size="small" type="password" autocomplete="off" placeholder="确认新密码" show-password
                      v-model="pwdForm.confirm_new_pwd"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.pwdLoading" @click="handleUpdateUserPwd">保存</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </d2-container>
</template>

<script>
import { mapActions } from 'vuex'
import { getAllUsers, updateUser, updateUserPwd } from '@/api/aries/user'

export default {
  name: 'user',
  data () {
    // 自定义校验函数
    const validateConfirmNewPwd = (rule, value, callback) => {
      if (this.pwdForm.new_pwd !== value) {
        callback(new Error('新密码和确认新密码不一致！'))
      } else {
        callback()
      }
    }
    const validateNewPwd = (rule, value, callback) => {
      if (this.pwdForm.old_pwd === value) {
        callback(new Error('新旧密码不能相同！'))
      } else {
        callback()
      }
    }
    return {
      tabPosition: 'top',
      btn: {
        userInfoLoading: false,
        pwdLoading: false
      },
      userForm: {
        ID: null,
        user_img: '',
        username: '',
        nickname: '',
        email: '',
        signature: ''
      },
      pwdForm: {
        username: '',
        old_pwd: '',
        new_pwd: '',
        confirm_new_pwd: ''
      },
      userRules: {
        username: [
          { required: true, trigger: 'blur', message: '请输入用户名' },
          { min: 3, max: 30, trigger: 'blur', message: '用户名长度必须为 3 到 30 位' }
        ],
        email: [
          { required: true, trigger: 'blur', message: '请输入邮箱' },
          { max: 30, trigger: 'blur', message: '邮箱长度不能超过 30 位' },
          { type: 'email', trigger: 'blur', message: '邮箱格式不正确' }
        ]
      },
      pwdRules: {
        username: [
          { required: true, trigger: 'blur', message: '请输入用户名' },
          { min: 3, max: 30, trigger: 'blur', message: '用户名长度必须为 3 到 30 位' }
        ],
        old_pwd: [
          { required: true, trigger: 'blur', message: '请输入旧密码' },
          { min: 6, max: 20, message: '旧密码长度需在 6~20 之间', trigger: 'blur' }
        ],
        new_pwd: [
          { required: true, trigger: 'blur', message: '请输入新密码' },
          { min: 6, max: 20, message: '新密码长度需在 6~20 之间', trigger: 'blur' },
          { validator: validateNewPwd, trigger: 'blur' }
        ],
        confirm_new_pwd: [
          { required: true, trigger: 'blur', message: '请输入确认新密码' },
          { min: 6, max: 20, message: '确认新密码长度需在 6~20 之间', trigger: 'blur' },
          { validator: validateConfirmNewPwd, trigger: 'blur' }
        ]
      }
    }
  },
  created () {
    this.fetchUserInfo()
  },
  methods: {
    ...mapActions('d2admin/account', [
      'logout'
    ]),
    // 获取用户信息
    fetchUserInfo () {
      getAllUsers()
        .then(res => {
          this.userForm.ID = res.data[0].ID
          this.userForm.user_img = res.data[0].user_img
          this.userForm.username = res.data[0].username
          this.userForm.nickname = res.data[0].nickname
          this.userForm.email = res.data[0].email
          this.userForm.signature = res.data[0].signature
          this.pwdForm.username = res.data[0].username
        })
        .catch(() => {
        })
    },
    // 修改用户信息事件
    handleUpdateUserInfo () {
      this.$refs.userForm.validate(valid => {
        if (valid) {
          this.btn.userInfoLoading = true
          setTimeout(() => {
            updateUser(this.userForm)
              .then(res => {
                this.$message.success(res.msg)
                this.fetchUserInfo()
              })
              .catch(() => {
              })
            this.btn.userInfoLoading = false
          }, 300)
        }
      })
    },
    // 修改密码事件
    handleUpdateUserPwd () {
      this.$refs.pwdForm.validate(valid => {
        if (valid) {
          this.btn.pwdLoading = true
          setTimeout(() => {
            updateUserPwd(this.pwdForm)
              .then(res => {
                this.$message.success(res.msg)
                // 注销
                this.logout({
                  confirm: false
                })
              })
              .catch(() => {
              })
            this.btn.pwdLoading = false
          }, 300)
        }
      })
    }
  }
}
</script>

<style scoped>
</style>
