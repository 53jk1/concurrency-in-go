benchmark:
	@echo "Running benchmark..."
	@cd queuing; go test -bench=. -benchmem