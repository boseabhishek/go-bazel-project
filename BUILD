load("@io_bazel_rules_go//go:def.bzl", "go_binary",)
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_binary(
    name = "firstgo",
    srcs =["firstgo.go"],
    deps = ["//second:secondgo"],
   
)

go_image(
    name = "firstgo_image",
    srcs = ["firstgo.go"],
    deps = ["//second:secondgo"],
)