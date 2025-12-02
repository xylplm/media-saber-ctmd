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

### 🤝 如何贡献

1. **发现问题**：如果您在使用 Media Saber 时发现TMDB数据有误，请在 [GitHub Issues](https://github.com/xylplm/media-saber-ctmd/issues) 上提交反馈，详细描述问题所在。

2. **提交修正**：欢迎通过 [Pull Request](https://github.com/xylplm/media-saber-ctmd/pulls) 直接提交更正后的元数据文件。请确保：
   - 数据准确无误
   - 遵循现有的JSON格式和命名规范
   - 在PR描述中说明修改原因和数据来源

3. **维护更新**：定期检查和更新数据，确保元数据的时效性和准确性。

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
