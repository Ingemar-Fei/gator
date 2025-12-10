package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/ingemar-fei/gator/internal/util"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// FetchFeed 从指定的URL获取RSS订阅源并解析
// 参数:
//
//	ctx - 用于请求控制的上下文，可以用于取消请求或设置超时
//	feedURL - RSS订阅源的URL地址
//
// 返回值:
//
//	*RSSFeed - 解析后的RSS订阅源结构体指针
//	error - 在获取或解析过程中出现的错误，成功时为nil
func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	/*
	 * 这段代码用于获取并解析RSS feed
	 * 主要功能包括：创建HTTP请求、发送请求、接收响应、解析XML并处理HTML实体
	 */
	// 创建一个带有10秒超时的上下文，并确保在函数返回前取消上下文
	requestCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 使用指定的上下文创建一个新的HTTP GET请求
	request, err := http.NewRequestWithContext(requestCtx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error new request: %v", err)
	}

	// 设置请求头中的User-Agent
	request.Header.Set("User-Agent", "gator")

	// 创建HTTP客户端并发送请求
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error send request: %v", err)
	}
	// 确保在函数返回前关闭响应体
	defer response.Body.Close()

	// 读取响应体内容
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error read body: %v", err)
	}

	// 将XML响应体解析为RSSFeed结构体
	var feed RSSFeed
	if err := xml.Unmarshal(responseBody, &feed); err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshal xml: %v", err)
	}

	// 如果处于调试模式，打印解析后的XML内容
	if util.DebugMode() {
		fmt.Printf("Parsed XML: %+v\n", feed)
	}
	// 解码HTML实体，确保标题和描述中的特殊字符能正确显示
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	for idx, item := range feed.Channel.Item {
		feed.Channel.Item[idx].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[idx].Description = html.UnescapeString(item.Description)
	}

	// 返回解析后的RSS feed
	return &feed, nil
}
