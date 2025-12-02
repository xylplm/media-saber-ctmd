@echo off
chcp 65001 >nul
REM TMDB Fetcher 交叉编译脚本 (Windows版本)
REM 编译多个平台的可执行文件

echo 开始编译 TMDB Fetcher...

REM Output directory (save to ../cli)
set OUTPUT_DIR=..\cli
if not exist %OUTPUT_DIR% mkdir %OUTPUT_DIR%

REM 项目名称
set APP_NAME=tmdb-fetcher

REM Windows AMD64
echo 编译 Windows AMD64...
set GOOS=windows
set GOARCH=amd64
go build -o "%OUTPUT_DIR%\%APP_NAME%-windows-amd64.exe" tmdb_fetcher.go

REM Windows ARM64
echo 编译 Windows ARM64...
set GOOS=windows
set GOARCH=arm64
go build -o "%OUTPUT_DIR%\%APP_NAME%-windows-arm64.exe" tmdb_fetcher.go

REM Linux AMD64
echo 编译 Linux AMD64...
set GOOS=linux
set GOARCH=amd64
go build -o "%OUTPUT_DIR%\%APP_NAME%-linux-amd64" tmdb_fetcher.go

REM Linux ARM64
echo 编译 Linux ARM64...
set GOOS=linux
set GOARCH=arm64
go build -o "%OUTPUT_DIR%\%APP_NAME%-linux-arm64" tmdb_fetcher.go

REM macOS AMD64 (Intel)
echo 编译 macOS AMD64...
set GOOS=darwin
set GOARCH=amd64
go build -o "%OUTPUT_DIR%\%APP_NAME%-macos-amd64" tmdb_fetcher.go

REM macOS ARM64 (Apple Silicon)
echo 编译 macOS ARM64...
set GOOS=darwin
set GOARCH=arm64
go build -o "%OUTPUT_DIR%\%APP_NAME%-macos-arm64" tmdb_fetcher.go

echo.
echo ✓ 编译完成！
echo 可执行文件已保存到: %OUTPUT_DIR%
echo.
dir %OUTPUT_DIR%

pause
