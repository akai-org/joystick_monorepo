<template>
    <div class="ping-meter" :class="pingStatus">
        <i class="fas fa-wifi" ></i>
        <p>{{latency}}ms</p>
    </div>
</template>

<script>

export default {
  name: 'PingMeter',
  data () {
    return { pingStatus: 'good', latency: 42 }
  },
  methods: {
    changePingStatus: function (x, y) {
      this.pingStatus = x
      this.latency = y
    }
  }
}
</script>

<style lang="scss" src="../styles/PingMeter.scss"></style>
