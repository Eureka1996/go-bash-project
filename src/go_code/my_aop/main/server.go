package main

/**
 * @author wufuqiang
 * @date 2022/7/3 16:56
 */

type Server interface {
}

type sdkHttpServer struct {
}

func NewHttpServer(name string, filters ...FilterBuilder) Server {

	return nil
}
