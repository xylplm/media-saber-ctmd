# Scripts 源代码目录

本目录包含 TMDB Manager 的源代码，仅供开发者使用。

## 📂 文件说明

- `tmdb_manager.go` - Go 程序源代码
- `go.mod` - Go 模块配置
- `build.bat` - Windows 交叉编译脚本
- `build.sh` - Linux/macOS 交叉编译脚本

## 🔨 编译

如需编译工具，首先安装 [Go 1.21+](https://golang.org/dl/)

**Windows:**
```bash
.\build.bat
```

**Linux/macOS:**
```bash
chmod +x build.sh
./build.sh
```

编译完成后的可执行文件会自动保存到项目根目录的 `cli/` 目录。

## 💡 本地开发

```bash
# 直接运行（需要在 cli 目录有 config.json）
go run tmdb_manager.go

# 单平台编译
go build -o tmdb-manager tmdb_manager.go
```

## 📝 修改编译输出目录

编译脚本中的 `OUTPUT_DIR` 默认指向 `../../cli`。如需修改，编辑对应脚本文件。

## 🔗 更多信息

- [用户文档](../README.md)
- [CLI 工具说明](../cli/README.md)
