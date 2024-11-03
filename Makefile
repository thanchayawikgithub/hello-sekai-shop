# Define Go executable and main file
GO := go
MAIN_FILE := main.go

# Define state
STATE := dev

# Define services
SERVICES := auth inventory item payment player

# Default target
.PHONY: all
all: $(SERVICES)

# Rule for each service
.PHONY: $(SERVICES)
$(SERVICES):
	@echo "Running service $@ in $(STATE) state..."
	$(GO) run $(MAIN_FILE) $(STATE) $@
