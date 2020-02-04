package main

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	proto "github.com/micro/go-plugins/config/source/grpc/proto"

	"github.com/f1renze/the-architect/common/utils/log"
)

func main() {
	var (
		mode = flag.Bool("dev", false, "decide which config will be used")
	)
	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			log.WarnF("Recovered from: %v", r)
		}
	}()

	log.Init()

	files := []string{"config.yml"}
	if *mode {
		log.InfoF("working in dev mode")
		files = []string{"config.dev.yml"}
	}

	if err := loadAndWatchConfigFile(files); err != nil {
		log.Fatal("加载监听配置失败", err)
	}

	srv := grpc.NewServer()
	proto.RegisterSourceServer(srv, new(SourceServer))

	addr := ":9689"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("config server: tcp 端口不可用", err)
	}

	log.InfoF("config server is launched, listening on: %s", addr)

	if err = srv.Serve(listener); err != nil {
		log.Fatal("config server: launch failed", err)
	}
}

// 加载并监听配置
func loadAndWatchConfigFile(files []string) (err error) {
	for i := range files {
		err = config.Load(file.NewSource(
			file.WithPath(files[i]),
		))
		if err != nil {
			log.Fatal("加载配置文件失败", err)
			return
		}
	}

	watcher, err := config.Watch()
	if err != nil {
		log.Fatal("无法启动配置监听器", err)
		return
	}

	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatal("监听配置时错误", err)
				return
			}
			log.Info("配置文件变动", log.Any{
				"change": string(v.Bytes()),
			})
		}
	}()

	return
}

// A gRPC source server should implement the Source proto interface.
type SourceServer struct{}

// 读取 yml 后存储的是 map, 根据 req.Path 作为最外层键取回 value
func (s *SourceServer) Read(ctx context.Context, req *proto.ReadRequest) (resp *proto.ReadResponse, err error) {
	bytes := config.Get(req.Path).Bytes()

	resp = &proto.ReadResponse{
		ChangeSet: &proto.ChangeSet{
			Data:      bytes,
			Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
			Format:    "yml",
			Source:    "file",
			Timestamp: time.Now().Unix(),
		},
	}
	return
}

func (s *SourceServer) Watch(req *proto.WatchRequest, server proto.Source_WatchServer) error {
	bytes := config.Get(req.Path).Bytes()

	resp := &proto.WatchResponse{
		ChangeSet: &proto.ChangeSet{
			Data:      bytes,
			Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
			Format:    "yml",
			Source:    "file",
			Timestamp: time.Now().Unix(),
		},
	}

	if err := server.Send(resp); err != nil {
		log.ErrorF("监听时出现错误", err)
		return err
	}
	return nil
}
