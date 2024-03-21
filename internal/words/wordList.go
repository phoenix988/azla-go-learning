package words

// create_wordlist creates a map of wordlists
func CreateWordlist() map[string]map[string]string {
	// Initially create map of worlist categories
	// Each category contain one map of words
	var wordlist = map[string]map[string]string{
		"Animals": {
			"Cat": "Pisik",
			"Dog": "It",
		},
		"Months": {
			"January":   "Yanvar",
			"February":  "Fevral",
			"March":     "Mart",
			"April":     "Aprel",
			"May":       "May",
			"June":      "İyun",
			"July":      "İyul",
			"August":    "Avqust",
			"September": "Sentyabr",
			"October":   "Oktyabr",
			"November":  "Noyabr",
			"December":  "Dekabr",
		},
	}

	// Words of Colors
	wordlist["Colors"] = map[string]string{
		"Colors":    "Rənglər",
		"Red":       "Qırmızı",
		"Black":     "Qara",
		"Yellow":    "Sarı",
		"Green":     "Yaşıl",
		"Purple":    "Bənövşəyi",
		"Blue":      "Mavi",
		"Orange":    "Narıncı",
		"Grey":      "Boz",
		"Brown":     "Qəhvəyi",
		"Pink":      "Çəhrayı",
		"Dark blue": "Göy",
		"White":     "Ağ",
		"Maroon":    "Həşər",
		"Dark red":  "Tünd qırmızı",
		"Cyan":      "Zövqəng",
		"Silver":    "Gümüşü",
		"Coral":     "Mərcanı",
		"Olive":     "Kəhrəba",
		"Mauve":     "Mürd",
		"Ashen":     "Bozçalı",
		"Pearl":     "Nacar",
	}

	// Words of body parts
	wordlist["Body Parts"] = map[string]string{
		"head":      "baş",
		"ear":       "Qulaq",
		"hair":      "Saç",
		"Mouth":     "Ağız",
		"Nose":      "Burun",
		"Eye":       "Göz",
		"lungs":     "Ağ ciyər",
		"liver":     "Qara ciyər",
		"neck":      "Sırğa",
		"arms":      "Qollar",
		"shoulders": "Qoşa",
		"elbows":    "Dirəklər",
		"hands":     "Əllər",
		"body":      "Bədən",
		"chest":     "Göğs",
		"stomach":   "Qarn",
		"waist":     "Bel",
		"hips":      "Göbələk",
		"ribcage":   "Qəfəs",
		"bone":      "Sümük",
		"heart":     "Ürək",
		"knee":      "Kürək",
		"foot":      "Ayaq",
	}

	// Words of Numbers
	wordlist["Numbers"] = map[string]string{
		"One":        "bir",
		"Two":        "iki",
		"Three":      "üç",
		"Four":       "dörd",
		"Five":       "beş",
		"Six":        "alti",
		"Seven":      "yeddi",
		"Eight":      "səkkiz",
		"Nine":       "doqquz",
		"Ten":        "on",
		"Twenty":     "iyirmi",
		"Thirty":     "otuz",
		"Fourty":     "qirx",
		"Fifty":      "əlli",
		"Sixty":      "altmiş",
		"Seventy":    "yetmiş",
		"Eighty":     "səksən",
		"Ninety":     "doxsan",
		"Hundred":    "yüz",
		"Thousand":   "min",
		"million":    "bir million",
		"today":      "bu gün",
		"tomorrow":   "sabah",
		"yesterday":  "dünən",
		"week":       "həftə",
		"month":      "ay",
		"year":       "il",
		"decade":     "onillik",
		"century":    "əsr",
		"millennium": "minillik",
	}

	// Adjective words
	wordlist["Adjectives"] = Adjectives
	wordlist["Verbs"] = Verbs
	wordlist["Nouns"] = Nouns

	// Return the map
	return wordlist
}
