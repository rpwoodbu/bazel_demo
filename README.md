# Bazel Demos

[Bazel](https://bazel.build/) isn't as hard as it used to be. This repo
endeavors to strip away all the cruft and illustrate the essentials. You'll find
that moving to Bazel (or better yet, _starting_ with it) is not that hard.

> [!NOTE]
> This Git repo contains multiple Bazel "repos" (naming is hard), each with its
> own `MODULE.bazel` file, focusing on a different language. This keeps
> things simple and focused. **But this is _not_ how projects should normally be
> structured!** Bazel shines when used in a [monorepo](https://monorepo.tools/).
> Bazel brings together multiple languages and ecosystems into one coherent
> experience. All these Bazel repos could, in principle, be combined into a
> single repo at the top of the source tree. If you use ideas from these repos,
> combine them into a single Bazel repo in your source tree, with a single
> `MODULE.bazel` file to rule them all.

## Try it out!

Once you've installed Bazel (or Bazelisk; see below) and cloned this repo, just
`cd` to any of the directories and try building/running any of the targets in
the `BUILD.bazel` files.

### Use Bazelisk

Instead of installing Bazel itself, you should actually install
[Bazelisk](https://github.com/bazelbuild/bazelisk), a transparent wrapper around
the `bazel` command which will load whatever version is specified in
`.bazelversion`, or the latest released version if that isn't present. That way
we can specify the whole toolchain, including Bazel itself, within the source
tree.

### Quick start

Impatient? Take 2 minutes and just do this on a Debian/Ubuntu machine:
```shell
cd /tmp
# Install Bazelisk.
curl -OL https://github.com/bazelbuild/bazelisk/releases/download/v1.28.1/bazelisk-amd64.deb
dpkg -i bazelisk-amd64.deb
# Clone the repo.
git clone https://github.com/rpwoodbu/bazel_demo
cd bazel_demo
# Go into the `cc` repo and run the C++ "Hello World".
cd cc
bazel run :hello_world
# Run all the tests while you're there (there's only one).
bazel test ...
# Extra credit: Try running the tests in the Bazel repos for the other languages.
```

## Hermetic out-of-the-box

All of these repos are configured to be fully hermetic. This means that the only
environmental dependency you have is Bazelisk. There is no need for a
locally-installed compiler or interpreter for any language, nor even Bazel
itself.

### Prove it!

```shell
docker run --name=bazel-demo -it --rm debian:13-slim /bin/bash
# In the container:
apt update && apt install -y curl git
curl -OL https://github.com/bazelbuild/bazelisk/releases/download/v1.28.1/bazelisk-amd64.deb
dpkg -i bazelisk-amd64.deb
git clone https://github.com/rpwoodbu/bazel_demo
cd bazel_demo
cd cc  # ...or any other repo you want.
bazel test ...
cd ../java  # ...or any other repo you want.
bazel test ...
# Rinse and repeat.
```

## But wait, there's more!

This demo is just meant to be a basic guide to the essentials. But Bazel can do
so much more than this: remote caching and execution, static analysis, linting
and formatting, packaging and pushing, etc., etc.

See the documentation for the ruleset for any language to learn more about how
to interface with their ecosystem. You can find all the rulesets and much more
at the [Bazel Central Registry](https://registry.bazel.build/).

Of course, you can find much more at https://bazel.build/.
