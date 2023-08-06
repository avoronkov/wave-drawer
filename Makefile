.PHONY: all build install clean

all: install

build:
	fyne package -os android/arm64 -icon coding.png -name fyneOnTermux -release -appID example.example.example

install: build
	cp -fT ./fyneOnTermux.apk ~/storage/downloads/fyneOnTermux.apk

clean:
	rm -f ./fyneOnTermux.apk
