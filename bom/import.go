package bom

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

func ImportBomRunes(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	deMapping := map[string]string{}
	enMapping := map[string]string{}
	for _, record := range records {
		character := record[0]
		en := record[1]
		de := record[2]

		deMapping[character] = de
		enMapping[character] = en
	}

	runesDe, err := json.Marshal(deMapping)
	if err != nil {
		return err
	}

	runesEn, err := json.Marshal(enMapping)
	if err != nil {
		return err
	}

	dataDe := "export default " + string(runesDe)
	dataEn := "export default " + string(runesEn)

	err = os.WriteFile("./public/bom/german.js", []byte(dataDe), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile("./public/bom/english.js", []byte(dataEn), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
