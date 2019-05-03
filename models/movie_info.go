package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"regexp"
	"strconv"
	"strings"
)



type MovieInfo struct{
	Id 						int64
	Movie_id 				int64
	Movie_name 				string
	Movie_pic 				string
	Movie_director 			string
	Movie_writer 			string
	Movie_country 			string
	Movie_language 			string
	Movie_main_character	string
	Movie_type 				string
	Movie_on_time 			string
	Movie_span 				string
	Movie_grade 			string
}
var db *gorm.DB
func Setup(db_type,db_user,db_pass,db_name,db_host string) {
	var err error
	db, err = gorm.Open(db_type,  fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",db_user,db_pass,db_host,db_name, ))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	//db.AutoMigrate(&MovieInfo{})
	fmt.Println("mysql init ok")
	//defer db.Close()
}



func AddMovie(movieInfo *MovieInfo){
	movieInfo.Id = 0
	fmt.Println("正在保存抓取到电影:",movieInfo.Movie_name)
	db.Table("movie_info").Create(&movieInfo)
}

func GetMovieDirector(movieHtml string) string{
	if movieHtml == ""{
		return ""
	}
	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieName(movieHtml string)string{
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieMainCharacters(movieHtml string)string{
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	mainCharacters := ""
	for _,v := range result{
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")
}

func GetMovieGrade(movieHtml string)string{
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}
	return string(result[0][1])
}

func GetMoviePic(movieHtml string)string{
	reg := regexp.MustCompile(`<img\ssrc="(.*?)"\stitle="点击看更多海报"\s.*?/>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0{
		return ""
	}
	return string(result[0][1])
}
func GetMovieid(movieHtml string)int64{
	reg := regexp.MustCompile(`<link\srel="alternate" href="(.*?)"\s/>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0{
		return 0
	}
	mid := strings.Split(string(result[0][1]),"/")
	aa, _ := strconv.ParseInt(mid[5],10, 64)
	return aa
}
func GetMoviewriter(movieHtml string)string{
	reg := regexp.MustCompile(`<span ><span class='pl'>编剧</span>:.*`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	result1 := fmt.Sprintf("%s",result)
	reg1 := regexp.MustCompile(`<a\shref=".*?">(.*?)</a>`)

	result = reg1.FindAllStringSubmatch(result1, -1)

	if len(result) == 0{
		return ""
	}
	mainCharacters := ""
	for _,v := range result{
		mainCharacters += v[1] + "/"
	}
	return strings.Trim(mainCharacters, "/")
}

func GetMovieCountry(movieHtml string)string{
	reg := regexp.MustCompile(`<span\sclass="pl">制片国家/地区:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0{
		return ""
	}
	return string(strings.TrimSpace(result[0][1]))
}
func GetMovieLanguage(movieHtml string)string{
	reg := regexp.MustCompile(`<span\sclass="pl">语言:</span>(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0{
		return ""
	}
	return string(strings.TrimSpace(result[0][1]))
}

func GetMovieGenre(movieHtml string)string{
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	movieGenre := ""
	for _,v := range result{
		movieGenre += v[1] + "/"
	}
	return strings.Trim(movieGenre, "/")
}

func GetMovieOnTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieRunningTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}


func GetMovieUrls(movieHtml string)[]string{
	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	var movieSets []string
	for _,v := range result{
		movieSets = append(movieSets, v[1])
	}

	return movieSets
}