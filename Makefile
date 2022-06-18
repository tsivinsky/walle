build:
	go build -o build/walle

install:
	cp build/walle /usr/local/bin/walle
