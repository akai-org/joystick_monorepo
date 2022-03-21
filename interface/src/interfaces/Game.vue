<template>
    <arrows-horizontal />
        <component
        @on-touchstart="handleTouchStart"
        @on-touchend="handleTouchEnd"
        :is="this.$store.state.gui" ></component>
</template>

<script>

import { keyActions as ka } from './Keys'
import { dispatchPressButton } from '../store/actionCreator'

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
    closeSocket () {
      this.$state.socket.close()
    }
  },
  mounted: function () {
    this.$store.dispatch('initWebsocketConnection')
    window.addEventListener('unload', this.closeSocket)
  },
  beforeUnmount: function () {
    window.removeEventListener('unload', this.closeSocket)
  }
}
</script>
