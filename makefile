# MAKEFILE FOR GIME
# =======================================================
# Prerequisite: 
# 	Wails CLI, upx, Go 1.20+, NPM (node 15+)
# 	Run wails doctor to check for missing dependencies
# =======================================================
TAGS = --tags "desktop production"
DEV_ARGS_1 = --appargs="--mime=video/x-youtube --file=input2.txt"
DEV_ARGS_2 = --appargs="--mime=image/jpeg --file=input.txt"
DEV_ARGS_3 = --appargs="--vid --file=input2.txt"
DEV_ARGS_4 = --appargs="--img --file=input.txt -p BL"

.PHONY: dev build build_win build_linux

dev:
	wails dev $(DEV_ARGS_4)
build:
	build_win build_linux
build_win:
	wails build --windowsconsole $(TAGS) --platform "windows" --upx
build_linux:
	wails build $(TAGS)  --platform "linux" --upx
