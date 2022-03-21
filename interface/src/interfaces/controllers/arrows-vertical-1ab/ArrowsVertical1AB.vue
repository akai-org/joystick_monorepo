<template>
    <div class="controller-arrows-vertical-1ab">
      <div class="col">
        <button class="ctrl-button" @touchstart="$emit(onTouchstart, keys.ARROW_UP)" @touchend="$emit(onTouchend, keys.ARROW_UP)">
          <i class="fas fa-arrow-up"></i>
        </button>

        <button class="ctrl-button" @touchstart="$emit(onTouchstart, keys.ARROW_DOWN)" @touchend="$emit(onTouchend, keys.ARROW_DOWN)">
          <i class="fas fa-arrow-down"></i>
         </button>
      </div>
      <div class="col">
        <button class="ctrl-button" id="a-btn" @touchstart="$emit(onTouchstart, keys.ACTION_BUTTON_1)" @touchend="$emit(onTouchend, keys.ACTION_BUTTON_1)">
          A
         </button>
        <button class="ctrl-button" id="b-btn" @touchstart="$emit(onTouchstart, keys.ACTION_BUTTON_2)" @touchend="$emit(onTouchend, keys.ACTION_BUTTON_2)">
          B
         </button>
      </div>
    </div>
    <div class="rotate-warning rotate-warning-portrait">
      <i class="fas fa-sync"></i>
      <h1>Rotate your device.</h1>
    </div>
    <div class="fullscreen-btn" @click="tryToFullscreen()">
      <i class="fas fa-expand-arrows-alt"></i>
      Click to go fullscreen
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
  name: 'ArrowsVertical1ABInterface',
  emits: [events.onTouchstart, events.onTouchend],
  data: function () {
    return { ...events, keys }
  },
  methods: {
    tryToFullscreen () {
      document.body.requestFullscreen()
    }
  }
}
</script>
