all: compile fmt build

fmt:
	find . -type f -iregex '.*\.go' -exec gofmt -l -s -w '{}' +

build:
	go generate ./...
	go build ./...

test:
	go test ./...

cover:
	go test ./... --cover

bench:
	go test -benchmem -bench=. ./internal/php5
	go test -benchmem -bench=. ./internal/php7

# wall-clock parse of a corpus, e.g. `make bench-corpus DIR=./laravel`
bench-corpus:
	go run ./benchmark $(DIR)

compile: ./internal/php5/php5.go ./internal/php7/php7.go ./internal/php8/php8.go ./internal/php8/scanner.go ./internal/scanner/scanner.go
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./internal/php5/php5.go
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./internal/php7/php7.go
	sed -i '' -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./internal/php8/php8.go
	sed -i '' -e 's/\/\/line/\/\/ line/g' ./internal/php5/php5.go
	sed -i '' -e 's/\/\/line/\/\/ line/g' ./internal/php7/php7.go
	sed -i '' -e 's/\/\/line/\/\/ line/g' ./internal/php8/php8.go
	sed -i '' -e 's/\/\/line/\/\/ line/g' ./internal/scanner/scanner.go
	sed -i '' -e 's/\/\/line/\/\/ line/g' ./internal/php8/scanner.go
	rm -f y.output

./internal/scanner/scanner.go: ./internal/scanner/scanner.rl
	ragel -Z -G2 -o $@ $<

./internal/php5/php5.go: ./internal/php5/php5.y
	goyacc -o $@ $<

./internal/php7/php7.go: ./internal/php7/php7.y
	goyacc -o $@ $<

./internal/php8/php8.go: ./internal/php8/php8.y
	goyacc -o $@ $<

./internal/php8/scanner.go: ./internal/php8/scanner.rl
	ragel -Z -G2 -o $@ $<
