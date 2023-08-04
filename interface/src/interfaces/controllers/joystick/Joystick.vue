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

function getDirection (radians) {
  const directions = [
    'left',
    'down-left',
    'down',
    'down-right',
    'right',
    'up-right',
    'up',
    'up-left'
  ]
  const index = Math.round((radians + Math.PI) / (Math.PI / 4)) % 8
  return directions[index]
}

function getKeyFromDir (directory) {
  let key
  switch (directory) {
    case 'up':
      key = [keys.ARROW_UP]
      break
    case 'down':
      key = [keys.ARROW_DOWN]
      break
    case 'left':
      key = [keys.ARROW_LEFT]
      break
    case 'right':
      key = [keys.ARROW_RIGHT]
      break
    case 'up-left':
      key = [keys.ARROW_UP, keys.ARROW_LEFT]
      break
    case 'up-right':
      key = [keys.ARROW_UP, keys.ARROW_RIGHT]
      break
    case 'down-left':
      key = [keys.ARROW_DOWN, keys.ARROW_LEFT]
      break
    case 'down-right':
      key = [keys.ARROW_DOWN, keys.ARROW_RIGHT]
      break
  }
  return key
}

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
  emits: [events.onTouchstart, events.onTouchend],
  data: function () {
    return { ...events, keys }
  },
  methods: {
    tryToFullscreen () {
      document.body.requestFullscreen()
    },

    compareDirectories (lastDir, newDir) {
      if (!lastDir || !newDir) return false

      if (lastDir.length !== newDir.length) return false

      return lastDir.every((d, i) => d === newDir[i])
    },

    move (dir) {
      console.log('move', dir)
      dir.forEach((d) => {
        this.$emit(events.onTouchstart, d)
      })
    },
    stop (dir) {
      dir.forEach((d) => {
        this.$emit(events.onTouchend, d)
      })
    }
  },
  mounted () {
    // Last joystick direction
    let lastDirectory

    const JoystickManager = nipplejs.create({
      zone: document.getElementById('joydiv'),
      threshold: 0.2,
      color: 'gray',
      restJoystick: true,
      mode: 'dynamic',
      size: 150,
      restOpacity: 0.5
    })

    // On joystick direction change
    JoystickManager.on('move', (_, nipple) => {
      const newDir = getDirection(nipple.angle.radian)
      const keyToPress = getKeyFromDir(newDir)

      if (this.compareDirectories(lastDirectory, keyToPress)) return

      if (lastDirectory) {
        this.stop(lastDirectory)
      }

      this.move(keyToPress)
      lastDirectory = keyToPress
    })

    // On joystick direction release
    JoystickManager.on('end', () => {
      this.stop(lastDirectory)
      lastDirectory = null
    })
  }
}
</script>
