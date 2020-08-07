<template>
  <d2-container>
    <el-tabs :tab-position="tabPosition" @tab-click="handleTabClick" type="border-card">
      <el-tab-pane label="网站设置" style="width: 500px">
        <el-form :model="siteForm" ref="siteForm" :rules="siteRules" label-width="130px">
          <el-form-item label="网站名称" prop="site_name">
            <el-input size="small" v-model="siteForm.site_name" type="text" autocomplete="off"
                      placeholder="网站名称"></el-input>
          </el-form-item>
          <el-form-item label="网站描述" prop="site_desc">
            <el-input size="small" v-model="siteForm.site_desc" type="text" autocomplete="off"
                      placeholder="网站描述"></el-input>
          </el-form-item>
          <el-form-item label="网站地址" prop="site_url">
            <el-input size="small" v-model="siteForm.site_url" type="text" autocomplete="off"
                      placeholder="网站地址"></el-input>
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
            <el-button size="small" type="primary" :loading="btn.siteSaveLoading" @click="saveSiteForm">保存</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="邮件设置" style="width: 600px">
        <el-tabs tab-position="left">
          <el-tab-pane label="SMTP 服务配置">
            <el-form :model="emailForm" :rules="emailRules" ref="emailForm" label-width="130px">
              <el-form-item label="SMTP 地址" prop="address">
                <el-input size="small" v-model="emailForm.address" type="text" autocomplete="off"
                          placeholder="SMTP 地址"></el-input>
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
                          placeholder="密码" show-password></el-input>
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
            <el-form :model="emailSendForm" :rules="emailSendRules" ref="emailSendForm" label-width="130px">
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

      <el-tab-pane label="图床设置" style="width: 500px">
        <el-alert
          title="提示"
          description="目前仅支持土豆图床： https://images.ac.cn/"
          type="success"
          :closable="false">
        </el-alert>
        <el-form :model="picBedForm" :rules="picBedFormRules" ref="picBedForm" label-width="130px">
          <el-form-item label="图床性质" prop="type">
            <el-select size="small" v-model="picBedForm.type" type="text" autocomplete="off"
                       placeholder="请选择图床性质" @change="bedTypeChange">
              <el-option value="0" label="公有云"></el-option>
              <el-option value="1" label="私有云"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="公有云图床类型" prop="api_type" v-show="is_public">
            <el-select size="small" v-model="picBedForm.api_type" placeholder="请选择公有云图床类型">
              <el-option
                v-for="item in publicBedTypes"
                :value="item.value"
                :key="item.value"
                :label="item.name">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="私有云存储类型" prop="private_storage" v-show="!is_public">
            <el-select size="small" v-model="picBedForm.private_storage" placeholder="请选择私有云存储类型">
              <el-option
                v-for="item in privateTypes"
                :value="item.value"
                :key="item.value"
                :label="item.name"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Token" prop="token">
            <el-input size="small" :rows="5" v-model="picBedForm.token" type="text" autocomplete="off"
                      placeholder="请输入Token"></el-input>
          </el-form-item>
          <el-form-item label="上传目录" prop="folder">
            <el-input size="small" :rows="5" v-model="picBedForm.folder" type="text" autocomplete="off"
                      placeholder="请输入上传目录"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.bedSaveLoading" @click="savePicBedForm">保存</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </d2-container>
</template>

<script>
import { getSysSettingItem, savePicBedSetting, saveSiteSetting, saveSMTPSetting, sendTestEmail } from '@/api/aries/sys'

export default {
  name: 'setting',
  data () {
    return {
      tabPosition: 'top',
      publicBedTypes: [
        { value: 'xiaomi', name: '小米' },
        { value: 'Catbox', name: 'Catbox' },
        { value: 'SuNing', name: '苏宁' },
        { value: 'juejin', name: '掘金论坛' },
        { value: 'Neteasy', name: '网易' },
        { value: 'toutiao', name: '头条' },
        { value: 'BaiDu', name: '百度' }
      ],
      privateTypes: [
        { value: 'ftp', name: 'FTP' },
        { value: 'upyun', name: '又拍云' },
        { value: 'qiniu', name: '七牛云' },
        { value: 'oos', name: '阿里云OSS' },
        { value: 'cos', name: '腾讯云COS' },
        { value: 'ufile', name: 'U-file' },
        { value: 'zzidc', name: '快云' }
      ],
      siteForm: {
        sys_id: null,
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
      picBedForm: {
        sys_id: null,
        type_name: '图床设置',
        type: '0',
        api_type: '',
        private_storage: '',
        token: '',
        folder: ''
      },
      siteRules: {
        site_name: [
          { required: true, message: '请输入网站名称', trigger: 'blur' }
        ],
        site_url: [
          { required: true, message: '请输入网站地址', trigger: 'blur' }
        ]
      },
      emailRules: {
        address: [
          { required: true, message: '请输入SMTP 地址', trigger: 'blur' }
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
      picBedFormRules: {
        type: [
          { required: true, message: '请选择图床性质', trigger: 'blur' }
        ],
        token: [
          { required: true, message: '请输入 Token', trigger: 'blur' }
        ],
        folder: [
          { required: true, message: '请输入上传目录', trigger: 'blur' }
        ]
      },
      btn: {
        siteSaveLoading: false,
        smtpSaveLoading: false,
        emailSendLoading: false,
        bedSaveLoading: false
      },
      is_public: true
    }
  },
  created () {
    this.getSysSetItem('网站设置', 'siteForm')
  },
  methods: {
    // tab 切换事件
    handleTabClick (tab) {
      if (tab.label === '邮件设置') {
        this.getSysSetItem(tab.label, 'emailForm')
      } else if (tab.label === '网站设置') {
        this.getSysSetItem(tab.label, 'siteForm')
      } else if (tab.label === '图床设置') {
        this.getSysSetItem(tab.label, 'picBedForm')
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
    // 保存网站设置
    saveSiteForm () {
      this.$refs.siteForm.validate(valid => {
        if (valid) {
          this.btn.siteSaveLoading = true
          setTimeout(() => {
            saveSiteSetting(this.siteForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('网站设置', 'siteForm')
              })
              .catch(() => {
              })
            this.btn.siteSaveLoading = false
          }, 300)
        }
      })
    },
    // 保存 SMTP 设置
    saveEmailForm () {
      this.$refs.emailForm.validate(valid => {
        if (valid) {
          this.btn.smtpSaveLoading = true
          setTimeout(() => {
            saveSMTPSetting(this.emailForm)
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
    // 发送邮件测试
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
    },
    // 监听图床性质
    bedTypeChange (value) {
      this.picBedForm.api_type = ''
      this.picBedForm.private_storage = ''
      this.is_public = value === '0'
    },
    // 保存图床设置
    savePicBedForm () {
      this.$refs.picBedForm.validate(valid => {
        if (valid) {
          this.btn.bedSaveLoading = true
          setTimeout(() => {
            savePicBedSetting(this.picBedForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('图床设置', 'picBedForm')
              })
              .catch(() => {
              })
            this.btn.bedSaveLoading = false
          }, 300)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
