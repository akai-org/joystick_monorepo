export const keyCodes = new Map([
  ['ARROW_UP', 0b00000000],
  ['ARROW_DOWN', 0b00000010],
  ['ARROW_RIGHT', 0b00000100],
  ['ARROW_LEFT', 0b00000110],
  ['ACTION_BUTTON_1', 0b00001000],
  ['ACTION_BUTTON_2', 0b00001010],
  ['ACTION_BUTTON_3', 0b00001100],
  ['ACTION_BUTTON_4', 0b00001110]
])

export const keyStates = new Map([
  ['KEY_UP', 0b00000000],
  ['KEY_DOWN', 0b00000001]
])

export const commandType = new Map([['button', 0b00000000]])
