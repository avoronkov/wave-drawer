.PHONY: all build install clean

all:
	@if [ "`uname -o`" = "Android" ]; then \
		make android; \
	else \
		make linux; \
	fi

build-android:
	fyne package -os android/arm64 -icon coding.png -name fyne-bezier -release -appID example.example.bezier

linux:
	env PKG_CONFIG_PATH=/usr/lib/pkgconfig go build .

android: build-android
	cp -fT ./fyne_bezier.apk ~/storage/downloads/fyne_bezier.apk

clean:
	rm -f ./fyne_bezier.apk
