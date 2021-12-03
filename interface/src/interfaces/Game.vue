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

export default {
  name: 'Game',
  data: function () {
    return {
      components: Components
    }
  },
  components: {
    ArrowsVertical,
    ArrowsHorizontal
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
    this.$store.dispatch('initWebsocketConnection')
  }
}
</script>
