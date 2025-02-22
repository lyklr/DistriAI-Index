package routes

import (
	"distriai-index-solana/handlers"
	"distriai-index-solana/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes unified management routes.
func RegisterRoutes(engine *gin.Engine) {
	engine.POST("/faucet", handlers.Faucet)
	engine.POST("/webhook", handlers.Webhook)
	mailbox := engine.Group("/mailbox")
	{
		mailbox.POST("/subscribe", handlers.Subscribe)
		mailbox.POST("/unsubscribe", handlers.Unsubscribe)
	}
	user := engine.Group("/user")
	{
		user.POST("/login", handlers.Login)
	}
	machine := engine.Group("/machine")
	{
		machine.POST("/filter", handlers.MachineFilter)
		machine.POST("/market", handlers.MachineMarket)
		machine.POST("/mine", handlers.MachineMine)
		machine.GET("/:Owner/:Uuid", handlers.MachineGet)
		machine.POST("/info/cached", handlers.MachineInfoCached)
	}
	order := engine.Group("/order")
	{
		order.POST("/mine", handlers.OrderMine)
		order.POST("/all", handlers.OrderAll)
		order.GET("/:Uuid", handlers.OrderGet)
	}
	reward := engine.Group("/reward")
	{
		reward.POST("/total", handlers.RewardTotal)
		reward.POST("/claimable/list", handlers.RewardClaimableList)
		reward.POST("/period/list", handlers.RewardPeriodList)
		reward.POST("/machine/list", handlers.RewardMachineList)
	}
	log := engine.Group("/log")
	{
		log.POST("/add", handlers.LogAdd)
		log.POST("/list", handlers.LogList)
	}
	model := engine.Group("/model")
	{
		model.POST("/list", handlers.ModelList)
		model.GET("/:Owner/:Name", handlers.ModelGet)
		model.POST("/download", handlers.ModelDownload)
		//model.GET("likecount", handlers.ModelLikeCount)
	}
	modelAuth := engine.Group("/model", middleware.Jwt())
	{
		modelAuth.POST("/likes", handlers.ModelLikes)
		modelAuth.POST("/presign", handlers.ModelPresign)
		modelAuth.POST("/like", handlers.ModelLike)
		modelAuth.GET("/islike", handlers.ModelIsLike)
	}
	dataset := engine.Group("/dataset")
	{
		dataset.POST("/list", handlers.DatasetList)
		dataset.GET("/:Owner/:Name", handlers.DatasetGet)
		dataset.POST("/download", handlers.DatasetDownload)
		//dataset.GET("/likecount", handlers.DatasetLikeCount)
	}
	datasetAuth := engine.Group("/dataset", middleware.Jwt())
	{
		datasetAuth.POST("/likes", handlers.DatasetLikes)
		datasetAuth.POST("/presign", handlers.DatasetPresign)
		datasetAuth.POST("/like", handlers.DatasetLike)
		datasetAuth.GET("/islike", handlers.DatasetIsLike)
	}
}
