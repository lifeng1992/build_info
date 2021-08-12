package build_info

type info struct {
	GitCommit string
	GoVersion string
	Os        string
	Arch      string
	BuildTime string `json:",omitempty"`
}

var (
	gitCommit string
	goVersion string
	os        string
	arch      string
	buildTime string
)

func Info() info {
	return info{
		GitCommit: gitCommit,
		GoVersion: goVersion,
		Os:        os,
		Arch:      arch,
		BuildTime: buildTime,
	}
}
