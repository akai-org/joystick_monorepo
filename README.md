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


## Umowy w komunikacji

### Serwer <-> host gry

#### Tworzenie nowej gry

1) **Host gry** - prośba o utworzenie nowego pokoju gry:
```
POST: http://servergo/create
{
    "interface": "nazwa interfejsu",
    "max_players": <maksymalna liczba graczy>(int)
}
```
2) **Serwer**:

a. odpowiedź w przypadku pomyślnego utworzenia pokoju:
```
kod 200 ok
{
    "room_code": "kod pokoju",
    "address": "http://servergo/conns/unikatowy_uri_do_komunikacji_przez_ws
}
```

b. odpowiedź w przypadku błędu w przesłanych danych:
```
kod 400 user error
{
    message: "jakaś wiadomość co poszło nie tak"
}
```

### Serwer <-> joystick

#### Rejestracja nowego gracza

1) **Gracz** - prośba o przypisanie do gry:
```
POST: http://servergo/join
{
    "nickname": "nick gracza",
    "room_code": "kod pokoju"
}
```

2) **Serwer**:

a. odpowiedź w przypadku pomyślnego zarejestrowania gracza:
```
kod 200 ok
{
    "interface": "nazwa interfejsu",
    "address": "http://servergo/conns/unikatowy_uri_do_komunikacji_przez_ws"
}
```

b. odpowiedź w przypadku błędu w przesłanych danych:
```
kod 400 user error
{
    message: "jakaś wiadomość co poszło nie tak"
}
```