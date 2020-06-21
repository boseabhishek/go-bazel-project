load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "secondgo",
    srcs =["secondgo.go"],
    importpath = "github.com/boseabhishek/go-bazel-project/second",
    visibility = ["//visibility:public"],
)

go_test(
    name = "secondgotest",
    srcs =["secondgo_test.go"],
    embed = [":secondgo"],
)