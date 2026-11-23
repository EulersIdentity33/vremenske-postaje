package vreme

import "fmt"

var MinTemp float64 = -50
var MaxTemp float64 = 50
var StMeritev int = 3

type Postaja struct {
	Ime     string
	Kraj    string
	Meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {
	_, obst := postaje[oznaka]
	if obst {
		fmt.Println("Postaja ze obstaja:", oznaka)
		return
	}
	postaje[oznaka] = Postaja{ime, kraj, []float64{}}
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	postaja, obst := postaje[oznaka]
	if !obst {
		fmt.Println("Ta postaja ne obstaja:", oznaka)
		return
	}
	if temperatura > MaxTemp || temperatura < MinTemp {
		fmt.Println("Neveljavna temperatura")
		return
	}
	postaja.Meritve = append(postaja.Meritve, temperatura)
	postaje[oznaka] = postaja
}

func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obst := postaje[oznaka]
	if !obst || len(postaja.Meritve) < StMeritev || len(postaja.Meritve) == 0 {
		return 42, 42, 42, false
	}

	najnizja = postaja.Meritve[0]
	najvisja = postaja.Meritve[0]

	for _, t := range postaja.Meritve {
		if t < najnizja {
			najnizja = t
		}
		if t > najvisja {
			najvisja = t
		}
		povprecna += t
	}

	povprecna /= float64(len(postaja.Meritve))
	return najnizja, najvisja, povprecna, true
}

func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.Ime, postaja.Kraj, postaja.Meritve)
	}
}

func IzpisPorocila(postaje map[string]Postaja) {
	for oznaka, postaja := range postaje {
		min, max, povprecje, ok := statistika(postaje, oznaka)
		if !ok {
			fmt.Printf("Premalo meritev\n")
			continue
		}

		var opis string
		switch {
		case povprecje >= 25:
			opis = "Vroce"
		case povprecje >= 15:
			opis = "Toplo"
		case povprecje >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprecje %.1f °C -> %s\n",
			postaja.Ime, min, max, povprecje, opis)
	}
}
