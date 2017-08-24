kushell:
	go build -o ${GOBIN}/kushell ku-funcs.go kushell.go

ku:
	go build -o ${GOBIN}/ku ku.go ku-funcs.go
