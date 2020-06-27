# go-bazel-project

Simple Golang application built using [Bazel](https://bazel.build/).

## Resources

- official bazel documentation 
- how to build some [other bazel projects](https://docs.bazel.build/versions/3.3.0/bazel-overview.html#tutorials)
- shameless plug: a [medium article I wrote](https://medium.com/@abhbose6/bazel-101-2b0272b15da8) on java based Bazel project 

## Recipes

### bazel build (usually a go_library or go_binary or even go_test target)

> bazel build :firstgo

> bazel build //second:secondgo

### bazel test

> bazel test //second:secondgotest

### bazel run

> bazel run :firstgo

### bazel query

> bazel query "deps(:firstgo)"

> bazel query "deps(//second:secondgo)"

### build docker image

> bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 :firstgo_image  

### run docker image

> bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 :firstgo_image

(it's impotant to provide the target os/architecture say for Linux cOntainer where you want to run the image on)
## Some popular Open Source repos which use bazel

- [Kubernetes](https://github.com/kubernetes/kubernetes)
- [Tensorflow](https://github.com/tensorflow/tensorflow)
- [Find many more here](https://github.com/bazelbuild/bazel/wiki/Bazel-Users)
