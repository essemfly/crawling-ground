package items

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type RequestPayload struct {
	ExceptAllCount bool `json:"exceptAllCount"`
	StartIndex     int  `json:"startIndex"`
}

type DetailRequestPayload struct {
	StoreId int `json:"storeId"`
}

// Define a struct to hold the response data
type BrandListResponse struct {
	Data struct {
		Brands []Brand `json:"allBrandList"`
	}
}

type StoreListResponse struct {
	Data AmondzStore
}

type Brand struct {
	StoreId   int    `json:"storeId"`
	StoreName string `json:"storeName"`
}

type AmondzStore struct {
	StoreId    int    `json:"storeId"`
	StoreName  string `json:"storeName"`
	Phone      string `json:"asPhone"`
	Addr       string `json:"companyAddress"`
	Email      string `json:"email"`
	Kakao      string `json:"kakaoAccount"`
	ReturnAddr string `json:"returnAddress"`
}

func CrawlAmondz() {
	startIndexes := []int{
		0, 30, 60, 90, 120, 150, 180, 210, 240, 270, 300,
	}
	stores := []Brand{}
	// Create a new request payload

	for _, startIndex := range startIndexes {
		payload := RequestPayload{
			ExceptAllCount: true,
			StartIndex:     startIndex,
		}

		// Convert the payload to a JSON string
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		// Create a new HTTP POST request with the payload
		req, err := http.NewRequest("POST", "https://api.amondz.com/api/brand/list/pagination/web/v1", bytes.NewBuffer(payloadBytes))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// Set the content type of the request to JSON
		req.Header.Set("Content-Type", "application/json")

		// Make the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		var brandList BrandListResponse
		err = json.NewDecoder(resp.Body).Decode(&brandList)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		for _, brand := range brandList.Data.Brands {
			stores = append(stores, brand)
		}

	}

	finalStores := []AmondzStore{}
	for _, store := range stores {
		detailUrl := "https://api.amondz.com/api/brand/detail/info/web/v1"
		detailPayload := DetailRequestPayload{
			StoreId: store.StoreId,
		}

		// Convert the payload to a JSON string
		detailPayloadBytes, err := json.Marshal(detailPayload)
		if err != nil {
			fmt.Println("Error marshaling payload:", err)
			return
		}

		// Create a new HTTP POST request with the payload
		newReq, err := http.NewRequest("POST", detailUrl, bytes.NewBuffer(detailPayloadBytes))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		newReq.Header.Set("Content-Type", "application/json")

		// Make the request
		client := &http.Client{}
		newResp, err := client.Do(newReq)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		var brand StoreListResponse
		err = json.NewDecoder(newResp.Body).Decode(&brand)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		finalStores = append(finalStores, AmondzStore{
			StoreId:    store.StoreId,
			StoreName:  store.StoreName,
			Phone:      brand.Data.Phone,
			Addr:       brand.Data.Addr,
			Email:      brand.Data.Email,
			Kakao:      brand.Data.Kakao,
			ReturnAddr: brand.Data.ReturnAddr,
		})
	}

	file := excelize.NewFile()

	// Create a new sheet
	sheetName := "Sheet1"
	file.NewSheet(sheetName)

	// Write header row
	file.SetCellValue(sheetName, "A1", "Name")
	file.SetCellValue(sheetName, "B1", "Phone")
	file.SetCellValue(sheetName, "C1", "Email")
	file.SetCellValue(sheetName, "D1", "Url")

	for i, p := range finalStores {
		rowIndex := i + 2
		file.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex), p.StoreName)
		file.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex), p.Phone)
		file.SetCellValue(sheetName, "C"+strconv.Itoa(rowIndex), p.Email)
		file.SetCellValue(sheetName, "D"+strconv.Itoa(rowIndex), "https://www.amondz.com/brand/"+strconv.Itoa(p.StoreId))
	}

	err := file.SaveAs("amondz.xlsx")
	if err != nil {
		panic(err)
	}
}
