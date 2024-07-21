#!/bin/bash

# Определение целевых платформ в формате GOOS_GOARCH
targets=("darwin_amd64" "darwin_arm64" "windows_amd64" "linux_amd64")

# Цикл для компиляции под каждую целевую платформу
for target in "${targets[@]}"; do
    # Разделение строки target на GOOS и GOARCH
    IFS='_' read -r GOOS GOARCH <<< "$target"

    # Установка имени выходного файла
    output="clicker-${GOOS}-${GOARCH}"
    [ "$GOOS" == "windows" ] && output+=".exe"

    # Выполнение команды компиляции
    echo "Building for $GOOS ($GOARCH)..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o "$output" || { echo "Failed to build for $GOOS ($GOARCH)"; exit 1; }
    echo "Done."
done