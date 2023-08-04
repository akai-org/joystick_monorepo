import { createStore } from 'vuex'
import { commandType, keyCodes, keyStates } from '../utils/keyCodes'
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
      socket.addEventListener('error', (event) => {
        console.log(event)
        alert('Connection to server failed')
      })
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

      const action = (keyCode | kState)
      const payload = new Uint8Array(2)
      payload[0] = commandType.get('button')
      payload[1] = action
      state.socket.send(payload.buffer)

      commit(PRESS_BUTTON)
    },
    // This action doesn't show up in devtools because it's called before devtools attach to the app
    initWebsocketConnection ({ commit, state }) {
      if (!state.socket) {
        const socket = new WebSocket(process.env.VUE_APP_SOCKET_ADDRESS)
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
