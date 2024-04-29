package handlers

import (
	"distriai-index-solana/common"
	"distriai-index-solana/model"
	"distriai-index-solana/utils/logs"
	"distriai-index-solana/utils/resp"
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"strings"
	"time"
)

type DatasetCreateReq struct {
	Name    string `binding:"required,max=50"`
	Scale   uint8  `binding:"required"`
	License uint8  `binding:"required"`
	Type1   uint32 `binging:"required"`
	Type2   uint32 `binding:"required"`
	Tags    string `binding:"max=128"`
}

// DatasetCreate user create a dataset.
func DatasetCreate(context *gin.Context) {
	account := getAuthAccount(context)
	var req DatasetCreateReq
	if err := context.ShouldBindJSON(&req); err != nil {
		resp.Fail(context, err.Error())
		return
	}

	var count int64
	tx := common.Db.Model(&model.Dataset{}).
		Where("owner = ?", account).
		Where("name = ?", req.Name).
		Count(&count)
	if tx.Error != nil {
		resp.Fail(context, "Database error")
		return
	}
	if count > 0 {
		resp.Fail(context, "Duplicate dataset name")
		return
	}

	//generate random number
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	likes := uint32(rnd.Int31n(100))
	// create new dataset struct
	dataSet := model.Dataset{
		Owner:     account,
		Name:      req.Name,
		Scale:     req.Scale,
		License:   req.License,
		Type1:     req.Type1,
		Type2:     req.Type2,
		Tags:      req.Tags,
		Downloads: likes + uint32(rnd.Int31n(1000)),
		Likes:     likes,
	}
	if err := common.Db.Create(&dataSet).Error; err != nil {
		logs.Error(fmt.Sprintf("Database error: %v \n", err))
		resp.Fail(context, "Database error")
		return
	}

	resp.Success(context, "")
}

type DatasetPresignReq struct {
	Id       uint   `binding:"required"`
	FilePath string `binding:"required"`
	Method   string `binding:"required,oneof= PUT DELETE"`
}

func DatasetPresign(context *gin.Context) {
	account := getAuthAccount(context)
	var req DatasetPresignReq
	if err := context.ShouldBindJSON(&req); err != nil {
		resp.Fail(context, err.Error())
		return
	}

	var dataset model.Dataset
	if err := common.Db.Where("id = ? AND owner = ?", req.Id, account).Model(&dataset).Error; err != nil {
		resp.Fail(context, "Dataset not found")
	}
	objectKey := fmt.Sprintf("dataset/%s/%s/%s", dataset.Owner, dataset.Name, req.FilePath)
	presignedPutRequest := new(v4.PresignedHTTPRequest)
	var err error
	switch req.Method {
	case "PUT":
		presignedPutRequest, err = common.S3Presigner.PutObject("distriai", objectKey, 3600)
	case "DELETE":
		presignedPutRequest, err = common.S3Presigner.DeleteObject("distriai", objectKey)
	}
	if err != nil {
		resp.Fail(context, "Create presigned URL error")
		return
	}
	resp.Success(context, presignedPutRequest.URL)
}

type DatasetListReq struct {
	Name    string
	Type1   uint32
	Type2   uint32
	OrderBy string
}

type DatasetListResponse struct {
	List []model.Dataset
	PageResp
}

func DatasetList(context *gin.Context) {
	var req DatasetListReq
	if err := context.ShouldBindBodyWith(&req, &binding.JSON); err != nil {
		resp.Fail(context, err.Error())
		return
	}

	var response DatasetListResponse
	dataset := model.Dataset{Type1: req.Type1, Type2: req.Type2}
	tx := common.Db.Model(&dataset).Where(&dataset)
	if "" != req.Name {
		name := strings.ReplaceAll(req.Name, "%", "\\%")
		tx.Where("name LIKE ?", "%"+name+"%")
	}
	if err := tx.Count(&response.Total).Error; err != nil {
		resp.Fail(context, "Database error")
		return
	}

	switch req.OrderBy {
	case "downloads DESC", "likes DESC":
		tx.Order(req.OrderBy)
	default:
		tx.Order("updated_at DESC")
	}
	if tx.Scopes(Paginate(context)).Find(&response.List).Error != nil {
		resp.Fail(context, "Database error")
		return
	}

	resp.Success(context, response)
}

type DatasetGetReq struct {
	Owner string `binding:"required"`
	Name  string `binding:"required"`
}

func DatasetGet(context *gin.Context) {
	var req DatasetGetReq
	if err := context.ShouldBindUri(&req); err != nil {
		resp.Fail(context, err.Error())
		return
	}
	var dataset model.Dataset
	tx := common.Db.
		Where("owner = ?", req.Owner).
		Where("name = ?", req.Name).
		Take(&dataset)
	if tx.Error != nil {
		resp.Fail(context, "Dataset not found")
		return
	}

	resp.Success(context, dataset)
}
