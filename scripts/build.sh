#!/bin/bash

# TMDB Manager 交叉编译脚本
# 编译多个平台的可执行文件

echo "开始编译 TMDB Manager..."

# Output directory (save to ../cli)
OUTPUT_DIR="../cli"
mkdir -p $OUTPUT_DIR

# 项目名称
APP_NAME="tmdb-manager"

# Windows AMD64
echo "编译 Windows AMD64..."
GOOS=windows GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-windows-amd64.exe" tmdb_manager.go

# Windows ARM64
echo "编译 Windows ARM64..."
GOOS=windows GOARCH=arm64 go build -o "${OUTPUT_DIR}/${APP_NAME}-windows-arm64.exe" tmdb_manager.go

# Linux AMD64
echo "编译 Linux AMD64..."
GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-linux-amd64" tmdb_manager.go

# Linux ARM64
echo "编译 Linux ARM64..."
GOOS=linux GOARCH=arm64 go build -o "${OUTPUT_DIR}/${APP_NAME}-linux-arm64" tmdb_manager.go

# macOS AMD64 (Intel)
echo "编译 macOS AMD64..."
GOOS=darwin GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-macos-amd64" tmdb_manager.go

# macOS ARM64 (Apple Silicon)
echo "编译 macOS ARM64..."
GOOS=darwin GOARCH=arm64 go build -o "${OUTPUT_DIR}/${APP_NAME}-macos-arm64" tmdb_manager.go

echo ""
echo "✓ 编译完成！"
echo "可执行文件已保存到: ${OUTPUT_DIR}"
echo ""
ls -lh ${OUTPUT_DIR}
