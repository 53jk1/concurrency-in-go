#run benchmark from queuing package
benchmark:
	@echo "Running benchmark..."
	# run from queuing package
	@cd queuing; go test -bench=. -benchmem