<template>
  <d2-container>
    <iframe :src="src" style="width: 100%;height: 100%;"></iframe>
  </d2-container>
</template>

<script>
export default {
  name: 'doc',
  data () {
    return {
      src: '/swagger/index.html'
    }
  }
}
</script>

<style scoped>

</style>
