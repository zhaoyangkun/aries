<template>
  <div class="addEditor-box">
    <div id="addEditor"></div>
  </div>
</template>

<script>
import Vditor from 'vditor'
import 'vditor/src/assets/scss/index.scss'
import { getSysSettingItem } from '@api/aries/sys'
import { uploadImgToAttachment } from '@/api/aries/picture'

export default {
  name: 'addEditor',
  data () {
    return {
      contentEditor: null,
      headers: {},
      uploadUrl: null
    }
  },
  props: {
    content: {
      type: String,
      default: ''
    }
  },
  mounted () {
    // await this.initUploadUrl()
    // await this.initHeaders()
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
    async initEditor () {
      this.contentEditor = new Vditor('addEditor', {
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
          // url: this.uploadUrl,
          // headers: this.headers,
          // withCredentials: true,
          filename: name =>
            name
              .replace(/[^(a-zA-Z0-9\u4e00-\u9fa5.)]/g, '')
              .replace(/[?\\/:|<>*[\]()$%{}@~]/g, '')
              .replace('/\\s/g', ''),
          // 自定义上传
          handler: files => {
            const formData = new FormData()
            for (const file of files) {
              formData.append('file[]', file)
            }
            uploadImgToAttachment(formData)
              .then(res => {
                console.log(res)
              })
              .catch(() => {
              })
          }
          // success (editor, data) {
          //   data = JSON.parse(data) // 将 json 字符串转换成 json
          //   // editor.innerHTML = img_text // 将图片链接写入编辑区
          // },
          // error (data) {
          //   console.log(data)
          //   alert('上传失败')
          // }
        },
        after: () => {
          this.setContent(this.content)
        }
      })
    },
    // 初始化请求头
    async initHeaders () {
      this.headers = {
        token: localStorage.getItem('token'),
        'Content-Type': 'multipart/form-data',
        Accept: 'application/json'
      }
    },
    // 初始化上传路径
    async initUploadUrl () {
      await getSysSettingItem('网站设置')
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
.vditor-toolbar--pin {
  position: relative;
}
</style>
