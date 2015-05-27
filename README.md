# vbox

This is a VirtualBox API client for Go, heavily inspired by
[vboxgo](https://github.com/th4t/vboxgo/). It is my first piece of Go code, so
don't expect too much from it.


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
