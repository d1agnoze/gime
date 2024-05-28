dev:
	wails dev --appargs="mime=image/png,file=input.txt"
build:
	wails build --windowsconsole --tags "desktop production"
