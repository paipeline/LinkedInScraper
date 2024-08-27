package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// 全局变量
var (
	mongoClient *mongo.Client
	database    *mongo.Database
)

// 初始化 MongoDB 连接
func initMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB 连接 URI
	uri := "mongodb://localhost:27017"

	// 创建 MongoDB 客户端
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("无法创建 MongoDB 客户端: %v", err)
	}

	// 检查连接
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("无法连接到 MongoDB: %v", err)
	}

	mongoClient = client
	database = client.Database("linkedin_scraper")

	log.Println("成功连接到 MongoDB!")
	return nil
}

// 设置 Gin 路由
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

// 主函数
func main() {
	// 初始化 MongoDB
	if err := initMongoDB(); err != nil {
		log.Fatalf("初始化 MongoDB 失败: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Printf("断开 MongoDB 连接时发生错误: %v", err)
		}
	}()

	// 获取 jobs 集合
	jobsCollection := database.Collection("jobs")
	log.Printf("使用数据库: %s, 集合: %s\n", database.Name(), jobsCollection.Name())

	// 设置 Gin 路由
	router := setupRouter()

	// 启动服务器
	log.Println("正在启动服务器...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
