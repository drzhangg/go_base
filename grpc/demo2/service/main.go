package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "path/to/proto/file" // 引入自动生成的 proto 代码
)

const (
	port = ":50051"
)

type server struct {
	orders []*pb.Order
}

// 实现 GetOrder 方法，获取单个订单信息
func (s *server) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.Order, error) {
	for _, order := range s.orders {
		if order.Id == req.Id {
			return order, nil
		}
	}
	return nil, fmt.Errorf("order not found")
}

// 实现 ListOrders 方法，获取所有订单信息
func (s *server) ListOrders(ctx context.Context, req *pb.Empty) (*pb.OrderList, error) {
	return &pb.OrderList{Orders: s.orders}, nil
}

// 实现 CreateOrder 方法，创建订单
func (s *server) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	req.Id = int64(len(s.orders) + 1)
	s.orders = append(s.orders, req)
	return req, nil
}

// 实现 UpdateOrder 方法，更新订单
func (s *server) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	for i, order := range s.orders {
		if order.Id == req.Id {
			s.orders[i] = req
			return req, nil
		}
	}
	return nil, fmt.Errorf("order not found")
}

func main() {
	// 初始化 server
	s := &server{}

	// 创建 gRPC Server 对象
	grpcServer := grpc.NewServer()

	// 注册 OrderService
	pb.RegisterOrderServiceServer(grpcServer, s)

	// 监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 启动 gRPC 服务
	log.Printf("server listening at %v")
}
