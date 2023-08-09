<template >
<div>
  <div class="controller-arrows-vertical">
    <button class="ctrl-button" @touchstart="$emit(onTouchstart, keys.ARROW_UP)" @touchend="$emit(onTouchend, keys.ARROW_UP)">
      <i class="fas fa-arrow-up"></i>
    </button>

    <button class="ctrl-button" @touchstart="$emit(onTouchstart, keys.ARROW_DOWN)" @touchend="$emit(onTouchend, keys.ARROW_DOWN)">
      <i class="fas fa-arrow-down"></i>
    </button>

    <div id="divider"></div>
  </div>

  <div class="rotate-warning rotate-warning-landscape">
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
const wantedOrientation = 'portrait'

export default {
  name: 'BasicTopBottomInterface',
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
