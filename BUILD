load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/boseabhishek/go-bazel-project
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["firstgo.go"],
    importpath = "github.com/boseabhishek/go-bazel-project",
    visibility = ["//visibility:private"],
    deps = ["//second:go_default_library"],
)

go_binary(
    name = "go-bazel-project",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
