@echo off
REM TMDB Manager 交叉编译脚本 (Windows版本)
REM 编译多个平台的可执行文件

REM Output directory (save to ../cli)
set OUTPUT_DIR=..\cli
if not exist %OUTPUT_DIR% mkdir %OUTPUT_DIR%

REM 项目名称
set APP_NAME=tmdb-manager

REM Windows AMD64
echo Compiling Windows AMD64...
set GOOS=windows
set GOARCH=amd64
go build -o "%OUTPUT_DIR%\%APP_NAME%-windows-amd64.exe" tmdb_manager.go

REM Windows ARM64
echo Compiling Windows ARM64...
set GOOS=windows
set GOARCH=arm64
go build -o "%OUTPUT_DIR%\%APP_NAME%-windows-arm64.exe" tmdb_manager.go

REM Linux AMD64
echo Compiling Linux AMD64...
set GOOS=linux
set GOARCH=amd64
go build -o "%OUTPUT_DIR%\%APP_NAME%-linux-amd64" tmdb_manager.go

REM Linux ARM64
echo Compiling Linux ARM64...
set GOOS=linux
set GOARCH=arm64
go build -o "%OUTPUT_DIR%\%APP_NAME%-linux-arm64" tmdb_manager.go

REM macOS AMD64 (Intel)
echo Compiling macOS AMD64...
set GOOS=darwin
set GOARCH=amd64
go build -o "%OUTPUT_DIR%\%APP_NAME%-macos-amd64" tmdb_manager.go

REM macOS ARM64 (Apple Silicon)
echo Compiling macOS ARM64...
set GOOS=darwin
set GOARCH=arm64
go build -o "%OUTPUT_DIR%\%APP_NAME%-macos-arm64" tmdb_manager.go

echo.
echo Build completed!
echo Executables saved to: %OUTPUT_DIR%
echo.
dir %OUTPUT_DIR%

pause
