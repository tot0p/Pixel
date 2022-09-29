# Pixel


## Introduction

Pixel is a esolang, was created by me, and is still in development.

## How to use

### Hello World

```Pixel
128,72,0,255
128,101,0,255
128,108,0,255
128,108,0,255
128,111,0,255
128,32,0,255
128,87,0,255
128,111,0,255
128,114,0,255
128,108,0,255
128,100,0,255
```

now compile it and we are this image (
![Hello World](demo/HelloWorld/HelloWorld.pix.png) )
and you can run it


### Syntaxe

Pixel is a esolang, so it's syntaxe is very simple , there are one type of data , utf-8 code

| R   | G     | B     | A   | Description                                                                                                      |
|-----|-------|-------|-----|------------------------------------------------------------------------------------------------------------------|
| 255 | 0-255 | 0-255 | 255 | Add to Stack the number G * B if G && B != 0 , if B == 0 add G and if G == 0 && B != 0 add b*255                 |
| 128 | 0-255 | 0-255 | 255 | if G && B == 0 Print the la number of the last stack and if G && B != 0 the same logic of the stacking was apply |

