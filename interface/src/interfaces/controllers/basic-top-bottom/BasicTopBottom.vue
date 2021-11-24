<template >
  <div>
    <div id="btn-container">
      <button @touchstart="$emit(arrowUpDown)" @touchend="$emit(arrowUpUp)">
          <i class="fas fa-arrow-up"></i>
      </button>
      
      <button @touchstart="$emit(arrowDownDown)" @touchend="$emit(arrowDownUp)">
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

export default {
  name: 'BasicTopBottomInterface',
  emits: [events.arrowUpDown, events.arrowUpUp, events.arrowDownDown, events.arrowDownUp],
  data: function () {
    return {...events}
  }
}
</script>

<style lang="scss" scoped>
  @import './basic-top-bottom.scss'
</style>
