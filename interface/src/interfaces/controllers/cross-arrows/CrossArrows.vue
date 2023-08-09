<template>
  <div>
  <div class="controller-crossarrows">
      <button class="ctrl-button" id="top-btn" @touchstart="$emit(onTouchstart, keys.ARROW_UP)" @touchend="$emit(onTouchend, keys.ARROW_UP)">
          <i class="fas fa-arrow-up"></i>
      </button>
      <button class="ctrl-button" id="left-btn" @touchstart="$emit(onTouchstart, keys.ARROW_LEFT)" @touchend="$emit(onTouchend, keys.ARROW_LEFT)">
          <i class="fas fa-arrow-left"></i>
      </button>
      <button class="ctrl-button" id="right-btn" @touchstart="$emit(onTouchstart, keys.ARROW_RIGHT)" @touchend="$emit(onTouchend, keys.ARROW_RIGHT)">
          <i class="fas fa-arrow-right"></i>
      </button>
      <button class="ctrl-button" id="bottom-btn" @touchstart="$emit(onTouchstart, keys.ARROW_DOWN)" @touchend="$emit(onTouchend, keys.ARROW_DOWN)">
          <i class="fas fa-arrow-down"></i>
      </button>
  </div>

  <div class="rotate-warning rotate-warning-portrait">
    <i class="fas fa-sync"></i>
    <h1>Rotate your device.</h1>
  </div>

  <div class="fullscreen-btn" @click="tryToFullscreen()">
    <i class="fas fa-expand-arrows-alt"></i>
    Click to go fullscreen
  </div>
  </div>
</template>

<script>

import { useInterfaceLifecycleEssentials } from '../../../utils/useInterfaceLifecycleEssentials'
import { events } from '../../InterfacesEvents'
import { keyActions as keys } from '../../Keys'

// when phone is rotated to this orientation we lock it
const wantedOrientation = 'landscape'

export default {
  name: 'CrossArrowsInterface',
  emits: [events.onTouchstart, events.onTouchend, events.onSetOrientation],
  setup (_, context) {
    useInterfaceLifecycleEssentials(context, wantedOrientation)
  },
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
