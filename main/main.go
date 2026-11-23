package main

import (
	"context"
	"fmt"
	"os"

	"github.com/EulersIdentity33/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	//construts: go run ./main --minTemp=-10 --maxTemp=30 --stMeritev=5
	cmd := &cli.Command{
		Flags:  []cli.Flag{&cli.Float64Flag{Name: "minTemp", Value: -50}, &cli.Float64Flag{Name: "maxTemp", Value: 50}, &cli.IntFlag{Name: "stMeritev", Value: 3}},
		Action: izvedi,
	}
	//exec
	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func izvedi(_ context.Context, cmd *cli.Command) error {
	vreme.MinTemp = cmd.Float64("minTemp")
	vreme.MaxTemp = cmd.Float64("maxTemp")
	vreme.StMeritev = cmd.Int("stMeritev")

	postaje := map[string]vreme.Postaja{
		"KR01": {"Kranj Center", "Kranj", []float64{20.5, 23, 26.1}},
		"LD01": {"Ledine", "Skuta", []float64{-2, -1.4, 0.5}},
		"BR01": {"Letalisce Brnik", "Zgornji Brnik", []float64{19.6, 28.5, 17.8, 15.5}},
		"SC01": {"Sencur", "Sencur", []float64{20.4}},
	}

	vreme.DodajPostajo(postaje, "AA00", "ABCD", "CDEF")

	vreme.DodajMeritev(postaje, "KR01", 23.4)
	vreme.DodajMeritev(postaje, "KR01", 60)
	vreme.DodajMeritev(postaje, "LJ01", 20)

	vreme.IzpisMeritev(postaje)
	vreme.IzpisPorocila(postaje)
	return nil
}
