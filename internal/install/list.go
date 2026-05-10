package install

import (
	"bufio"
	"os"
	"sort"
	"strings"

	"github.com/Fedsam/fjm/internal/paths"
	"github.com/hashicorp/go-version"
)

type SortDirection int

const (
	ASC SortDirection = iota
	DESC
)

func ListVersions(direction SortDirection) ([]*version.Version, error) {
	versionsDir := paths.VersionsDir()

	entries, err := os.ReadDir(versionsDir)
	if err != nil {
		return nil, err
	}

	var versions []*version.Version

	for _, e := range entries {
		if e.IsDir() {
			parsedVersion, err := version.NewVersion(e.Name())
			if err != nil {
				continue
			}

			versions = append(versions, parsedVersion)
		}
	}
	if direction == ASC {
		sort.Sort(version.Collection(versions))
	} else {
		sort.Slice(versions, func(i, j int) bool {
			return versions[i].GreaterThan(versions[j])
		})
	}
	return versions, nil
}

func ParseRelease(filename string) (map[string]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var res map[string]string
	res = make(map[string]string)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		s := strings.SplitN(line, "=", 2)
		value := strings.Trim(s[1], "\"")
		if len(value) == 0 {
			continue
		}
		res[s[0]] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
