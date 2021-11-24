import store from './index'
import { PRESS_BUTTON } from './actions'

export const dispatchPressButton = ({ key, keyState }) => {
  store.dispatch(PRESS_BUTTON, { key, keyState })
}
