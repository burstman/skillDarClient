.PHONY: help build-android build-android-all install-android push-apk clean-android run test logs

# Default target
help:
	@echo "SkillDar Client - Makefile Commands"
	@echo ""
	@echo "Build Commands:"
	@echo "  make build-android       - Build Android APK (ARM64 only)"
	@echo "  make build-android-all   - Build Android APK (all architectures)"
	@echo ""
	@echo "Install Commands:"
	@echo "  make install-android     - Build and install APK on connected device"
	@echo "  make quick-push          - Install existing APK (no rebuild - FAST!)"
	@echo "  make push-apk            - Push APK to device Downloads folder"
	@echo ""
	@echo "Debug Commands:"
	@echo "  make logs                - Show worker loading logs from Android device"
	@echo ""
	@echo "Development Commands:"
	@echo "  make run                - Run the app locally"
	@echo "  make test               - Run tests"
	@echo "  make clean-android      - Clean Android build artifacts"
	@echo ""
	@echo "Quick Commands:"
	@echo "  make android            - Build and install in one command"

# App configuration
APP_ID := com.skilldar.client
APP_ICON := Icon.png
APK_PATH := skillDar.apk
DEVICE_APK_PATH := /data/local/tmp/skillDar.apk

# Build Android APK (ARM64 only - faster, works on most devices)
build-android:
	@echo "Building Android APK with fyne package..."
	fyne package -os android --app-id $(APP_ID) --icon $(APP_ICON)
	@echo "Build complete! APK: $(APK_PATH)"

# Build Android APK for all architectures (ARM64, ARM, AMD64, 386)
build-android-all:
	@echo "Building Android APK with fyne package..."
	fyne package -os android -appID $(APP_ID) -icon $(APP_ICON)
	@echo "Build complete! APK: $(APK_PATH)"

# Quick push - use existing APK without rebuilding
quick-push:
	@echo "Pushing existing APK to device..."
	@if [ ! -f $(APK_PATH) ]; then \
		echo "Error: APK not found. Run 'make build-android' first."; \
		exit 1; \
	fi
	adb push $(APK_PATH) $(DEVICE_APK_PATH)
	@echo "Installing APK on device..."
	adb shell pm install -r $(DEVICE_APK_PATH)
	@echo "Installation complete!"

# Push APK to device /data/local/tmp (system accessible)
push-apk: build-android
	@echo "Pushing APK to device..."
	adb push $(APK_PATH) $(DEVICE_APK_PATH)

# Install APK on connected Android device
install-android: push-apk
	@echo "Installing APK on device..."
	@echo "If Google Play Protect warning appears, tap 'Install anyway'"
	adb shell pm install -r $(DEVICE_APK_PATH)
	@echo "Installation complete!"

# Push APK to Downloads folder for manual installation
push-download: build-android
	@echo "Pushing APK to Downloads folder..."
	adb push $(APK_PATH) /sdcard/Download/skillDarClient.apk
	@echo "APK copied to Downloads. Open Files app on phone to install."

# Quick build and install
android: install-android

# Run locally (desktop)
run:
	@echo "Running SkillDar Client..."
	go run .

# Build locally
build:
	@echo "Building locally..."
	go build -v

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean Android build artifacts
clean-android:
	@echo "Cleaning Android build artifacts..."
	rm -f $(APK_PATH)
	rm -rf fyne-cross/
	@echo "Clean complete!"

# Clean all build artifacts
clean: clean-android
	@echo "Cleaning all build artifacts..."
	rm -f skillDar
	rm -f skillDarClient
	rm -f build.log
	@echo "All clean!"

# Check if device is connected
check-device:
	@echo "Checking for connected devices..."
	@adb devices -l
	@echo ""
	@adb devices | grep -q "device$$" || (echo "ERROR: No device connected!" && exit 1)
	@echo "Device connected ✓"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go install fyne.io/tools/cmd/fyne@latest
	go install github.com/fyne-io/fyne-cross@latest
	@echo "Dependencies installed!"

# Show device info
device-info: check-device
	@echo "Device Information:"
	@echo "-------------------"
	@echo "Model: $$(adb shell getprop ro.product.model)"
	@echo "Android Version: $$(adb shell getprop ro.build.version.release)"
	@echo "SDK Version: $$(adb shell getprop ro.build.version.sdk)"
	@echo "Architecture: $$(adb shell getprop ro.product.cpu.abi)"

# Show worker loading and map logs from Android device
logs:
	@echo "Showing app logs (Press Ctrl+C to stop)..."
	@echo "---------------------------------------------------"
	adb logcat | grep -E "Loading workers|GetWorkers|Failed to load|Successfully loaded|Map|map widget|map area|Creating map|Map configured|Map spacer|Map container|tap"
