import store from './index'
import { JOYSTICK_MOVE, PRESS_BUTTON, SAVE_INTERFACE } from './actions'

export const dispatchPressButton = ({ key, keyState }) => {
  store.dispatch(PRESS_BUTTON, { key, keyState })
}

export const dispatchSaveInterface = (gui) => {
  store.dispatch(SAVE_INTERFACE, gui)
}

export const dispatchJoystickMove = ({ joystickType, payload }) => {
  store.dispatch(JOYSTICK_MOVE, { joystickType, payload })
}
