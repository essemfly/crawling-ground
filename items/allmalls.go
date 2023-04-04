package items

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gocolly/colly"
)

const (
	SBTH       = "b23e2e647174560fbb09c6c1344bd802ce823c05000889649d5f9de88b37296b6d94a52b9da39aa394a47ef6e7419207"
	REFERER    = "https://search.shopping.naver.com/allmall"
	numWorkers = 1
)

type SmartStoreListResponse struct {
	MallList   []*SmartStoreInfo `json:"mallList"`
	TotalCount int               `json:"totalCount"`
}

func CrawlStarter() {

	// mallTpNms := []string{
	// 	"DEPARTMENT_AND_MART_AND_HOME_SHOPPING",
	// 	"TOTALMALL",
	// 	"BRANDSTORE",
	// 	"SOHO",
	// }

	repCatNms := []string{
		"CLOTHING",  // 56245
		"SHOES",     // 54306
		"COSMETICS", // 22361
		"LIVING",    // 42275
		"FOOD",      // 73139
		"PARENTING", // 18744
		"SPORTS",    // 30122
		"DIGITAL",   // 37965
		"ETC",       // 166706
	}

	workers := make(chan bool, numWorkers)
	done := make(chan bool, numWorkers)

	for c := 0; c < numWorkers; c++ {
		done <- true
	}

	for _, cat := range repCatNms {
		workers <- true
		<-done
		go CrawlAllmalls(workers, done, cat)
	}

	for c := 0; c < numWorkers; c++ {
		<-done
	}
}

func CrawlAllmalls(worker chan bool, done chan bool, category string) {
	smartStores := []*SmartStoreInfo{}
	j := 1

	for j <= 100 {
		log.Println("Crawling category: "+category+" page:", j)
		smartStores = append(smartStores, CrawlAllmallPage(j, category)...)
		j++
	}
	WriteStoreExtendExcel(smartStores, "allmalls-"+category)
	<-worker
	done <- true
}

func CrawlAllmallPage(pageNum int, catName string) []*SmartStoreInfo {
	c := colly.NewCollector()
	reqUrl := "https://search.shopping.naver.com/allmall/api/allmall?mallTpNm=SOHO&isSmartStore=Y&sortingOrder=prodClk&page=" + strconv.Itoa(pageNum) + "&repCatNm=" + catName

	log.Println("req url", reqUrl)
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
	log.Println("Total store count:", storeList.TotalCount)
	for _, store := range storeList.MallList {
		alreadyCheck := false
		c.OnHTML("div#pc-storeNameWidget div._1bplHci37r div._3TZha2IPoQ a._2yPVRArtDH", func(e *colly.HTMLElement) {
			if alreadyCheck || e.Attr("href") == "" {
				return
			}

			store.CrUrl = "https://smartstore.naver.com" + e.Attr("href")
			storeInfo := CrawlSmartStoreDetail(store.CrUrl)
			store.StoreInfo = *storeInfo
			smartStores = append(smartStores, store)
			alreadyCheck = true
		})

		c.Visit(store.CrUrl)
	}

	return smartStores
}
