package controller

import (
	"log"
	"net/http"
	"product-service/repository"
	"product-service/request"
	"product-service/response"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
}

type productController struct {
	prodRepo repository.ProductRepo
}

// UpdateProduct implements ProductController.
func (productController) UpdateProduct(*gin.Context) {
	panic("unimplemented")
}

// AddProduct implements ProductController.
func (p productController) AddProduct(ctx *gin.Context) {
	var req request.RequestAddProduct
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := p.prodRepo.AddProduct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	relasi, err := p.prodRepo.AddProductCategory(req, data.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Log Relasi : ", relasi)

	res := response.ResponseSucessProduct{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Msg:         "Data Berhasil di Tambahkan",
	}

	ctx.JSON(http.StatusOK, res)

}

func NewProductRepo(repo repository.ProductRepo) ProductController {
	return productController{
		prodRepo: repo,
	}
}
