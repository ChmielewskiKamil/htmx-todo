test:
	@echo "[INFO] Generating the templates..."
	templ generate
	@echo "[INFO] Running the tests..."
	gotestsum --format testname

run:
	@echo "[INFO] Generating the templates..."
	templ generate
	@echo "[INFO] Building the app..."
	go build -o bin/
	@echo "[INFO] Launching the server..."
	./bin/htmx-todo
