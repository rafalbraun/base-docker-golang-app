## https://unix.stackexchange.com/questions/656437/export-env-variable-does-not-work-from-makefile
export DATABASE_HOST := 127.0.0.1

run:
	cd webserver-app && go run webserver.go

