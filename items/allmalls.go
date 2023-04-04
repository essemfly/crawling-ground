package items

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	SBTH    = "b23e2e647174560fbb09c6c1344bd802ce823c05000889649d5f9de88b37296b6d94a52b9da39aa394a47ef6e7419207"
	REFERER = "https://search.shopping.naver.com/allmall"
)

type SmartStoreListResponse struct {
	MallList   []*SmartStoreInfo `json:"mallList"`
	TotalCount int               `json:"totalCount"`
}

func CrawlAllmalls() {
	smartStores := []*SmartStoreInfo{}
	pageCount := 1

	smartStores = append(smartStores, CrawlAllmallPage(pageCount)...)

	WriteStoreExtendExcel(smartStores, "allmalls")
}

func CrawlAllmallPage(pageNum int) []*SmartStoreInfo {
	reqUrl := "https://search.shopping.naver.com/allmall/api/allmall?isSmartStore=Y&sortingOrder=prodClk&page=" + strconv.Itoa(pageNum)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}
	req.Header.Set("Referer", REFERER)
	req.Header.Set("sbth", SBTH)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}

	var storeList SmartStoreListResponse
	err = json.NewDecoder(resp.Body).Decode(&storeList)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}

	smartStores := []*SmartStoreInfo{}
	for _, store := range storeList.MallList {

		storeInfo := CrawlSmartStoreDetail(store.CrUrl)
		store.StoreInfo = *storeInfo
		smartStores = append(smartStores, store)
	}

	return smartStores
}
