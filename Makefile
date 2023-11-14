camera:
	go run ./demo/showvideo/main.go 1

clean:
	rm -rf build

builddir:
	mkdir -p build

build/hellolife-tinygo:
	tinygo build -o ./build/hellolife-tinygo ./demo/hellolife/

hellolife-tinygo: builddir build/hellolife-tinygo
	./build/hellolife-tinygo

build/hellolife-go:
	go build -o ./build/hellolife-go ./demo/hellolife/

hellolife-biggo	: builddir build/hellolife-go
	./build/hellolife-go

hellolife-size:
	ls -l ./build/hellolife-tinygo ./build/hellolife-go

hellolife: hellolife-biggo

badgelife:
	tinygo flash -size short -target=gopher-badge ./demo/badgelife/

cubetest:
	tinygo flash -size short -target=itsybitsy-m4 -opt=2 ./demo/cubetest/

panellife:
	tinygo flash -size short -target=itsybitsy-m4 -opt=2 ./demo/panellife/

cubelife:
	tinygo flash -size short -target=itsybitsy-m4 -opt=2 ./demo/cubelife/

showlife:
	tinygo flash -size short -target=itsybitsy-m4 -opt=2 ./demo/showlife/
