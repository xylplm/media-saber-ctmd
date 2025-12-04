<div align="center">
  <img src="logo.png" alt="logo-blue" width="300"/>
</div>

# 🎬 Media Saber 媒体库管理工具

## 📦 项目介绍

Media Saber 是一款适合自己人用的媒体管理工具。

由多位热心Pter共同努力，立足于开发自用，完全自主开发，做一款全新、高效、稳定、快速、方便好用、功能强大的媒体管理工具！

详情见 [这里](https://wiki.msaber.fun) 。


## 📚 仓库介绍

本仓库是 Media Saber 的元数据自维护库。主要目的是解决TMDB库被恶意修改导致识别错误的问题，通过社区协作维护高质量的媒体元数据。

### 🔄 自动更新

本仓库通过 GitHub Actions 自动打包和发布元数据：
- 📦 每次 `tmdb_config` 文件夹有更新时，自动生成压缩包并发布
- 📥 Media Saber 程序在配置开启后，会自动订阅和更新最新的元数据
- ✓ 支持完整性校验（SHA256）

无需手动配置，程序会自动拉取最新数据。

详见 [Release 页面](https://github.com/xylplm/media-saber-ctmd/releases)

## 📁 目录结构

```
tmdb_config/
├── movie/                          # 电影元数据
│   └── {tmdb_id}/                  # TMDB电影ID目录
│       ├── details.json            # 电影基本信息
│       └── release_dates.json       # 电影发布日期信息
└── tv/                             # 电视剧元数据
    └── {tmdb_id}/                  # TMDB电视剧ID目录
        ├── details.json            # 电视剧基本信息
        ├── content_ratings.json     # 电视剧内容分级信息
        └── season/                 # 季度目录（如需要）
            └── {season_number}/
                └── episode/        # 集数目录（如需要）
```

## 🚀 快速开始

### 📥 获取TMDB数据

我们提供了预编译的命令行工具，**无需任何环境**，下载即用！

**第一步：配置API Key**

```bash
cd cli
copy config.example.json config.json
# 编辑 config.json 填入你的 TMDB API Key
```

在 [TMDB网站](https://www.themoviedb.org/settings/api) 申请API Key

**第二步：运行工具**

根据你的操作系统，选择对应的可执行文件：

**Windows:**
```bash
# 双击运行或命令行执行
.\cli\tmdb-manager-windows-amd64.exe
```

**Linux:**
```bash
chmod +x ./cli/tmdb-manager-linux-amd64
./cli/tmdb-manager-linux-amd64
```

**macOS:**
```bash
# Intel 芯片
chmod +x ./cli/tmdb-manager-macos-amd64
./cli/tmdb-manager-macos-amd64

# Apple Silicon (M1/M2/M3)
chmod +x ./cli/tmdb-manager-macos-arm64
./cli/tmdb-manager-macos-arm64
```

**第三步：按提示操作**

1. 选择媒体类型（电影或电视剧）
2. 输入TMDB ID（可从TMDB网站的URL中获取，如 `https://www.themoviedb.org/movie/842675` 中的 `842675`）
3. 工具会自动获取并保存数据到 `tmdb_config` 目录

**示例 - 获取电影《流浪地球2》(ID: 842675)：**

```
请选择媒体类型:
  1. 电影 (Movie)
  2. 电视剧 (TV Show)
  q. 退出

请输入选项 (1/2/q): 1

请输入TMDB ID (或输入 'q' 退出): 842675

开始获取电影 ID: 842675 的数据...
正在请求: /movie/842675
已保存: ../tmdb_config/movie/842675/details.json
正在请求: /movie/842675/release_dates
已保存: ../tmdb_config/movie/842675/release_dates.json

✓ 电影数据获取完成!
  标题: 流浪地球2
  目录: ../tmdb_config/movie/842675
```

数据将自动保存到 `tmdb_config/movie/842675/` 目录下。

#### 配置说明

配置文件 `cli/config.json` 可以自定义以下选项：

```json
{
  "tmdb_api_key": "your_api_key_here",           // TMDB API Key（必填）
  "language": "zh-CN",                           // 响应语言（默认中文）
  "proxy": {
    "enabled": true,                             // 是否启用代理
    "url": "http://127.0.0.1:7890"              // 代理服务器地址
  }
}
```

**配置说明：**
- `tmdb_api_key` - 必须填写，从 [TMDB 设置](https://www.themoviedb.org/settings/api) 申请
- `language` - 响应数据的语言代码，默认 `zh-CN`（简体中文）
- `proxy.enabled` - 如果在中国大陆，建议设为 `true`
- `proxy.url` - 代理服务器地址，根据实际情况修改

#### 自己编译（开发者）

如果你想自己编译工具，需要安装 [Go 1.21+](https://golang.org/dl/)：

```bash
cd scripts
# Windows: 运行 build.bat
# Linux/macOS: ./build.sh
```

编译后的文件会保存到 `cli/` 目录。详见 [scripts/README.md](scripts/README.md)

#### 获取的数据内容

**电影数据包含：**
- `details.json` - 包含完整的电影信息、演职人员、其他片名、翻译、外部ID等
- `release_dates.json` - 各国发行日期和分级信息

**电视剧数据包含：**
- `details.json` - 包含完整的电视剧信息、演职人员、其他剧名、翻译、外部ID等
- `content_ratings.json` - 各国内容分级信息

生成的JSON文件可以直接用于后续的维护和修改。

## 🛠️ TMDB元数据维护指南

### 🧱 数据结构说明

#### movie/{tmdb_id}/details.json
包含电影的基本信息，如标题、描述、发布日期、评分等TMDB官方数据的修正或补充。

#### movie/{tmdb_id}/release_dates.json
包含电影在不同国家/地区的发布日期和类型。

#### tv/{tmdb_id}/details.json
包含电视剧的基本信息，如标题、描述、首播日期、评分等。

#### tv/{tmdb_id}/content_ratings.json
包含电视剧在不同国家/地区的内容分级信息。

### ✏️ 维护和修改元数据

使用脚本生成的JSON文件后，你可以根据需要对元数据进行修改和维护：

1. **找到生成的文件**
   - 电影：`tmdb_config/movie/{tmdb_id}/details.json` 和 `release_dates.json`
   - 电视剧：`tmdb_config/tv/{tmdb_id}/details.json` 和 `content_ratings.json`

2. **修改内容**
   - 使用任何文本编辑器打开JSON文件
   - 修正错误的标题、描述、日期等信息
   - 添加缺失的翻译或其他语言版本
   - 更正演职人员信息
   - 修改分级信息等

3. **保持格式**
   - 确保JSON格式正确（可使用在线JSON验证工具）
   - 保持与原始TMDB数据结构一致
   - 注意中文编码使用UTF-8

4. **常见修改场景**
   - 修正被恶意篡改的中文译名
   - 补充缺失的中文描述
   - 更正错误的发行日期
   - 添加准确的分级信息

### 🤝 如何贡献

1. **发现问题**：如果您在使用 Media Saber 时发现TMDB数据有误，请在 [GitHub Issues](https://github.com/xylplm/media-saber-ctmd/issues) 上提交反馈，详细描述问题所在。

2. **使用脚本获取并修正**：
   - 运行 `cli/` 目录中的工具获取原始TMDB数据
   - 在生成的JSON文件中修正错误或补充信息
   
3. **一键提交PR**（推荐）：
   - 使用内置的 **一键提交PR功能** 方便快捷地提交修改
   - 运行工具后选择菜单中的 "2. 一键提交修改到PR"
   - 支持两种模式：
     - **模式1**：创建新分支提交新的PR
     - **模式2**：提交修改到已有的PR（只需推送即可）
   - 工具会自动处理git操作，无需手动执行git命令

4. **手动提交PR**：
   - 如果不使用自动工具，可以通过 [Pull Request](https://github.com/xylplm/media-saber-ctmd/pulls) 手动提交修正后的文件
   - 请确保：
     - 数据准确无误
     - 遵循现有的JSON格式和命名规范
     - 在PR描述中说明修改原因和数据来源

5. **维护更新**：定期检查和更新数据，确保元数据的时效性和准确性。

## 🙋 参与方式

- **提交Issue**：发现数据问题？请在 [GitHub Issues](https://github.com/xylplm/media-saber-ctmd/issues) 中提出，我们会及时处理。
- **提交PR**：如果您有能力修复问题或补充数据，欢迎提交 [Pull Request](https://github.com/xylplm/media-saber-ctmd/pulls)，我们会认真审阅您的贡献。
- **讨论建议**：有任何改进建议？欢迎在 Issues 中分享您的想法。

## ⚠️ 免责声明

1. **数据来源声明**：本仓库中的元数据来自于TMDB (The Movie Database) 及其他公开来源，仅用于修正因恶意修改或错误输入导致的数据问题。

2. **使用目的限制**：本仓库仅供学习交流使用，所有修正数据仅作为辅助工具简化用户手工操作，对用户的行为及内容毫不知情，使用本仓库数据产生的任何责任需由使用者本人承担。

3. **数据准确性**：虽然我们尽力确保数据的准确性，但不能保证所有数据完全无误。如发现数据错误，请通过Issue或PR反馈，我们会及时修正。

4. **非商业用途**：本项目为非商业性项目，旨在为社区服务。本项目没有在任何地方发布捐赠信息页面，也不会接受捐赠或进行收费，请仔细辨别避免误导。

5. **知识产权**：本仓库遵循许可证条款。TMDB相关数据受其官方条款约束，本仓库仅为修正和维护之用。

## 📄 许可证

详见 [LICENSE](LICENSE) 文件。
