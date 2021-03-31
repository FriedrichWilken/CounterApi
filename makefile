test:
	go test -timeout 30s  ./repository
	go test -timeout 30s  
run:
	go run counter
get:
	go get github.com/gin-gonic/gin
	go get github.com/steinfletcher/apitest