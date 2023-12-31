package functions

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const TechMax float64 = 100
const ConMax float64 = 500

type ScoreSlice []Score
type Score struct {
	Name                                       string
	Image, Link                                string
	ConScore, TechScore                        int
	ConStars, TechStars, TechRating, ConRating string
}

// võta selle pikkus
func (s ScoreSlice) Len() int {
	return len(s)
}

// võrdleb kahte elementi
func (s ScoreSlice) Less(i, j int) bool {
	sum1 := s[i].TechScore*5 + s[i].ConScore
	sum2 := s[j].TechScore*5 + s[j].ConScore //*5 sest tech max on 100 ja con max on 500

	if sum1 != sum2 {
		return sum1 < sum2
	}
	if s[i].ConScore != s[j].ConScore {
		return s[i].ConScore < s[j].ConScore
	}
	return s[i].Name < s[j].Name
}

// vahetab kahe asja kohad
func (s ScoreSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

/* type Score2 struct {
	field map[string]int
	image string
} */

func ScoreCalc(inputKon, inputTeh []string) ScoreSlice {
	furnituurid := map[string]map[string]int{
		"Puittüübel keerdsooneline": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    100,
			"avatav":         0,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"mittenahtav":    200,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 50,
			"ServaPuurEI":  0,
			"CNCjah":       50,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Läätsveeder": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         150,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    100,
			"avatav":         0,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"mittenahtav":    200,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  0,
			"CNCjah":       0,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    -100,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Konstruktorkruvi": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      50,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 100,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 50,
			"ServaPuurEI":  0,
			"CNCjah":       50,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Mööblipolt + tapimutter": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      50,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 50,
			"Seest Nahtav":   50,

			"ServaPuurJAH": 50,
			"ServaPuurEI":  0,
			"CNCjah":       50,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Eksentriktõmmits + tüübel": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 50,
			"ServaPuurEI":  0,
			"CNCjah":       50,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Maxifix + tüübel": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 50,
			"ServaPuurEI":  0,
			"CNCjah":       50,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"INVIS": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      0,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       50,
			"Nahtav":         0,
			"mittenahtav":    200,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 50,
			"ServaPuurEI":  0,
			"CNCjah":       50,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Rafix": {
			"Kandev Kulg":    50,
			"Lagi Peal":      0,
			"Eerung":         0,
			"Vaheriiul":      150,
			"Vahesein":       0,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  0,
			"CNCjah":       100,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Clamex P “LAMELLO”": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         150,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  0,
			"CNCjah":       0,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    -100,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"Tenso P “LAMELLO”": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         150,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    100,
			"avatav":         0,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       50,
			"Nahtav":         0,
			"mittenahtav":    200,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  0,
			"CNCjah":       0,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    -100,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"“DOMINO” puit": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    100,
			"avatav":         0,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"mittenahtav":    200,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  0,
			"CNCjah":       0,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    100,
			"dominoEI":     -100,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"“DOMINO” KV-LR32": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  0,
			"CNCjah":       0,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    100,
			"dominoEI":     -100,
			"ovvoJah":      0,
			"ovvoEI":       0,
		},
		"”OVVO” Freesitav avatav": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         50,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         0,
			"mittenahtav":    200,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  0,
			"CNCjah":       0,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     0,
			"ovvoJah":      100,
			"ovvoEI":       -100,
		},
	}
	//kõik valikud mida kasutaja hakkab valima.
	valitudKonstrukt := map[string]bool{
		"Kandev Kulg":    false,
		"Lagi Peal":      false,
		"Eerung":         false,
		"Vaheriiul":      false,
		"Vahesein":       false,
		"mitteavatav":    false,
		"avatav":         false,
		"Liimiga":        false,
		"Liimita":        false,
		"Monteeritud":    false,
		"Flatpack":       false,
		"Nahtav":         false,
		"mittenahtav":    false,
		"Valjast Nahtav": false,
		"Seest Nahtav":   false,
	}

	valitudTehnoloog := map[string]bool{
		"ServaPuurJAH": false,
		"ServaPuurEI":  false,
		"CNCjah":       false,
		"CNCei":        false,
		"LamelloJAH":   false,
		"LamelloEI":    false,
		"dominoJAH":    false,
		"dominoEI":     false,
		"ovvoJah":      false,
		"ovvoEI":       false,
	}

	for _, ch := range inputKon {
		valitudKonstrukt[ch] = true

	}
	for _, ch := range inputTeh {
		valitudTehnoloog[ch] = true

	}
	//tühjad mapid selleks et skoori arvutada
	konstruktsioon := make(map[string]int)
	tehnoloogia := make(map[string]int)
	doit := false
	if valitudKonstrukt["Valjast Nahtav"] && valitudKonstrukt["Seest Nahtav"] {
		fmt.Println("tra")
		valitudKonstrukt["Valjast Nahtav"] = false
		valitudKonstrukt["Seest Nahtav"] = false
		doit = true
	}
	//kui on true..
	for valik, onValitud := range valitudKonstrukt {
		if onValitud {
			//lisab skoori juurde
			for furnituur, skoor := range furnituurid {
				if furnituur == "Mööblipolt + tapimutter" && doit {
					konstruktsioon[furnituur] += 100
					fmt.Println("works")
					doit = false
				}
				konstruktsioon[furnituur] += skoor[valik]
			}
		} /* else {
			//paneb ühe juurde kui ei ole valitud
			for furnituur := range furnituurid {
				konstruktsioon[furnituur]++
			}
		} */
	}
	for valik, onValitud := range valitudTehnoloog {
		if onValitud {
			for furnituur, skoor := range furnituurid {
				tehnoloogia[furnituur] += skoor[valik]

			}
		} /* else {
			for furnituur := range furnituurid {
				tehnoloogia[furnituur]++
			}
		} */

	}
	//teeb tühja  skoor slice
	var list = ScoreSlice{}
	//võtab kõik furnituuri nimed
	for nimi := range furnituurid {

		img := "/static/" + nimi + ".jpg"
		t, c := ConvertStars(tehnoloogia[nimi], konstruktsioon[nimi])
		techS, conS := ConvertScore(tehnoloogia[nimi], konstruktsioon[nimi])
		var s = Score{
			Name:       nimi,
			TechScore:  tehnoloogia[nimi],
			ConScore:   konstruktsioon[nimi],
			TechStars:  t,
			ConStars:   c,
			TechRating: techS,
			ConRating:  conS,
			Image:      img,
			Link:       "static/Eksentrik.pdf",
		}
		list = append(list, s)
	}
	sort.Sort(list)
	sort.Sort(sort.Reverse(list))
	return list

}

func ConvertStars(tech, con int) (string, string) {
	conStars := ""
	techStars := ""
	conFullStars := con / 100
	techFullStars := tech / 20

	if tech%20 != 0 { //kui ei tule taisarv jagamisel siis sisuliselt ymardab yles (tegelt lisab alla ymardatud numbrile +1)
		techFullStars += 1
	}

	if con%100 != 0 { //kui ei tule taisarv jagamisel siis sisuliselt ymardab yles (tegelt lisab alla ymardatud numbrile +1)
		conFullStars += 1
	}

	if tech < 0 { //ei soovi negatiivseid skoore aga nad tekivad hetkese systeemiga
		techFullStars = 0
	}

	conStars += strings.Join(make([]string, conFullStars+1), "\u2605")
	techStars += strings.Join(make([]string, techFullStars+1), "\u2605")

	for len(conStars) < 15 { // kui on vahem kui 5 t2rni siis lisab tyhja t2rni. 15 sest selle taheelemendi pikkus on 3
		conStars += "\u2606"
	}
	for len(techStars) < 15 {

		techStars += "\u2606"
	}
	return techStars, conStars
}
func ConvertScore(tech, con int) (string, string) {
	/*
		const TechMax int = 200
		const ConMax int = 500
	*/

	techFloat := float64(tech)
	conFloat := float64(con)
	techPoint := 0.0
	conPoint := 0.0

	techPoint = roundFloat((techFloat/TechMax)*5, 1)

	conPoint = roundFloat((conFloat/ConMax)*5, 1)
	return (fmt.Sprintf("%g", techPoint) + "/5"), (fmt.Sprintf("%g", conPoint) + "/5")
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
