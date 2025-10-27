package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/ingemar-fei/gator/internal/util"
	"io"
	"net/http"
	"time"
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

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// 1. 构造上下文，方便做超时控制
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 2. 创建请求
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error new request: %v", err)
	}

	// 3. 设置 User-Agent
	req.Header.Set("User-Agent", "gator")

	// 4. 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error send request: %v", err)
	}
	defer resp.Body.Close()

	// 5. 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error read body: %v", err)
	}

	// 6. 解析 XML
	var result RSSFeed
	if err := xml.Unmarshal(body, &result); err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshal xml: %v", err)
	}

	// 7. 使用解析结果
	if util.DebugMode() {
		fmt.Printf("Parsed XML: %+v\n", result)
	}

	return &result, nil
}
