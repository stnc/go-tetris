<img src='https://raw.githubusercontent.com/medvednikov/go-tetris/screenshot/screenshot.png' width=250>




I will continue to improve the tetris game I am working on with Golang as time goes.

Golang ile yapmaya çalıştığım tetris oyunu zaman buldukça geliştirmeye devam edeceğim 

## Windows install Opengl 

First gcc install http://www.msys2.org/

pacman -Su

pacman -S  gcc cmake

pacman -S base-devel gcc cmake

## or 

pacman -S mingw-w64-x86_64-gcc  --disable-download-timeout

pacman -S mingw-w64-x86_64-make  --disable-download-timeout

pacman -S mingw-w64-x86_64-gdb  --disable-download-timeout

MSYS2 ve GCC'yi güncelleme için MSYS2 terminaline ara ara şu kodu girip kontrol edin:

pacman -Syu --disable-download-timeout

make -f Makefile


Video Tutorials 
https://www.youtube.com/watch?v=aeHfqk0cVOE

https://stackoverflow.com/questions/30069830/how-to-install-mingw-w64-and-msys2#30071634

<pre>
go get github.com/go-gl/gl/v2.1/gl

go get github.com/go-gl/glfw/v3.1/glfw
</pre>

forked by https://github.com/medvednikov/go-tetris


