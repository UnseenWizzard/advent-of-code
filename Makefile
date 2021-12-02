DATE=$$(date --iso)

check:
	go fmt ./...
	go vet ./...
	golint ./...

test:
	go test ./...

run-day:
	go run riedmann.dev/aoc21/$(DATE) $(auth)
	go run riedmann.dev/aoc21/$(DATE)-part2 $(auth)

folder:
	mkdir $(DATE)
	cp solution.template $(DATE)/solution.go
	echo "const day = \"$$(date '+%-d')\"" >> $(DATE)/solution.go
	cp solution_test.template $(DATE)/solution_test.go

part2:
	cp -r $(DATE)/ $(DATE)-part2/