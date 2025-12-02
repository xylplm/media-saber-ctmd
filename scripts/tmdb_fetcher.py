#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
TMDB数据获取脚本
用于从TMDB API获取电影或电视剧的详细信息，并按照指定格式保存到本地
"""

import json
import os
import sys
import requests
from typing import Dict, Optional


class TMDBFetcher:
    """TMDB数据获取器"""
    
    BASE_URL = "https://api.themoviedb.org/3"
    
    def __init__(self, config_path: str = "config.json"):
        """
        初始化TMDB获取器
        
        Args:
            config_path: 配置文件路径
        """
        self.config = self._load_config(config_path)
        self.api_key = self.config.get("tmdb_api_key")
        self.language = self.config.get("language", "zh-CN")
        self.session = self._create_session()
        
        if not self.api_key:
            raise ValueError("请在配置文件中设置TMDB API Key")
    
    def _load_config(self, config_path: str) -> Dict:
        """加载配置文件"""
        if not os.path.exists(config_path):
            print(f"错误: 配置文件 '{config_path}' 不存在")
            print("请复制 'config.example.json' 为 'config.json' 并填写您的API Key")
            sys.exit(1)
        
        try:
            with open(config_path, 'r', encoding='utf-8') as f:
                return json.load(f)
        except json.JSONDecodeError as e:
            print(f"错误: 配置文件格式错误 - {e}")
            sys.exit(1)
    
    def _create_session(self) -> requests.Session:
        """创建HTTP会话，配置代理"""
        session = requests.Session()
        
        # 配置代理
        proxy_config = self.config.get("proxy", {})
        if proxy_config.get("enabled", False):
            proxies = {
                "http": proxy_config.get("http"),
                "https": proxy_config.get("https")
            }
            session.proxies.update(proxies)
            print(f"已启用代理: {proxies}")
        
        return session
    
    def _make_request(self, endpoint: str, params: Optional[Dict] = None) -> Dict:
        """
        发起API请求
        
        Args:
            endpoint: API端点
            params: 请求参数
            
        Returns:
            API响应的JSON数据
        """
        if params is None:
            params = {}
        
        params["api_key"] = self.api_key
        params["language"] = self.language
        
        url = f"{self.BASE_URL}{endpoint}"
        
        try:
            print(f"正在请求: {endpoint}")
            response = self.session.get(url, params=params, timeout=30)
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"请求失败: {e}")
            sys.exit(1)
    
    def fetch_movie_details(self, movie_id: int) -> Dict:
        """
        获取电影详细信息
        
        Args:
            movie_id: 电影ID
            
        Returns:
            电影详细信息
        """
        endpoint = f"/movie/{movie_id}"
        params = {
            "append_to_response": "credits,alternative_titles,translations,external_ids"
        }
        return self._make_request(endpoint, params)
    
    def fetch_movie_release_dates(self, movie_id: int) -> Dict:
        """
        获取电影发行日期信息
        
        Args:
            movie_id: 电影ID
            
        Returns:
            电影发行日期信息
        """
        endpoint = f"/movie/{movie_id}/release_dates"
        return self._make_request(endpoint)
    
    def fetch_tv_details(self, tv_id: int) -> Dict:
        """
        获取电视剧详细信息
        
        Args:
            tv_id: 电视剧ID
            
        Returns:
            电视剧详细信息
        """
        endpoint = f"/tv/{tv_id}"
        params = {
            "append_to_response": "credits,alternative_titles,translations,external_ids,aggregate_credits"
        }
        return self._make_request(endpoint, params)
    
    def fetch_tv_content_ratings(self, tv_id: int) -> Dict:
        """
        获取电视剧内容分级信息
        
        Args:
            tv_id: 电视剧ID
            
        Returns:
            电视剧内容分级信息
        """
        endpoint = f"/tv/{tv_id}/content_ratings"
        return self._make_request(endpoint)
    
    def save_json(self, data: Dict, file_path: str) -> None:
        """
        保存JSON数据到文件
        
        Args:
            data: 要保存的数据
            file_path: 文件路径
        """
        # 确保目录存在
        os.makedirs(os.path.dirname(file_path), exist_ok=True)
        
        # 保存JSON文件
        with open(file_path, 'w', encoding='utf-8') as f:
            json.dump(data, f, ensure_ascii=False, indent=2)
        
        print(f"已保存: {file_path}")
    
    def check_directory_exists(self, base_dir: str) -> bool:
        """
        检查目录是否已存在
        
        Args:
            base_dir: 目录路径
            
        Returns:
            目录是否存在
        """
        return os.path.exists(base_dir) and os.path.isdir(base_dir)
    
    def fetch_and_save_movie(self, movie_id: int) -> None:
        """
        获取并保存电影相关数据
        
        Args:
            movie_id: 电影ID
        """
        print(f"\n开始获取电影 ID: {movie_id} 的数据...")
        
        # 创建目录 (保存到上级目录的tmdb_config文件夹)
        base_dir = os.path.join("..", "tmdb_config", "movie", str(movie_id))
        
        # 检查目录是否已存在
        if self.check_directory_exists(base_dir):
            print(f"\n⚠️  警告: 目录已存在: {base_dir}")
            print("该电影数据已经生成，为防止覆盖已维护的元数据，操作已取消。")
            print("\n如需重新生成，请先手动删除该目录:")
            print(f"  rmdir /s \"{os.path.abspath(base_dir)}\"")
            return
        
        # 获取并保存详细信息
        details = self.fetch_movie_details(movie_id)
        self.save_json(details, os.path.join(base_dir, "details.json"))
        
        # 获取并保存发行日期
        release_dates = self.fetch_movie_release_dates(movie_id)
        self.save_json(release_dates, os.path.join(base_dir, "release_dates.json"))
        
        print(f"\n✓ 电影数据获取完成!")
        print(f"  标题: {details.get('title', details.get('original_title', 'N/A'))}")
        print(f"  目录: {base_dir}")
    
    def fetch_and_save_tv(self, tv_id: int) -> None:
        """
        获取并保存电视剧相关数据
        
        Args:
            tv_id: 电视剧ID
        """
        print(f"\n开始获取电视剧 ID: {tv_id} 的数据...")
        
        # 创建目录 (保存到上级目录的tmdb_config文件夹)
        base_dir = os.path.join("..", "tmdb_config", "tv", str(tv_id))
        
        # 检查目录是否已存在
        if self.check_directory_exists(base_dir):
            print(f"\n⚠️  警告: 目录已存在: {base_dir}")
            print("该电视剧数据已经生成，为防止覆盖已维护的元数据，操作已取消。")
            print("\n如需重新生成，请先手动删除该目录:")
            print(f"  rmdir /s \"{os.path.abspath(base_dir)}\"")
            return
        
        # 获取并保存详细信息
        details = self.fetch_tv_details(tv_id)
        self.save_json(details, os.path.join(base_dir, "details.json"))
        
        # 获取并保存内容分级
        content_ratings = self.fetch_tv_content_ratings(tv_id)
        self.save_json(content_ratings, os.path.join(base_dir, "content_ratings.json"))
        
        print(f"\n✓ 电视剧数据获取完成!")
        print(f"  标题: {details.get('name', details.get('original_name', 'N/A'))}")
        print(f"  目录: {base_dir}")


def print_banner():
    """打印欢迎横幅"""
    print("=" * 60)
    print("  TMDB 数据获取工具")
    print("  从TMDB API获取电影/电视剧数据并按格式保存")
    print("=" * 60)
    print()


def get_media_type() -> str:
    """获取媒体类型"""
    while True:
        print("请选择媒体类型:")
        print("  1. 电影 (Movie)")
        print("  2. 电视剧 (TV Show)")
        print("  q. 退出")
        
        choice = input("\n请输入选项 (1/2/q): ").strip().lower()
        
        if choice == '1':
            return 'movie'
        elif choice == '2':
            return 'tv'
        elif choice == 'q':
            print("再见!")
            sys.exit(0)
        else:
            print("无效的选项，请重新输入\n")


def get_media_id() -> int:
    """获取媒体ID"""
    while True:
        try:
            media_id = input("\n请输入TMDB ID (或输入 'q' 退出): ").strip()
            
            if media_id.lower() == 'q':
                print("再见!")
                sys.exit(0)
            
            media_id = int(media_id)
            
            if media_id <= 0:
                print("ID必须是正整数，请重新输入")
                continue
            
            return media_id
        except ValueError:
            print("无效的ID格式，请输入数字")


def main():
    """主函数"""
    print_banner()
    
    try:
        # 初始化获取器
        fetcher = TMDBFetcher()
        
        while True:
            # 获取媒体类型
            media_type = get_media_type()
            
            # 获取媒体ID
            media_id = get_media_id()
            
            # 获取并保存数据
            try:
                if media_type == 'movie':
                    fetcher.fetch_and_save_movie(media_id)
                else:
                    fetcher.fetch_and_save_tv(media_id)
                
                # 询问是否继续
                print("\n" + "=" * 60)
                continue_choice = input("\n是否继续获取其他数据? (y/n): ").strip().lower()
                if continue_choice != 'y':
                    print("\n感谢使用，再见!")
                    break
                print()
                
            except Exception as e:
                print(f"\n错误: {e}")
                retry = input("是否重试? (y/n): ").strip().lower()
                if retry != 'y':
                    break
    
    except KeyboardInterrupt:
        print("\n\n程序已被用户中断")
        sys.exit(0)
    except Exception as e:
        print(f"\n发生错误: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()
