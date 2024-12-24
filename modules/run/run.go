package run

import (
	"PassGet/modules/browser"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Run() {
	app := &cli.App{
		Name:            "PassGet",
		Usage:           "A Tool For Windows Post-exploitation Password Crawler",
		UsageText:       "[PassGet.exe (browser/navicat/finalshell/winscp/filezilla/sunlogin/todesk/wifi...)]\nExport password data in windwos\nGithub Link: https://github.com/adeljck/PassGet",
		Version:         "0.0.1b",
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			if err := runAll(); err != nil {
				return err
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "browser",
				Usage: "Get browser data",
				Action: func(c *cli.Context) error {
					err := browser.Get()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("run app error %v", err)
	}
}
func runAll() error {
	fmt.Println("run all....")
	return nil
}
