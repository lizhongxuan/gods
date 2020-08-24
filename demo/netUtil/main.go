package main
var app_route *AppRoute
func init() {
	app_route = NewAppRoute()
}

func main() {

	CreateNet(TCP,socket_io_address, tls_address,
		cert_file, key_file)

	CreateNet(WEBSOCKET,ws_address, wss_address,
		cert_file, key_file)
	if have := HaveNet() ; !have{
		panic("not net")
	}
	select {}
}