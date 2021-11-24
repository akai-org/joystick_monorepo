import store from './index'

export const dispatchPressButton = ({ key, keyState }) => {
  store.dispatch('pressButton', { key, keyState })
}
