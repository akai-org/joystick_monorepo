import { keyCodes, keyStates as kStates } from '../utils/keyCodes.js'

export const keyActions = {}

for (const [key] of keyCodes) {
  keyActions[key] = key
}

for (const [key] of kStates) {
  keyActions[key] = key
}
