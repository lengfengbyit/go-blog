package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gotour/blog-service/global"
	"gotour/blog-service/internal/middleware"
	"gotour/blog-service/internal/routers/api"
	v1 "gotour/blog-service/internal/routers/api/v1"
	_ "gotour/blog-service/docs"
	"gotour/blog-service/pkg/limiter"
	"net/http"
	"time"
)

func setupLimiter() limiter.LimiterIface {

	paths := []string{
		"/api/v1/articles",
	}

	methodLimiter := limiter.NewMethodLimiter()

	var rules []limiter.LimiterBucketRule
	for _, path := range paths {
		rule := limiter.LimiterBucketRule{
			Key: path,
			FillInterval: 1 * time.Second,
			Capacity: 2,
			Quantum:1,
		}
		rules = append(rules, rule)
	}

	methodLimiter.AddBuckets(rules...)

	return methodLimiter
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), middleware.Recovery(), middleware.AccessLog())
	r.Use(middleware.Translations()) // 注册中间件
	r.Use(middleware.RateLimiter(setupLimiter()))

	//url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置静态文件目录
	r.StaticFS("./" + global.UploadSetting.ServerUrl, http.Dir(global.UploadSetting.SavePath))

	tag := v1.NewTag()
	article := v1.NewArticle()
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

	}

	// 上传文件
	upload := api.NewUpload()
	r.POST("/upload/file",upload.UploadFile)

	// JWT 获取 token
	r.POST("/auth", api.GetAuth)


	return r
}
