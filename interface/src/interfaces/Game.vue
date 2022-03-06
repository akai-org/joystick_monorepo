<template>
        <component
        @on-touchstart="handleTouchStart"
        @on-touchend="handleTouchEnd"
        :is="this.components.get(this.$store.state.gui)" ></component>
</template>

<script>

import { Components } from './components'
import { keyActions as ka } from './Keys'
import { dispatchPressButton } from '../store/actionCreator'

import ArrowsVertical from './controllers/basic-left-right/ArrowsVertical.vue'
import ArrowsHorizontal from './controllers/basic-top-bottom/ArrowsHorizontal.vue'
import ArrowsVertical1AB from './controllers/arrows-vertical-1ab/ArrowsVertical1AB.vue'
import CrossArrows from './controllers/cross-arrows/CrossArrows.vue'
import CrossArrows1AB from './controllers/cross-arrows-1ab/CrossArrows1AB.vue'

export default {
  name: 'Game',
  data: function () {
    return {
      components: Components
    }
  },
  components: {
    ArrowsVertical,
    ArrowsHorizontal,
    ArrowsVertical1AB,
    CrossArrows,
    CrossArrows1AB
  },
  methods: {
    handleTouchStart (key) {
      dispatchPressButton({ key, keyState: ka.KEY_DOWN })
    },
    handleTouchEnd (key) {
      dispatchPressButton({ key, keyState: ka.KEY_UP })
    }
  },
  mounted: function () {
    console.log(this.$store.state.gui)
    this.$store.dispatch('initWebsocketConnection')
  }
}
</script>
