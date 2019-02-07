/*
	Package sql defines models and APIs for the spire backend database.

	Building without cgo

	The default database driver is sqlite, an embedded single-file database
	written in C, and is currently the only dependency that requires cgo.
	However, many installations will not need sqlite, and can therefore exclude
	that dependency. This results in a much more portable binary and is easily
	achieved by setting the CGO_ENABLED environment variable when building. For
	example:

		CGO_ENABLED=0 make build

	Note that this will have affect Go stdlib packages that work differently
	depending upon the availability of cgo. For instance, os/user will use the
	internal implementation that parses /etc/passwd and /etc/group, rather than
	the cgo-based implementation that relies on the C standard library. In the
	case that this behavior is undesirable, the nosqlite build tag is available
	to only disable sqlite.
*/
