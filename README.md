# vbox

This is a VirtualBox API client for Go, heavily inspired by
[vboxgo](https://github.com/th4t/vboxgo/). It is my first piece of Go code, so
code style feedback is welcome.


## Usage

The package should build out of a clean checkout. The `samples` directory
contains reasonable starting code for a new library user.

The package generally follows the VirtualBox XPCOM API, meaning that it is
rather cumbersome. Most users will prefer using a higher-level library, or
building their own abstractions over this library.


## Prerequisites

Building this package requires a [cgo](https://golang.org/cmd/cgo)-enabled Go
installation. Most notably, the cgo requirement seems to preclude
cross-compilation.

The package dynamically loads `VBoxXPCOMC`, a library that implements the
VirtualBox XPCOM/COM API. The library is included with standard VirtualBox
installations.

On systems where VirtualBox is installed at a non-standard location, the
`VBOX_APP_HOME` environment variable must be set to point to the installation
location. The following example accomplishes that on 64-bit Fedora.

```bash
export VBOX_APP_HOME=/usr/lib64/virtualbox
```


## Debugging

When debugging failing tests, it is useful to start the `VBoxSVC` process in a
console, and inspect its console output. VirtualBox and the API client start
the process automatically, but it dies after 5 seconds of inactivity. So,
keeping the VirtualBox UI closed for 5 seconds should get rid of the existing
process.

The following environment variables enable logging in Release builds of
VboxSVC, which are included in the downloadable packages and most
distributions. The variables were listed off of the
[https://www.virtualbox.org/wiki/VBoxMainLogging](VirtualBox wiki).

```bash
export VBOXSVC_RELEASE_LOG=main.e.l.f+gui.e.l.f
export VBOXSVC_RELEASE_LOG_FLAGS="time tid thread"
export VBOXSVC_RELEASE_LOG_DEST=stdout
```


## Vendored VirtualBox SDK

The package contains a patched version of VirtualBox's SDK, under
`third_party/VirtualBoxSDK`. The original version was obtained by unzipping the
SDK package on the
[VirtualBox downloads page](https://www.virtualbox.org/wiki/Downloads).

The vendored version removes all the files that are not related to the C
bindings, which was necessary to keep the repository small.

The patches are minimal changes that work around a
[code generation bug](https://www.virtualbox.org/ticket/14149). The bug
description contains an example of a change. All other changes follow the same
pattern, and can be discovered by trying to build this package with the
original SDK.


## Copyright and Licensing

The licensing situation of this package is complicated due to issues outside of
the author's control. Briefly, the package's contents is licensed under a mix
of the MIT license and the LGPL v2.

All files outside of `third_party/` are (C) Victor Costan 2015, and made
available under the MIT license, which is contained in the `LICENSE` file.

The vendored VirtualBox has most files (under `bindings/c/glue`) licensed under
the MIT license. However, one header file (under `bindings/c/include`) is
licensed under the LGPL v2.

This package is compiled against the LGPL-licensed header, which causes the
LGPL's viral infection clause to kick in, and extends the LGPL provisions over
the entire package.

However, the author wishes that the package could be MIT-licensed. Therefore,
all the files written by the author are made available under the MIT license.
This is intended to help an interested party produce a MIT-licensed package,
given some work for replacing the LGPL-licensed header.
