<template>
  <div>
    <div id="btn-container">
      <button @touchstart="$emit(onTouchstart, keys.ARROW_LEFT)" @touchend="$emit(onTouchend, keys.ARROW_LEFT)">
          <i class="fas fa-arrow-left"></i>
      </button>
      <button @touchstart="$emit(onTouchstart, keys.ARROW_RIGHT)" @touchend="$emit(onTouchend, keys.ARROW_RIGHT)">
          <i class="fas fa-arrow-right"></i>
      </button>
    </div>

    <div id="rotate-warning">
      <i class="fas fa-sync"></i>
      <h1>Rotate your device.</h1>
    </div>

    <div id="divider"></div>
    <div class="fullscreen-btn" @click="tryToFullscreen()">
      <i class="fas fa-expand-arrows-alt"></i>
      Click to go fullscreen
    </div>
  </div>
</template>

<script>

import { events } from '../../InterfacesEvents'
import { keyActions as keys } from '../../Keys'

// when phone is rotated to this orientation we lock it
const wantedOrientation = 'landscape'

screen.orientation.addEventListener('change', function (e) {
  const orientation = e.target.type
  if (orientation === `${wantedOrientation}-primary` || orientation === `${wantedOrientation}-secondary`) {
    screen.orientation.lock(orientation)
  }
})

export default {
  name: 'BasicLeftRightInterface',
  emits: [events.onTouchstart, events.onTouchend],
  data: function () {
    return {
      ...events,
      keys
    }
  },
  methods: {
    tryToFullscreen () {
      document.body.requestFullscreen()
    }
  }
}
</script>

<style lang="scss" scoped>
  @import './ArrowsHorizontal.scss'
</style>
