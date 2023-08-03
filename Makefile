
build:
	cd ./cmd/linkhub-notify && go build -o ../../bin/websocketServer.exe

run:
	cd bin && ./websocketServer.exe
	
