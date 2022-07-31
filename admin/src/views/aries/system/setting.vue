<template>
  <d2-container>
    <template slot="header">系统 / 设置</template>
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
          <el-form-item label="静态资源根路径" prop="static_root">
            <el-input size="small" v-model="siteForm.static_root" type="text" autocomplete="off"
                      placeholder="静态资源根路径"></el-input>
          </el-form-item>
          <el-form-item label="Logo" prop="site_logo">
            <el-input size="small" v-model="siteForm.site_logo" type="text" autocomplete="off"
                      placeholder="图标地址"></el-input>
          </el-form-item>
          <el-form-item label="SEO 关键词" prop="seo_key_words">
            <el-input size="small" v-model="siteForm.seo_key_words" type="text" autocomplete="off"
                      placeholder="关键词以逗号隔开，如 Java,python,Golang"></el-input>
          </el-form-item>
          <el-form-item label="备案号" prop="record_number">
            <el-input size="small" :rows="5" v-model="siteForm.record_number" type="text" autocomplete="off"
                      placeholder="备案号"></el-input>
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

        <el-form :model="qubuForm" :rules="qubuFormRules" ref="qubuForm" label-width="130px"
                 v-show="storageForm.storage_type==='7bu'">
          <el-form-item label="Token" prop="token">
            <el-input size="small" v-model="qubuForm.token" type="text" autocomplete="off"
                      placeholder="Token"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.bedSaveLoading" @click="saveQubuForm">保存</el-button>
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

        <el-form :model="imgbbForm" :rules="imgbbFormRules" ref="imgbbForm" label-width="130px"
                 v-show="storageForm.storage_type==='imgbb'">
          <el-form-item label="Token" prop="token">
            <el-input size="small" v-model="imgbbForm.token" type="text" autocomplete="off"
                      placeholder="Token"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.bedSaveLoading" @click="saveImgbbForm">保存</el-button>
          </el-form-item>
        </el-form>

        <el-form :model="tencentCosForm" ref="tencentCosForm" :rules="tencentCosFormRules" label-width="130px"
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
        <el-form :model="commentPlugInForm" :rules="commentPlugInFormRules" ref="commentPlugInForm" label-width="130px">
          <el-form-item label="评论组件" prop="plug_in">
            <el-select size="small" v-model="commentPlugInForm.plug_in" placeholder="请选择评论组件"
                       @change="commentPlugInSelectChange">
              <el-option
                v-for="item in commentPlugIns"
                :value="item.value"
                :key="item.value"
                :label="item.name">
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>

        <el-form ref="localCommentForm" :model="localCommentForm" :rules="localCommentFormRules"
                 v-show="commentPlugInForm.plug_in==='local-comment'" label-width="130px">
          <el-form-item label="评论功能" prop="is_on">
            <el-switch size="small" active-value="1" inactive-value="0" v-model="localCommentForm.is_on"></el-switch>
          </el-form-item>
          <el-form-item label="评论审核" prop="is_review_on">
            <el-switch size="small" active-value="1" inactive-value="0"
                       v-model="localCommentForm.is_review_on"></el-switch>
          </el-form-item>
          <el-form-item label="回复后邮件通知" prop="is_reply_on">
            <el-switch size="small" active-value="1" inactive-value="0"
                       v-model="localCommentForm.is_reply_on"></el-switch>
          </el-form-item>
          <el-form-item label="每页评论个数" prop="page_size">
            <el-input size="small" v-model="localCommentForm.page_size" type="number" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.localCommentSaveLoading"
                       @click="saveLocalCommentForm">保存
            </el-button>
          </el-form-item>
        </el-form>

        <el-form ref="twikooForm" :model="twikooForm" v-show="commentPlugInForm.plug_in==='twikoo-comment'"
                 :rules="twikooFormRules" label-width="130px">
          <el-form-item label="Twikoo 部署教程:">
            <label style="color: gray">
              <a style="color: dodgerblue" href="https://twikoo.js.org/quick-start.html"
                 target="_blank">https://twikoo.js.org/quick-start.html</a>
            </label>
          </el-form-item>
          <el-form-item label="环境 id" prop="env_id">
            <el-input size="small" type="text" v-model="twikooForm.env_id" placeholder="env ID"></el-input>
          </el-form-item>
          <el-form-item label="区域" prop="region">
            <el-input size="small" type="text" v-model="twikooForm.region" placeholder="region"></el-input>
          </el-form-item>
          <el-form-item label="路径" prop="path">
            <el-input size="small" type="text" v-model="twikooForm.path" placeholder="path"></el-input>
          </el-form-item>
          <el-form-item label="语言" prop="lang">
            <el-input size="small" type="text" v-model="twikooForm.lang" placeholder="lang"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.twikooSaveLoading"
                       @click="saveTwikooForm">保存
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="参数设置" style="width: 500px">
        <el-form ref="paramForm" :model="paramForm" :rules="paramFormRules" label-width="130px">
          <el-form-item label="首页每页条数" prop="index_page_size">
            <el-input size="small" v-model="paramForm.index_page_size" type="number" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="归档页每页条数" prop="archive_page_size">
            <el-input size="small" v-model="paramForm.archive_page_size" type="number" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="站点地图每页条数" prop="site_map_page_size">
            <el-input size="small" v-model="paramForm.site_map_page_size" type="number" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.paramSaveLoading"
                       @click="saveParamForm">保存
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="社交信息" style="width: 500px">
        <el-form ref="socialForm" :model="socialForm" label-width="130px">
          <el-form-item label="QQ" prop="qq">
            <el-input size="small" v-model="socialForm.qq" type="text" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="微信" prop="wechat">
            <el-input size="small" v-model="socialForm.wechat" type="text" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="github" prop="github">
            <el-input size="small" v-model="socialForm.github" type="text" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="微博" prop="weibo">
            <el-input size="small" v-model="socialForm.weibo" type="text" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="知乎" prop="zhihu">
            <el-input size="small" v-model="socialForm.zhihu" type="text" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="small" type="primary" :loading="btn.socialSaveLoading"
                       @click="saveSocialInfoForm">保存
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
  saveImgbbSetting,
  saveLocalCommentSetting,
  saveParamSetting,
  saveQubuSetting,
  saveSiteSetting,
  saveSmmsSetting,
  saveSMTPSetting,
  saveSocialInfo,
  saveTencentCosSetting,
  saveTwikooSetting,
  sendTestEmail
} from '@/api/aries/sys'

export default {
  name: 'setting',
  data () {
    return {
      tabPosition: 'top',
      storageTypes: [
        { value: '7bu', name: '去不图床' },
        { value: 'sm.ms', name: 'sm.ms' },
        { value: 'imgbb', name: 'imgbb' },
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
      commentPlugIns: [
        { name: '本地评论组件', value: 'local-comment' },
        { name: 'twikoo 评论组件', value: 'twikoo-comment' }
      ],
      siteForm: {
        sys_id: null,
        type_name: '网站设置',
        site_name: '',
        site_desc: '',
        site_url: '',
        static_root: '',
        site_logo: '',
        seo_key_words: '',
        record_number: '',
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
      qubuForm: {
        sys_id: null,
        storage_type: '7bu',
        token: ''
      },
      smmsForm: {
        sys_id: null,
        storage_type: 'sm.ms',
        token: ''
      },
      imgbbForm: {
        sys_id: null,
        storage_type: 'imgbb',
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
      commentPlugInForm: {
        plug_in: 'local-comment'
      },
      localCommentForm: {
        sys_id: '',
        plug_in: 'local-comment',
        is_on: '1',
        is_review_on: '1',
        is_reply_on: '0',
        page_size: '10'
      },
      twikooForm: {
        sys_id: '',
        plug_in: 'twikoo-comment',
        env_id: '',
        region: '',
        path: 'window.location.pathname',
        lang: 'zh-CN'
      },
      paramForm: {
        sys_id: '',
        type_name: '参数设置',
        index_page_size: null,
        archive_page_size: null,
        site_map_page_size: null
      },
      socialForm: {
        sys_id: '',
        type_name: '社交信息',
        qq: '',
        wechat: '',
        github: '',
        weibo: '',
        zhihu: ''
      },
      siteRules: {
        site_name: [
          { required: true, message: '请输入网站名称', trigger: 'blur' },
          { max: 50, message: '网站名称长度不能超过 50', trigger: 'blur' }
        ],
        site_url: [
          { required: true, message: '请输入网站地址', trigger: 'blur' },
          { max: 255, message: '网站地址长度不能超过 255', trigger: 'blur' },
          { type: 'url', message: '请输入正确的 URL', trigger: 'blur' }
        ],
        static_root: [
          { max: 255, message: '静态资源根目录长度不能超过 255', trigger: 'blur' },
          { type: 'url', message: '静态资源根目录必须为有效的 URL', trigger: 'blur' }
        ],
        site_desc: [
          { max: 255, message: '网站描述长度不能超过 255', trigger: 'blur' }
        ],
        site_logo: [
          { max: 255, message: '网站 logo 长度不能超过 255', trigger: 'blur' }
        ],
        seo_key_words: [
          { max: 255, message: 'SEO 关键词长度不能超过 255', trigger: 'blur' }
        ],
        head_content: [
          { max: 1000, message: '全局 head 脚本长度不能超过 1000', trigger: 'blur' }
        ],
        footer_content: [
          { max: 1000, message: '全局 footer 脚本长度不能超过 1000', trigger: 'blur' }
        ]
      },
      emailRules: {
        address: [
          { required: true, message: '请输入SMTP 地址', trigger: 'blur' },
          { max: 50, message: 'SMTP 地址长度不能超过 50', trigger: 'blur' }
        ],
        port: [
          { required: true, message: '请输入端口', trigger: 'blur' },
          { max: 3, message: '端口长度不能超过 3', trigger: 'blur' }
        ],
        account: [
          { required: true, message: '请输入邮箱帐号', trigger: 'blur' },
          { max: 30, message: '邮箱帐号长度不能超过 30', trigger: 'blur' }
        ],
        pwd: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { max: 30, message: '密码长度不能超过 30', trigger: 'blur' }
        ],
        sender: [
          { required: true, message: '请输入发件人', trigger: 'blur' },
          { max: 30, message: '发件人长度不能超过 30', trigger: 'blur' }
        ]
      },
      emailSendRules: {
        sender: [
          { required: true, message: '请输入发送人', trigger: 'blur' },
          { max: 30, message: '发送人长度不能超过 30', trigger: 'blur' }
        ],
        receive_email: [
          { required: true, message: '请输入接收邮箱', trigger: 'blur' },
          { max: 30, message: '邮箱长度不能超过 30', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱帐号', trigger: 'blur' }
        ],
        title: [
          { required: true, message: '请输入邮件标题', trigger: 'blur' },
          { max: 100, message: '邮件标题长度不能超过 100', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '请输入邮件内容', trigger: 'blur' },
          { max: 1000, message: '邮件内容长度不能超过 1000', trigger: 'blur' }
        ]
      },
      storageFormRules: {
        storageType: [
          { required: true, message: '请选择存储类型', trigger: 'blur' }
        ]
      },
      qubuFormRules: {
        token: [
          { required: true, message: '请输入 Token', trigger: 'blur' },
          { max: 100, message: 'Token 长度不能超过 100', trigger: 'blur' }
        ]
      },
      smmsFormRules: {
        token: [
          { required: true, message: '请输入 Token', trigger: 'blur' },
          { max: 100, message: 'Token 长度不能超过 100', trigger: 'blur' }
        ]
      },
      imgbbFormRules: {
        token: [
          { required: true, message: '请输入 Token', trigger: 'blur' },
          { max: 100, message: 'Token 长度不能超过 100', trigger: 'blur' }
        ]
      },
      tencentCosFormRules: {
        host: [
          { required: true, message: '请输入存储桶地址', trigger: 'blur' },
          { max: 255, message: '存储桶地址长度不能超过 255 ', trigger: 'blur' }
        ],
        scheme: [
          { required: true, message: '请选择传输协议', trigger: 'blur' },
          { max: 5, message: '传输协议只能为 http 或者 https', trigger: 'blur' }
        ],
        region: [
          { required: true, message: '请选择区域', trigger: 'blur' },
          { max: 20, message: '区域长度不能超过 20 ', trigger: 'blur' }
        ],
        secret_id: [
          { required: true, message: '请输入 secret_id', trigger: 'blur' },
          { max: 255, message: 'secret_id 长度不能超过 255 ', trigger: 'blur' }
        ],
        secret_key: [
          { required: true, message: '请输入 secret_key', trigger: 'blur' },
          { max: 255, message: 'secret_key 长度不能超过 255 ', trigger: 'blur' }
        ],
        folder_path: [
          { required: true, message: '请输入上传目录', trigger: 'blur' },
          { max: 255, message: '上传目录长度不能超过 255 ', trigger: 'blur' }
        ]
      },
      commentPlugInFormRules: {
        plug_in: [
          { required: true, message: '请选择评论组件', trigger: 'blur' }
        ]
      },
      localCommentFormRules: {
        is_on: [
          { required: true, message: '请选择是否开启评论功能', trigger: 'blur' }
        ],
        is_review_on: [
          { required: true, message: '请选择是否开启评论审核功能', trigger: 'blur' }
        ]
      },
      twikooFormRules: {
        env_id: [
          { required: true, message: '请输入环境 id', trigger: 'blur' }
        ]
      },
      paramFormRules: {},
      btn: {
        siteSaveLoading: false,
        smtpSaveLoading: false,
        emailSendLoading: false,
        bedSaveLoading: false,
        localCommentSaveLoading: false,
        twikooSaveLoading: false,
        paramSaveLoading: false,
        socialSaveLoading: false
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
      } else if (tab.label === '评论设置') {
        this.initCommentSetting()
      } else if (tab.label === '参数设置') {
        this.getSysSetItem(tab.label, 'paramForm')
      } else if (tab.label === '社交信息') {
        this.getSysSetItem(tab.label, 'socialForm')
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
    // 初始化评论配置
    async initCommentSetting () {
      await this.getSysSetItem('评论组件设置', 'commentPlugInForm')
      await this.commentPlugInSelectChange()
    },
    // 图床类型切换
    async storageTypeSelectChange () {
      switch (this.storageForm.storage_type) {
        case '7bu':
          await this.getSysSetItem('7bu', 'qubuForm')
          break
        case 'sm.ms':
          await this.getSysSetItem('sm.ms', 'smmsForm')
          break
        case 'imgbb':
          await this.getSysSetItem('imgbb', 'imgbbForm')
          break
        case 'cos':
          await this.getSysSetItem('cos', 'tencentCosForm')
          break
      }
    },
    // 评论组件切换
    async commentPlugInSelectChange () {
      switch (this.commentPlugInForm.plug_in) {
        case 'local-comment':
          await this.getSysSetItem('local-comment', 'localCommentForm')
          break
        case 'twikoo-comment':
          await this.getSysSetItem('twikoo-comment', 'twikooForm')
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
    // 保存去不图床设置
    saveQubuForm () {
      this.$refs.qubuForm.validate(valid => {
        if (valid) {
          this.btn.bedSaveLoading = true
          setTimeout(() => {
            saveQubuSetting(this.qubuForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('7bu', 'qubuForm')
              })
              .catch(() => {
              })
            this.btn.bedSaveLoading = false
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
    // 保存 imgbb 设置
    saveImgbbForm () {
      this.$refs.imgbbForm.validate(valid => {
        if (valid) {
          this.btn.bedSaveLoading = true
          setTimeout(() => {
            saveImgbbSetting(this.imgbbForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('imgbb', 'imgbbForm')
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
    },
    // 保存评论设置
    saveLocalCommentForm () {
      this.$refs.localCommentForm.validate(valid => {
        if (valid) {
          this.btn.localCommentSaveLoading = true
          setTimeout(() => {
            saveLocalCommentSetting(this.localCommentForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('local-comment', 'localCommentForm')
              })
              .catch(() => {
              })
            this.btn.localCommentSaveLoading = false
          }, 300)
        }
      })
    },
    saveTwikooForm () {
      this.$refs.twikooForm.validate(valid => {
        if (valid) {
          this.btn.twikooSaveLoading = true
          setTimeout(() => {
            saveTwikooSetting(this.twikooForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('twikoo-comment', 'twikooForm')
              })
              .catch(() => {
              })
            this.btn.twikooSaveLoading = false
          }, 300)
        }
      })
    },
    // 保存参数设置
    saveParamForm () {
      this.$refs.paramForm.validate(valid => {
        if (valid) {
          this.btn.paramSaveLoading = true
          setTimeout(() => {
            saveParamSetting(this.paramForm)
              .then(res => {
                this.$message.success(res.msg)
                this.getSysSetItem('参数设置', 'paramForm')
              })
              .catch(() => {
              })
            this.btn.paramSaveLoading = false
          }, 300)
        }
      })
    },
    // 保存社交信息
    saveSocialInfoForm () {
      this.btn.socialSaveLoading = true
      setTimeout(() => {
        saveSocialInfo(this.socialForm)
          .then(res => {
            this.$message.success(res.msg)
            this.getSysSetItem('社交信息', 'socialForm')
          })
          .catch(() => {
          })
        this.btn.socialSaveLoading = false
      }, 300)
    }
  }
}
</script>

<style scoped>

</style>
