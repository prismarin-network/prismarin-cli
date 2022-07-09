package scaffold

import "strings"

type Scaffold interface {
	GetName() string
	Download()
}

var Scaffolds map[string][]Scaffold = map[string][]Scaffold{}

func Init() {
	nextjs := []Scaffold{
		TsukiScaffold{
			Name:        "minimal",
			DownloadUrl: "minimal.zip",
		},
		TsukiScaffold{
			Name:        "with-tailwind",
			DownloadUrl: "with-tailwind.zip",
		},
		TsukiScaffold{
			Name:        "with-tailwind-mantine",
			DownloadUrl: "with-tailwind-mantine.zip",
		},
		TsukiScaffold{
			Name:        "with-tailwind-daisyui",
			DownloadUrl: "with-tailwind-daisyui.zip",
		},
		TsukiScaffold{
			Name:        "with-tailwind-flowbite",
			DownloadUrl: "with-tailwind-flowbite.zip",
		},
	}
	Scaffolds["nextjs"] = nextjs
}

func ExistCategory(category string) bool {
	return Scaffolds[strings.ToLower(category)] != nil
}

func GetCategory(category string) []Scaffold {
	return Scaffolds[strings.ToLower(category)]
}

func ExistScaffold(scaffolds []Scaffold, name string) bool {
	return GetScaffold(scaffolds, name) != nil
}

func GetScaffold(scaffolds []Scaffold, name string) Scaffold {
	for _, scaffold := range scaffolds {
		if strings.ToLower(scaffold.GetName()) == strings.ToLower(name) {
			return scaffold
		}
	}
	return nil
}
