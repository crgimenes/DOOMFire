all: build

build:
	CGO_ENABLED=1 go build -a -ldflags "-X 'main.GitTag=$(GIT_TAG)' -extldflags '-s -w'" -o doomfire

clean:
	rm -f doomfire

