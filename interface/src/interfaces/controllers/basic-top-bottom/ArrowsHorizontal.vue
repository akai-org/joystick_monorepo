<template >
  <div>
    <div id="btn-container">
      <button @touchstart="$emit(onTouchstart, keys.ARROW_UP)" @touchend="$emit(onTouchend, keys.ARROW_UP)">
          <i class="fas fa-arrow-up"></i>
      </button>

      <button @touchstart="$emit(onTouchstart, keys.ARROW_DOWN)" @touchend="$emit(onTouchend, keys.ARROW_DOWN)">
          <i class="fas fa-arrow-down"></i>
      </button>
    </div>

    <div id="rotate-warning">
      <i class="fas fa-sync"></i>
      <h1>Rotate your device.</h1>
    </div>

    <div id="divider"></div>
  </div>
</template>

<script>

import { events } from '../../InterfacesEvents'
import { keyActions as keys } from '../../Keys'

export default {
  name: 'BasicTopBottomInterface',
  emits: [events.onTouchstart, events.onTouchend],
  data: function () {
    return { ...events, keys }
  }
}
</script>

<style lang="scss" scoped>
  @import './basic-top-bottom.scss'
</style>
