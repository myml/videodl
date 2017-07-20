// 视频下载 project main.go
package main

import (
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-macaron/macaron"
	"github.com/tidwall/gjson"
)

type SUB struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type VideoInfo struct {
	Title     string `json:"title"`
	VideoID   string `json:"videoID"`
	SourceURL string `json:"sourceURL"`
	PlayCount int64  `json:"playCount"`
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Panic(err)
	}

	http.DefaultClient.Jar = jar

	m := macaron.Classic()
	m.Use(macaron.Renderer())

	pwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	gopath := os.Getenv("GOPATH")
	if gopath != "" && strings.Contains(pwd, os.Getenv("GOPATH")) {
		log.Println("调试")
		m.Use(macaron.Static("html"))
	} else {
		log.Println("在浏览器输入127.0.0.1:8080")
		if runtime.GOOS == "windows" {
			exec.Command("start", "http://127.0.0.1:8080").CombinedOutput()
		} else {
			exec.Command("xdg-open", "http://127.0.0.1:8080").CombinedOutput()
		}
		m.Use(macaron.Static("", macaron.StaticOptions{
			FileSystem: assetFS(),
		}))
	}

	m.Get("/host", func(ctx *macaron.Context) {
		out := make([]SUB, 0)
		out = append(out, SUB{"toutiao", "今日头条"})
		out = append(out, SUB{"360", "360视频"})
		out = append(out, SUB{"miaopai", "秒拍"})
		out = append(out, SUB{"vmovier", "V视频"})
		ctx.JSON(200, out)
	})
	m.Get("/video", func(ctx *macaron.Context) {
		src := ctx.Query("src")
		log.Println(src)
		req, err := http.NewRequest("GET", src, nil)
		if err != nil {
			log.Panic(err)
		}
		for k, v := range ctx.Req.Header {
			if k != "Host" && k != "Referer" {
				req.Header[k] = v
			}
		}
		req.Header.Set("Referer", src)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Panic(err)
		}
		defer resp.Body.Close()
		h := ctx.Header()
		for k, v := range resp.Header {
			h[k] = v
		}
		h.Set("Content-Disposition", fmt.Sprintf("attachment; filename=video.mp4"))
		if ctx.Query("name") != "" {
			h.Set("Content-type", "application/octet-stream")
			h.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.mp4", ctx.Query("name")))
		} else {
			h.Set("Content-type", "video/mp4")
		}
		ctx.WriteHeader(resp.StatusCode)
		io.Copy(ctx, resp.Body)
	})

	//V视频
	m.Group("/host/vmovier", func() {
		m.Get("/", func(ctx *macaron.Context) {
			doc, err := goquery.NewDocument("http://www.vmovier.com/channel")
			if err != nil {
				log.Panic(err)
			}
			var out []SUB
			doc.Find(".channel-img").Each(func(i int, s *goquery.Selection) {
				id := s.AttrOr("href", "")[len("/channel/"):]
				title := s.Find("img").AttrOr("alt", "")
				out = append(out, SUB{id, title})
			})
			ctx.JSON(200, out)
		})
		m.Get("/sub/:sub", func(ctx *macaron.Context) {
			doc, err := goquery.NewDocument("http://www.vmovier.com/channel/" + ctx.Params(":sub"))
			if err != nil {
				log.Panic(err)
			}
			var out []VideoInfo
			doc.Find(".works-text").Each(func(i int, s *goquery.Selection) {
				a := s.Find("a")
				src := a.AttrOr("href", "")
				like, _ := strconv.Atoi(s.Find("span[title='喜欢数']").Text())
				out = append(out, VideoInfo{
					Title:     a.Text(),
					VideoID:   src[1:strings.Index(src, "?")],
					SourceURL: "http://www.vmovier.com" + src,
					PlayCount: int64(like),
				})
			})
			ctx.JSON(200, out)
		})

		m.Get("/video/:vid/", func(ctx *macaron.Context) {
			resp, err := http.Get("http://www.vmovier.com/post/getViewData?id=" + ctx.Params("vid"))
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}

			out := make(map[string]interface{})
			out["src"] = gjson.ParseBytes(b).Get("data.qiniu_url").String()
			ctx.JSON(200, out)
		})
	})

	//秒拍
	m.Group("/host/miaopai", func() {
		m.Get("/", func(ctx *macaron.Context) {
			out := []SUB{{"128", "搞笑"}, {"144", "宝宝"}, {"172", "音乐"}, {"168", "旅游"}, {"28", "美食"}, {"124", "明星"}, {"156", "时尚"}, {"106", "韩娱"}, {"160", "牛人"}, {"140", "萌宠"}, {"114", "汽车"}, {"136", "现场"}, {"148", "奥运"}}
			ctx.JSON(200, out)
		})
		m.Get("/sub/:sub", func(ctx *macaron.Context) {
			resp, err := http.Get("http://api.miaopai.com/m/cate2_channel.json?cateid=" + ctx.Params("sub"))
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}
			var out []VideoInfo
			for i, result := range gjson.ParseBytes(b).Get("result.#.channel").Array() {
				out = append(out, VideoInfo{
					Title:     result.Get("ext.t").String(),
					VideoID:   result.Get("scid").String(),
					SourceURL: result.Get("stream.base").String(),
					PlayCount: result.Get("stat.vcnt").Int()})
				if len([]rune(out[i].Title)) > 40 {
					out[i].Title = string([]rune(out[i].Title)[:40]) + "..."
				}
			}
			ctx.JSON(200, out)
		})

		m.Get("/video/:vid/", func(ctx *macaron.Context) {
			out := make(map[string]interface{})
			out["src"] = fmt.Sprintf("http://gslb.miaopai.com/stream/%s.mp4?vend=miaopai", ctx.Params("vid"))
			ctx.JSON(200, out)
		})
	})
	//360
	m.Group("/host/360", func() {
		m.Get("/", func(ctx *macaron.Context) {
			doc, err := goquery.NewDocument("http://v.sj.360.cn/")
			if err != nil {
				log.Panic(err)
			}
			out := make([]SUB, 0)
			doc.Find(".nav").Each(func(i int, s *goquery.Selection) {
				out = append(out, SUB{})
				id := s.AttrOr("data-id", "")
				if id != "" {
					out[i].ID = id
					out[i].Title = s.Text()
				}
			})
			ctx.JSON(200, out)
		})

		m.Get("/sub/:sub", func(ctx *macaron.Context) {
			resp, err := http.Get("http://v.sj.360.cn/pc/list?f=json&channel_id=" + ctx.Params("sub"))
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}
			var out []VideoInfo
			for _, result := range gjson.ParseBytes(b).Get("data.res").Array() {
				out = append(out, VideoInfo{
					Title:     result.Get("t").String(),
					SourceURL: result.Get("videoUrl").String(),
					VideoID:   result.Get("exData.code").String(),
					PlayCount: result.Get("playcnt").Int(),
				})
			}
			ctx.JSON(200, out)
		})

		m.Get("/video/:vid", func(ctx *macaron.Context) {
			resp, err := http.Get("http://v.sj.360.cn/pc/play?f=json&id=" + ctx.Params("vid"))
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}
			out := make(map[string]interface{})
			out["src"] = gjson.Get(string(b), "data.url").String()
			ctx.JSON(200, out)
		})
	})

	//头条
	m.Group("/host/toutiao", func() {
		m.Get("/", func(ctx *macaron.Context) {
			resp, err := http.Get("http://www.toutiao.com/ch/video/")
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}
			body := string(b)
			//			log.Println(body)
			start, end := "subchannels = ", "}];"
			//			log.Println(start, end)
			subchannels := body[strings.Index(body, start)+len(start) : strings.Index(body, end)+len(end)]
			//			log.Println(subchannels)
			out := make([]SUB, 0)
			for i, result := range gjson.Parse(subchannels).Array() {
				out = append(out, SUB{})
				out[i].ID = result.Get("category").String()
				out[i].Title = result.Get("name").String()
				if out[i].Title == "中国新唱将" {
					out[i].ID = out[i].Title
				}
			}
			ctx.JSON(200, out)
		})
		m.Get("/sub/:sub", func(ctx *macaron.Context) {
			url := "http://www.toutiao.com/api/pc/feed/?category=" + ctx.Params("sub")
			resp, err := http.Get(url)
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}
			var out []VideoInfo
			for _, result := range gjson.ParseBytes(b).Get("data").Array() {
				out = append(out, VideoInfo{
					Title:     result.Get("title").String(),
					VideoID:   result.Get("video_id").String(),
					SourceURL: "http://www.toutiao.com" + result.Get("source_url").String(),
					PlayCount: result.Get("video_play_count").Int(),
				})
			}
			ctx.JSON(200, out)
		})
		m.Get("/video/:vid/", func(ctx *macaron.Context) {
			vid := ctx.Params("vid")
			r := strconv.Itoa(rand.Int())[:12]
			s := crc32.ChecksumIEEE([]byte(fmt.Sprintf("/video/urls/v/1/toutiao/mp4/%s?r=%s", vid, r)))
			url := fmt.Sprintf("http://i.snssdk.com/video/urls/v/1/toutiao/mp4/%s?r=%s&s=%d", vid, r, s)
			resp, err := http.Get(url)
			if err != nil {
				log.Panic(err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Panic(err)
			}
			log.Println(string(b))
			result := gjson.ParseBytes(b).Get("data.video_list")
			for i := 3; i > 0; i-- {
				k := fmt.Sprintf("video_%d", i)
				if result.Get(k).Exists() {
					b, err := base64.StdEncoding.DecodeString(result.Get(k).Get("main_url").String())
					if err != nil {
						log.Panic(err)
					}
					out := make(map[string]interface{})
					out["src"] = string(b)
					out["size"] = result.Get(k).Get("size").Int()
					ctx.JSON(200, out)
					return
				}
			}
		})
	})

	m.Run(8080)
}

//	//微博
//	m.Group("/weibo", func() {
//		m.Get("/sub", func(ctx *macaron.Context) {
//			doc, err := goquery.NewDocument("http://weibo.com/tv")
//			if err != nil {
//				log.Panic(err)
//			}
//			log.Println(doc.Html())
//			out := make(map[string]string)
//			doc.Find(".tab_item").Each(func(i int, s *goquery.Selection) {
//				href := s.AttrOr("href", "")
//				if href != "" {
//					out[filepath.Base(href)] = s.Text()
//				}

//				log.Println(s.AttrOr("href", ""))
//			})
//			ctx.JSON(200, out)
//		})
//		m.Get("/:sub/list", func(ctx *macaron.Context) {
//			doc, err := goquery.NewDocument("http://weibo.com/tv/" + ctx.Params("sub"))
//			if err != nil {
//				log.Panic(err)
//			}
//			var out []map[string]string
//			doc.Find(".weibo_tv_frame a").Each(func(i int, s *goquery.Selection) {
//				out = append(out, make(map[string]string))
//				out[i]["title"] = strings.TrimSpace(s.Find(".txt_cut").Text())
//				vid := strings.Split(filepath.Base(s.AttrOr("href", "")), "?")[0]
//				out[i]["videoID"] = vid
//				out[i]["sourceURL"] = "http://weibo.com/tv/v/" + vid
//			})
//			ctx.JSON(200, out)
//		})
//		m.Get("/:vid/info/", func(ctx *macaron.Context) {
//			doc, err := goquery.NewDocument("http://weibo.com/tv/v/" + ctx.Params("vid"))
//			if err != nil {
//				log.Panic(err)
//			}
//			data := doc.Find("#playerRoom div").First().AttrOr("action-data", "")
//			v, err := url.ParseQuery(data)
//			if err != nil {
//				log.Panic(err)
//			}
//			out := make(map[string]string)
//			out["src"], err = url.QueryUnescape(v["video_src"][0])
//			if err != nil {
//				log.Panic(err)
//			}
//			ctx.JSON(200, out)
//		})
//	})
