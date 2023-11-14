package functions

import (
	"sort"
)

type ScoreSlice []Score
type Score struct {
	Name                string
	ConScore, TechScore int
}

// võta selle pikkus
func (s ScoreSlice) Len() int {
	return len(s)
}

// võrdleb kahte elementi
func (s ScoreSlice) Less(i, j int) bool {
	// Calculate the sum of attributes for comparison
	sum1 := s[i].TechScore + s[i].ConScore
	sum2 := s[j].TechScore + s[j].ConScore
	return sum1 < sum2 // Change to > for descending order
}

// vahetab kahe asja kohad
func (s ScoreSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

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
			"Liimiga":        100,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 100,
			"ServaPuurEI":  -100,
			"CNCjah":       100,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},

		"Läätsveeder": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         100,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    100,
			"avatav":         0,
			"Liimiga":        100,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 100,
			"ServaPuurEI":  0,
			"CNCjah":       150,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    -100,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},

		"Konstruktorkruvi": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 100,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 100,
			"ServaPuurEI":  -100,
			"CNCjah":       100,
			"CNCei":        100,
			"LamelloJAH":   0,
			"LamelloEI":    100,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},
		"Mööblipolt + tapimutter": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 100,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 100,
			"ServaPuurEI":  -100,
			"CNCjah":       100,
			"CNCei":        100,
			"LamelloJAH":   0,
			"LamelloEI":    100,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},
		"Minifix + tüübel": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 100,
			"ServaPuurEI":  -100,
			"CNCjah":       100,
			"CNCei":        100,
			"LamelloJAH":   0,
			"LamelloEI":    100,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
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
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 100,
			"ServaPuurEI":  -100,
			"CNCjah":       100,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		}, "INVIS": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      0,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   000,

			"ServaPuurJAH": 100,
			"ServaPuurEI":  -100,
			"CNCjah":       100,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},
		"Rafix": {
			"Kandev Kulg":    100,
			"Lagi Peal":      0,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       0,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  100,
			"CNCjah":       100,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    0,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},
		"Clamex P “LAMELLO”": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         100,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  100,
			"CNCjah":       150,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    -100,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},
		"Tenso P “LAMELLO”": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         100,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    100,
			"avatav":         0,
			"Liimiga":        100,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         0,
			"mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  100,
			"CNCjah":       150,
			"CNCei":        0,
			"LamelloJAH":   100,
			"LamelloEI":    -100,
			"dominoJAH":    0,
			"dominoEI":     100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},
		"“DOMINO” puit": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    100,
			"avatav":         0,
			"Liimiga":        100,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         0,
			"mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  100,
			"CNCjah":       150,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    100,
			"dominoJAH":    100,
			"dominoEI":     -100,
			"ovvoJah":      0,
			"ovvoEI":       100,
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
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  100,
			"CNCjah":       150,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    100,
			"dominoJAH":    100,
			"dominoEI":     -100,
			"ovvoJah":      0,
			"ovvoEI":       100,
		},
		"”OVVO” Freesitav avatav": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         100,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"mitteavatav":    0,
			"avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         0,
			"mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"ServaPuurJAH": 0,
			"ServaPuurEI":  100,
			"CNCjah":       150,
			"CNCei":        0,
			"LamelloJAH":   0,
			"LamelloEI":    100,
			"dominoJAH":    0,
			"dominoEI":     100,
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
	//kui on true..
	for valik, onValitud := range valitudKonstrukt {
		if onValitud {
			//lisab skoori juurde
			for furnituur, skoor := range furnituurid {
				konstruktsioon[furnituur] += skoor[valik]
			}
		} else {
			//paneb ühe juurde kui ei ole valitud
			for furnituur := range furnituurid {
				konstruktsioon[furnituur]++
			}
		}
	}
	for valik, onValitud := range valitudTehnoloog {
		if onValitud {
			for furnituur, skoor := range furnituurid {
				tehnoloogia[furnituur] += skoor[valik]

			}
		} else {
			for furnituur := range furnituurid {
				tehnoloogia[furnituur]++
			}
		}
	}
	//teeb tühja  skoor slice
	var list = ScoreSlice{}
	//võtab kõik furnituuri nimed
	for nimi := range furnituurid {

		var s = Score{
			Name:      nimi,
			TechScore: tehnoloogia[nimi],
			ConScore:  konstruktsioon[nimi],
		}
		list = append(list, s)
	}

	sort.Sort(sort.Reverse(list))
	return list

}
