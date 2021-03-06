package cmd

import (
	"fmt"
	"github.com/j3ssie/goverview/core"
	"github.com/j3ssie/goverview/utils"
	"github.com/panjf2000/ants"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

func init() {
	var screenCmd = &cobra.Command{
		Use:   "screen",
		Short: "Do Screenshot on target",
		RunE:  runScreen,
	}

	// screen options
	screenCmd.Flags().BoolVar(&options.AbsPath, "A", false, "Use Absolute path in summary")
	screenCmd.Flags().BoolVar(&options.Screen.UseChromedp, "cdp", false, "Use old chromedp instead of rod")
	screenCmd.Flags().IntVar(&options.Screen.ScreenTimeout, "screen-timeout", 40, "screenshot timeout")
	screenCmd.Flags().IntVar(&options.Screen.ImgHeight, "height", 0, "Height screenshot")
	screenCmd.Flags().IntVar(&options.Screen.ImgWidth, "width", 0, "Width screenshot")
	RootCmd.AddCommand(screenCmd)
}

func runScreen(_ *cobra.Command, _ []string) error {
	// prepare output
	prepareOutput()
	err := core.LoadTechs(options)
	if err == nil {
		options.Fin.Loaded = true
	}

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(options.Concurrency, func(i interface{}) {
		defer wg.Done()
		job := i.(string)

		if strings.TrimSpace(job) == "" {
			return
		}
		if options.InputAsBurp {
			job = core.ParseBurpRequest(job)
		}

		utils.InforF("[screenshot] %v", job)
		var out string
		if options.Screen.UseChromedp {
			out = core.DoScreenshot(options, job)
		} else {
			out = core.NewDoScreenshot(options, job)
		}
		if out != "" {
			fmt.Println(out)
			core.AppendTo(options.ScreenShotFile, out)
		}
	}, ants.WithPreAlloc(true))
	defer p.Release()

	for _, raw := range inputs {
		wg.Add(1)
		err := p.Invoke(raw)
		if err != nil {
			utils.ErrorF("Invoke error: %s", err)
		}
	}
	wg.Wait()
	printOutput()
	return nil
}
