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
import BasicLeftRight from './controllers/basic-left-right/BasicLeftRight.vue'
import { ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT, ARROW_UP, KEY_DOWN, KEY_UP } from './Keys'

export default {
  name: 'Game',
  components: {
    // BasicTopBottom,
    BasicLeftRight
  },
  methods: {
    arrowUpDown () {
      this.$store.dispatch('pressButton', { key: ARROW_UP, keyState: KEY_DOWN })
    },
    arrowUpUp () {
      this.$store.dispatch('pressButton', { key: ARROW_UP, keyState: KEY_UP })
    },
    arrowDownDown () {
      this.$store.dispatch('pressButton', { key: ARROW_DOWN, keyState: KEY_DOWN })
    },
    arrowDownUp () {
      this.$store.dispatch('pressButton', { key: ARROW_DOWN, keyState: KEY_UP })
    },
    arrowLeftDown () {
      this.$store.dispatch('pressButton', { key: ARROW_LEFT, keyState: KEY_DOWN })
    },
    arrowLeftUp () {
      this.$store.dispatch('pressButton', { key: ARROW_LEFT, keyState: KEY_UP })
    },
    arrowRightDown () {
      this.$store.dispatch('pressButton', { key: ARROW_RIGHT, keyState: KEY_DOWN })
    },
    arrowRightUp () {
      this.$store.dispatch('pressButton', { key: ARROW_RIGHT, keyState: KEY_UP })
    }
  },
  mounted: function () {
    this.$store.dispatch('initWebsocketConnection')
  }
}
</script>
