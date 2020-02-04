package registry

import "context"

type Registry interface {
	// 插件名字
	Name() string
	// 初始化
	Init(ctx context.Context, opts ...Options) (err error)
	// 服务注册
	Register(ctx context.Context, service *Service) (err error)
	// 服务反注册，例如服务端停了，注册列表销毁
	UnRegister(ctx context.Context, service *Service) (err error)
	// 服务发现(ip port[] string)
	GetService(ctx context.Context, name string) (service *Service, err error)
}
