package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	components "github.com/Mohammed785/goBlog/components/post"
	"github.com/Mohammed785/goBlog/database/sqlc"
	"github.com/Mohammed785/goBlog/helpers"
	"github.com/gofiber/fiber/v3"
)

type PostController struct {
	queries   *sqlc.Queries
	validator helpers.Validator
}

func NewPostController(queries *sqlc.Queries, validator helpers.Validator) *PostController {
	return &PostController{
		queries:   queries,
		validator: validator,
	}
}

func (pc *PostController) FindOne(c fiber.Ctx) error {
	param := c.Params("id", "")
	view:=c.Query("update","")==""
	if c.Locals("isAdmin")!=true{
		view=true
	}
	pid, err := strconv.Atoi(param)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "invalid post id")
		return c.SendStatus(http.StatusBadRequest)
	}
	post,err:=pc.queries.FindPostById(context.Background(),int32(pid))
	if err!=nil{
		helpers.SendMsg(c, helpers.ERROR_MSG, err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	return helpers.Render(c,components.PostViewPage(view,&post))
}

func (pc *PostController) List(c fiber.Ctx) error {
	p := c.Query("page", "0")
	page, err := strconv.Atoi(p)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "invalid page")
		return c.SendStatus(http.StatusBadRequest)
	}
	posts,err:=pc.queries.ListPosts(context.Background(),int32(page*25))
	if err!=nil{
		helpers.SendMsg(c, helpers.ERROR_MSG, err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	if p!="0"{
		return helpers.Render(c,components.PostsList(posts,page+1))
	}
	return helpers.Render(c,components.PostsListPage(posts,1))
}

func (pc *PostController) Create(c fiber.Ctx) error {
	var postData sqlc.CreatePostParams
	if err := c.Bind().Form(&postData); err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "please validate your request body")
		return c.SendStatus(http.StatusBadRequest)
	}
	if err := pc.validator.ValidateStruct(postData); err != nil {
		return helpers.Render(c, components.PostForm(nil,pc.validator.ParseValidationError(err)))
	}
	pid, err := pc.queries.CreatePost(context.Background(), postData)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.Redirect().To(fmt.Sprintf("/post/%d", pid))
}

func (pc *PostController) Update(c fiber.Ctx) error {
	param := c.Params("id", "")
	pid, err := strconv.Atoi(param)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "invalid post id")
		return c.SendStatus(http.StatusBadRequest)
	}
	postData := sqlc.UpdatePostParams{Pid: int32(pid)}
	if err := c.Bind().Form(&postData); err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "please validate your request body")
		return c.SendStatus(http.StatusBadRequest)
	}
	if err := pc.validator.ValidateStruct(postData); err != nil {
		return helpers.Render(c, components.PostForm(nil,pc.validator.ParseValidationError(err)))
	}
	err = pc.queries.UpdatePost(context.Background(), postData)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.Redirect().To(fmt.Sprintf("/post/%d", pid))
}

func (pc *PostController) Delete(c fiber.Ctx) error {
	param := c.Params("id", "")
	id, err := strconv.Atoi(param)
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "invalid post id")
		return c.SendStatus(http.StatusBadRequest)
	}
	err = pc.queries.DeletePost(context.Background(), int32(id))
	if err != nil {
		helpers.SendMsg(c, helpers.ERROR_MSG, "couldn't delete post")
		return c.SendStatus(http.StatusBadRequest)
	}
	helpers.SendMsg(c, helpers.SUCCESS_MSG, "post deleted successfully")
	return c.SendStatus(http.StatusNoContent)
}
