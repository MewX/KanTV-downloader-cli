module github.com/MewX/KanTV-downloader-cli

go 1.14

// To sync changes to WORKSPACE, run:
// $ bazel run :gazelle -- update-repos -from_file=go.mod
require (
	github.com/grafov/m3u8 v0.11.0
	github.com/urfave/cli/v2 v2.2.0
)
