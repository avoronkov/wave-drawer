.PHONY: all build install clean

all: android

build-android:
	fyne package -os android/arm64 -icon coding.png -name fyneOnTermux -release -appID example.example.example

linux:
	env PKG_CONFIG_PATH=/usr/lib/pkgconfig go build .

android: build-android
	cp -fT ./fyneOnTermux.apk ~/storage/downloads/fyneOnTermux.apk

clean:
	rm -f ./fyneOnTermux.apk
