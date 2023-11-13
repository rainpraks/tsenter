/* package main

import "fmt"

func main() {

	furnituurid := map[string]map[string]bool{
		"Puittüübel keerdsooneline": {
			"Kandev Kulg":    true,
			"Lagi Peal":      true,
			"Eerung":         false,
			"Vaheriiul":      true,
			"Vahesein":       true,
			"Ei Ava":         true,
			"Avatav":         false,
			"Liimiga":        true,
			"Liimita":        false,
			"Monteeritud":    true,
			"Flatpack":       false,
			"Nahtav":         false,
			"Mittenahtav":    true,
			"Valjast Nahtav": false,
			"Seest Nahtav":   false,
			"Serva Puur":
		},
		"Läätsveeder": {
			"Kandev Kulg":    true,
			"Lagi Peal":      true,
			"Eerung":         true,
			"Vaheriiul":      true,
			"Vahesein":       true,
			"Ei Ava":         true,
			"Avatav":         false,
			"Liimiga":        true,
			"Liimita":        false,
			"Monteeritud":    true,
			"Flatpack":       false,
			"Nahtav":         false,
			"Mittenahtav":    true,
			"Valjast Nahtav": false,
			"Seest Nahtav":   false,
		},
		"Maxifix + tüübel": {
			"Kandev Kulg":    true,
			"Lagi Peal":      true,
			"Eerung":         false,
			"Vaheriiul":      true,
			"Vahesein":       true,
			"Ei Ava":         false,
			"Avatav":         true,
			"Liimiga":        false,
			"Liimita":        true,
			"Monteeritud":    true,
			"Flatpack":       true,
			"Nahtav":         true,
			"Mittenahtav":    false,
			"Valjast Nahtav": false,
			"Seest Nahtav":   true,
		},
	}

	valitud := map[string]bool{
		"Kandev Kulg":    false,
		"Lagi Peal":      false,
		"Eerung":         false,
		"Vaheriiul":      false,
		"Vahesein":       false,
		"Ei Ava":         false,
		"Avatav":         false,
		"Liimiga":        false,
		"Liimita":        false,
		"Monteeritud":    false,
		"Flatpack":       false,
		"Nahtav":         false,
		"Mittenahtav":    false,
		"Valjast Nahtav": false,
		"Seest Nahtav":   false,
	}

	//vota frontendi valikud siia
	valitud["Eerung"] = true
	valitud["Avatav"] = true
	valitud["Liimita"] = true
	valitud["Monteeritud"] = true
	valitud["Nahtav"] = true
	valitud["Valjast Nahtav"] = true
	// *****************************

	tulemus := make(map[string]int)

	for valik, onValitud := range valitud {
		if onValitud {
			for furnituur, tabel := range furnituurid {
				if tabel[valik] == true {
					tulemus[furnituur] += 100
				}
			}
		}

	}
	fmt.Print(tulemus)
}
*/

package main

import "fmt"

func main() {

	furnituurid := map[string]map[string]int{
		"Puittüübel keerdsooneline": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"Ei Ava":         100,
			"Avatav":         0,
			"Liimiga":        100,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"Mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"On Serva Puur":  100,
			"Ei Ole Serva":   -100,
			"On CNC":         100,
			"Ei Ole CNC":     0,
			"On Lamello":     100,
			"Ei Ole Lamello": 0,
			"On Domino":      100,
			"Ei Ole Domino":  0,
			"On Ovvo":        100,
			"Ei Ole Ovvo":    0,
		},

		"Läätsveeder": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         100,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"Ei Ava":         100,
			"Avatav":         0,
			"Liimiga":        100,
			"Liimita":        0,
			"Monteeritud":    100,
			"Flatpack":       0,
			"Nahtav":         0,
			"Mittenahtav":    100,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   0,

			"On Serva Puur":  100,
			"Ei Ole Serva":   0,
			"On CNC":         150,
			"Ei Ole CNC":     0,
			"On Lamello":     100,
			"Ei Ole Lamello": -100,
			"On Domino":      100,
			"Ei Ole Domino":  0,
			"On Ovvo":        100,
			"Ei Ole Ovvo":    0,
		},

		"Maxifix + tüübel": {
			"Kandev Kulg":    100,
			"Lagi Peal":      100,
			"Eerung":         0,
			"Vaheriiul":      100,
			"Vahesein":       100,
			"Ei Ava":         0,
			"Avatav":         100,
			"Liimiga":        0,
			"Liimita":        100,
			"Monteeritud":    100,
			"Flatpack":       100,
			"Nahtav":         100,
			"Mittenahtav":    0,
			"Valjast Nahtav": 0,
			"Seest Nahtav":   100,

			"On Serva Puur":  100,
			"Ei Ole Serva":   -100,
			"On CNC":         100,
			"Ei Ole CNC":     0,
			"On Lamello":     100,
			"Ei Ole Lamello": 0,
			"On Domino":      100,
			"Ei Ole Domino":  0,
			"On Ovvo":        100,
			"Ei Ole Ovvo":    0,
		},
	}

	valitudKonstrukt := map[string]bool{
		"Kandev Kulg":    false,
		"Lagi Peal":      false,
		"Eerung":         false,
		"Vaheriiul":      false,
		"Vahesein":       false,
		"Ei Ava":         false,
		"Avatav":         false,
		"Liimiga":        false,
		"Liimita":        false,
		"Monteeritud":    false,
		"Flatpack":       false,
		"Nahtav":         false,
		"Mittenahtav":    false,
		"Valjast Nahtav": false,
		"Seest Nahtav":   false,
	}

	valitudTehnoloog := map[string]bool{
		"On Serva Puur":  false,
		"Ei Ole Serva":   false,
		"On CNC":         false,
		"Ei Ole CNC":     false,
		"On Lamello":     false,
		"Ei Ole Lamello": false,
		"On Domino":      false,
		"Ei Ole Domino":  false,
		"On Ovvo":        false,
		"Ei Ole Ovvo":    false,
	}

	//vota frontendi valikud siia
	valitudKonstrukt["Eerung"] = true
	valitudKonstrukt["Avatav"] = true
	valitudKonstrukt["Liimita"] = true
	valitudKonstrukt["Monteeritud"] = true
	valitudKonstrukt["Nahtav"] = true
	valitudKonstrukt["Valjast Nahtav"] = true

	valitudTehnoloog["On CNC"] = true
	valitudTehnoloog["Ei Ole Serva"] = true
	valitudTehnoloog["On Lamello"] = true
	valitudTehnoloog["On Domino"] = true
	valitudTehnoloog["On Ovvo"] = true
	// *****************************

	konstruktsioon := make(map[string]int)
	tehnoloogia := make(map[string]int)

	for valik, onValitud := range valitudKonstrukt {
		if onValitud {
			for furnituur, skoor := range furnituurid {
				konstruktsioon[furnituur] += skoor[valik]
			}
		}
	}
	for valik, onValitud := range valitudTehnoloog {
		if onValitud {
			for furnituur, skoor := range furnituurid {
				tehnoloogia[furnituur] += skoor[valik]
			}
		}
	}

	fmt.Println(konstruktsioon)
	fmt.Println(tehnoloogia)
}
