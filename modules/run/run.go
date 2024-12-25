package run

import (
	"PassGet/modules/browser"
	"PassGet/modules/finalshell"
	"PassGet/modules/sunlogin"
	"PassGet/modules/todesk"
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
			}, {
				Name:  "nav",
				Usage: "Get navicat data",
				Action: func(c *cli.Context) error {
					err := browser.Get()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "scp",
				Usage: "Get winscp data",
				Action: func(c *cli.Context) error {
					err := browser.Get()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "filez",
				Usage: "Get filezilla data",
				Action: func(c *cli.Context) error {
					err := browser.Get()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "wifi",
				Usage: "Get wifi data",
				Action: func(c *cli.Context) error {
					err := browser.Get()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "sun",
				Usage: "Get sunlogin data",
				Action: func(c *cli.Context) error {
					err := sunlogin.Get()
					if err == nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "tdesk",
				Usage: "Get todesk data",
				Action: func(c *cli.Context) error {
					err := todesk.Get()
					if err == nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "fshell",
				Usage: "Get finalshell data",
				Action: func(c *cli.Context) error {
					_, ServerDetails := finalshell.Get("")
					if ServerDetails == nil {
						log.Fatalf("get finalshell data error")
					}
					fmt.Println(ServerDetails)
					return nil
				},
			}, {
				Name:  "svn",
				Usage: "Get TortoiseSVN data",
				Action: func(c *cli.Context) error {
					err := browser.Get()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "xman",
				Usage: "Get Xmanager data",
				Action: func(c *cli.Context) error {
					err := browser.Get()
					if err != nil {
						log.Fatalf("get browser data error %v", err)
					}
					return nil
				},
			}, {
				Name:  "mxterm",
				Usage: "Get MobaltXterm data",
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
