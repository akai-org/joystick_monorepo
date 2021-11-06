# Projekt Joystick

## Przyciski

### Lista przycisków i ich reprezentacja binarna:
```
- ARROW_UP        = 00000000
- ARROW_DOWN      = 00000010
- ARROW_RIGHT     = 00000100
- ARROW_LEFT      = 00000110
- ACTION_BUTTON_1 = 00001000
- ACTION_BUTTON_2 = 00001010
- ACTION_BUTTON_3 = 00001100
- ACTION_BUTTON_4 = 00001110
```

### Zmiany stanu przycisku
```
- KEY_UP          = 00000000
- KEY_DOWN        = 00000001
```

### Kompozycja informacji nt. przycisków (przykład):
`KEY_INFO = KEY_DOWN | ARROW_UP`
