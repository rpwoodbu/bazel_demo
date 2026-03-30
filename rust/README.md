# Rust

## Caching
The sysroot seems to have a bad interaction when using a remote cache (even
`--disk_cache`). It is likely implemented as a TreeArtifact, which can have
subtle bugs. A different sysroot may be necessary (or no sysroot, depending on
the local system, which is non-hermetic).

## Hermeticity
This one isn't _perfectly_ hermetic. It requires the use of `toolchains_llvm`,
which uses the main LLVM Clang binary distribution. While Bazel fetches this
hermetically, the included binaries are dynamically linked against `libxml` and
`libtinfo`, so these must be installed on your system.
