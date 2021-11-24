<script setup>
import BasicLeftRight from './controllers/basic-left-right/BasicLeftRight.vue'
</script>
<template>
    <!-- <BasicTopBottom
        @arrow-up-down="arrowUpDown"
        @arrow-up-up="arrowUpUp"
        @arrow-down-down="arrowDownDown"
        @arrow-down-up="arrowDownUp" /> -->
    <BasicLeftRight
        @arrow-left-down="arrowLeftDown"
        @arrow-left-up="arrowLeftUp"
        @arrow-right-down="arrowRightDown"
        @arrow-right-up="arrowRightUp" />
</template>

<script>
// import BasicTopBottom from './controllers/basic-top-bottom/BasicTopBottom.vue'
import { keyActions as ka } from './Keys'
import { dispatchPressButton } from '../store/actionCreator'

export default {
  name: 'Game',
  methods: {
    arrowUpDown () {
      dispatchPressButton({ key: ka.ARROW_UP, keyState: ka.KEY_DOWN })
    },
    arrowUpUp () {
      dispatchPressButton({ key: ka.ARROW_UP, keyState: ka.KEY_UP })
    },
    arrowDownDown () {
      dispatchPressButton({ key: ka.ARROW_DOWN, keyState: ka.KEY_DOWN })
    },
    arrowDownUp () {
      dispatchPressButton({ key: ka.ARROW_DOWN, keyState: ka.KEY_UP })
    },
    arrowLeftDown () {
      dispatchPressButton({ key: ka.ARROW_LEFT, keyState: ka.KEY_DOWN })
    },
    arrowLeftUp () {
      dispatchPressButton({ key: ka.ARROW_LEFT, keyState: ka.KEY_UP })
    },
    arrowRightDown () {
      dispatchPressButton({ key: ka.ARROW_RIGHT, keyState: ka.KEY_DOWN })
    },
    arrowRightUp () {
      dispatchPressButton({ key: ka.ARROW_RIGHT, keyState: ka.KEY_UP })
    }
  },
  mounted: function () {
    this.$store.dispatch('initWebsocketConnection')
  }
}
</script>
