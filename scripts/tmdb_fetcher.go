package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Config 配置文件结构
type Config struct {
	TMDBAPIKey string `json:"tmdb_api_key"`
	Language   string `json:"language"`
	Proxy      struct {
		Enabled bool   `json:"enabled"`
		URL     string `json:"url"`
	} `json:"proxy"`
}

// TMDBFetcher TMDB数据获取器
type TMDBFetcher struct {
	config     Config
	httpClient *http.Client
	baseURL    string
}

const banner = `============================================================
  TMDB 数据获取工具
  从TMDB API获取电影/电视剧数据并按格式保存
============================================================
`

// NewTMDBFetcher 创建新的TMDB获取器
func NewTMDBFetcher(configPath string) (*TMDBFetcher, error) {
	config, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	if config.TMDBAPIKey == "" || config.TMDBAPIKey == "your_tmdb_api_key_here" {
		return nil, fmt.Errorf("请在配置文件中设置TMDB API Key")
	}

	// 设置默认语言
	if config.Language == "" {
		config.Language = "zh-CN"
	}

	fetcher := &TMDBFetcher{
		config:  config,
		baseURL: "https://api.themoviedb.org/3",
	}

	// 创建HTTP客户端
	fetcher.httpClient = createHTTPClient(config)

	return fetcher, nil
}

// loadConfig 加载配置文件
func loadConfig(configPath string) (Config, error) {
	var config Config

	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, fmt.Errorf("配置文件 '%s' 不存在\n请复制 'config.example.json' 为 'config.json' 并填写您的API Key", configPath)
		}
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("配置文件格式错误: %v", err)
	}

	return config, nil
}

// createHTTPClient 创建HTTP客户端
func createHTTPClient(config Config) *http.Client {
	transport := &http.Transport{}

	// 配置代理
	if config.Proxy.Enabled && config.Proxy.URL != "" {
		proxyURL, err := url.Parse(config.Proxy.URL)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
			fmt.Printf("已启用代理: %s\n", config.Proxy.URL)
		}
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
}

// makeRequest 发起API请求
func (f *TMDBFetcher) makeRequest(endpoint string, params map[string]string) (map[string]interface{}, error) {
	if params == nil {
		params = make(map[string]string)
	}

	params["api_key"] = f.config.TMDBAPIKey
	params["language"] = f.config.Language

	// 构建URL
	reqURL := f.baseURL + endpoint
	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		reqURL += "?" + values.Encode()
	}

	fmt.Printf("正在请求: %s\n", endpoint)

	resp, err := f.httpClient.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API返回错误 %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	return result, nil
}

// fetchMovieDetails 获取电影详细信息
func (f *TMDBFetcher) fetchMovieDetails(movieID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/movie/%s", movieID)
	params := map[string]string{
		"append_to_response": "credits,alternative_titles,translations,external_ids",
	}
	return f.makeRequest(endpoint, params)
}

// fetchMovieReleaseDates 获取电影发行日期
func (f *TMDBFetcher) fetchMovieReleaseDates(movieID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/movie/%s/release_dates", movieID)
	return f.makeRequest(endpoint, nil)
}

// fetchTVDetails 获取电视剧详细信息
func (f *TMDBFetcher) fetchTVDetails(tvID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/tv/%s", tvID)
	params := map[string]string{
		"append_to_response": "credits,alternative_titles,translations,external_ids,aggregate_credits",
	}
	return f.makeRequest(endpoint, params)
}

// fetchTVContentRatings 获取电视剧内容分级
func (f *TMDBFetcher) fetchTVContentRatings(tvID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/tv/%s/content_ratings", tvID)
	return f.makeRequest(endpoint, nil)
}

// saveJSON 保存JSON数据到文件
func saveJSON(data map[string]interface{}, filePath string) error {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 保存JSON文件
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	fmt.Printf("已保存: %s\n", filePath)
	return nil
}

// checkDirectoryExists 检查目录是否存在
func checkDirectoryExists(dir string) bool {
	info, err := os.Stat(dir)
	return err == nil && info.IsDir()
}

// fetchAndSaveMovie 获取并保存电影数据
func (f *TMDBFetcher) fetchAndSaveMovie(movieID string) error {
	fmt.Printf("\n开始获取电影 ID: %s 的数据...\n", movieID)

	// 创建目录 (保存到上级目录的tmdb_config文件夹)
	baseDir := filepath.Join("..", "tmdb_config", "movie", movieID)

	// 检查目录是否已存在
	if checkDirectoryExists(baseDir) {
		absPath, _ := filepath.Abs(baseDir)
		fmt.Printf("\n⚠️  警告: 目录已存在: %s\n", baseDir)
		fmt.Println("该电影数据已经生成，为防止覆盖已维护的元数据，操作已取消。")
		fmt.Println("\n如需重新生成，请先手动删除该目录:")
		fmt.Printf("  rmdir /s \"%s\"\n", absPath)
		return nil
	}

	// 获取并保存详细信息
	details, err := f.fetchMovieDetails(movieID)
	if err != nil {
		return err
	}
	if err := saveJSON(details, filepath.Join(baseDir, "details.json")); err != nil {
		return err
	}

	// 获取并保存发行日期
	releaseDates, err := f.fetchMovieReleaseDates(movieID)
	if err != nil {
		return err
	}
	if err := saveJSON(releaseDates, filepath.Join(baseDir, "release_dates.json")); err != nil {
		return err
	}

	fmt.Println("\n✓ 电影数据获取完成!")
	if title, ok := details["title"].(string); ok {
		fmt.Printf("  标题: %s\n", title)
	} else if origTitle, ok := details["original_title"].(string); ok {
		fmt.Printf("  标题: %s\n", origTitle)
	}
	fmt.Printf("  目录: %s\n", baseDir)

	return nil
}

// fetchAndSaveTV 获取并保存电视剧数据
func (f *TMDBFetcher) fetchAndSaveTV(tvID string) error {
	fmt.Printf("\n开始获取电视剧 ID: %s 的数据...\n", tvID)

	// 创建目录 (保存到上级目录的tmdb_config文件夹)
	baseDir := filepath.Join("..", "tmdb_config", "tv", tvID)

	// 检查目录是否已存在
	if checkDirectoryExists(baseDir) {
		absPath, _ := filepath.Abs(baseDir)
		fmt.Printf("\n⚠️  警告: 目录已存在: %s\n", baseDir)
		fmt.Println("该电视剧数据已经生成，为防止覆盖已维护的元数据，操作已取消。")
		fmt.Println("\n如需重新生成，请先手动删除该目录:")
		fmt.Printf("  rmdir /s \"%s\"\n", absPath)
		return nil
	}

	// 获取并保存详细信息
	details, err := f.fetchTVDetails(tvID)
	if err != nil {
		return err
	}
	if err := saveJSON(details, filepath.Join(baseDir, "details.json")); err != nil {
		return err
	}

	// 获取并保存内容分级
	contentRatings, err := f.fetchTVContentRatings(tvID)
	if err != nil {
		return err
	}
	if err := saveJSON(contentRatings, filepath.Join(baseDir, "content_ratings.json")); err != nil {
		return err
	}

	fmt.Println("\n✓ 电视剧数据获取完成!")
	if name, ok := details["name"].(string); ok {
		fmt.Printf("  标题: %s\n", name)
	} else if origName, ok := details["original_name"].(string); ok {
		fmt.Printf("  标题: %s\n", origName)
	}
	fmt.Printf("  目录: %s\n", baseDir)

	return nil
}

// getMediaType 获取媒体类型
func getMediaType(reader *bufio.Reader) (string, error) {
	for {
		fmt.Println("\n请选择媒体类型:")
		fmt.Println("  1. 电影 (Movie)")
		fmt.Println("  2. 电视剧 (TV Show)")
		fmt.Println("  q. 退出")
		fmt.Print("\n请输入选项 (1/2/q): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		input = strings.TrimSpace(input)

		switch input {
		case "1":
			return "movie", nil
		case "2":
			return "tv", nil
		case "q", "Q":
			return "quit", nil
		default:
			fmt.Println("无效的选项，请重新输入")
		}
	}
}

// getMediaID 获取媒体ID
func getMediaID(reader *bufio.Reader) (string, error) {
	for {
		fmt.Print("\n请输入TMDB ID (或输入 'q' 退出): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		input = strings.TrimSpace(input)

		if input == "q" || input == "Q" {
			return "quit", nil
		}

		if input == "" {
			fmt.Println("ID不能为空，请重新输入")
			continue
		}

		return input, nil
	}
}

// getContinue 询问是否继续
func getContinue(reader *bufio.Reader) bool {
	fmt.Print("\n是否继续获取其他数据? (y/n): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y" || input == "yes"
}

func main() {
	fmt.Println(banner)

	// 初始化获取器
	fetcher, err := NewTMDBFetcher("../cli/config.json")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		fmt.Println("\n按回车键退出...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// 获取媒体类型
		mediaType, err := getMediaType(reader)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			break
		}
		if mediaType == "quit" {
			fmt.Println("\n再见!")
			break
		}

		// 获取媒体ID
		mediaID, err := getMediaID(reader)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			break
		}
		if mediaID == "quit" {
			fmt.Println("\n再见!")
			break
		}

		// 获取并保存数据
		var fetchErr error
		if mediaType == "movie" {
			fetchErr = fetcher.fetchAndSaveMovie(mediaID)
		} else {
			fetchErr = fetcher.fetchAndSaveTV(mediaID)
		}

		if fetchErr != nil {
			fmt.Printf("\n错误: %v\n", fetchErr)
			fmt.Print("是否重试? (y/n): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToLower(input))
			if input != "y" && input != "yes" {
				break
			}
			continue
		}

		// 询问是否继续
		fmt.Println("\n" + strings.Repeat("=", 60))
		if !getContinue(reader) {
			fmt.Println("\n感谢使用，再见!")
			break
		}
	}
}
