package simple

type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc
