package main

import (
	"crawl_douban_movie/models"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
	"github.com/astaxie/beego/config"
)

func main() {
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	db_type := iniconf.String("database::Type")
	db_user := iniconf.String("database::User")
	db_pass := iniconf.String("database::Password")
	db_host := iniconf.String("database::Host")
	db_name := iniconf.String("database::Name")
	redis_addr := iniconf.String("redis::Addr")
	//redis_pass := iniconf.String("database::Password")
	models.Setup(db_type,db_user,db_pass,db_name,db_host)
	//连接到redis
	models.ConnectRedis(redis_addr)
	//models.RedisAuth(redis_pass)
	
	fmt.Println("crawl start...")
	var movieInfo models.MovieInfo
	//爬虫入口url
	sUrl := "https://movie.douban.com/"
	models.PutinQueue(sUrl)
	for{
		length := models.GetQueueLength()
		if length == 0{
			break //如果url队列为空 则退出当前循环
		}
		sUrl = models.PopfromQueue()
		//我们应当判断sUrl是否应该被访问过
		if models.IsVisit(sUrl){
			continue
		}
		rsp := httplib.Get(sUrl)
		//设置User-agent以及cookie
		rsp.Header("User-Agent", models.GetUserAgent())
		rsp.Header("Cookie", `bid=gFP9qSgGTfA; __utma=30149280.1124851270.1482153600.1483055851.1483064193.8; __utmz=30149280.1482971588.4.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; ll="118221"; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1483064193%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=5afcf5e5496eab22.1482413017.7.1483066280.1483057909.; __utma=223695111.1636117731.1482413017.1483055857.1483064193.7; __utmz=223695111.1483055857.6.5.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=BDC2DBEDF8958EC838F9D9394CC5D9A0|2cc6ef7952be8c2d5408cb7c8cce2684; ap=1; viewed="1006073"; gr_user_id=e5c932fc-2af6-4861-8a4f-5d696f34570b; __utmc=30149280; __utmc=223695111; _pk_ses.100001.4cf6=*; __utmb=30149280.0.10.1483064193; __utmb=223695111.0.10.1483064193`)
		sMovieHtml,err := rsp.String()
		if err != nil{
			panic(err)
		}
		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		if movieInfo.Movie_name != ""{
			movieInfo.Movie_director        = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character  = models.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_type            = models.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_on_time         = models.GetMovieOnTime(sMovieHtml)
			movieInfo.Movie_grade           = models.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span            = models.GetMovieRunningTime(sMovieHtml)
			movieInfo.Movie_country			= models.GetMovieCountry(sMovieHtml)
			movieInfo.Movie_language		= models.GetMovieLanguage(sMovieHtml)
			movieInfo.Movie_pic				= models.GetMoviePic(sMovieHtml)
			movieInfo.Movie_id				= models.GetMovieid(sMovieHtml)
			movieInfo.Movie_writer			= models.GetMoviewriter(sMovieHtml)
			models.AddMovie(&movieInfo)
		}

		//提取该页面的所有连接
		urls := models.GetMovieUrls(sMovieHtml)

		for _,url := range urls{
			models.PutinQueue(url)
			//fmt.Println(url)
		}
		//sUrl 应当记录到 访问set中
		models.AddToSet(sUrl)

		time.Sleep(time.Second)
	}
	fmt.Println("end of crawl!")
}

