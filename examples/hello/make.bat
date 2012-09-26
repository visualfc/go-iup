windres rc/goiup.rc -o temp-rc.o

set GOIUPPKG=../../../../../../pkg/windows_386
set SRC=main.go
set OUT=hello

go tool 8g -I%GOIUPPKG% %SRC%

go tool pack grc _go_.8 main.8 temp-rc.o

go tool 8l -L%GOIUPPKG% -s -Hwindowsgui -o %OUT%.exe _go_.8

rm *.8 *.o