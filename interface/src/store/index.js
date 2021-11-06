import { createStore } from 'vuex'

export default createStore({
  state: {
  },
  mutations: {
    pressButton () {
      console.log('button pressed')
    }
  },
  actions: {
    pressButton ({ commit, state }, key) {
      commit('pressButton')
    }
  },
  modules: {
  }
})
