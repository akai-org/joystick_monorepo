<template src="./basic-top-bottom.html" ></template>

<script>

import { events } from '../../InterfacesEvents'

export default {
  name: 'BasicTopBottomInterface',
  emits: [events.arrowUpDown, events.arrowUpUp, events.arrowDownDown, events.arrowDownUp]
}
</script>

<style lang="scss" scoped>
  @import './basic-top-bottom.scss'
</style>
