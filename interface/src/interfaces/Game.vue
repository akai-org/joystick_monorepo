<template>
        <component
        @on-touchstart="handleTouchStart"
        @on-touchend="handleTouchEnd"
        @on-set-orientation="handleOrientationChange"
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

async function onChangeOrientation (e, wantedOrientation) {
  console.log('orientation changed')
  const orientation = e.target.type
  if (orientation === `${wantedOrientation}-primary` || orientation === `${wantedOrientation}-secondary`) {
    try {
      await screen.orientation.lock(wantedOrientation)
    } catch (error) {
      console.log(error)
    }
  }
}

export default {
  name: 'GameDisplay',
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
      this.$store.state.socket.close()
      this.$store.state.socket = null
    },
    handleOrientationChange (orientation) {
      screen.orientation.removeEventListener('change', (e) => onChangeOrientation(e, orientation))
      if (orientation) {
        screen.orientation.addEventListener('change', (e) => onChangeOrientation(e, orientation))
      }
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
