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

https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/hello_world.cc#L1-L7

https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/BUILD.bazel#L1-L5

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

https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/src/main/java/com/example/HelloWorld.java#L1-L8

https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/src/main/java/com/example/BUILD.bazel#L1-L5

You're done! You can build and run your code:
```shell
bazel run //src/main/java/com/example:HelloWorld --java_runtime_version=remotejdk_11
```

If you'd like not to have to use that flag all the time (and I know you don't), create a file at the top of the repo (next to `WORKSPACE`) called `.bazelrc`, and place this line in it:
https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/.bazelrc#L1

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
https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/WORKSPACE#L1-L44

Edit the `.bazelrc` file (or create it if you didn't), and place this in it:
https://github.com/rpwoodbu/bazel_demo/blob/c5ebcb8adee3bf31e41b0e0610a8a57d89759342/.bazelrc#L3-L6

Now do `bazel run` as before, and see it automatically download everything you need!

### Java

Surprise! It was already hermetic! Didn't you notice that we never installed Java?
