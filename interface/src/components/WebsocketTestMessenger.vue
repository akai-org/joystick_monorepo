<template>
  <div class="hello">
    <label for="message">
      {{ msg }}
      <input type="text" name="message" @keydown.enter="send" v-model="toSend">
    </label>
  </div>
</template>

<script>
export default {
  name: 'WebsocketTestMessenger',
  props: {
    msg: String
  },
  mounted () {
    console.log("mounted")
    this.$store.dispatch('initWebsocketConnection')
  },
  data () {
    return {
      toSend: ''
    }
  },
  methods: {
    send () {
      console.log('sending ', this.toSend)
      this.$store.dispatch('pressButton', { key: 'ARROW_RIGHT', keyState: 'KEY_DOWN' })
    }
  }

}
</script>
