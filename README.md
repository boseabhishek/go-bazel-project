# go-bazel-project

### bazel query

> 

### build docker image

> bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 :firstgo_image  

### run docker image

> bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 :firstgo_image