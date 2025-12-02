# 📜 Scripts 目录

本目录包含用于从 TMDB API 获取媒体数据的脚本，支持多种编程语言。

## 📁 目录结构

```
scripts/
├── config.example.json  # 配置文件模板（所有语言共用）
├── python/              # Python 实现
│   ├── tmdb_fetcher.py
│   └── requirements.txt
└── (其他语言实现...)
```

## 🚀 使用说明

### 第一步：配置 API Key

所有语言版本共用同一个配置文件：

```bash
cd scripts
copy config.example.json config.json
# 编辑 config.json 填入你的 TMDB API Key
```

### Python 版本

```bash
cd python
pip install -r requirements.txt
python tmdb_fetcher.py
```

详细使用说明请参考主项目的 [README.md](../README.md)。

## 🔮 未来计划

我们计划添加更多语言的实现：
- Go
- Node.js
- Rust
- ...

欢迎贡献其他语言版本的实现！
