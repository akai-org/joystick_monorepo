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
    this.connectToWs()
  },
  data () {
    return {
      toSend: '',
      socket: null
    }
  },
  methods: {
    connectToWs () {
      this.socket = new WebSocket('ws://127.0.0.1:8081/socket')
    },
    send () {
      console.log('sending ', this.toSend)
      this.socket.send(this.toSend)
    }
  }

}
</script>
