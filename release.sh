#!/bin/bash
# The script for building all cross-platform binaries.
versionCode=$(cat kantv/kantv_cli.go | grep Version | cut -d '"' -f2)
versionName="v$versionCode"
echo "Current release version is $versionName"

# Release one platform.
# $1 platform name
# $2 binary suffix
# $3 final binary platform name
# $4 extra parameters, e.g. for CC, etc.
release () {
    echo "Releasing $1"
    # Build the file.
    mkdir release
    bazel build --platforms "@io_bazel_rules_go//go/toolchain:$1" :kantv
    targetFile="release/kantv-$versionName-$3$2"
    cp -f "bazel-bin/$1_pure_stripped/kantv$2" "$targetFile"

    # Generate a checksum.
    sha1sum -b "$targetFile" > "$targetFile.sha1"
    echo "Finished $1"
}

# Command to list all supported platforms:
#   $ bazel query 'kind(config_setting, @io_bazel_rules_go//go/platform:all)'
# x86
release "linux_386" "" "linux-x86" ""
release "windows_386" ".exe" "windows-x86" ""
release "darwin_386" "" "macos-x86" ""

# x64
release "linux_amd64" "" "linux-x64" ""
release "windows_amd64" ".exe" "windows-x64" ""
release "darwin_amd64" "" "macos-x64" ""

# ARM32

# ARM64
