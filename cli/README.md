# 📦 TMDB Fetcher 命令行工具

这里存放预编译的 TMDB 数据获取工具，**无需安装任何环境**，下载即用！

## ✨ 特点

- ✅ **零依赖** - 无需安装 Python、Go 等运行时环境
- ✅ **跨平台** - 支持 Windows、Linux、macOS
- ✅ **多架构** - 支持 Intel/AMD 和 ARM 处理器
- ✅ **单文件** - 独立可执行，拷贝即用
- ✅ **快速** - 原生编译，启动迅速

## 🚀 使用方法

### 第一步：配置 API Key

```bash
cd cli
copy config.example.json config.json
```

编辑 `cli/config.json` 文件，填入你的 TMDB API Key。

在 [TMDB 设置页面](https://www.themoviedb.org/settings/api) 申请 API Key。

### 第二步：运行工具

根据你的操作系统选择对应的可执行文件：

#### Windows

```bash
# Intel/AMD 处理器（64位）
.\cli\tmdb-fetcher-windows-amd64.exe

# ARM 处理器（如 Surface Pro X）
.\cli\tmdb-fetcher-windows-arm64.exe
```

💡 **提示**: 可以直接双击运行 `.exe` 文件

#### Linux

```bash
# Intel/AMD 处理器
chmod +x ./cli/tmdb-fetcher-linux-amd64
./cli/tmdb-fetcher-linux-amd64

# ARM 处理器（如树莓派）
chmod +x ./cli/tmdb-fetcher-linux-arm64
./cli/tmdb-fetcher-linux-arm64
```

#### macOS

```bash
# Intel 芯片（2020年及之前的 Mac）
chmod +x ./cli/tmdb-fetcher-macos-amd64
./cli/tmdb-fetcher-macos-amd64

# Apple Silicon 芯片（M1/M2/M3）
chmod +x ./cli/tmdb-fetcher-macos-arm64
./cli/tmdb-fetcher-macos-arm64
```

### 第三步：按提示操作

1. 选择媒体类型（电影或电视剧）
2. 输入 TMDB ID
3. 数据会自动保存到 `tmdb_config/` 目录

## 📋 可用文件

| 文件名 | 平台 | 架构 | 文件大小 |
|--------|------|------|----------|
| `tmdb-fetcher-windows-amd64.exe` | Windows | Intel/AMD 64位 | ~8MB |
| `tmdb-fetcher-windows-arm64.exe` | Windows | ARM 64位 | ~8MB |
| `tmdb-fetcher-linux-amd64` | Linux | Intel/AMD 64位 | ~8MB |
| `tmdb-fetcher-linux-arm64` | Linux | ARM 64位 | ~8MB |
| `tmdb-fetcher-macos-amd64` | macOS | Intel | ~8MB |
| `tmdb-fetcher-macos-arm64` | macOS | Apple Silicon | ~8MB |

## 🔨 自己编译

如果你想自己编译工具，可以运行：

**Windows:**
```bash
cd scripts
.\build.bat
```

**Linux/macOS:**
```bash
cd scripts
chmod +x build.sh
./build.sh
```

编译后的文件会自动保存到这个 `cli/` 目录。

## 📖 使用示例

```bash
# 运行工具（以 Windows 为例）
> .\cli\tmdb-fetcher-windows-amd64.exe

============================================================
  TMDB 数据获取工具
  从TMDB API获取电影/电视剧数据并按格式保存
============================================================

请选择媒体类型:
  1. 电影 (Movie)
  2. 电视剧 (TV Show)
  q. 退出

请输入选项 (1/2/q): 1

请输入TMDB ID (或输入 'q' 退出): 842675

开始获取电影 ID: 842675 的数据...
正在请求: /movie/842675
已保存: ../../tmdb_config/movie/842675/details.json
正在请求: /movie/842675/release_dates
已保存: ../../tmdb_config/movie/842675/release_dates.json

✓ 电影数据获取完成!
  标题: 流浪地球2
  目录: ../../tmdb_config/movie/842675
```

## ❓ 常见问题

**Q: 如何获取 TMDB ID？**

A: 访问 TMDB 网站查找电影或电视剧，URL 中的数字就是 ID。例如 `https://www.themoviedb.org/movie/842675` 中的 `842675`。

**Q: 配置文件在哪里？**

A: 配置文件位于 `cli/config.json`，需要从 `cli/config.example.json` 复制并填入 API Key。

**Q: 工具会覆盖已有数据吗？**

A: 不会。如果目标目录已存在，工具会提示并拒绝覆盖，保护已维护的元数据。

**Q: 需要代理吗？**

A: 如果你在中国大陆，建议在配置文件中启用代理（默认已启用）。

**Q: macOS 提示"无法打开，因为无法验证开发者"怎么办？**

A: 右键点击文件，选择"打开"，或在终端运行 `xattr -d com.apple.quarantine ./cli/tmdb-fetcher-macos-*`

## 🔗 相关链接

- [主文档](../README.md)
- [源代码](../scripts/README.md)
- [TMDB API 文档](https://developers.themoviedb.org/3)
