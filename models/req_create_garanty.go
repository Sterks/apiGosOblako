package models

// ReqCreateGaranty ...
type ReqCreateGaranty struct {
	InternalID           string  `json:"internal_id"`             // Номер заявки
	Completeness         string  `json:"completeness"`            // Полнота данных
	EntityINN            string  `json:"entity_inn"`              // ИНН Организации-заявителя
	CustomerINN          string  `json:"customer_inn"`            // ИНН заказчика
	OKOPF                string  `json:"okopf"`                   // ОКОПФ
	Purchase             string  `json:"purchase"`                // Номер закупки
	ProductType          string  `json:"product_type"`            // БГ: Гарантия исполнения обязательств по контракту/ Гарантия по возврату аванса / Гарантия, обеспечивающая обязательства в период гарантийного срока/ Гарантия на участие
	BGStart              string  `json:"bg_start"`                // Дата выдачи
	BGEnd                string  `json:"bg_end"`                  // Дата окончания
	BGStartTender        string  `json:"bg_start_tender"`         // Дата выдачи Гарантии на участие
	BGEndTender          string  `json:"bg_end_tender"`           // Дата окончания Гарантии на участие
	BGEndContract        string  `json:"bg_end_contract"`         // Дата окончания контракта
	BegPrice             float32 `json:"beg_price"`               // НМЦ
	FinalPrice           float32 `json:"final_price"`             // Предложенная цена контракта
	BiGPartSum           float32 `json:"bid_partSum"`             // Требуемая сумма //обязательно, если bg_type не равен "Гарантия на участие"
	BiGPartSumTender     float32 `json:"bid_partSum_tender"`      // Требуемая сумма Гарантии на участие //обязательно, если bg_type равен "Гарантия на участие"
	IsAvans              string  `json:"is_avans"`                // Значения: Да/Нет. По умолчанию - Нет
	IsBeneficiaryBG      string  `json:"is_beneficiary_bg"`       // БГ по форме бенефициара // Значения: Да/Нет. По умолчанию - Нет
	IsNondisscusWriteOff string  `json:"is_nondisscus_write_off"` // Условие бесспорного списания // Значения: Да/Нет. По умолчанию - Нет
	TaxMethod            string  `json:"tax_method"`              // Форма налогообложения // Значения: Общая/УСН. По умолчанию - Нет
	FinPeriod            string  `json:"fin_period"`              // Период отчетности
	Account              string  `json:"account"`                 // Расчетный счет
	BIK                  string  `json:"bik"`                     // BIK
	BankName             string  `json:"bank_name"`               // Наименование банка
	BGSpecLNmae          string  `json:"bg_spec_l_name"`          // Контактное лицо имя
	BGSpecFName          string  `json:"bg_spec_f_name"`          // Контактное лицо фамилия
	BGSpecSName          string  `json:"bg_spec_s_name"`          // Контактное лицо отчество
	BGSpecLPost          string  `json:"bg_spec_l_post"`          // Контактное лицо должность
	BGSpecPhone          string  `json:"bg_spec_phone"`           // Контактное лицо рабочий телефон
	BGSpecExtraPhone     string  `json:"bg_spec_extra_phone"`     // Контактное лицо Мобильный телефон
	BGSpecMail           string  `json:"bg_spec_mail"`
	FounderFLData        FounderFLData
}

// FounderFLData ...
type FounderFLData struct {
	TypeP              string `json:"TypeP"`                 // Физическое лицо"
	INN                string `json:"INN"`                   // ИНН
	SharePercent       string `json:"SharePercent"`          // Доля в уставном капитале
	Surname            string `json:"Surname"`               // Фамилия
	Name               string `json:"Name"`                  // Имя
	Patronomic         string `json:"Patronomic"`            // Отчество
	Post               string `json:"Post"`                  // Должность
	Founder            string `json:"Founder"`               // Является ли учредителем?
	Manager            string `json:"Manager"`               // Является руководителем организации?
	Sig                string `json:"sig"`                   // Имеет право подписи финансовых документов?
	Birth              string `json:"birth"`                 // Дата рождения
	PBirthPlace        string `json:"pbirthplace"`           // Место рождения
	Gender             string `json:"gender"`                // Пол
	ResidendRF         string `json:"residend_rf"`           // Резидент
	DocumentType       string `json:"document_type"`         // Документ
	Pseries            string `json:"pseries"`               // Серия
	PNumber            string `json:"pnumber"`               // Номер документа
	PProducted         string `json:"pproducted"`            // Кем выдан
	PDate              string `json:"pdate"`                 // Дата выдачи
	PCode              string `json:"pcode"`                 // Код подразделения
	MigrNum            string `json:"migr_num"`              // Миграционная карта, номер
	MigrDate1          string `json:"migr_date1"`            // Миграционная карта, дата с
	MigrDate2          string `json:"migr_date2"`            // Миграционная карта, дата по
	DocumentSeries     string `json:"doc_series"`            // Удостоверениче личности, серия
	DocNum             string `json:"doc_num"`               // Удостоверениче личности, номер
	DocDate1           string `json:"doc_date1"`             // Удостоверение личности, дата с
	DocDate2           string `json:"doc_date2"`             // Удостоверение личности, дата по
	PostCode           string `json:"post_code"`             // Индекс
	RegionCode         string `json:"region_code"`           // Код региона
	Region             string `json:"region"`                // Регион
	RType              string `json:"rtype"`                 // Район (вид)
	Reign              string `json:"reign"`                 // Район
	CType              string `json:"ctype"`                 // Город (вид)
	City               string `json:"city"`                  // Город
	TType              string `json:"ttype"`                 // Населенный пункт (вид)
	Town               string `json:"town"`                  // Населенный пункт
	OKTMO              string `json:"oktmo"`                 // ОКТМО
	Stype              string `json:"stype"`                 // Улица (вид)
	Street             string `json:"street"`                // Улица (название)
	House              string `json:"house"`                 // Дом
	Korp               string `json:"korp"`                  // Корпус
	Building           string `json:"building"`              // Строение
	Flat               string `json:"flat"`                  // Квартира (офис)
	OKATO              string `json:"okato"`                 // OKATO
	FactualPostCode    string `json:"factual_post_code"`     // Индекс
	FFactualRegionCode string `json:"f_factual_region_code"` // Код региона
	FactualRegion      string `json:"factual_region"`        // Регион
	FactualRType       string `json:"factual_rtype"`         // Район (вид)
	FactualRegign      string `json:"factual_reign"`         // Район (вид)
	FactualCType       string `json:"factual_ctype"`         // Район
	FactualCity        string `json:"factual_city"`          //  Город (вид)
	FactualTType       string `json:"factual_ttype"`         // Населенный пункт (вид)
	FactualTown        string `json:"factual_town"`          // Населенный пункт
	FactualOKTMO       string `json:"factual_oktmo"`         // ОКТМО
	FactualStype       string `json:"factual_stype"`         // Улица (вид)
	FactualStreet      string `json:"factual_street"`        // Улица (название)
	FactualHouse       string `json:"factual_house"`         // Дом
	FactualKorp        string `json:"factual_korp"`          // Корпус
	FactualBuilding    string `json:"factual_building"`      // Строение
	FactualFlat        string `json:"factual_flat"`          // Квартира (офис)
	FactualOKATO       string `json:"factual_okato"`         // ОКАТО
}
