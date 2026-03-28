# Zero to Hermetic Bazel in 5 minutes

It's fast and easy to start a new project with Bazel. But did you know it's also easy for your new project to be _fully hermetic_, such that the toolchains themselves are fully specified in your source tree?

## Container setup

```shell
docker run --name=bazel-demo -it --rm debian:13-slim /bin/bash
```

Inside the container:
```shell
apt update && apt install -y wget vim
```

Note that we didn't install Bazel. We'll instead install [Bazelisk]([url](https://github.com/bazelbuild/bazelisk)), a transparent wrapper around Bazel which will load whatever version is specified in `.bazelversion`, or the latest released version if that isn't present. That way we can express the whole toolchain, including Bazel itself, within the source tree.
```shell
wget https://github.com/bazelbuild/bazelisk/releases/download/v1.28.1/bazelisk-amd64.deb
dpkg -i bazelisk-amd64.deb
```

## Create the source tree

```shell
mkdir hello_world
cd hello_world
touch MODULE.bazel  # This tells Bazel this is the top of the repo.
```

## C++

Now open `vim` (or install and run your favorite `$EDITOR`), and edit `MODULE.bazel`. Add the dependencies for the C++ toolchain and ruleset, and register the toolchain:

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/MODULE.bazel#L1-L4

Then create these two files:

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/hello_world.cc#L1-L7

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/BUILD.bazel#L1-L7

You're done! You can build and run your code:
```shell
bazel run //:hello_world
```
... or the short form if you're in the directory with the BUILD file:
```
bazel run :hello_world
```

### Pro-tip

To be 100% sure that Bazel doesn't accidentally use a locally-installed C++ toolchain due to a misconfiguration (prefer it to fail instead), put this in `.bazelrc` next to the `MODULE.bazel` file:

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/.bazelrc#L3-L4

## Java

Open `vim` (or install and run your favorite `$EDITOR`), and edit `MODULE.bazel`. Add the dependencies for the Java ruleset:

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/MODULE.bazel#L6-L7

Then put this in `.bazelrc` next to the `MODULE.bazel` file to bring in the Java toolchain (JDK):

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/.bazelrc#L1

Create the Java directory structure:
```shell
mkdir -p src/main/java/com/example
cd src/main/java/com/example
```

Create these two files:

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/src/main/java/com/example/HelloWorld.java#L1-L8

https://github.com/rpwoodbu/bazel_demo/blob/27dfdc36cf3f60f78ce5b965ff912e9010083db7/src/main/java/com/example/BUILD.bazel#L1-L7

You're done! You can build and run your code:
```shell
bazel run //src/main/java/com/example:HelloWorld
```

... or the short form if you're in the directory with the BUILD file:
```shell
bazel run :HelloWorld
```
