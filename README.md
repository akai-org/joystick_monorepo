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
    "global_id": "someglobalid"
}
```

b. odpowiedź w przypadku błędu w przesłanych danych:
```
kod 400 user error
{
    message: "jakaś wiadomość co poszło nie tak"
}
```
---
#### Połączenie się gry

1) **Host gry**:
podłączenie się do ws, wysłanie wiadomości autoryzacyjnej:
```
WS: /game/socket
{"global_id": "someglobalid"}
```
wiadomość musi być typu tekstowego, musi być jsonem i
zawierać taką treść, w miejscu someglobalid powinien być przesłany
global_id uzyskany podczas rejestracji gry. Jeżeli wiadomość będzie inna
lub niepoprawna to serwer zamknie połączenie ws i trzeba próbować jeszcze raz.

2) **Serwer** - jeżeli wiadomość jest poprawna to utrzymuje połączenie.
Kolejne wiadomości są wysyłane przez serwer:
- W formacie binarnym przyjmuje
informacje o akcjach graczy, tj. 16 bitowe informacje, z czego pierwsze
8 bitow jest dla identyfikacji gracza, a drugie dla identyfikacji akcji.
8-bitowy identyfikator gracza jest unikatowy w obrębie jednej gry.
- W formacie tekstowym  przyjmuje informacje o wyjątkowych zdarzeniach. 
  - dołączenie gracza
    ```json
        {"event":  "player_joined", "id":  22}
    ```
  - odejście gracza
    ```json
        {"event":  "player_left", "id":  22}
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
    "gui": "nazwa interfejsu",
    "global_id": "someglobalidforplayer"
}
```

b. odpowiedź w przypadku błędu w przesłanych danych:
```
kod 400 user error
{
    message: "jakaś wiadomość co poszło nie tak"
}
```
---
#### Połączenie się gracza 

1) **Gracz**:
podłączenie się do ws, wysłanie wiadomości autoryzacyjnej:
```
WS: /player/socket
{"global_id": "someglobalid"}
```
wiadomość musi być typu tekstowego, musi być jsonem i 
zawierać taką treść, w miejscu someglobalid powinien być przesłany
global_id uzyskany podczas rejestracji gracza. Jeżeli wiadomość będzie inna
lub niepoprawna to serwer zamknie połączenie ws i trzeba próbować jeszcze raz.

2) **Serwer**:
jeżeli wiadomość jest poprawna to utrzymuje połączenie. 
Kolejne wysyłane przez gracza wiadomości muszą być typu binarnego,
zawierać informację nt przycisku.