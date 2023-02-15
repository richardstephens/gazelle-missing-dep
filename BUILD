load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/example/project
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "project_lib",
    srcs = ["main.go"],
    importpath = "github.com/example/project",
    visibility = ["//visibility:private"],
    deps = ["@com_github_census_instrumentation_opencensus_proto//gen-go/trace/v1:trace"],
)

go_binary(
    name = "project",
    embed = [":project_lib"],
    visibility = ["//visibility:public"],
)
