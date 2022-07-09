package scaffold

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"os"
	"prismarin/utils"
	"strings"
)

type TsukiScaffold struct {
	Name        string
	DownloadUrl string
}

func (t TsukiScaffold) GetName() string {
	return t.Name
}

func (t TsukiScaffold) Download() {
	agent := fiber.AcquireAgent()

	req := agent.Request()
	SetHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.SetRequestURI("https://api.github.com/repos/zoey-kaiser/nextjs-ts-scaffolds/releases/latest")

	agent.Parse()
	code, body, errs := agent.Bytes()
	if code != 200 || errs != nil {
		fmt.Println("There was an error while trying to retrieve a scaffold")
		return
	}
	release := &GithubRelease{}
	err := json.Unmarshal(body, release)
	if err != nil {
		fmt.Println("There was an error while trying to retrieve a scaffold")
		return
	}

	fmt.Println("Retrieved latest tag: " + release.TagName)
	fmt.Println("Starting download request")
	out, err := os.Create(t.DownloadUrl)
	resp, err := http.Get("https://github.com/zoey-kaiser/nextjs-ts-scaffolds/releases/download/" + release.TagName + "/" + t.DownloadUrl)
	io.Copy(out, resp.Body)

	out.Close()
	resp.Body.Close()

	utils.UnzipSource(t.DownloadUrl, "")
	err = os.Remove(t.DownloadUrl)
	if err != nil {
		fmt.Println("Cannot delete old zip file")
		return
	}
	fmt.Println("Successfully downloaded scaffold")

	fmt.Println("Please enter your project name:")
	var projectName string
	fmt.Scanln(&projectName)

	fileName := strings.Replace(t.DownloadUrl, ".zip", "", -1)
	os.Rename(fileName, projectName)
	fmt.Println("Your project was successfully created")
}

func SetHeaders(req *fiber.Request) {
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Mobile Safari/537.36")

}

type GithubRelease struct {
	TagName string `json:"tag_name"`
}
