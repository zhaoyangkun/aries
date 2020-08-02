<template>
  <d2-container>
    <el-tabs :tab-position="tabPosition" type="border-card">
      <el-tab-pane label="个人信息" style="width: 500px">
        <el-form :model="userForm" ref="userForm" :rules="userRules" label-width="100px">
          <el-form-item label="头像" prop="user_img">
            <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
          </el-form-item>
          <el-form-item label="用户名" prop="username">
            <el-input size="small" type="text" autocomplete="off" placeholder="用户名" disabled
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
                      v-model="userForm.nickname"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary">保存</el-button>
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
          <el-form-item label="确认密码" prop="confirm_new_pwd">
            <el-input size="small" type="password" autocomplete="off" placeholder="确认密码" show-password
                      v-model="pwdForm.confirm_new_pwd"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary">保存</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </d2-container>
</template>

<script>
export default {
  name: 'user',
  data () {
    return {
      tabPosition: 'top',
      userForm: {
        username: '',
        nickname: '',
        email: '',
        signature: ''
      },
      pwdForm: {
        old_pwd: '',
        new_pwd: '',
        confirm_new_pwd: ''
      },
      userRules: {},
      pwdRules: {}
    }
  }
}
</script>

<style scoped>

</style>
