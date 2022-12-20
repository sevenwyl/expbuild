module github.com/expbuild/expbuild

go 1.19

replace github.com/bazelbuild/remote-apis/build/bazel/semver => github.com/expbuild/expbuild/proto/build/bazel/semver v0.0.0

require (
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	cloud.google.com/go/longrunning v0.3.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20221014081412-f15817d10f9b // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.4.0 // indirect
)