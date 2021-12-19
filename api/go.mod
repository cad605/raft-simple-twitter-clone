module simple-twitter/api

go 1.16

replace simple-twitter.com/database/rpc/proto => ../backend

replace simple-twitter.com/database/concensus => ../backend

replace simple-twitter.com/database/server => ../backend

replace simple-twitter.com/backend => ../backend

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-sockaddr v1.0.2
	github.com/ogier/pflag v0.0.1
	google.golang.org/grpc v1.42.0
	simple-twitter.com/backend v0.0.0-00010101000000-000000000000
)
