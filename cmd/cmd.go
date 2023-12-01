package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func globalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Value: false, Usage: "debug mode"},
		&cli.IntFlag{Name: "scan_num", Aliases: []string{"n"}, Value: 1000, Usage: "scan num"},
		&cli.IntFlag{Name: "timeout", Aliases: []string{"t"}, Value: 5, Usage: "timeout"},
		&cli.StringFlag{Name: "filename", Aliases: []string{"f"}, Value: "input_proxy_list.txt", Usage: "filename"},
		&cli.StringFlag{Name: "output_file", Aliases: []string{"o"}, Value: "proxyscan.txt", Usage: "output_file"},
	}
}

func Main() {
	app := &cli.App{
		Name:     "Proxysan",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Loki",
				Email: "Loki@github.com",
			},
		},
		Usage: "SOCKS4/SOCKS4a/SOCKS5/HTTP/HTTPS fast proxy scanner",
		Flags: globalFlags(),
		Commands: []*cli.Command{
			cmdscan(),
		},
		//Commands:             cmd.Scan,
		EnableBashCompletion: true,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v run error: %v\n", app.Name, err)
		//log.Fatal(err)
	}
}

//var Scan = []*cli.Command{
//	{
//		Name:        "scan",
//		Usage:       "let's scan proxy",
//		Description: "let's scan proxy",
//		Action:      util.Scan,
//		Flags: []cli.Flag{
//			&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Value: false, Usage: "debug mode"},
//			&cli.IntFlag{Name: "scan_num", Aliases: []string{"n"}, Value: 1000, Usage: "scan num"},
//			&cli.IntFlag{Name: "timeout", Aliases: []string{"t"}, Value: 5, Usage: "timeout"},
//			&cli.StringFlag{Name: "filename", Aliases: []string{"f"}, Value: "input_proxy_list.txt", Usage: "filename"},
//			&cli.StringFlag{Name: "output_file", Aliases: []string{"o"}, Value: "proxyscan.txt", Usage: "output_file"},
//		},
//	},
//}
