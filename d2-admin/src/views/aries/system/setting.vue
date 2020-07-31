<template>
  <d2-container>
    <el-tabs :tab-position="tabPosition" @tab-click="handleTabClick">
      <el-tab-pane label="网站设置" style="width: 500px">
        <el-form :model="siteForm" ref="siteForm" label-width="100px">
          <el-form-item label="网站名称" prop="site_name">
            <el-input size="small" v-model="siteForm.site_name" type="text" autocomplete="off"
                      placeholder="网站名称"></el-input>
          </el-form-item>
          <el-form-item label="网站描述" prop="site_desc">
            <el-input size="small" v-model="siteForm.site_desc" type="text" autocomplete="off"
                      placeholder="网站描述"></el-input>
          </el-form-item>
          <el-form-item label="网站链接" prop="site_url">
            <el-input size="small" v-model="siteForm.site_url" type="text" autocomplete="off"
                      placeholder="网站链接"></el-input>
          </el-form-item>
          <el-form-item label="Logo" prop="site_logo">
            <el-input size="small" v-model="siteForm.site_logo" type="text" autocomplete="off"
                      placeholder="图标地址"></el-input>
          </el-form-item>
          <el-form-item label="SEO 关键词" prop="seo_key_words">
            <el-input size="small" v-model="siteForm.seo_key_words" type="text" autocomplete="off"
                      placeholder="关键词以逗号隔开，如 Java,python,Golang"></el-input>
          </el-form-item>
          <el-form-item label="全局 head" prop="head_content">
            <el-input size="small" :rows="5" v-model="siteForm.head_content" type="textarea" autocomplete="off"
                      placeholder="全局 head"></el-input>
          </el-form-item>
          <el-form-item label="全局 footer" prop="footer_content">
            <el-input size="small" :rows="5" v-model="siteForm.footer_content" type="textarea" autocomplete="off"
                      placeholder="全局 footer"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary">保存</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="邮件设置" style="width: 600px">
        <el-tabs tab-position="left">
          <el-tab-pane label="SMTP 服务配置">
            <el-form :model="emailForm" :rules="emailRules" ref="emailForm" label-width="100px">
              <el-form-item label="SMTP 地址" prop="address">
                <el-input size="small" v-model="emailForm.address" type="text" autocomplete="off"
                          placeholder="SMTP 地址"></el-input>
              </el-form-item>
              <el-form-item label="协议" prop="protocol">
                <el-input size="small" v-model="emailForm.protocol" type="text" autocomplete="off"
                          placeholder="协议"></el-input>
              </el-form-item>
              <el-form-item label="端口" prop="port">
                <el-input size="small" v-model="emailForm.port" type="number" autocomplete="off"
                          placeholder="端口"></el-input>
              </el-form-item>
              <el-form-item label="邮箱帐号" prop="account">
                <el-input size="small" v-model="emailForm.account" type="text" autocomplete="off"
                          placeholder="邮箱帐号"></el-input>
              </el-form-item>
              <el-form-item label="密码" prop="pwd">
                <el-input size="small" v-model="emailForm.pwd" type="password" autocomplete="off"
                          placeholder="密码"></el-input>
              </el-form-item>
              <el-form-item label="发件人" prop="sender">
                <el-input size="small" v-model="emailForm.sender" type="text" autocomplete="off"
                          placeholder="发件人"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button size="small" type="primary" :loading="btn.smtpSaveLoading" @click="saveEmailForm">保存
                </el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="邮件发送测试">
            <el-form :model="emailSendForm" :rules="emailSendRules" ref="emailSendForm" label-width="100px">
              <el-form-item label="发送人" prop="sender">
                <el-input size="small" v-model="emailSendForm.sender" type="text" autocomplete="off"
                          placeholder="发送人"></el-input>
              </el-form-item>
              <el-form-item label="接收邮箱" prop="receive_email">
                <el-input size="small" v-model="emailSendForm.receive_email" type="email" autocomplete="off"
                          placeholder="接收邮箱"></el-input>
              </el-form-item>
              <el-form-item label="邮件标题" prop="title">
                <el-input size="small" v-model="emailSendForm.title" type="text" autocomplete="off"
                          placeholder="邮件标题"></el-input>
              </el-form-item>
              <el-form-item label="邮件内容" prop="content">
                <el-input size="small" :rows="5" v-model="emailSendForm.content" type="textarea" autocomplete="off"
                          placeholder="邮件内容"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button size="small" type="primary" :loading="btn.emailSendLoading" @click="saveEmailSendForm">发送
                </el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>
      </el-tab-pane>
      <el-tab-pane label="图床设置">图床设置</el-tab-pane>
    </el-tabs>
  </d2-container>
</template>

<script>

import { getSysSettingItem, saveSMTP, sendTestEmail } from '@/api/aries/sys'

export default {
  name: 'setting',
  data () {
    return {
      tabPosition: 'top',
      siteForm: {
        type_name: '网站设置',
        site_name: '',
        site_desc: '',
        site_url: '',
        site_logo: '',
        seo_key_words: '',
        head_content: '',
        footer_content: ''
      },
      emailForm: {
        sys_id: null,
        type_name: '邮件设置',
        address: '',
        protocol: '',
        port: null,
        account: '',
        pwd: '',
        sender: ''
      },
      emailSendForm: {
        sender: '',
        receive_email: '',
        title: '',
        content: ''
      },
      siteRules: {},
      emailRules: {
        address: [
          { required: true, message: '请输入SMTP 地址', trigger: 'blur' }
        ],
        protocol: [
          { required: true, message: '请输入协议', trigger: 'blur' }
        ],
        port: [
          { required: true, message: '请输入端口', trigger: 'blur' }
        ],
        account: [
          { required: true, message: '请输入邮箱帐号', trigger: 'blur' }
        ],
        pwd: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ],
        sender: [
          { required: true, message: '请输入发件人', trigger: 'blur' }
        ]
      },
      emailSendRules: {
        sender: [
          { required: true, message: '请输入发送人', trigger: 'blur' }
        ],
        receive_email: [
          { required: true, message: '请输入接收邮箱', trigger: 'blur' }
        ],
        title: [
          { required: true, message: '请输入邮件标题', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '请输入邮件内容', trigger: 'blur' }
        ]
      },
      btn: {
        smtpSaveLoading: false,
        emailSendLoading: false
      }
    }
  },
  created () {
  },
  methods: {
    // tab 切换事件
    handleTabClick (tab) {
      if (tab.label === '邮件设置') {
        this.getSysSetItem(tab.label, 'emailForm')
      }
    },
    // 获取设置条目
    getSysSetItem (name, form) {
      getSysSettingItem(name)
        .then(res => {
          if (Object.keys(res.data).length > 0) {
            this[form] = res.data
          }
        })
        .catch(() => {
        })
    },
    // 保存 SMTP 表单事件
    saveEmailForm () {
      this.$refs.emailForm.validate(valid => {
        if (valid) {
          this.btn.smtpSaveLoading = true
          setTimeout(() => {
            saveSMTP(this.emailForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('邮件设置', 'emailForm')
              })
              .catch(() => {
              })
            this.btn.smtpSaveLoading = false
          }, 300)
        }
      })
    },
    // 发送邮件测试事件
    saveEmailSendForm () {
      this.$refs.emailSendForm.validate(valid => {
        if (valid) {
          this.btn.emailSendLoading = true
          setTimeout(() => {
            sendTestEmail(this.emailSendForm)
              .then(res => {
                this.$message.success(res.msg)
              })
              .catch(() => {
              })
            this.btn.emailSendLoading = false
          }, 300)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
