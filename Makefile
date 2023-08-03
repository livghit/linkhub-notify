
build:
	cd ./cmd/linkhub-notify && go build -b ../../bin/websocketServer.exe

run:
	cd ./bin & ./websocketServer.exe
	
