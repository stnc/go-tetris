Golang ile yapmaya çalıştığım tetris oyunu zaman buldukça geliştirmeye devam edeceğim 

Windows install Opengl 

First gcc install http://www.msys2.org/
pacman -Su
pacman -S base-devel gcc vim cmake
or 
pacman -S mingw-w64-x86_64-gcc  --disable-download-timeout
pacman -S mingw-w64-x86_64-make  --disable-download-timeout
pacman -S mingw-w64-x86_64-gdb  --disable-download-timeout

MSYS2 ve GCC'yi güncelleme için MSYS2 terminaline ara ara şu kodu girip kontrol edin:
pacman -Syu --disable-download-timeout


Video Tutorials 
https://www.youtube.com/watch?v=aeHfqk0cVOE
<pre>
go get github.com/go-gl/gl/v2.1/gl

go get github.com/go-gl/glfw/v3.1/glfw
</pre>

forked by https://github.com/medvednikov/go-tetris


<img src='https://raw.githubusercontent.com/medvednikov/go-tetris/screenshot/screenshot.png' width=250>
