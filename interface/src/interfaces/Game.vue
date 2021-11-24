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
import { keyActions as ka } from './Keys'

export default {
  name: 'Game',
  components: {
    // BasicTopBottom,
    BasicLeftRight
  },
  data: function () {
    return {
      interface: 'some interface'
    }
  },
  methods: {
    arrowUpDown () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_UP, keyState: ka.KEY_DOWN })
    },
    arrowUpUp () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_UP, keyState: ka.KEY_UP })
    },
    arrowDownDown () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_DOWN, keyState: ka.KEY_DOWN })
    },
    arrowDownUp () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_DOWN, keyState: ka.KEY_UP })
    },
    arrowLeftDown () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_LEFT, keyState: ka.KEY_DOWN })
    },
    arrowLeftUp () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_LEFT, keyState: ka.KEY_DOWN })
    },
    arrowRightDown () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_RIGHT, keyState: ka.KEY_DOWN })
    },
    arrowRightUp () {
      this.$store.dispatch('pressButton', { key: ka.ARROW_RIGHT, keyState: ka.KEY_UP })
    }
  },
  mounted: function () {
    this.$store.dispatch('initWebsocketConnection')
  }
}
</script>
