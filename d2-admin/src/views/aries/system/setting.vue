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
        <el-form :model="storageForm" :rules="storageFormRules" ref="storageForm" label-width="130px">
          <el-form-item label="存储类型" prop="storage_type">
            <el-select size="small" v-model="storageForm.storage_type" placeholder="请选择存储类型"
                       @change="storageTypeSelectChange">
              <el-option
                v-for="item in storageTypes"
                :value="item.value"
                :key="item.value"
                :label="item.name">
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <el-form :model="smmsForm" :rules="smmsFormRules" ref="smmsForm" label-width="130px"
                 v-show="storageForm.storage_type==='sm.ms'">
          <el-form-item label="Token" prop="token">
            <el-input size="small" v-model="smmsForm.token" type="text" autocomplete="off"
                      placeholder="Token"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.bedSaveLoading" @click="saveSmmsForm">保存</el-button>
          </el-form-item>
        </el-form>
        <el-form :model="tencentCosForm" ref="tencentCosForm" label-width="130px"
                 v-show="storageForm.storage_type==='cos'">
          <el-form-item label="存储桶地址" prop="host">
            <el-input size="small" v-model="tencentCosForm.host" type="text" autocomplete="off"
                      placeholder="存储桶地址"></el-input>
          </el-form-item>
          <el-form-item label="传输协议" prop="scheme">
            <el-radio v-model="tencentCosForm.scheme" label="http">http</el-radio>
            <el-radio v-model="tencentCosForm.scheme" label="https">https</el-radio>
          </el-form-item>
          <el-form-item label="区域" prop="region">
            <el-select size="small" v-model="tencentCosForm.region" placeholder="请选择区域">
              <el-option
                v-for="item in cosRegions"
                :key="item.value"
                :label="item.name"
                :value="item.value"
              >
                <span style="float: left">{{ item.name }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="secret_id" prop="secret_id">
            <el-input size="small" v-model="tencentCosForm.secret_id" type="text" autocomplete="off"
                      placeholder="secret_id"></el-input>
          </el-form-item>
          <el-form-item label="secret_key" prop="secret_key">
            <el-input size="small" v-model="tencentCosForm.secret_key" type="text" autocomplete="off"
                      placeholder="secret_key"></el-input>
          </el-form-item>
          <el-form-item label="上传目录" prop="folder_path">
            <el-input size="small" v-model="tencentCosForm.folder_path" type="text" autocomplete="off"
                      placeholder="上传目录"></el-input>
          </el-form-item>
          <el-form-item label="图片处理" prop="img_process">
            <el-input size="small" v-model="tencentCosForm.img_process" type="text" autocomplete="off"
                      placeholder="请参照腾讯云 COS 图片处理文档"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.bedSaveLoading" @click="saveTencentCosForm">保存
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="评论设置" style="width: 500px">
        <el-form :model="tencentCosForm" ref="tencentCosForm" label-width="130px"
                 v-show="storageForm.storage_type==='cos'">
          <el-form-item label="存储桶地址" prop="host">
            <el-input size="small" v-model="tencentCosForm.host" type="text" autocomplete="off"
                      placeholder="存储桶地址"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.bedSaveLoading">保存
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </d2-container>
</template>

<script>
import {
  getSysSettingItem,
  saveSiteSetting,
  saveSmmsSetting,
  saveSMTPSetting,
  saveTencentCosSetting,
  sendTestEmail
} from '@/api/aries/sys'

export default {
  name: 'setting',
  data () {
    return {
      tabPosition: 'top',
      storageTypes: [
        { value: 'sm.ms', name: 'sm.ms' },
        { value: 'cos', name: '腾讯云' }
      ],
      cosRegions: [
        { name: '北京一区', value: 'ap-beijing-1' },
        { name: '北京', value: 'ap-beijing' },
        { name: '南京', value: 'ap-nanjing' },
        { name: '上海', value: 'ap-shanghai' },
        { name: '广州', value: 'ap-guangzhou' },
        { name: '成都', value: 'ap-chengdu' },
        { name: '重庆', value: 'ap-chongqing' },
        { name: '深圳金融', value: 'ap-shenzhen-fsi' },
        { name: '上海金融', value: 'ap-shanghai-fsi' },
        { name: '北京金融', value: 'ap-beijing-fsi' }
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
      storageForm: {
        storage_type: 'sm.ms'
      },
      smmsForm: {
        sys_id: null,
        storage_type: 'sm.ms',
        token: ''
      },
      tencentCosForm: {
        sys_id: null,
        storage_type: 'cos',
        host: '',
        scheme: 'https',
        region: '',
        secret_id: '',
        secret_key: '',
        folder_path: '',
        img_process: ''
      },
      commentForm: {

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
      storageFormRules: {
        storageType: [
          { required: true, message: '请选择存储类型', trigger: 'blur' }
        ]
      },
      smmsFormRules: {
        token: [
          { required: true, message: '请输入 Token', trigger: 'blur' }
        ]
      },
      tencentCosFormRules: {},
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
        this.initPicBedSetting()
      }
    },
    // 获取设置条目
    async getSysSetItem (name, form) {
      await getSysSettingItem(name)
        .then(res => {
          if (Object.keys(res.data).length > 0) {
            this[form] = res.data
          }
        })
        .catch(() => {
        })
    },
    // 初始化图床配置
    async initPicBedSetting () {
      await this.getSysSetItem('图床设置', 'storageForm')
      await this.storageTypeSelectChange()
    },
    // 图床类型切换
    async storageTypeSelectChange () {
      switch (this.storageForm.storage_type) {
        case 'sm.ms':
          await this.getSysSetItem('sm.ms', 'smmsForm')
          break
        case 'cos':
          await this.getSysSetItem('cos', 'tencentCosForm')
          break
      }
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
    // 保存 sm.ms 设置
    saveSmmsForm () {
      this.$refs.smmsForm.validate(valid => {
        if (valid) {
          this.btn.bedSaveLoading = true
          setTimeout(() => {
            saveSmmsSetting(this.smmsForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('sm.ms', 'smmsForm')
              })
              .catch(() => {
              })
            this.btn.bedSaveLoading = false
          }, 300)
        }
      })
    },
    // 保存腾讯云 COS 设置
    saveTencentCosForm () {
      this.$refs.tencentCosForm.validate(valid => {
        if (valid) {
          this.btn.bedSaveLoading = true
          setTimeout(() => {
            saveTencentCosSetting(this.tencentCosForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('cos', 'tencentCosForm')
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
