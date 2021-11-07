import { createStore } from 'vuex'
import { keyCodes, keyStates } from '../utils/keyCodes'

export default createStore({
  state: {
    socket: null
  },
  mutations: {
    pressButton () {
    },
    initWebsocketConnectionSuccess (state, socket) {
      state.socket = socket
    },
    websocketConnectionAlreadyEstablished () {

    }
  },
  actions: {
    pressButton ({ commit, state }, { key, keyState }) {
      const keyCode = keyCodes.get(key)
      const kState = keyStates.get(keyState)
      const keyStateMarker = parseInt(kState, 2)
      const endOfStringOffset = 1

      // zero is marking beggining of the string
      const message = keyCode.slice(0, keyCode.length - endOfStringOffset) + keyStateMarker

      state.socket.send(message)

      commit('pressButton')
    },
    // This action doesn't show up in devtools because it's called before devtools attach to the app
    initWebsocketConnection ({ commit, state }) {
      if (!state.socket) {
        const socket = new WebSocket('ws://127.0.0.1:8081/socket')
        commit('initWebsocketConnectionSuccess', socket)
      } else {
        commit('websocketConnectionAlreadyEstablished')
      }
    }
  },
  modules: {
  }
})
