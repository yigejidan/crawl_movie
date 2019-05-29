package controllers

import (
	"github.com/astaxie/beego"
	"crawl_movie/models"
	_"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}
/**
	目前这个爬虫只能爬取静态数据，对于像京东的部分动态数据，无法爬取
	对于动态数据，可以采用一个组件，phantomjs
*/

func (c *CrawlMovieController) CrawlMovie() {
	var movieInfo models.MovieInfo

	//连接到redis
	models.ConnentRedis("127.0.0.1:6379")
	//爬虫入口url
	sUrl := "https://movie.douban.com/subject/26891256/"
	models.PutinQueue(sUrl)
	
	for {
		length := models.GetQueueLength()
		if length == 0 {
			break //如果url队列为空，则退出当前循环
		}
		
		sUrl = models.PopfromQueue()
		//判断sUrl是否已经访问过
		if models.IsVisit(sUrl) {
			continue
		}
		rsp := httplib.Get(sUrl)
		sMovieHtml,err := rsp.String()
		if err != nil {
			panic(err)
		}

		movieInfo.Movie_name           = models.GetMovieName(sMovieHtml)
		//记录电影信息
		if movieInfo.Movie_name != "" {
			// movieInfo.Id                   = id
			movieInfo.Movie_director       = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_country        = models.GetMovieCountry(sMovieHtml)
			movieInfo.Movie_language       = models.GetMovieLanguage(sMovieHtml)
			movieInfo.Movie_writer         = models.GetMovieWriter(sMovieHtml)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_grade          = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_type           = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_on_time        = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_span           = models.GetMovieRunningTime(sMovieHtml)
			_,err := models.AddMovie(&movieInfo)
			if err != nil {
				panic(err)
			}
		}
		//提取该页面的所有链接
		urls := models.GetMovieUrls(sMovieHtml)

		for _,url := range urls {
			models.PutinQueue(url)
			c.Ctx.WriteString("<br>" + url + "</br>")
		}

		//sUrl应当记录到访问set中
		models.AddToSet(sUrl)

		time.Sleep(time.Second)
		//c.Ctx.WriteString(fmt.Sprintf("%v",urls))
	}
	c.Ctx.WriteString("end of crawl!")
}
