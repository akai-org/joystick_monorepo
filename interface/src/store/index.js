import { createStore } from 'vuex'
import { keyCodes, keyStates } from '../utils/keyCodes'

export default createStore({
  state: {
  },
  mutations: {
    pressButton () {
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

      commit('pressButton')
    }
  },
  modules: {
  }
})
