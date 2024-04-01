package server

import (
	"log"
)

type MyServer struct {
	Logger *log.Logger
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

func NewMyServer()MyServer{
	return MyServer{}
}