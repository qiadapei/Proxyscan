package cmd

import (
	"github.com/qiadapei/Proxyscan/util"
	"github.com/urfave/cli/v2"
)

func cmdscan() *cli.Command {
	return &cli.Command{
		Name:        "scan",
		Usage:       "let's scan proxy",
		Description: "let's scan proxy",
		Action:      util.Scan,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Value: false, Usage: "debug mode"},
			&cli.IntFlag{Name: "scan_num", Aliases: []string{"n"}, Value: 1000, Usage: "scan num"},
			&cli.IntFlag{Name: "timeout", Aliases: []string{"t"}, Value: 5, Usage: "timeout"},
			&cli.StringFlag{Name: "filename", Aliases: []string{"f"}, Value: "input_proxy_list.txt", Usage: "filename"},
			&cli.StringFlag{Name: "output_file", Aliases: []string{"o"}, Value: "proxyscan.txt", Usage: "output_file"},
		},
	}
}
func scan(ctx *cli.Context) error {
	return nil
}
