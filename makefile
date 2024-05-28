TAGS = --tags "desktop production"
DEV_ARGS_1 = ---appargs="-mime=video/x-youtube --file=input2.txt"
DEV_ARGS_2 = ---appargs="-mime=image/jpeg --file=input.txt"

.PHONY: dev build build_win build_linux

dev:
	wails dev $(DEV_ARGS_2)
build:
	build_win build_linux
build_win:
	wails build --windowsconsole $(TAGS) --platform "windows"
build_linux:
	wails build $(TAGS)  --platform "linux"
