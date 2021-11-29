<template>
  <div>
    <div id="btn-container">
      <button @touchstart="$emit(arrowLeftDown)" @touchend="$emit(arrowLeftUp)">
          <i class="fas fa-arrow-left"></i>
      </button>
      <button @touchstart="$emit(arrowRightDown)" @touchend="$emit(arrowRightUp)">
          <i class="fas fa-arrow-right"></i>
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

export default {
  name: 'BasicLeftRightInterface',
  emits: [events.arrowLeftDown, events.arrowLeftUp, events.arrowRightDown, events.arrowRightUp],
  data: function () {
    return {
      ...events
    }
  }
}
</script>

<style lang="scss" scoped>
  @import './basic-left-right.scss'
</style>
