package api

//接口api继承
type BaseApi interface {
	//回调绑定方法
	CallBindingApi()
}

type BaseApiObserver interface {
	//添加Api
	addApi(baseApis []BaseApi)
	//绑定api method方法
	bindingApi(baseApi BaseApi, f func(map[string]interface{}) map[string]interface{}, requestMethod string, path []string)
}
