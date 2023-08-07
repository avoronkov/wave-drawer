.PHONY: all build install clean

all: install

build:
	fyne package -os android/arm64 -icon coding.png -name fyneOnTermux -release -appID example.example.example

build-linux:
	env PKG_CONFIG_PATH=/usr/lib/pkgconfig go build .

install: build
	cp -fT ./fyneOnTermux.apk ~/storage/downloads/fyneOnTermux.apk

clean:
	rm -f ./fyneOnTermux.apk
