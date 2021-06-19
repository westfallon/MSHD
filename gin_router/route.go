package gin_router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/westfallon/naseum/db"
	"github.com/westfallon/naseum/service"
	"io/ioutil"
	"time"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/DeathStatistics", func(context *gin.Context) {
		values := service.GetDeathStatistics()
		context.JSON(200, values)
	})
	router.GET("/CivilStructure", func(context *gin.Context) {
		values := service.GetCivilStructure()
		context.JSON(200, values)
	})
	router.GET("/CommDisaster", func(context *gin.Context) {
		values := service.GetCommDisaster()
		context.JSON(200, values)
	})
	router.GET("/CollapseRecord", func(context *gin.Context) {
		values := service.GetCollapseRecord()
		context.JSON(200, values)
	})
	router.GET("/DisasterInfo", func(context *gin.Context) {
		values := service.GetDisasterInfo()
		context.JSON(200, values)
	})
	router.GET("/DisasterRequest", func(context *gin.Context) {
		values := service.GetDisasterRequest()
		context.JSON(200, values)
	})
	router.POST("/request", func(context *gin.Context) {
		var request db.DisasterRequest
		err := context.ShouldBindJSON(&request)
		if err != nil {
			fmt.Println(err)
			context.String(400, "请求失败")
			return
		}
		request.Date = time.Now().Format("2006-01-02 15:04:05")
		request.Id = service.RandSeq(19)
		fmt.Println(request)
		err = db.GetDB().Create(request).Error
		if err != nil {
			fmt.Println(err)
			context.String(400, "请求失败")
			return
		}
		context.String(200, "请求成功")
	})
	router.POST("/upload", func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				context.String(400, "文件类型错误")
			}
		}()
		rFile, err := context.FormFile("file")
		if err != nil {
			panic(err)
		}
		file, err := rFile.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		var requests []*db.DeathStatistics
		err = json.Unmarshal(content, &requests)
		if err != nil {
			panic(err)
		}
		for _, request := range requests {
			request.Id = service.RandSeq(19)
			err = db.GetDB().Create(request).Error
			if err != nil {
				panic(err)
			}
		}
		context.String(200, "上传成功")
	})

	err := router.Run(":8181")
	if err != nil {
		return
	}
}
