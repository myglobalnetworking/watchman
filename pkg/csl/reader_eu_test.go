package csl

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEU(t *testing.T) {
	euCSL, euCSLMap, err := ReadEUFile(filepath.Join("..", "..", "test", "testdata", "eu_csl.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if euCSL == nil || euCSLMap == nil {
		t.Fatal("failed to parse eu_csl.csv")
	}

	if len(euCSL) == 0 {
		t.Fatal("failed to convert map to sheet")
	}

	testLogicalID := 13

	if euCSLMap[testLogicalID] == nil {
		t.Fatalf("expected a record at %d and got nil", testLogicalID)
	}
	expectedFileGenerationDate := "28/10/2022"
	expectedReferenceNumber := "EU.27.28"
	expectedEntityRemark := "(UNSC RESOLUTION 1483)"
	expectedClassificationCode := "person"
	expectedPublicationURL := "http://eur-lex.europa.eu/LexUriServ/LexUriServ.do?uri=OJ:L:2003:169:0006:0023:EN:PDF"

	// Name alias
	expectedNameAliasWholeName1 := "Saddam Hussein Al-Tikriti"
	expectedNameAliasWholeName2 := "Abu Ali"
	expectedNameAliasWholeName3 := "Abou Ali"
	expectedNameAliasTitle := ""

	// Address
	// No address found for this record

	expectedBirthDate := "1937-04-28"
	expectedBirthCity := "al-Awja, near Tikrit"
	expectedBirthCountryDescription := "IRAQ"

	assert.Greater(t, len(euCSL), 0)
	assert.NotNil(t, euCSLMap[testLogicalID].Entity)
	assert.NotNil(t, euCSLMap[testLogicalID].NameAliases)
	assert.NotNil(t, euCSLMap[testLogicalID].Addresses)
	assert.NotNil(t, euCSLMap[testLogicalID].BirthDates)
	assert.NotNil(t, euCSLMap[testLogicalID].Identifications)

	assert.Equal(t, euCSLMap[testLogicalID].FileGenerationDate, expectedFileGenerationDate)

	// Entity
	assert.Equal(t, expectedReferenceNumber, euCSLMap[testLogicalID].Entity.ReferenceNumber)
	assert.Equal(t, expectedEntityRemark, euCSLMap[testLogicalID].Entity.Remark)
	assert.Equal(t, expectedClassificationCode, euCSLMap[testLogicalID].Entity.SubjectType.ClassificationCode)
	assert.Equal(t, expectedPublicationURL, euCSLMap[testLogicalID].Entity.Regulation.PublicationURL)

	// Name Alias
	assert.Equal(t, expectedNameAliasWholeName1, euCSLMap[testLogicalID].NameAliases[0].WholeName)
	assert.Equal(t, expectedNameAliasWholeName2, euCSLMap[testLogicalID].NameAliases[1].WholeName)
	assert.Equal(t, expectedNameAliasWholeName3, euCSLMap[testLogicalID].NameAliases[2].WholeName)
	assert.Equal(t, expectedNameAliasTitle, euCSLMap[testLogicalID].NameAliases[0].Title)

	// Address
	assert.Len(t, euCSLMap[testLogicalID].Addresses, 0)

	// BirthDate
	assert.Equal(t, expectedBirthDate, euCSLMap[testLogicalID].BirthDates[0].BirthDate)
	assert.Equal(t, expectedBirthCity, euCSLMap[testLogicalID].BirthDates[0].City)
	assert.Equal(t, expectedBirthCountryDescription, euCSLMap[testLogicalID].BirthDates[0].CountryDescription)

	// Identification
	assert.Len(t, euCSLMap[testLogicalID].Identifications, 0)
}
