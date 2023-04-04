package items

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type StoreInfo struct {
	StoreName           string
	OwnerName           string
	CustomerCenterPhone string
	Email               string
	RawInfo             string
	Url                 string
}

type SmartStoreInfo struct {
	CategoryFilter    string `json:"categoryFilter"`
	ChnlSeq           string `json:"chnlSeq"`
	CrUrl             string `json:"crUrl"`
	DefaultPayType    string `json:"defaultPayType"`
	IsSmartStore      string `json:"isSmartStore"`
	KeepCnt           int    `json:"keepCnt"`
	MallDesc          string `json:"mallDesc"`
	MallGrade         string `json:"mallGrade"`
	MallLogo          string `json:"mallLogo"`
	MallName          string `json:"mallName"`
	MallNameAndDesc   string `json:"mallNameAndDesc"`
	MallSeq           string `json:"mallSeq"`
	MallTypeFilter    string `json:"mallTypeFilter"`
	NaverPaySaveRatio int    `json:"naverPaySaveRatio"`
	ProdCnt           string `json:"prodCnt"`
	RepCatNm          string `json:"repCatNm"`
	StoreInfo         StoreInfo
}

func WriteStoreBriefExcel(stores []*StoreInfo, outputFile string) {
	file := excelize.NewFile()

	sheetName := "Sheet1"
	file.NewSheet(sheetName)

	// Write header row
	file.SetCellValue(sheetName, "A1", "Name")
	file.SetCellValue(sheetName, "B1", "Phone")
	file.SetCellValue(sheetName, "C1", "Email")
	file.SetCellValue(sheetName, "D1", "Url")

	for i, p := range stores {
		rowIndex := i + 2
		file.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex), p.StoreName)
		file.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex), p.CustomerCenterPhone)
		file.SetCellValue(sheetName, "C"+strconv.Itoa(rowIndex), p.Email)
		file.SetCellValue(sheetName, "D"+strconv.Itoa(rowIndex), p.Url)
	}

	err := file.SaveAs("outputs/" + outputFile + ".xlsx")
	if err != nil {
		panic(err)
	}
}

func WriteStoreExtendExcel(stores []*SmartStoreInfo, outputFile string) {
	file := excelize.NewFile()

	sheetName := "Sheet1"
	file.NewSheet(sheetName)

	// Write header row
	file.SetCellValue(sheetName, "A1", "Name")
	file.SetCellValue(sheetName, "B1", "Phone")
	file.SetCellValue(sheetName, "C1", "Email")
	file.SetCellValue(sheetName, "D1", "Url")
	file.SetCellValue(sheetName, "E1", "CategoryFilter")
	file.SetCellValue(sheetName, "F1", "ChnlSeq")
	file.SetCellValue(sheetName, "G1", "CrUrl")
	file.SetCellValue(sheetName, "H1", "DefaultPayType")
	file.SetCellValue(sheetName, "I1", "IsSmartStore")
	file.SetCellValue(sheetName, "J1", "KeepCnt")
	file.SetCellValue(sheetName, "K1", "MallDesc")
	file.SetCellValue(sheetName, "L1", "MallGrade")
	file.SetCellValue(sheetName, "M1", "MallLogo")
	file.SetCellValue(sheetName, "N1", "MallName")
	file.SetCellValue(sheetName, "O1", "MallNameAndDesc")
	file.SetCellValue(sheetName, "P1", "MallSeq")
	file.SetCellValue(sheetName, "Q1", "MallTypeFilter")
	file.SetCellValue(sheetName, "R1", "NaverPaySaveRatio")
	file.SetCellValue(sheetName, "S1", "ProdCnt")
	file.SetCellValue(sheetName, "T1", "RepCatNm")

	for i, p := range stores {
		rowIndex := i + 2
		file.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex), p.StoreInfo.StoreName)
		file.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex), p.StoreInfo.CustomerCenterPhone)
		file.SetCellValue(sheetName, "C"+strconv.Itoa(rowIndex), p.StoreInfo.Email)
		file.SetCellValue(sheetName, "D"+strconv.Itoa(rowIndex), p.StoreInfo.Url)
		file.SetCellValue(sheetName, "E"+strconv.Itoa(rowIndex), p.CategoryFilter)
		file.SetCellValue(sheetName, "F"+strconv.Itoa(rowIndex), p.ChnlSeq)
		file.SetCellValue(sheetName, "G"+strconv.Itoa(rowIndex), p.CrUrl)
		file.SetCellValue(sheetName, "H"+strconv.Itoa(rowIndex), p.DefaultPayType)
		file.SetCellValue(sheetName, "I"+strconv.Itoa(rowIndex), p.IsSmartStore)
		file.SetCellValue(sheetName, "J"+strconv.Itoa(rowIndex), p.KeepCnt)
		file.SetCellValue(sheetName, "K"+strconv.Itoa(rowIndex), p.MallDesc)
		file.SetCellValue(sheetName, "L"+strconv.Itoa(rowIndex), p.MallGrade)
		file.SetCellValue(sheetName, "M"+strconv.Itoa(rowIndex), p.MallLogo)
		file.SetCellValue(sheetName, "N"+strconv.Itoa(rowIndex), p.MallName)
		file.SetCellValue(sheetName, "O"+strconv.Itoa(rowIndex), p.MallNameAndDesc)
		file.SetCellValue(sheetName, "P"+strconv.Itoa(rowIndex), p.MallSeq)
		file.SetCellValue(sheetName, "Q"+strconv.Itoa(rowIndex), p.MallTypeFilter)
		file.SetCellValue(sheetName, "R"+strconv.Itoa(rowIndex), p.NaverPaySaveRatio)
		file.SetCellValue(sheetName, "S"+strconv.Itoa(rowIndex), p.ProdCnt)
		file.SetCellValue(sheetName, "T"+strconv.Itoa(rowIndex), p.RepCatNm)
	}

	err := file.SaveAs("outputs/" + outputFile + ".xlsx")
	if err != nil {
		panic(err)
	}
}
