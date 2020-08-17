package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sterks/apiGosCloud/models"
)

func main() {
	token := GetAutorize()
	fmt.Println(token.AccessToken)
	// AiAddMessages(token)
	AiCreate(token)
}

// GetAutorize Авторизация
func GetAutorize() models.Token {
	var user models.User
	var token models.Token
	user.ClientID = "nekrasov"
	user.Password = "h4nqQffEWHsP9"

	us, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err2 := http.NewRequest("POST", "http://esb-dev.gosoblako.ru/services/tokens", bytes.NewReader(us))
	if err2 != nil {
		log.Fatal(err2)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err3 := client.Do(req)
	if err3 != nil {
		log.Fatal(err3)
	}

	bb, err4 := ioutil.ReadAll(res.Body)
	if err4 != nil {
		log.Fatal(err4)
	}

	if err5 := json.Unmarshal(bb, &token); err5 != nil {
		log.Fatal(err5)
	}
	return token
}

//AiAddMessages Предназначен для выгрузки сообщений  Агенту из чата  в ГО по текущей заявке
func AiAddMessages(token models.Token) {

	var req models.ReqMessages
	var res models.ResMessages
	req.InternalID = "2d3abb1cfcdf4da09cc4a8bcd2b81b7f"
	layout := "2006-01-02 15:04:05"
	str1 := "2020-05-07 18:00:39"
	str2 := "2020-05-14 12:53:05"
	t1, err := time.Parse(layout, str1)
	t2, err := time.Parse(layout, str2)
	if err != nil {
		fmt.Println(err)
	}
	req.ResultM.Messages1.Date = t1
	req.ResultM.Messages1.ID = "681325"
	req.ResultM.Messages1.Mess = "Тестовое сообщение"
	req.ResultM.Messages1.User = "GOB"

	req.ResultM.Messages2.Date = t2
	req.ResultM.Messages2.ID = "681333"
	req.ResultM.Messages2.Mess = "Тестовое сообщение"
	req.ResultM.Messages2.User = "GOB"

	bb, err := json.Marshal(req)
	fmt.Println(string(bb))

	client := &http.Client{}
	r, err := http.NewRequest("POST", "http://esb-dev.gosoblako.ru/services/ai_add_messages", bytes.NewReader(bb))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("x-auth-token", token.AccessToken)
	rs, err := client.Do(r)

	br, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(br))
	if err5 := json.Unmarshal(br, &res); err5 != nil {
		log.Fatal(err5)
	}
	fmt.Println(res)
}

//AiCreate Сервис для создания в Базе данных заявки от источника заявок Платформа Агент.
func AiCreate(token models.Token) {

	var aiCreate models.ReqCreateGaranty
	var res models.ResCreateGaranty

	aiCreate.InternalID = "226-2020"
	aiCreate.Completeness = "2"
	aiCreate.EntityINN = "5263093100"
	aiCreate.CustomerINN = "7708503727"
	aiCreate.OKOPF = "12200"
	aiCreate.Purchase = "31908157568"
	aiCreate.ProductType = "Гарантия по возврату аванса"
	aiCreate.BGStart = "14.08.2019"
	aiCreate.BGEnd = "31.01.2021"
	aiCreate.BGStartTender = "31.01.2021"
	aiCreate.BGEndTender = "31.01.2021"
	aiCreate.BGEndContract = "31.01.2021"
	aiCreate.BegPrice = 38420800.03
	aiCreate.FinalPrice = 29489000.00
	aiCreate.BegPrice = 8846700.00
	aiCreate.BiGPartSum = 8846700.00
	aiCreate.IsAvans = "Нет"
	aiCreate.IsBeneficiaryBG = "Нет"
	aiCreate.IsNondisscusWriteOff = "Да"
	aiCreate.TaxMethod = "Общая"
	aiCreate.FinPeriod = "3 мес."
	aiCreate.Account = "40702810302680000204"
	aiCreate.BIK = "044525593"
	aiCreate.BankName = "АО АЛЬФА-БАНК"
	aiCreate.BGSpecFName = "Борисенков"
	aiCreate.BGSpecLNmae = "Роман"
	aiCreate.BGSpecSName = "Александрович"
	aiCreate.BGSpecLPost = "Генеральный директор"
	aiCreate.BGSpecPhone = "89163451354"
	aiCreate.BGSpecExtraPhone = "89163451354"
	aiCreate.BGSpecMail = "akolmykova@mcpuet.ru"
	aiCreate.FounderFLData.TypeP = "Физическое лицо"
	aiCreate.FounderFLData.INN = ""
	aiCreate.FounderFLData.SharePercent = "100"
	aiCreate.FounderFLData.Surname = "Борисенков"
	aiCreate.FounderFLData.Name = "Роман"
	aiCreate.FounderFLData.Patronomic = "Александрович"
	aiCreate.FounderFLData.Post = "ГЕНЕРАЛЬНЫЙ ДИРЕКТОР"
	aiCreate.FounderFLData.Founder = "Да"
	aiCreate.FounderFLData.Manager = "Да"
	aiCreate.FounderFLData.Sig = "Да"
	aiCreate.FounderFLData.Birth = "09.11.1984"
	aiCreate.FounderFLData.PBirthPlace = "ГОР. ГОРЬКИЙ"
	aiCreate.FounderFLData.Gender = "мужской"
	aiCreate.FounderFLData.ResidendRF = "Да"
	aiCreate.FounderFLData.DocumentType = "паспорт РФ"
	aiCreate.FounderFLData.Pseries = "2204/447368"
	aiCreate.FounderFLData.PNumber = "2204/447368"
	aiCreate.FounderFLData.PProducted = "УПРАВЛЕНИЕМ ВНУТРЕННИХ ДЕЛ НИЖЕГОРОДСКОГО РАЙОНА ГОРОДА НИЖНЕГО НОВГОРОДА"
	aiCreate.FounderFLData.PDate = "12.03.2005"
	aiCreate.FounderFLData.PCode = "522-005"
	aiCreate.FounderFLData.MigrNum = ""
	aiCreate.FounderFLData.MigrDate1 = ""
	aiCreate.FounderFLData.MigrDate2 = ""
	aiCreate.FounderFLData.DocumentSeries = ""
	aiCreate.FounderFLData.DocNum = ""
	aiCreate.FounderFLData.DocDate1 = ""
	aiCreate.FounderFLData.DocDate2 = ""
	aiCreate.FounderFLData.PostCode = "603081"
	aiCreate.FounderFLData.RegionCode = "52"
	aiCreate.FounderFLData.Region = "Нижегородская область"
	aiCreate.FounderFLData.RType = ""
	aiCreate.FounderFLData.Region = ""
	aiCreate.FounderFLData.CType = "Город"
	aiCreate.FounderFLData.City = "Нижний Новгород"
	aiCreate.FounderFLData.TType = ""
	aiCreate.FounderFLData.Town = ""
	aiCreate.FounderFLData.OKTMO = ""
	aiCreate.FounderFLData.Stype = "Улица"
	aiCreate.FounderFLData.Street = "Краснозвездная"
	aiCreate.FounderFLData.House = "7А"
	aiCreate.FounderFLData.Korp = ""
	aiCreate.FounderFLData.Building = ""
	aiCreate.FounderFLData.Flat = "76"
	aiCreate.FounderFLData.OKATO = "22401379000"
	aiCreate.FounderFLData.FactualPostCode = "603081"
	aiCreate.FounderFLData.FFactualRegionCode = ""
	aiCreate.FounderFLData.FactualRegion = "Нижегородская область"
	aiCreate.FounderFLData.FactualRType = ""
	aiCreate.FounderFLData.FactualRegign = ""
	aiCreate.FounderFLData.FactualCType = ""
	aiCreate.FounderFLData.FactualCity = "г. Нижний Новгород"
	aiCreate.FounderFLData.FactualTType = ""
	aiCreate.FounderFLData.FactualTown = ""
	aiCreate.FounderFLData.FactualTown = ""
	aiCreate.FounderFLData.FactualStype = ""
	aiCreate.FounderFLData.FactualStreet = "ул. Краснозвездная"
	aiCreate.FounderFLData.FactualHouse = "д. 7А"
	aiCreate.FounderFLData.FactualKorp = ""
	aiCreate.FounderFLData.FactualBuilding = ""
	aiCreate.FounderFLData.FactualFlat = ""
	aiCreate.FounderFLData.FactualOKATO = "22401379000"

	bb, err := json.Marshal(aiCreate)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(bb))
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://esb-dev.gosoblako.ru/services/ai_create", bytes.NewReader(bb))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-auth-token", token.AccessToken)
	response, err := client.Do(req)
	br, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(br))
	if err5 := json.Unmarshal(br, &res); err5 != nil {
		log.Fatal(err5)
	}
	fmt.Println(res)
}
