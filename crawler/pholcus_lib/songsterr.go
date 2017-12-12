package pholcus_lib

// 基础包
import (
	//必需

	"regexp"

	"github.com/henrylee2cn/pholcus/app/downloader/request"
	. "github.com/henrylee2cn/pholcus/app/spider"
	"github.com/henrylee2cn/pholcus/common/goquery"
	"github.com/henrylee2cn/pholcus/logs"
)

var (
	BaseURL = "https://www.songsterr.com"
)

func init() {
	Songsterr.Register()
}

// Songsterr search
var Songsterr = &Spider{
	Name:         "Songsterr",
	Description:  "Songsterr [https://www.songsterr.com/a/wa/search?pattern=Joe]",
	Keyin:        KEYIN,
	Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.Aid(map[string]interface{}{"loop": [2]int{0, 1}, "Rule": "生成请求"}, "生成请求")
		},
		Trunk: map[string]*Rule{

			"生成请求": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					for loop := aid["loop"].([2]int); loop[0] < loop[1]; loop[0]++ {
						ctx.AddQueue(
							&request.Request{
								Url:  BaseURL + "/a/wa/search?pattern=" + ctx.GetKeyin(),
								Rule: aid["Rule"].(string),
							},
						)
					}
					return nil
				},
				ItemFields: []string{
					"name",
					"artist",
					"target",
					"gtp",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("a.SongsItem").Each(func(i int, s *goquery.Selection) {
						a, _ := s.Attr("href")
						name := s.Find(".SongsItem-name").First()
						artist := s.Find(".SongsItem-artist").First()
						temp := ctx.CreatItem(map[int]interface{}{
							0: name,
							1: artist,
							2: a,
						})
						logs.Log.Critical("[消息提示：| 艺术家：%v | 专辑：%v | 地址：%v | 文件: %v]\n", temp["0"], temp["1"], temp["2"], "")
						// ctx.AddQueue(&request.Request{
						// 	Url:      BaseURL + a,
						// 	Rule:     "详情",
						// 	Temp:     temp,
						// 	Priority: 1,
						// })
					})
				},
			},

			"详情": {
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					temp := ctx.CopyTemps()
					query.Find("script").Each(func(arg0 int, s *goquery.Selection) {
						src := s.Text()
						re, _ := regexp.Compile(`source:"[\s\S]*"`)
						source := re.FindString(src)
						urlStr := source[8 : len(source)-1]
						temp["3"] = urlStr
						logs.Log.Critical("[消息提示：| 艺术家：%v | 专辑：%v | 地址：%v | 文件: %v]\n", temp["0"], temp["1"], temp["2"], temp["3"])
						ctx.Output(temp)
					})
				},
			},
		},
	},
}
