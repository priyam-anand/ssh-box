BINARY_NAME = ssh_box
BUILD_DIR = ./build
INSTALL_DIR = /usr/local/bin

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

run:
	go run main.go

install:
	install -m 0755 $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)

clean:
	rm -rf $(BUILD_DIR)

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

