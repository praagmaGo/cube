
** Animated cube icon generator **

This program generates animated pictures with a 3D cube moving inside. Looking at the code, it's easy to change size, and moving speeds in every three directions.

This was initially coded in 2015. This code is renamed original.go
It uses the library "code.google.com/p/draw2d/draw2d" which seems no more up to date.

I've just upgraded the code with a new similar library:  "github.com/llgcode/draw2d/draw2dimg"
A few changes have been made in the code to adapt.

To produce an animated gif from all the ouput files on linux:

```
ffmpeg -f image2 -framerate 10 -i cube%04d.png -loop 0 output.gif
```

I've taken the ouput test16.gif and renamed it to favicon.png and placed it on a local copy of OpenJSCAD.org . On firefox on windows, the icon is animated. All the others are not.

The original idea for this code, was to be able to generate a custom icon for every page request. Colors or turning speeds would include IP and time information.
