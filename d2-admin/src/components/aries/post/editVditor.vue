<template>
  <div class="editEditor-box">
    <div id="editEditor"></div>
  </div>
</template>

<script>
import Vditor from 'vditor'
import 'vditor/src/assets/scss/index.scss'
import { getSysSettingItem } from '@api/aries/sys'

export default {
  name: 'editEditor',
  data () {
    return {
      contentEditor: null,
      headers: {},
      uploadUrl: ''
    }
  },
  props: {
    content: {
      type: String,
      default: ''
    }
  },
  created () {
    this.initHeaders()
    this.initUploadUrl()
  },
  mounted () {
    this.initEditor()
  },
  watch: {
    // 监听编辑器内容变化
    content (newValue) {
      this.setContent(newValue)
    }
  },
  methods: {
    // 初始化 vditor
    initEditor () {
      const _this = this
      console.log(_this.uploadUrl)
      this.contentEditor = new Vditor('editEditor', {
        height: 380,
        toolbarConfig: {
          pin: true
        },
        cache: {
          enable: false
        },
        hint: {
          emoji: {
            pray: '🙏',
            broken_heart: '💔',
            ok_hand: '👌',
            smile: '😄',
            laughing: '😆',
            smirk: '😏',
            heart_eyes: '😍',
            grin: '😁',
            stuck_out_tongue: '😛',
            expressionless: '😑',
            unamused: '😒',
            sob: '😭',
            joy: '😂',
            tired_face: '😫',
            blush: '😊',
            cry: '😢',
            fearful: '😨'
          }
        },
        counter: 100000,
        upload: {
          accept: '.jpg,.png,.gif,.jpeg',
          max: 2 * 1024 * 1024,
          url: _this.uploadUrl,
          headers: _this.headers,
          filename: name =>
            name
              .replace(/[^(a-zA-Z0-9\u4e00-\u9fa5.)]/g, '')
              .replace(/[?\\/:|<>*[\]()$%{}@~]/g, '')
              .replace('/\\s/g', ''),
          success (editor, data) {
            data = JSON.parse(data) // 将 json 字符串转换成 json
            // editor.innerHTML = img_text // 将图片链接写入编辑区
          },
          error (data) {
            console.log(data)
            alert('上传失败')
          }
        },
        after: () => {
          this.setContent(this.content)
        }
      })
    },
    // 初始化请求头
    initHeaders () {
      this.headers = {
        token: localStorage.getItem('token'),
        'Content-Type': 'multipart/form-data',
        Accept: 'application/json'
      }
    },
    // 初始化上传路径
    initUploadUrl () {
      getSysSettingItem('网站设置')
        .then(res => {
          const siteUrl = res.data.site_url
          if (siteUrl.substr(-1) === '/') {
            this.uploadUrl = siteUrl + 'images/attachment/upload'
          } else {
            this.uploadUrl = siteUrl + '/images/attachment/upload'
          }
        })
        .catch(() => {
        })
    },
    // 获取编辑器文本
    getContent () {
      return this.contentEditor.getValue()
    },
    // 获取 HTML
    getHTML () {
      // 获取 markdown 的 HTML 内容
      return this.contentEditor.getHTML()
    },
    // 设置编辑器文本
    setContent (val) {
      this.contentEditor.setValue(val)
    }
  }
}
</script>

<style lang="scss">
</style>
