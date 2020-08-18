package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sterks/apiGosCloud/common"

	"github.com/sterks/apiGosCloud/models"
)

// ProxyURLDebug ...
const ProxyURLDebug = "http://127.0.0.1:8888"

func main() {
	tokenObj := GetAutorize()
	AiCreate(tokenObj)
	res := AiStatuses(tokenObj)
	var status string
	for status != "3" {
		for _, r := range res {
			if r.Status == "3" {
				status = r.Status
				break
			} else {
				status = r.Status
				continue
			}
		}
		fmt.Printf("Статус равен : %v \n", status)
		if status == "3" {
			break
		}
		time.Sleep(10 * time.Second)
	}
	fmt.Println("Готов к отправке файлов")
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

	proxyURL, err := url.Parse(ProxyURLDebug)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
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

	var aiCreate models.ReqCreateGaranty   // Запрос
	var res models.ResCreateGaranty        // Ответ
	var founderFLDate models.FounderFLData // Подзапрос структуры

	// ownerTT := make(map[string]string)
	// var ffff models.FounderFLData

	aiCreate.InternalID = "226-2025"
	aiCreate.Completeness = "2"
	aiCreate.EntityINN = "7723475739"
	aiCreate.CustomerINN = "5009043833"
	aiCreate.OKOPF = "12200"
	aiCreate.Purchase = "32009353877"
	aiCreate.ProductType = "Гарантия исполнения обязательств по контракту"
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
	aiCreate.BGSpecPhone = "79163451354"
	aiCreate.BGSpecExtraPhone = "79163451354"
	aiCreate.BGSpecMail = "akolmykova@mcpuet.ru"
	aiCreate.USResident = "Нет"
	aiCreate.NameGoverning = "Нет"

	// ownerTT["OWNER_TYPE_STATE"] = "Государство или государственные органы"
	// ownerTT["OWNER_TYPE_INTERNATIONAL"] = "Международная организация"
	// ownerTT["OWNER_TYPE_JUR"] = "Юридическое лицо"
	// ownerTT["OWNER_TYPE_PHYS"] = "Физические лица"
	// ownerTT["OWNER_TYPE_OTHER"] = "Прочее"
	var ownerT = []string{"Государство или государственные органы", "Международная организация", "Юридическое лицо", "Физические лица", "Прочее"}
	aiCreate.OwnerType = ownerT
	aiCreate.USTaxResidents = "Нет"

	founderFLDate.TypeP = "Физическое лицо"
	founderFLDate.INN = ""
	founderFLDate.SharePercent = "100"
	founderFLDate.Surname = "Борисенков"
	founderFLDate.Name = "Роман"
	founderFLDate.Patronomic = "Александрович"
	founderFLDate.Post = "ГЕНЕРАЛЬНЫЙ ДИРЕКТОР"
	founderFLDate.Founder = "Да"
	founderFLDate.Manager = "Да"
	founderFLDate.Sig = "Да"
	founderFLDate.Birth = "09.11.1984"
	founderFLDate.PBirthPlace = "ГОР. ГОРЬКИЙ"
	founderFLDate.Gender = "мужской"
	founderFLDate.ResidendRF = "Да"
	founderFLDate.DocumentType = "паспорт РФ"
	founderFLDate.Pseries = "2204/447368"
	founderFLDate.PNumber = "2204/447368"
	founderFLDate.PProducted = "УПРАВЛЕНИЕМ ВНУТРЕННИХ ДЕЛ НИЖЕГОРОДСКОГО РАЙОНА ГОРОДА НИЖНЕГО НОВГОРОДА"
	founderFLDate.PDate = "12.03.2005"
	founderFLDate.PCode = "522-005"
	founderFLDate.MigrNum = ""
	founderFLDate.MigrDate1 = ""
	founderFLDate.MigrDate2 = ""
	founderFLDate.DocumentSeries = ""
	founderFLDate.DocNum = ""
	founderFLDate.DocDate1 = ""
	founderFLDate.DocDate2 = ""
	founderFLDate.PostCode = "603081"
	founderFLDate.RegionCode = "52"
	founderFLDate.Region = "Нижегородская область"
	founderFLDate.RType = ""
	founderFLDate.Region = ""
	founderFLDate.CType = "Город"
	founderFLDate.City = "Нижний Новгород"
	founderFLDate.CitiZenShip = "Да"
	founderFLDate.Region = "Да"
	founderFLDate.TType = ""
	founderFLDate.Town = ""
	founderFLDate.OKTMO = ""
	founderFLDate.Stype = "Улица"
	founderFLDate.Street = "Краснозвездная"
	founderFLDate.House = "7А"
	founderFLDate.Korp = ""
	founderFLDate.Building = ""
	founderFLDate.Flat = "76"
	founderFLDate.OKATO = "22401379000"
	founderFLDate.FactualPostCode = "603081"
	founderFLDate.FFactualRegionCode = "Да"
	founderFLDate.FactualRegion = "Нижегородская область"
	founderFLDate.FactualRType = ""
	founderFLDate.FactualRegign = ""
	founderFLDate.FactualCType = ""
	founderFLDate.FactualCity = "г. Нижний Новгород"
	founderFLDate.FactualTType = ""
	founderFLDate.FactualTown = ""
	founderFLDate.FactualTown = ""
	founderFLDate.FactualStype = ""
	founderFLDate.FactualStreet = "ул. Краснозвездная"
	founderFLDate.FactualHouse = "д. 7А"
	founderFLDate.FactualKorp = ""
	founderFLDate.FactualBuilding = ""
	founderFLDate.FactualFlat = ""
	founderFLDate.FactualOKATO = "22401379000"
	aiCreate.FounderFLData = append(aiCreate.FounderFLData, founderFLDate)

	bb, err := json.Marshal(aiCreate)
	if err != nil {
		log.Fatal(err)
	}
	common.JSONPretty(bb)
	proxyURL, err := url.Parse(ProxyURLDebug)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
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
	if err5 := json.Unmarshal(br, &res); err5 != nil {
		log.Fatal(err5)
	}
	common.JSONPretty(br)
	if err := ioutil.WriteFile("Task.txt", []byte(aiCreate.InternalID), 0777); err != nil {
		log.Fatal(err)
	}
}

// AiStatuses Получение статуса
func AiStatuses(token models.Token) []models.ResStatus {
	InternalID, err := ioutil.ReadFile("Task.txt")
	var arTask []string
	var res []models.ResStatus
	arTask = append(arTask, string(InternalID))
	bb, err := json.Marshal(arTask)
	proxyURL, err := url.Parse(ProxyURLDebug)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	req, err := http.NewRequest("POST", "http://esb-dev.gosoblako.ru/services/ai_statuses", bytes.NewReader(bb))
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
	if err5 := json.Unmarshal(br, &res); err5 != nil {
		log.Fatal(err5)
	}
	common.JSONPretty(br)
	return res
}
