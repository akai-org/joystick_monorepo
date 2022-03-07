import { createStore } from 'vuex'
import { keyCodes, keyStates } from '../utils/keyCodes'
import { onOpenHandler } from '../utils/socketHelpers'
import { PRESS_BUTTON, SAVE_INTERFACE } from './actions'

export default createStore({
  state: {
    socket: null,
    gui: null
  },
  mutations: {
    pressButton () {
    },
    initWebsocketConnectionSuccess (state, socket) {
      state.socket = socket
      socket.addEventListener('open', onOpenHandler)
    },
    websocketConnectionAlreadyEstablished () {

    },
    saveInterface (state, gui) {
      state.gui = gui
    }
  },
  actions: {
    pressButton ({ commit, state }, { key, keyState }) {
      const keyCode = keyCodes.get(key)
      const kState = keyStates.get(keyState)

      const message = (keyCode | kState)
      const payload = new Uint8Array(1)
      payload[0] = message
      state.socket.send(payload.buffer)

      commit(PRESS_BUTTON)
    },
    // This action doesn't show up in devtools because it's called before devtools attach to the app
    initWebsocketConnection ({ commit, state }) {
      if (!state.socket) {
        const socket = new WebSocket('ws://192.168.1.104:8081/player/socket')
        commit('initWebsocketConnectionSuccess', socket)
      } else {
        commit('websocketConnectionAlreadyEstablished')
      }
    },
    saveInterface ({ commit }, gui) {
      commit(SAVE_INTERFACE, gui)
    }
  },
  modules: {
  }
})
