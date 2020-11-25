package v1

import (
	"github.com/gin-gonic/gin"
	"gotour/blog-service/global"
	"gotour/blog-service/internal/service"
	"gotour/blog-service/pkg/app"
	"gotour/blog-service/pkg/convert"
	"gotour/blog-service/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取一个文章
// @Produce json
// @Param id path int true "文章ID"
// @Param title query string false "文章标题" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	response := app.NewResponse(c)

	id := convert.StrTo(c.Param("id")).MustUInt32()
	if id <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	ser := service.New(c.Request.Context())
	art, err := ser.GetArticle(id)
	if err != nil {
		global.Logger.Errorf(c, err.Error())
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(art)
	return
}

// @Summary 获取多个文章
// @Produce json
// @Param title query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := &service.ListArticleRequest{}
	response := app.NewResponse(c)

	// 参数校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "Article List app.BindAndValid err: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	ser := service.New(c.Request.Context())
	paper := &app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}

	total, err := ser.CountArticle(&service.CountArticleRequest{
		Title: param.Title,
		State: param.State,
	})

	if err != nil {
		global.Logger.Errorf(c, "service.CountArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}

	articleList, err := ser.GetArticleList(param, paper)
	if err != nil {
		global.Logger.Errorf(c, "Article List service.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}

	response.ToResponseList(articleList, total)
	return
}

// @Summary 创建文章
// @Produce json
// @Param title body string true "文章名称" maxlength(100)
// @Param desc body string true "文章描述" maxlength(100)
// @Param cover_image_url body string true "封图" maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := &service.CreateArticleRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	ser := service.New(c.Request.Context())
	err := ser.CreateArticle(param)
	if err != nil {
		global.Logger.Errorf(c, "service.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 修改文章
// @Produce json
// @Param id path int true "文章ID"
// @Param title body string true "文章名称" maxlength(100)
// @Param desc body string true "文章描述" maxlength(100)
// @Param cover_image_url body string true "封图" maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := &service.UpdateArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
		err := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(err)
		return
	}

	ser := service.New(c.Request.Context())
	err := ser.UpdateArticle(param)
	if err != nil {
		global.Logger.Error(c, "service.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {

	param := &service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
		err := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(err)
		return
	}

	ser := service.New(c.Request.Context())
	err := ser.DeleteArticle(param)
	if err != nil {
		global.Logger.Errorf(c,"service.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}
