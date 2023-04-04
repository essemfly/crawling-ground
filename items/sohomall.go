package items

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func CrawlSohomall() {
	sites := []string{
		"https://smartstore.naver.com/purestream",
		"https://smartstore.naver.com/tamemall",
		"https://smartstore.naver.com/ninewolf",
		"https://smartstore.naver.com/namdaemoonqueen",
		"https://smartstore.naver.com/kim082408",
		"https://smartstore.naver.com/shinjintex",
		"https://smartstore.naver.com/sunsports",
		"https://smartstore.naver.com/ajiyoutong",
		"https://smartstore.naver.com/haewhadang",
		"https://smartstore.naver.com/howmuchmall",
		"https://smartstore.naver.com/dearl",
		"https://smartstore.naver.com/jejuhetsal",
		"https://smartstore.naver.com/vaeracompany",
		"https://smartstore.naver.com/saleshop1",
		"https://smartstore.naver.com/danurimu",
		"https://smartstore.naver.com/sang_sang",
		"https://smartstore.naver.com/2019healthup",
		"https://smartstore.naver.com/gorebobcom",
		"https://smartstore.naver.com/pandalife",
		"https://smartstore.naver.com/lallalashop",
		"https://smartstore.naver.com/novaliving",
		"https://smartstore.naver.com/hgcosstore",
		"https://smartstore.naver.com/ckmart",
		"https://smartstore.naver.com/ak1860",
		"https://smartstore.naver.com/hmbag",
		"https://smartstore.naver.com/arthome",
		"https://smartstore.naver.com/sodanana",
		"https://smartstore.naver.com/thenaturefarmers",
		"https://smartstore.naver.com/anymall2",
		"https://smartstore.naver.com/lifeaddspecial",
		"https://smartstore.naver.com/chaligo",
		"https://smartstore.naver.com/sekwangpowertools",
		"https://smartstore.naver.com/altair17",
		"https://smartstore.naver.com/wonwoo",
		"https://smartstore.naver.com/foodnplan",
		"https://smartstore.naver.com/bathfriends",
		"https://smartstore.naver.com/factory-ppokppogi",
		"https://smartstore.naver.com/bungmatmall",
		"https://smartstore.naver.com/k-nutra",
		"https://smartstore.naver.com/bok2nefood",
		"https://smartstore.naver.com/qvr",
		"https://smartstore.naver.com/pepecandle",
		"https://smartstore.naver.com/lottesuw98",
		"https://smartstore.naver.com/feelplant",
		"https://smartstore.naver.com/globalseller07",
		"https://smartstore.naver.com/sharpbedding",
		"https://smartstore.naver.com/choisejeong",
		"https://smartstore.naver.com/babience",
		"https://smartstore.naver.com/dahanmall",
		"https://smartstore.naver.com/mintgreen",
		"https://smartstore.naver.com/hyun",
		"https://smartstore.naver.com/maendrong",
	}

	stores := []*StoreInfo{}
	for _, site := range sites {
		stores = append(stores, CrawlSmartStoreDetail(site))
	}

	WriteStoreBriefExcel(stores, "sohomall")
}

func CrawlSmartStoreDetail(storeUrl string) *StoreInfo {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
		r.Headers.Set("Cookie", "SBC=b85971e8-6313-4433-9b40-85932ee4d5be")
	})

	storeInfo := &StoreInfo{}

	c.OnHTML("._3MuEQCqxSb", func(e *colly.HTMLElement) {
		naverShopName := e.ChildText("strong.name")

		rawInfo := e.ChildText("strong._3rOxlIskeS + div")

		phoneNumberRawStr := e.ChildText("div._2PXb_kpdRh:first-of-type")

		phoneNumber := phoneNumberRawStr
		idx := strings.Index(phoneNumberRawStr, "인증")
		if idx != -1 {
			phoneNumber = phoneNumberRawStr[0:idx]
		}

		email := ""
		emailIdx := strings.Index(rawInfo, "e-mail")
		if idx != -1 {
			email = rawInfo[emailIdx+6:]
		}

		storeInfo.Email = email
		storeInfo.CustomerCenterPhone = phoneNumber
		storeInfo.StoreName = naverShopName
		storeInfo.RawInfo = rawInfo
		storeInfo.Url = storeUrl + "/profile"
	})

	profileUrl := storeUrl + "/profile"
	err := c.Visit(profileUrl)
	if err != nil {
		log.Println("ERROR on crawling detail: ", err)
	}

	return storeInfo
}
