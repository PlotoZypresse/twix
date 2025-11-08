.PHONY: run build install clean

run:
	go run . test_images

build:
	go build twix.go

install:
	go install 

clean:
	rm -f twix