# LED cube

This is code for a LED cube, inspired by a LED cube I've seen at
[CCC](https://en.wikipedia.org/wiki/Chaos_Communication_Congress) and the
[SquareWave Dot](https://squarewave.io/) cube.

Hardware:

  * 6 32x32 hub75 panels, for example [these](https://www.aliexpress.com/item/33023473536.html).
  * [Adafruit ItsyBitsy M4](https://www.adafruit.com/product/3800).
  * A power source, for example a USB power bank.
  * Something to make it a cube, for example wood (you can remove the plastic
    frame and use the small screws that remain).
  * Some wires and stuff to connect things internally.


Cube layout:

- bottommost - panel 1 (looking inside, with connector arrows going LR)
- right side - panel 2 (3 oclock)
- bottom side -panel 3 (6 oclock)
- left side - panel 4 (9 oclock)
- top side - panel 5 (12 oclock)
- topmost - panel

Ribbon connectors:

panel 1 out to panel 2 in
panel 2 out to panel 3 in
panel 3 out to panel 4 in
panel 4 out to panel 5 in
panel 5 out to panel 6 in

panel 1 in to controller board 1
panel 6 out to controller board 6

Power connectors:
- Each board needs a power connector plugged in from power octopus
- Power octopus plugged into battery output 1, or wall supply
- MCU plugged into battery output 2, or computer, or wall supply


Ribbon pinout
Pin 1 - R1
Pin 2 - G1
Pin 3 - B1
Pin 4 - GND
Pin 5 - R2
Pin 6 - G2
Pin 7 - B2
Pin 8 - GND
Pin 9 - A
Pin 10 - B
Pin 11 - C
Pin 12 - D
Pin 13 - CLK
Pin 14 - LAT
Pin 15 - OE
Pin 16 - GND

Wires from the ItsyBitsy to the first panel:
SCK -> CLK
MO -> R1
D5 -> LAT
D7 -> OE
D9 -> A
D10 -> B
D11 -> C
D12 -> D

Wires from Panel 6 to Panel 1
R1 -> R2
R2 -> G1
G1 -> G2
G2 -> B1
B1 -> B2

Assembly:

- Place bottommost LED panel face down, on top of some protective pad. Orient LR.
- Plug in controller board ribbon to panel 1 in
- Plug in power/ribbon out to panel 1
- Plug in power/ribbon in to panel 2
- Plug in ribbon out from 2 to ribbon in on panel 3

- Use velcro to position panels 2 and 3 to make a corner
- Plug in power to panel 3 and 4
- Plug in ribbon out from panel 3 to ribbon in on panel 4
- plug in ribbon out on panel 4
- Use velcro to position panels 4 to make a corner

- Plug in power to panel 5
- Plug in ribbon output from panel 4 to input on panel 5
- plug in ribbon output from panel 5, leave hanging
- Use velcro to position panel 5 to complete base

- Plug in power to panel 6
- Plug in ribbon output from panel 5 to input on panel 6
- Plug in controller board ribbon 2 to panel 6 output

- connect power from controller to battery
- connect power from MCU to battery
