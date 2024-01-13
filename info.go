package build_info

import (
	"encoding/json"
	"runtime/debug"
)

type info struct {
	GitCommit string
	GoVersion string
	Os        string
	Arch      string
	BuildTime string `json:",omitempty"`
}

var (
	gitCommit = "unknown"
	goVersion = "unknown"
	os        = "unknown"
	arch      = "unknown"
	buildTime = "unknown"
)

func Info() info {
	build, ok := debug.ReadBuildInfo()
	if ok {
		goVersion = build.GoVersion

		settings := map[string]string{}
		for _, s := range build.Settings {
			settings[s.Key] = s.Value
		}

		if v, ok := settings["GOOS"]; ok {
			os = v
		}
		if v, ok := settings["GOARCH"]; ok {
			arch = v
		}

		if v, ok := settings["vcs"]; ok && v == "git" {
			gitCommit = settings["vcs.revision"][:8]
			if modified, ok := settings["vcs.modified"]; ok && modified == "true" {
				gitCommit += "-dirty"
			}
		}
	}

	return info{
		GitCommit: gitCommit,
		GoVersion: goVersion,
		Os:        os,
		Arch:      arch,
		BuildTime: buildTime,
	}
}

func (i info) String() string {
	b, _ := json.MarshalIndent(i, "", "  ")
	return string(b)
}
