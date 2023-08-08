<template>
        <component
        @on-touchstart="handleTouchStart"
        @on-touchend="handleTouchEnd"
        @on-joystick="handleJoystick"
        :is="this.$store.state.gui" ></component>
</template>

<script>

import { keyActions as ka } from './Keys'
import { dispatchJoystickMove, dispatchPressButton } from '../store/actionCreator'

import ArrowsVertical from './controllers/arrows-vertical/ArrowsVertical.vue'
import ArrowsHorizontal from './controllers/arrows-horizontal/ArrowsHorizontal.vue'
import ArrowsVertical1AB from './controllers/arrows-vertical-1ab/ArrowsVertical1AB.vue'
import CrossArrows from './controllers/cross-arrows/CrossArrows.vue'
import CrossArrows1AB from './controllers/cross-arrows-1ab/CrossArrows1AB.vue'
import Joystick from './controllers/joystick/Joystick.vue'

export default {
  name: 'Game',
  components: {
    ArrowsVertical,
    ArrowsHorizontal,
    ArrowsVertical1AB,
    CrossArrows,
    CrossArrows1AB,
    Joystick
  },
  methods: {
    handleTouchStart (key) {
      dispatchPressButton({ key, keyState: ka.KEY_DOWN })
    },
    handleTouchEnd (key) {
      dispatchPressButton({ key, keyState: ka.KEY_UP })
    },
    handleJoystick (joystickType, payload) {
      dispatchJoystickMove({ joystickType, payload })
    },
    closeSocket () {
      this.$store.state.socket.close()
      this.$store.state.socket = null
    }
  },
  mounted: function () {
    this.$store.dispatch('initWebsocketConnection')
  },
  beforeRouteLeave: function (to, from, next) {
    this.closeSocket()
    next()
  }
}
</script>
