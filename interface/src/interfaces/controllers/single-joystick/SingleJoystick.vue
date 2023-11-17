<template>
    <joystick @on-joystick="(type, payload) => $emit(onJoystick, type, payload)" joystick-type="left" />
</template>

<script>
import { events } from '../../InterfacesEvents'
import Joystick from '../../../components/joystick/Joystick.vue'

export default {
  name: 'SingleJoystickInterface',
  emits: [events.onJoystick],
  data () {
    return { ...events }
  },
  components: {
    Joystick
  }
}
</script>
