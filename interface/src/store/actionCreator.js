import store from './index'
import { PRESS_BUTTON, SAVE_INTERFACE } from './actions'

export const dispatchPressButton = ({ key, keyState }) => {
  store.dispatch(PRESS_BUTTON, { key, keyState })
}

export const dispatchSaveInterface = (gui) => {
  store.dispatch(SAVE_INTERFACE, gui)
}
