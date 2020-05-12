bin: bin/webserver bin/pyaml

install: clean bin
	@echo ---- installing pyaml on your system ----
	sudo cp bin/pyaml /usr/local/bin/pyaml

bin/webserver:
	@echo ---- building cmd/webserver/main.go binary ----
	go build -o bin/webserver cmd/blog/main.go

bin/pyaml:
	@echo ---- building cmd/pyaml/main.go binary ----
	go build -o bin/pyaml cmd/pyaml/main.go

test:
	@echo ---- performing test ----
	go test internal/app/blog/*.go -v
	go test internal/app/pyaml/*.go -v

clean:
	@echo ---- cleaning bin directory ----
	rm -fvr bin

server: bin
	@echo ---- running web server ----
	bash server.sh

