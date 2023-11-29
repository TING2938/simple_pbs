

qsub:
	go run cmd/qsub/main.go -task test -cmd ". tests/echo.sh > output.dat"

qdel:
	go run cmd/qdel/dele.go -task test