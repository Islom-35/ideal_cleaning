package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	grpc_client "example.com/m/internal/api-gateway/delivery/grpc"
	pb "example.com/m/internal/genproto/product/pb"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ProductInput struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
	Count uint32 `json:"count"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ProductResponse struct {
	ID                   int32      `json:"ID"`
	Name                 string     `json:"name"`
	Price                int32      `json:"price"`
	Count                int32      `json:"count"`
	CreatedAt            time.Time 	`json:"created_at"`
	UpdatedAt            time.Time 	`json:"updated_at"`
}

type Handler struct {
	productClient grpc_client.Client
}

func NewHandler(client grpc_client.Client) *Handler {
	return &Handler{
		productClient: client,
	}
}

// @Summary Create a new product
// @Security ApiKeyAuth
// @Description Create a new product with the provided JSON data
// @Tags products
// @Accept json
// @Produce json
// @Param product body ProductInput true "Product object that needs to be created"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var productInp ProductInput

	if err = json.Unmarshal(reqBytes, &productInp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if productInp.Count==0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Count of products can not be nol"})
		return
	}
	if productInp.Price==0{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price can not be nol"})
		return
	}
	if productInp.Name==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name can not be empty"})
		return
	}

	product := pb.ProductRequest{
		Name:  productInp.Name,
		Price: productInp.Price,
		Count: productInp.Count,
	}

	err = h.productClient.CreateProduct(ctx, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

// @Summary Get a product by ID
// @Security ApiKeyAuth
// @Description Get product details by providing its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
func (h *Handler) GetProductByID(c *gin.Context) {
	id, err := getIdFromRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inp :=pb.ID{
		ID: id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product,err := h.productClient.GetProductByID(ctx,&inp)

	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func getIdFromRequest(c *gin.Context) (uint32, error) {
	idStr := c.Param("id")
	if idStr == "" {
		return 0, errors.New("id must be provided")
	}

	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	if id64 == 0 {
		return 0, errors.New("id can't be 0")
	}

	id := uint32(id64)
	return id, nil
}
