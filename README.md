# `buildinfo`

This project is a convenience module to wrap `debug.buildinfo` and provide easy to access version information for a running program.  Programs built using this module will derive their version information from one of the following:

* The `github.com/packetspace/buildinfo.version` variable, when set using the `-X` build flag
* The version target of a `go install` build
* A default string (currently `(unknown)`)

Version information is always available to a running program at `buildinfo.Default.AppVersion()`, with extra information available from additional methods.

## Contributing

Contributions of code and documentation are welcome, as are bug reports and suggestions!  Please see the project's [CONTRIBUTING](CONTRIBUTING.md) page for more information.
