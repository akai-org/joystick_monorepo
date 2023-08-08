<template>
  <div>
    <div class="controller-joystick">
      <div id="joydiv"></div>
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
import { events } from '../../InterfacesEvents'
import { keyActions as keys } from '../../Keys'
import nipplejs from 'nipplejs'

// when phone is rotated to this orientation we lock it
const wantedOrientation = 'portrait'

screen.orientation.addEventListener('change', function (e) {
  const orientation = e.target.type
  if (
    orientation === `${wantedOrientation}-primary` ||
    orientation === `${wantedOrientation}-secondary`
  ) {
    screen.orientation.lock(orientation)
  }
})

export default {
  name: 'JoystickInterface',
  emits: [events.onTouchstart, events.onTouchend, events.onJoystick],
  data: function () {
    return { ...events, keys }
  },
  methods: {
    tryToFullscreen () {
      document.body.requestFullscreen()
    }
  },
  mounted () {
    let lastPayload = 0b00000000
    const dragResolution = 4
    const maxRadius = 75
    // 2 to the power of 6 because we have 6 bits available to send the angle
    const angleResolution = 2 * Math.PI.toPrecision(3) / ((1 << 6) - 1)
    console.log(`angle resolution: ${angleResolution}`)

    var JoystickManager = nipplejs.create({
      zone: document.getElementById('joydiv'),
      threshold: 0.2,
      color: 'gray',
      restJoystick: true,
      mode: 'dynamic',
      size: maxRadius * 2,
      restOpacity: 0.5
    })

    // On joystick direction change
    JoystickManager.on('move', (_, nipple) => {
      const angle = nipple.angle.radian
      const drag = nipple.distance

      const parsedDrag = Math.ceil((drag * dragResolution / maxRadius)) - 1

      // multiplied by 64 to put drag info in two most significant bits
      const payloadDrag = (parsedDrag * 1 << 6)

      const payloadAngle = Math.round(angle / angleResolution)

      const payload = payloadDrag | payloadAngle
      if (payload !== lastPayload) {
        this.$emit(events.onJoystick, 'left', payload)
        lastPayload = payload
      }
    })

    // On joystick direction release
    JoystickManager.on('end', () => {
      const payload = 0b00000000

      this.$emit(events.onJoystick, 'left', payload)
    })
  }
}
</script>
