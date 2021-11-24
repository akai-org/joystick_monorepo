<template src="./basic-left-right.html"></template>

<script>

import { events } from '../../InterfacesEvents'

export default {
  name: 'BasicLeftRightInterface',
  emits: [events.arrowLeftDown, events.arrowLeftUp, events.arrowRightDown, events.arrowRightUp]
}
</script>

<style lang="scss" scoped>
  @import './basic-left-right.scss'
</style>
