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

function getKeyFromDir (dir) {
  let k
  switch (dir) {
    case 'up':
      k = keys.ARROW_UP
      break
    case 'down':
      k = keys.ARROW_DOWN
      break
    case 'left':
      k = keys.ARROW_LEFT
      break
    case 'right':
      k = keys.ARROW_RIGHT
      break
  }
  return k
}

// when phone is rotated to this orientation we lock it
const wantedOrientation = 'portrait'

screen.orientation.addEventListener('change', function (e) {
  const orientation = e.target.type
  if (orientation === `${wantedOrientation}-primary` || orientation === `${wantedOrientation}-secondary`) {
    screen.orientation.lock(orientation)
  }
})

export default {
  name: 'JoystickInterface',
  emits: [events.onTouchstart, events.onTouchend],
  data: function () {
    return { ...events, keys }
  },
  methods: {
    tryToFullscreen () {
      document.body.requestFullscreen()
    },
    stopLastKey (key) {
      this.$emit(events.onTouchend, key)
    },
    startLastKey (key) {
      this.$emit(events.onTouchstart, key)
    }
  },
  mounted () {
    // Last joystick direction
    let lastDir

    const comp = this

    var JoystickManager = nipplejs.create({
      zone: document.getElementById('joydiv'),
      threshold: 0.2,
      color: 'gray',
      restJoystick: true,
      mode: 'static',
      position: { left: '50%', bottom: '150px' },
      size: 150,
      restOpacity: 0.5
    })

    // On joystick direction change
    JoystickManager.on('dir', function (event, nipple) {
      const newDir = nipple.direction.angle
      const keyToPress = getKeyFromDir(newDir)

      if (lastDir) {
        comp.stopLastKey(lastDir)
      }

      comp.startLastKey(keyToPress)
      lastDir = keyToPress
    })
  }
}
</script>
