package seeder

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func SeedingRegionStructure(db *gorm.DB) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	mSourcePath := fmt.Sprintf("%v/infrastructure/datastore/seeder/place", path)
	fmt.Println("SEED PLACE => ", mSourcePath)
	//Seed Nationality
	if db.Migrator().HasTable(&model.Nationality{}) {
		if err := db.First(&model.Nationality{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			// Open jsonFile
			jsonFile := fmt.Sprintf("%s/%s", mSourcePath, "nationality.json")
			jsonNationality, errFile := os.Open(jsonFile)

			if errFile != nil {
				panic("Migration file nationality.json not found")
			}
			fmt.Println("Successfully Opened nationality.json")
			defer jsonNationality.Close()
			byteValue, _ := ioutil.ReadAll(jsonNationality)
			var resultNationality []model.Nationality
			json.Unmarshal([]byte(byteValue), &resultNationality)
			fmt.Println("Nationality seed migration loaded")
			err = db.Create(resultNationality).Error
			if err != nil {
				panic(err)
			}
		}
	}

	//Seed Region
	if db.Migrator().HasTable(&model.Province{}) { //Propinsi
		if err := db.First(&model.Province{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

			jsonFile := fmt.Sprintf("%s/%s", mSourcePath, "provinces.json")
			jsonProvinces, errFile := os.Open(jsonFile)
			if errFile != nil {
				panic("Migration file provinces.json not found")
			}
			fmt.Println("Successfully Opened provinces.json")
			defer jsonProvinces.Close()

			var resultProvinces []model.Province
			byteValue, _ := ioutil.ReadAll(jsonProvinces)
			json.Unmarshal([]byte(byteValue), &resultProvinces)
			fmt.Println("Provinces seed migration loaded")
			err = db.Create(resultProvinces).Error
			if err != nil {
				panic(err)
			}

			nationality := model.Nationality{}
			db.Where("code = 'ID'").First(&nationality)
			result := db.Model(model.Province{}).Where("1=1").Updates(model.Province{ParentID: nationality.ID})
			fmt.Println("Update Province: ", result.RowsAffected)
		}
	}

	if db.Migrator().HasTable(&model.City{}) { // Kota
		if err := db.First(&model.City{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			jsonFile := fmt.Sprintf("%s/%s", mSourcePath, "city.json")
			jsonDistricts, errFile := os.Open(jsonFile)
			if errFile != nil {
				panic("Migration file districts.json not found")
			}
			fmt.Println("Successfully Opened city.json")
			defer jsonDistricts.Close()

			var resultDistrict []model.City
			byteValue, _ := ioutil.ReadAll(jsonDistricts)
			json.Unmarshal([]byte(byteValue), &resultDistrict)
			fmt.Println("Districts seed migration loaded")
			err = db.Create(resultDistrict).Error
			if err != nil {
				panic(err)
			}
		}
	}

	if db.Migrator().HasTable(&model.District{}) { //Kecamatan
		if err := db.First(&model.District{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			jsonFile := fmt.Sprintf("%s/%s", mSourcePath, "districts.json")
			jsonSubDistricts, errFile := os.Open(jsonFile)
			if errFile != nil {
				panic("Migration file sub_districts.json not found")
			}
			fmt.Println("Successfully Opened districts.json")
			defer jsonSubDistricts.Close()

			var resultSubDistrict []model.District
			byteValue, _ := ioutil.ReadAll(jsonSubDistricts)
			json.Unmarshal([]byte(byteValue), &resultSubDistrict)
			fmt.Println("Sub Districts seed migration loaded")
			err = db.Create(resultSubDistrict).Error
			if err != nil {
				panic(err)
			}
		}
	}

	if db.Migrator().HasTable(&model.Village{}) { //Desa
		if err := db.First(&model.Village{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			csvFile := fmt.Sprintf("%s/%s", mSourcePath, "villages.csv")
			csvVillages, errFile := os.Open(csvFile)
			if errFile != nil {
				panic("Migration file villages.json not found")
			}
			var header []string
			//var villageRaw map[string]interface{}
			var resultVillages []model.Village
			r := csv.NewReader(csvVillages)
			for {
				record, err2 := r.Read()
				if err2 == io.EOF {
					break
				}
				if err2 != nil {
					panic(err2)
				}
				if header == nil {
					header = record
				} else {
					village := new(model.Village)
					x := reflect.ValueOf(village)
					v := x.Elem()
					fieldNames := map[string]int{}
					findJsonName := func(t reflect.StructTag) (string, error) {
						if jt, ok := t.Lookup("json"); ok {
							return strings.Split(jt, ",")[0], nil
						}
						return "", fmt.Errorf("tag provided does not define a json tag")
					}

					for i := 0; i < v.NumField(); i++ {
						typeField := v.Type().Field(i)
						tag := typeField.Tag
						jname, _ := findJsonName(tag)
						fieldNames[jname] = i
					}

					for i := range header {
						fieldNum, ok := fieldNames[header[i]]
						if !ok {
							continue
						}
						fieldVal := v.Field(fieldNum)
						//vData := record[i]
						fieldType := fieldVal.Kind()
						switch fieldType.String() {
						case "string":
							fieldVal.SetString(record[i])
						case "int", "int64", "int32", "int16", "int8":
							valx, _ := strconv.Atoi(record[i])
							fieldVal.SetInt(int64(valx))
						case "float64", "float32", "float":
							valx, _ := strconv.ParseFloat(record[i], 64)
							fieldVal.SetFloat(valx)
						case "bool":
							valx, _ := strconv.ParseBool(record[i])
							fieldVal.SetBool(valx)
						}
						//fmt.Println(fieldType)
						//fmt.Println("Set field: ", header[i], " to: ", record[i])
					}
					resultVillages = append(resultVillages, *village)
				}
			}
			err = db.CreateInBatches(resultVillages, 500).Error
			if err != nil {
				panic(err)
			}
		}
	}

}
