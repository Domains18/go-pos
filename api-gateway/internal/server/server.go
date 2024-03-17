package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

type MyServer struct {
	Logger *log.Logger
}


func (s *MyServer)GetCustomers(ctx *gin.Context){

	ctx.JSON(200,gin.H{"customers:":"list of customers"})
}


//func Server(){
//	router:=gin.Default()
//	myserver=MyServer{}
//
//	options:=handlers.GinServerOptions{
//		BaseURL: '/api',
//		Middlewares: []gin.HandlerFunc{
//			//middlewares
//		},
//		ErrrorHandler: func(c *gin.Context,err error,statuscode int){
//			//error handling logic implementation
//			c.JSON(statuscode,gin.H{"error":err.Error()})
//		},
//	}
//	RegisteredHandlersWithOptions(routes,&myserver,options)
//	router.Run(":8080")
//}

//func NewMyServer()MyServer{
//	return MyServer{}
//}