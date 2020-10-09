package main

import "context"

type HttpServer struct {
	pool *Pool
}
//define request
type Request struct {
	// all request parameters like header/body content/
}

//define request handling
func (request Request)handle(ctx context.Context) {
	//handle request, may take a long time
}
func NewHttpServer(workerNum int,jobQueueNum int)*HttpServer{
	return &HttpServer{pool:NewPool(workerNum,jobQueueNum)}
}

//queue: A channel consists  of all incoming HTTP request with type
func (server *HttpServer)serve(c chan *Request) {
	for request :=range c{
		server.pool.JobQueue<- func() {
			request.handle(context.TODO())
		}
	}
}
func (server *HttpServer)Shutdown(){
	server.pool.WaitAll()
	server.pool.Release()
}

