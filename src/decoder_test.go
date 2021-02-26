package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoundClassesDecoder_Decode(t *testing.T) {
	testCases := map[string][]string{
		"*múina":             {"MN"},
		"*pápí":              {"PP"},
		"*pàrà":              {"PR"},
		"*ə̀pə̀-":            {"HP"},
		"*tə́rí":             {"TR"},
		"*kàm-":              {"KM"},
		"*kùruà-":            {"KR"},
		"*tí":                {"TH"},
		"*pə̀nià":            {"PN"},
		"*múnà-i":            {"MN"},
		"*dák-":              {"TK"},
		"*túmái":             {"TM"},
		"*kùmua":             {"KM"},
		"*túmV-":             {"TM"},
		"*kə̀-":              {"KH"},
		"*sín-":              {"SN"},
		"*ìnù":               {"HN"},
		"*nə̀m-":             {"NM"},
		"*kàwà-":             {"KH"},
		"*mìmì":              {"MM"},
		"*tùtì":              {"TT"},
		"*kùp-":              {"KP"},
		"*mà-i":              {"MH"},
		"*àmpùr-à":           {"HMPR"},
		"*pánái":             {"PN"},
		"*pə̀-i":             {"PH"},
		"*íwuá":              {"HH"},
		"*tə́mp-":            {"TMP"},
		"*àsì":               {"HS"},
		"*mìt-":              {"MT"},
		"*átá-pá- ~ *kúrá-":  {"HT", "KR"},
		"*də̀-":              {"TH"},
		"*àwə̀-":             {"HH"},
		"*kàmì":              {"KM"},
		"*tà-i":              {"TH"},
		"*kàsìrà":            {"KSR"},
		"*kí-k-":             {"KH"},
		"*kə̀kə̀rə":          {"KKR"},
		"*tùnuá":             {"TN"},
		"*a":                 {"HH"},
		"*kə́rə-s-":          {"KR"},
		"*pínsá":             {"PNS"},
		"*sír-":              {"SR"},
		"*pá":                {"PH"},
		"*na-":               {"NH"},
		"*kìmuà":             {"KM"},
		"*nànkà-":            {"NNK"},
		"*sìrámí":            {"SRM"},
		"*mana-i-":           {"MN"},
		"*sìsì":              {"SS"},
		"*tùkù-i":            {"TK"},
		"*dàmà":              {"TM"},
		"*kútí":              {"KT"},
		"*náŋ":               {"NN"},
		"*kúmpí":             {"KMP"},
		"*nípí-":             {"NP"},
		"*duà":               {"TH"},
		"*páná":              {"PN"},
		"*na- ~ *=an-":       {"NH", "HN"},
		"*pitə":              {"PT"},
		"*pítə̀":             {"PT"},
		"*àmâi":              {"HM"},
		"*áká-":              {"HK"},
		"*mítí":              {"MT"},
		"*nài ~ *mə̀tə̀":     {"NH", "MT"},
		"*márə́":             {"MR"},
		"*sú-ná":             {"SH"},
		"*íp-":               {"HP"},
		"*mì-":               {"MH"},
		"*tànài":             {"TN"},
		"*bí-":               {"PH"},
		"*kapa":              {"KP"},
		"*úi-na":             {"HH"},
		"*tìpì-sà-":          {"TP"},
		"*kái-n-púri":        {"KH"},
		"*tàt-":              {"TT"},
		"*pə́sí":             {"PS"},
		"*ísì":               {"HS"},
		"*pí":                {"PH"},
		"*ə̀yə̀-nk-":         {"HH"},
		"*bə̀":               {"PH"},
		"*sə́-":              {"SH"},
		"*á- ~ *ká-":         {"HH", "KH"},
		"*kə́-":              {"KH"},
		"*ná":                {"NH"},
		"*sìtà":              {"ST"},
		"*pa":                {"PH"},
		"*kə̀i":              {"KH"},
		"*puta-":             {"PT"},
		"*dúk-":              {"TK"},
		"*àtà-taka-":         {"HT"},
		"*mí-n":              {"MH"},
		"*bà-":               {"PH"},
		"*nà-":               {"NH"},
		"*sìruà-":            {"SR"},
		"*tá":                {"TH"},
		"*bəmina":            {"PMN"},
		"*kú-i":              {"KH"},
		"*tə́pə́-":           {"TP"},
		"*ə́mə́- ~ *ə́mpə́-": {"HM", "HMP"},
		"*tìkà-":             {"TK"},
		"*sìpə̀":             {"SP"},
		"*mìnsì-kà-":         {"MNS"},
		"*pàimV ~ *pàimpV":   {"PM", "PMP"},
		"*pə̀sə̀-":           {"PS"},
		"*úsú-":              {"HS"},
		"*kánsá-i":           {"KNS"},
		"*músí":              {"MS"},
		"*tə̀sì":             {"TS"},
	}

	decoder, err := NewSoundClassesDecoder("../data/sounds.xlsx")
	assert.NoError(t, err)

	for forms, expected := range testCases {
		_, decoded := decoder.decodeForm(forms)
		assert.Equal(t, expected, decoded)
	}
}
