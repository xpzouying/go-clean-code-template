package constant

var (
	// Version of the application.
	Version string

	// BuildTime is the time the application was built. Inject by makefile.
	BuildTime   string
	GitBranch   string
	GitRevision string
)
