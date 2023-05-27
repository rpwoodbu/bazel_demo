# Zero to Hermetic Bazel in 5 minutes

It's fast and easy to start a new project with Bazel. But did you know it's also easy for your new project to be _fully hermetic_, such that the toolchains themselves are fully specified in your source tree?

## Non-hermetic toolchain

Let's start with a non-hermetic toolchain, just to get Bazel off the ground. To prove that there are very few dependencies, we'll do our work in a slimmed-down container.

### Container setup

```shell
docker run --name=bazel-demo -it --rm debian:11-slim /bin/bash
```

Inside the container:
```shell
apt-get update && apt-get install -y wget vim g++
```

Note that we didn't install Bazel. We'll instead install [Bazelisk]([url](https://github.com/bazelbuild/bazelisk)), a transparent wrapper around Bazel which will load whatever version is specified in `.bazelversion`, or the latest released version if that isn't present. That way we can express the whole toolchain, including Bazel itself, within the source tree.
```shell
wget https://github.com/bazelbuild/bazelisk/releases/download/v1.17.0/bazelisk-linux-amd64
chmod +x bazelisk-linux-amd64
mv bazelisk-linux-amd64 /usr/local/bin/bazel
```

### Create the workspace

```shell
mkdir hello_world
cd hello_world
touch WORKSPACE  # This tells Bazel this is the top of the repo.
```

### C++

Now open `vim` (or install and run your favorite `$EDITOR`) and create these two files:

https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/hello_world.cc

```starlark
# BUILD.bazel
cc_binary(
    name = "hello_world",
    srcs = ["hello_world.cc"],
)
```

You're done! You can build and run your code:
```shell
bazel run //:hello_world
```
... or the short form if you're in the directory with the BUILD file:
```
bazel run :hello_world
```

### Java

Create the Java directory structure:
```shell
mkdir -p src/main/java/com/example
cd src/main/java/com/example
```

Create these two files:

```java
// HelloWorld.java
package com.example;

class HelloWorld {
  public static void main(String[] args) {
    System.out.println("Hello, Bazel world!");
  }
}
```

```starlark
# BUILD.bazel
java_binary(
    name = "HelloWorld",
    srcs = ["HelloWorld.java"],
)
```

You're done! You can build and run your code:
```shell
bazel run //src/main/java/com/example:HelloWorld --java_runtime_version=remotejdk_11
```

If you'd like not to have to use that flag all the time (and I know you don't), create a file at the top of the repo (next to `WORKSPACE`) called `.bazelrc`, and place this line in it:
```
build --java_runtime_version=remotejdk_11
```

Now you're _really_ done! You can build and run your code:
```shell
bazel run //src/main/java/com/example:HelloWorld
```
... or the short form if you're in the directory with the BUILD file:
```shell
bazel run :HelloWorld
```

## Hermetic toolchain

Now lets get a hermetic toolchain set, so we don't have to install `g++`. We (sadly) _will_ need `python3` installed to satisfy this hermetic toolchain's needs. (TODO: Write up why it needs this.)
```shell
apt-get autoremove -y g++
apt-get install -y python3
bazel clean  # Be sure it must rebuild everything.
```
(Optionally you may start over and avoid installing `g++` to begin with.)

### C++

Edit the (empty) `WORKSPACE` file and place this in it:
```starlark
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

BAZEL_TOOLCHAIN_TAG = "0.7.2"
BAZEL_TOOLCHAIN_SHA = "f7aa8e59c9d3cafde6edb372d9bd25fb4ee7293ab20b916d867cd0baaa642529"

http_archive(
    name = "com_grail_bazel_toolchain",
    sha256 = BAZEL_TOOLCHAIN_SHA,
    strip_prefix = "bazel-toolchain-{tag}".format(tag = BAZEL_TOOLCHAIN_TAG),
    canonical_id = BAZEL_TOOLCHAIN_TAG,
    url = "https://github.com/grailbio/bazel-toolchain/archive/refs/tags/{tag}.tar.gz".format(tag = BAZEL_TOOLCHAIN_TAG),
)

load("@com_grail_bazel_toolchain//toolchain:deps.bzl", "bazel_toolchain_dependencies")

bazel_toolchain_dependencies()

# This sysroot is used by github.com/vsco/bazel-toolchains.
http_archive(
    name = "org_chromium_sysroot_linux_x64",
    build_file_content = """
filegroup(
  name = "sysroot",
  srcs = glob(["*/**"]),
  visibility = ["//visibility:public"],
)
""",
    sha256 = "84656a6df544ecef62169cfe3ab6e41bb4346a62d3ba2a045dc5a0a2ecea94a3",
    urls = ["https://commondatastorage.googleapis.com/chrome-linux-sysroot/toolchain/2202c161310ffde63729f29d27fe7bb24a0bc540/debian_stretch_amd64_sysroot.tar.xz"],
)

load("@com_grail_bazel_toolchain//toolchain:rules.bzl", "llvm_toolchain")

llvm_toolchain(
    name = "llvm_toolchain",
    llvm_version = "14.0.0",
    sysroot = {
        "linux-x86_64": "@org_chromium_sysroot_linux_x64//:sysroot",
    },    
)

load("@llvm_toolchain//:toolchains.bzl", "llvm_register_toolchains")

llvm_register_toolchains()
```

Edit the `.bazelrc` file (or create it if you didn't), and place this in it:
```
# Don't even look for a local toolchain.
build --repo_env=BAZEL_DO_NOT_DETECT_CPP_TOOLCHAIN=1

build --incompatible_enable_cc_toolchain_resolution
```

Now do `bazel run` as before, and see it automatically download everything you need!

### Java

Surprise! It was already hermetic! Didn't you notice that we never installed Java?
