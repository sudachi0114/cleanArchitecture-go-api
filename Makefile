.PHONY: run
run:
	go run src/app/main.go

.PHONY: compose/up
compose/up: compose/down
	docker-compose up -d

.PHONY: compose/down
compose/down:
	docker-compose down

.PHONY: clean
clean: compose/down
	rm main