package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type FibreGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarLevels = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevel = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func (energy EnergyKJ) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(energy), energyLevelsBeverage)
	}

	return getPointsFromRange(float64(energy), energyLevels)
}

func (sugarGr SugarGram) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(sugarGr), sugarLevelsBeverage)
	}

	return getPointsFromRange(float64(sugarGr), sugarLevels)
}

func (satFattyAcids SaturatedFattyAcids) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(satFattyAcids), saturatedFattyAcidLevels)
}

func (sodiumMg SodiumMilligram) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(sodiumMg), saturatedFattyAcidLevels)
}

func (fibreGr FibreGram) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(fibreGr), saturatedFattyAcidLevels)
}

func (protein ProteinGram) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(protein), saturatedFattyAcidLevels)
}

func (fruits FruitsPercent) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(fruits), saturatedFattyAcidLevels)
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

func GetNutritionalScore(nutritionalData NutritionalData, scoreType ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	if scoreType != Water {
		fruitPoints := nutritionalData.Fruits.GetPoints(scoreType)
		fibrePoints := nutritionalData.Fibre.GetPoints(scoreType)

		negative = nutritionalData.Energy.GetPoints(scoreType) + nutritionalData.Sugars.GetPoints(scoreType) + nutritionalData.SaturatedFattyAcids.GetPoints(scoreType) + nutritionalData.Sodium.GetPoints(scoreType)
		positive = fruitPoints + fibrePoints + nutritionalData.Protein.GetPoints(scoreType)
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: scoreType,
	}
}

func getPointsFromRange(value float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if value > l {
			return lenSteps - i
		}
	}

	return 0
}
