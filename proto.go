package main

type Message struct {
	Title   string
	Content string
}

type ConfigOption struct {
	Modals         []string `yaml:"modals"`
	NotifyUrl      []string `yaml:"notifyUrl"`
	Location       string   `yaml:"location"`
	SearchInterval int      `yaml:"searchInterval"`
}

type SearchResponse struct {
	// Head SearchRespHead `json:"head"`
	Body SearchRespBody `json:"body"`
}

type SearchRespHead struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type SearchRespBody struct {
	Content Content `json:"content"`
}

type Content struct {
	PickupMessage PickupMessage1 `json:"pickupMessage"`
}

type PickupMessage1 struct {
	Stores []Store `json:"stores"`
}

type Store struct {
	StoreName         string            `json:"storeName"`
	PartsAvailability PartsAvailability `json:"partsAvailability"`
}

type PartsAvailability map[string]PartsAvailabilityValue //型号 => info

type PartsAvailabilityValue struct {
	PickupSearchQuote       string `json:"pickupSearchQuote"`       //可取货
	StorePickupProductTitle string `json:"storePickupProductTitle"` //iPhone 13 512GB 粉色
}

//有货
//{"head":{"status":"200","data":{}},"body":{"content":{"pickupMessage":{"stores":[{"storeEmail":"holidayplazashenzhen@apple.com","storeName":"深圳益田假日广场","reservationUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","makeReservationUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","state":"广东","storeImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R484.png?resize=828:*&output-format=jpg","country":"CN","city":"深圳","storeNumber":"R484","partsAvailability":{"MLE93CH/A":{"storePickEligible":true,"storeSearchEnabled":true,"storeSelectionEnabled":true,"storePickupQuote":"今天；Apple 深圳益田假日广场","storePickupQuote2_0":"<span class=\"as-pickup-quote-availability-quote\">今天</span>; <button type=\"button\" class=\"rf-pickup-quote-overlay-trigger as-retailavailabilitytrigger-infobutton retail-availability-search-trigger as-buttonlink\" data-ase-overlay=\"buac-overlay\" data-ase-click=\"show\">Apple 深圳益田假日广场</button>","pickupSearchQuote":"今天可取货","storePickupLabel":"取货 (店内)：","partNumber":"MLE93CH/A","purchaseOption":"","ctoOptions":"","storePickupLinkText":"查看供货情况","storePickupProductTitle":"iPhone 13 512GB 粉色","pickupDisplay":"available","pickupType":"店内取货"}},"phoneNumber":"4006171254","pickupTypeAvailabilityText":"此地点提供店内取货服务。","address":{"address":"Apple 深圳益田假日广场","address3":null,"address2":"深圳市南山区深南大道 9028 号益田假日广场","postalCode":"518000"},"hoursUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","storeHours":{"storeHoursText":"Store Hours","bopisPickupDays":"Days","bopisPickupHours":"Hours","hours":[{"storeTimings":"10:00 - 22:30","storeDays":"周五-周六:"},{"storeTimings":"10:00 - 22:00","storeDays":"周一-周四, 周日:"}]},"specialHours":{"specialHoursText":"特别营业时间：","bopisPickupDays":"Days","bopisPickupHours":"Hours","specialHoursData":[{"specialDays":"Sep 25:","specialTimings":"09:00 - 22:30"},{"specialDays":"Sep 26:","specialTimings":"09:00 - 22:00"},{"specialDays":"Sep 30:","specialTimings":"10:00 - 22:30"}],"viewAllSpecialHours":true},"storelatitude":22.53986,"storelongitude":113.97079,"storedistance":2.6,"storeDistanceWithUnit":"2.6 km","storeDistanceVoText":"2.6 km from 518000","retailStore":{"storeNumber":"R484","storeUniqueId":"R484","name":"深圳益田假日广场","storeTypeKey":"1","storeSubTypeKey":"0","storeType":"APPLESTORE_DEFAULT","phoneNumber":"4006171254","email":"holidayplazashenzhen@apple.com","carrierCode":null,"locationType":null,"latitude":22.53986,"longitude":113.97079,"address":{"city":"深圳","companyName":"Apple 深圳益田假日广场","countryCode":"CN","county":null,"district":"深圳","geoCode":null,"label":null,"languageCode":"zh-cn","mailStop":null,"postalCode":"518000","province":null,"state":"广东","street":"深圳市南山区深南大道 9028 号益田假日广场","street2":null,"street3":null,"suburb":null,"type":"SHIPPING","addrSourceType":null,"outsideCityFlag":null,"daytimePhoneAreaCode":null,"eveningPhoneAreaCode":null,"daytimePhone":"4006171254","fullPhoneNumber":null,"eveningPhone":null,"emailAddress":null,"firstName":null,"lastName":null,"suffix":null,"lastNamePhonetic":null,"firstNamePhonetic":null,"title":null,"businessAddress":false,"uuid":"e03dbe5a-a065-4fcb-899e-e5366e896e70","mobilePhone":null,"mobilePhoneAreaCode":null,"cityStateZip":null,"daytimePhoneSelected":false,"middleName":null,"residenceStatus":null,"moveInDate":null,"subscriberId":null,"locationType":null,"carrierCode":null,"metadata":{},"verificationState":"UN_VERIFIED","expiration":null,"markForDeletion":false,"primaryAddress":false,"fullEveningPhone":null,"fullDaytimePhone":"4006171254","correctionResult":null,"twoLineAddress":"深圳市南山区深南大道 9028 号益田假日广场,\n深圳, 广东, 518000","addressVerified":false},"urlKey":null,"directionsUrl":null,"storeImageUrl":"http://rtlimages.apple.com/cmc/dieter/store/4_3/R484.png?resize=828:*&output-format=jpg","makeReservationUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","hoursAndInfoUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","storeHours":[{"storeDays":"周五-周六","voStoreDays":null,"storeTimings":"10:00 - 22:30 "},{"storeDays":"周一-周四, 周日","voStoreDays":null,"storeTimings":"10:00 - 22:00 "}],"storeHolidays":[{"date":"Sep 25","description":"Fire Alarm","hours":"09:00 - 22:30 ","comments":"New Product Launch Day Two.","closed":false},{"date":"Sep 26","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":"New Product Launch Day Three.","closed":false},{"date":"Sep 30","description":"Fire Alarm","hours":"10:00 - 22:30 ","comments":"National Holiday","closed":false}],"secureStoreImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R484.png?resize=828:*&output-format=jpg","distance":2.6,"distanceUnit":"km","distanceWithUnit":"2.6 km","timezone":"Asia/Shanghai","storeIsActive":true,"lastUpdated":0.0,"lastFetched":1632534119615,"dateStamp":"24-Sep-2021","distanceSeparator":".","nextAvailableDate":null,"storeHolidayLookAheadWindow":7,"driveDistanceWithUnit":null,"driveDistanceInMeters":null,"dynamicAttributes":{},"storePickupMethodByType":{"INSTORE":{"type":"INSTORE","services":["APU"],"typeCoordinate":{"lat":22.53986,"lon":113.97079},"typeDirection":{"directionByLocale":null},"typeMeetupLocation":{"meetingLocationByLocale":null}}},"storeTimings":null,"availableNow":true},"storelistnumber":1,"storeListNumber":1,"pickupOptionsDetails":{"whatToExpectAtPickup":"<span class=\"as-pickupmethods-intro-header\">取货须知</span><br />当你的订单准备就绪后，我们会向你发送详细的取货说明电子邮件。有关新设备设置，你可以预约免费在线辅导，让 Specialist 专家为你提供指导。","comparePickupOptionsLink":"<a href=\"https://www.apple.com.cn/shop/shipping-pickup\" data-feature-name=\"Astro Link\" data-display-name=\"AOS: shop/shipping-pickup\" target=\"_blank\">进一步了解送货<br />和取货<span class=\"icon icon-after icon-chevronright\"></span><span class=\"visuallyhidden\">(在新窗口中打开)</span></a>","pickupOptions":[{"pickupOptionTitle":"店内","pickupOptionDescription":"提取你的在线订单商品。你可以获得设置帮助，还能选购配件。可能需要测量体温并佩戴口罩。","index":1}]},"rank":1,"pickupEncodedUpperDateString":"20210925"},{"storeEmail":"zhujiangnewtown@apple.com","storeName":"珠江新城","reservationUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","makeReservationUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","state":"广东","storeImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R639.png?resize=828:*&output-format=jpg","country":"CN","city":"广州","storeNumber":"R639","partsAvailability":{"MLE93CH/A":{"storePickEligible":true,"storeSearchEnabled":true,"storeSelectionEnabled":true,"storePickupQuote":"星期一 2021/09/27；Apple 珠江新城","storePickupQuote2_0":"<span class=\"as-pickup-quote-availability-quote\">星期一 2021/09/27</span>; <button type=\"button\" class=\"rf-pickup-quote-overlay-trigger as-retailavailabilitytrigger-infobutton retail-availability-search-trigger as-buttonlink\" data-ase-overlay=\"buac-overlay\" data-ase-click=\"show\">Apple 珠江新城</button>","pickupSearchQuote":"星期一 2021/09/27可取货","storePickupLabel":"取货 (店内)：","partNumber":"MLE93CH/A","purchaseOption":"","ctoOptions":"","storePickupLinkText":"查看供货情况","storePickupProductTitle":"iPhone 13 512GB 粉色","pickupDisplay":"available","pickupType":"店内取货"}},"phoneNumber":"4006393601","pickupTypeAvailabilityText":"此地点提供店内取货服务。","address":{"address":"Apple 珠江新城","address3":"天汇广场 1 层","address2":"广州珠江新城兴民路 222 号","postalCode":"510623"},"hoursUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","storeHours":{"storeHoursText":"Store Hours","bopisPickupDays":"Days","bopisPickupHours":"Hours","hours":[{"storeTimings":"10:00 - 22:00","storeDays":"周一-周日:"}]},"specialHours":{"specialHoursText":"特别营业时间：","bopisPickupDays":"Days","bopisPickupHours":"Hours","specialHoursData":[{"specialDays":"Sep 25:","specialTimings":"09:00 - 22:00"},{"specialDays":"Sep 26:","specialTimings":"09:00 - 22:00"}],"viewAllSpecialHours":false},"storelatitude":23.11843672960212,"storelongitude":113.32728981971741,"storedistance":89.59,"storeDistanceWithUnit":"89.59 km","storeDistanceVoText":"89.59 km from 510623","retailStore":{"storeNumber":"R639","storeUniqueId":"R639","name":"珠江新城","storeTypeKey":"1","storeSubTypeKey":"0","storeType":"APPLESTORE_DEFAULT","phoneNumber":"4006393601","email":"zhujiangnewtown@apple.com","carrierCode":null,"locationType":null,"latitude":23.11843672960212,"longitude":113.32728981971741,"address":{"city":"广州","companyName":"Apple 珠江新城","countryCode":"CN","county":null,"district":"广州","geoCode":null,"label":null,"languageCode":"zh-cn","mailStop":null,"postalCode":"510623","province":null,"state":"广东","street":"广州珠江新城兴民路 222 号","street2":"天汇广场 1 层","street3":null,"suburb":null,"type":"SHIPPING","addrSourceType":null,"outsideCityFlag":null,"daytimePhoneAreaCode":null,"eveningPhoneAreaCode":null,"daytimePhone":"4006393601","fullPhoneNumber":null,"eveningPhone":null,"emailAddress":null,"firstName":null,"lastName":null,"suffix":null,"lastNamePhonetic":null,"firstNamePhonetic":null,"title":null,"businessAddress":false,"uuid":"eba45599-9daa-4738-b700-818fd4e312e5","mobilePhone":null,"mobilePhoneAreaCode":null,"cityStateZip":null,"daytimePhoneSelected":false,"middleName":null,"residenceStatus":null,"moveInDate":null,"subscriberId":null,"locationType":null,"carrierCode":null,"metadata":{},"verificationState":"UN_VERIFIED","expiration":null,"markForDeletion":false,"primaryAddress":false,"fullEveningPhone":null,"fullDaytimePhone":"4006393601","correctionResult":null,"twoLineAddress":"广州珠江新城兴民路 222 号, 天汇广场 1 层,\n广州, 广东, 510623","addressVerified":false},"urlKey":null,"directionsUrl":null,"storeImageUrl":"http://rtlimages.apple.com/cmc/dieter/store/4_3/R639.png?resize=828:*&output-format=jpg","makeReservationUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","hoursAndInfoUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","storeHours":[{"storeDays":"周一-周日","voStoreDays":null,"storeTimings":"10:00 - 22:00 "}],"storeHolidays":[{"date":"Sep 25","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":"iPhone 13 Launch Day 2","closed":false},{"date":"Sep 26","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":"iPhone 13 Launch Day 3","closed":false}],"secureStoreImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R639.png?resize=828:*&output-format=jpg","distance":89.59,"distanceUnit":"km","distanceWithUnit":"89.59 km","timezone":"Asia/Shanghai","storeIsActive":true,"lastUpdated":0.0,"lastFetched":1632534119615,"dateStamp":"24-Sep-2021","distanceSeparator":".","nextAvailableDate":null,"storeHolidayLookAheadWindow":7,"driveDistanceWithUnit":null,"driveDistanceInMeters":null,"dynamicAttributes":{},"storePickupMethodByType":{"INSTORE":{"type":"INSTORE","services":["APU"],"typeCoordinate":{"lat":23.11843672960212,"lon":113.32728981971741},"typeDirection":{"directionByLocale":null},"typeMeetupLocation":{"meetingLocationByLocale":null}}},"storeTimings":null,"availableNow":true},"storelistnumber":2,"storeListNumber":2,"pickupOptionsDetails":{"whatToExpectAtPickup":"<span class=\"as-pickupmethods-intro-header\">取货须知</span><br />当你的订单准备就绪后，我们会向你发送详细的取货说明电子邮件。有关新设备设置，你可以预约免费在线辅导，让 Specialist 专家为你提供指导。","comparePickupOptionsLink":"<a href=\"https://www.apple.com.cn/shop/shipping-pickup\" data-feature-name=\"Astro Link\" data-display-name=\"AOS: shop/shipping-pickup\" target=\"_blank\">进一步了解送货<br />和取货<span class=\"icon icon-after icon-chevronright\"></span><span class=\"visuallyhidden\">(在新窗口中打开)</span></a>","pickupOptions":[{"pickupOptionTitle":"店内","pickupOptionDescription":"提取你的在线订单商品。你可以获得设置帮助，还能选购配件。可能需要测量体温并佩戴口罩。","index":1}]},"rank":2,"pickupEncodedUpperDateString":"20210927"},{"storeEmail":"parccentral@apple.com","storeName":"天环广场 ","reservationUrl":"http://www.apple.com/cn/retail/parccentral","makeReservationUrl":"http://www.apple.com/cn/retail/parccentral","state":"广东","storeImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R577.png?resize=828:*&output-format=jpg","country":"CN","city":"广州","storeNumber":"R577","partsAvailability":{"MLE93CH/A":{"storePickEligible":true,"storeSearchEnabled":true,"storeSelectionEnabled":false,"storePickupQuote":"Apple 天环广场目前不可取货","storePickupQuote2_0":"<button type=\"button\" class=\"rf-pickup-quote-overlay-trigger as-retailavailabilitytrigger-infobutton retail-availability-search-trigger as-buttonlink\" data-ase-overlay=\"buac-overlay\" data-ase-click=\"show\">Apple 天环广场</button> 目前无货","pickupSearchQuote":"暂无供应","storePickupLabel":"取货：","partNumber":"MLE93CH/A","purchaseOption":"","ctoOptions":"","storePickupLinkText":"查看供货情况","storePickupProductTitle":"iPhone 13 512GB 粉色","pickupDisplay":"unavailable","pickupType":"店内取货"}},"phoneNumber":"4006139742","pickupTypeAvailabilityText":"此地点提供店内取货服务。","address":{"address":"Apple 天环广场","address3":null,"address2":"广州市天河区天河路 218 号","postalCode":"510000"},"hoursUrl":"http://www.apple.com/cn/retail/parccentral","storeHours":{"storeHoursText":"Store Hours","bopisPickupDays":"Days","bopisPickupHours":"Hours","hours":[{"storeTimings":"10:00 - 22:00","storeDays":"周一-周日:"}]},"specialHours":{"specialHoursText":"特别营业时间：","bopisPickupDays":"Days","bopisPickupHours":"Hours","specialHoursData":[{"specialDays":"Sep 25:","specialTimings":"09:00 - 22:00"},{"specialDays":"Sep 26:","specialTimings":"09:00 - 22:00"}],"viewAllSpecialHours":false},"storelatitude":23.135,"storelongitude":113.31916,"storedistance":91.47,"storeDistanceWithUnit":"91.47 km","storeDistanceVoText":"91.47 km from 510000","retailStore":{"storeNumber":"R577","storeUniqueId":"R577","name":"天环广场 ","storeTypeKey":"1","storeSubTypeKey":"0","storeType":"APPLESTORE_DEFAULT","phoneNumber":"4006139742","email":"parccentral@apple.com","carrierCode":null,"locationType":null,"latitude":23.135,"longitude":113.31916,"address":{"city":"广州","companyName":"Apple 天环广场","countryCode":"CN","county":null,"district":"广州","geoCode":null,"label":null,"languageCode":"zh-cn","mailStop":null,"postalCode":"510000","province":null,"state":"广东","street":"广州市天河区天河路 218 号","street2":null,"street3":null,"suburb":null,"type":"SHIPPING","addrSourceType":null,"outsideCityFlag":null,"daytimePhoneAreaCode":null,"eveningPhoneAreaCode":null,"daytimePhone":"4006139742","fullPhoneNumber":null,"eveningPhone":null,"emailAddress":null,"firstName":null,"lastName":null,"suffix":null,"lastNamePhonetic":null,"firstNamePhonetic":null,"title":null,"businessAddress":false,"uuid":"c110c3c1-d688-478f-b63f-250b0de1a3dc","mobilePhone":null,"mobilePhoneAreaCode":null,"cityStateZip":null,"daytimePhoneSelected":false,"middleName":null,"residenceStatus":null,"moveInDate":null,"subscriberId":null,"locationType":null,"carrierCode":null,"metadata":{},"verificationState":"UN_VERIFIED","expiration":null,"markForDeletion":false,"primaryAddress":false,"fullEveningPhone":null,"fullDaytimePhone":"4006139742","correctionResult":null,"twoLineAddress":"广州市天河区天河路 218 号,\n广州, 广东, 510000","addressVerified":false},"urlKey":null,"directionsUrl":null,"storeImageUrl":"http://rtlimages.apple.com/cmc/dieter/store/4_3/R577.png?resize=828:*&output-format=jpg","makeReservationUrl":"http://www.apple.com/cn/retail/parccentral","hoursAndInfoUrl":"http://www.apple.com/cn/retail/parccentral","storeHours":[{"storeDays":"周一-周日","voStoreDays":null,"storeTimings":"10:00 - 22:00 "}],"storeHolidays":[{"date":"Sep 25","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":null,"closed":false},{"date":"Sep 26","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":null,"closed":false}],"secureStoreImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R577.png?resize=828:*&output-format=jpg","distance":91.47,"distanceUnit":"km","distanceWithUnit":"91.47 km","timezone":"Asia/Shanghai","storeIsActive":true,"lastUpdated":0.0,"lastFetched":1632534119615,"dateStamp":"24-Sep-2021","distanceSeparator":".","nextAvailableDate":null,"storeHolidayLookAheadWindow":7,"driveDistanceWithUnit":null,"driveDistanceInMeters":null,"dynamicAttributes":{},"storePickupMethodByType":{"INSTORE":{"type":"INSTORE","services":["APU"],"typeCoordinate":{"lat":23.135,"lon":113.31916},"typeDirection":{"directionByLocale":null},"typeMeetupLocation":{"meetingLocationByLocale":null}}},"storeTimings":null,"availableNow":true},"storelistnumber":3,"storeListNumber":3,"pickupOptionsDetails":{"whatToExpectAtPickup":"<span class=\"as-pickupmethods-intro-header\">取货须知</span><br />当你的订单准备就绪后，我们会向你发送详细的取货说明电子邮件。有关新设备设置，你可以预约免费在线辅导，让 Specialist 专家为你提供指导。","comparePickupOptionsLink":"<a href=\"https://www.apple.com.cn/shop/shipping-pickup\" data-feature-name=\"Astro Link\" data-display-name=\"AOS: shop/shipping-pickup\" target=\"_blank\">进一步了解送货<br />和取货<span class=\"icon icon-after icon-chevronright\"></span><span class=\"visuallyhidden\">(在新窗口中打开)</span></a>","pickupOptions":[{"pickupOptionTitle":"店内","pickupOptionDescription":"提取你的在线订单商品。你可以获得设置帮助，还能选购配件。可能需要测量体温并佩戴口罩。","index":1}]},"rank":3}],"overlayInitiatedFromWarmStart":false,"viewMoreHoursLinkText":"查看更多时段","storesCount":"3 stores found near 广东 深圳 南山区","little":false,"location":"广东 深圳 南山区","notAvailableNearby":"距离最近的 [X] 家零售店今日无货。","notAvailableNearOneStore":"距离最近的零售店今天不可取货。","warmDudeWithAPU":false,"viewMoreHoursVoText":"(在新窗口中打开)","availability":{"isComingSoon":false},"viewDetailsText":"查看详情"},"deliveryMessage":{"geoLocated":false,"messageType":"compact","geoEnabled":false,"dudeCookieSet":false,"MLE93CH/A":{"label":"预计发货日期：","quote":"7-10 个工作日","address":{},"showDeliveryOptionsLink":false,"messageType":"Ship","basePartNumber":"MLE93","commitCodeId":"7","inHomeSetup":false,"idl":false,"defaultLocationEnabled":false,"isBuyable":true,"isElectronic":false},"processing":"","contentloaded":"","deliveryLocationLabel":"查看到你所在地点的交货日期","locationCookieValueFoundForThisCountry":false,"dudeLocated":false,"accessibilityDeliveryOptions":"送货选项"}}}}

//无货
//{"head":{"status":"200","data":{}},"body":{"content":{"pickupMessage":{"stores":[{"storeEmail":"holidayplazashenzhen@apple.com","storeName":"深圳益田假日广场","reservationUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","makeReservationUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","state":"广东","storeImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R484.png?resize=828:*&output-format=jpg","country":"CN","city":"深圳","storeNumber":"R484","partsAvailability":{"MLDW3CH/A":{"storePickEligible":true,"storeSearchEnabled":true,"storeSelectionEnabled":false,"storePickupQuote":"Apple 深圳益田假日广场目前不可取货","storePickupQuote2_0":"<button type=\"button\" class=\"rf-pickup-quote-overlay-trigger as-retailavailabilitytrigger-infobutton retail-availability-search-trigger as-buttonlink\" data-ase-overlay=\"buac-overlay\" data-ase-click=\"show\">Apple 深圳益田假日广场</button> 目前无货","pickupSearchQuote":"暂无供应","storePickupLabel":"取货：","partNumber":"MLDW3CH/A","purchaseOption":"","ctoOptions":"","storePickupLinkText":"查看供货情况","storePickupProductTitle":"iPhone 13 128GB 粉色","pickupDisplay":"unavailable","pickupType":"店内取货"}},"phoneNumber":"4006171254","pickupTypeAvailabilityText":"此地点提供店内取货服务。","address":{"address":"Apple 深圳益田假日广场","address3":null,"address2":"深圳市南山区深南大道 9028 号益田假日广场","postalCode":"518000"},"hoursUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","storeHours":{"storeHoursText":"Store Hours","bopisPickupDays":"Days","bopisPickupHours":"Hours","hours":[{"storeTimings":"10:00 - 22:30","storeDays":"周五-周六:"},{"storeTimings":"10:00 - 22:00","storeDays":"周一-周四, 周日:"}]},"specialHours":{"specialHoursText":"特别营业时间：","bopisPickupDays":"Days","bopisPickupHours":"Hours","specialHoursData":[{"specialDays":"Sep 25:","specialTimings":"09:00 - 22:30"},{"specialDays":"Sep 26:","specialTimings":"09:00 - 22:00"},{"specialDays":"Sep 30:","specialTimings":"10:00 - 22:30"}],"viewAllSpecialHours":true},"storelatitude":22.53986,"storelongitude":113.97079,"storedistance":2.6,"storeDistanceWithUnit":"2.6 km","storeDistanceVoText":"2.6 km from 518000","retailStore":{"storeNumber":"R484","storeUniqueId":"R484","name":"深圳益田假日广场","storeTypeKey":"1","storeSubTypeKey":"0","storeType":"APPLESTORE_DEFAULT","phoneNumber":"4006171254","email":"holidayplazashenzhen@apple.com","carrierCode":null,"locationType":null,"latitude":22.53986,"longitude":113.97079,"address":{"city":"深圳","companyName":"Apple 深圳益田假日广场","countryCode":"CN","county":null,"district":"深圳","geoCode":null,"label":null,"languageCode":"zh-cn","mailStop":null,"postalCode":"518000","province":null,"state":"广东","street":"深圳市南山区深南大道 9028 号益田假日广场","street2":null,"street3":null,"suburb":null,"type":"SHIPPING","addrSourceType":null,"outsideCityFlag":null,"daytimePhoneAreaCode":null,"eveningPhoneAreaCode":null,"daytimePhone":"4006171254","fullPhoneNumber":null,"eveningPhone":null,"emailAddress":null,"firstName":null,"lastName":null,"suffix":null,"lastNamePhonetic":null,"firstNamePhonetic":null,"title":null,"businessAddress":false,"uuid":"29a0a0cd-273b-4499-a8d2-f764e701b5d3","mobilePhone":null,"mobilePhoneAreaCode":null,"cityStateZip":null,"daytimePhoneSelected":false,"middleName":null,"residenceStatus":null,"moveInDate":null,"subscriberId":null,"locationType":null,"carrierCode":null,"metadata":{},"verificationState":"UN_VERIFIED","expiration":null,"markForDeletion":false,"primaryAddress":false,"fullEveningPhone":null,"fullDaytimePhone":"4006171254","correctionResult":null,"twoLineAddress":"深圳市南山区深南大道 9028 号益田假日广场,\n深圳, 广东, 518000","addressVerified":false},"urlKey":null,"directionsUrl":null,"storeImageUrl":"http://rtlimages.apple.com/cmc/dieter/store/4_3/R484.png?resize=828:*&output-format=jpg","makeReservationUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","hoursAndInfoUrl":"http://www.apple.com.cn/retail/holidayplazashenzhen","storeHours":[{"storeDays":"周五-周六","voStoreDays":null,"storeTimings":"10:00 - 22:30 "},{"storeDays":"周一-周四, 周日","voStoreDays":null,"storeTimings":"10:00 - 22:00 "}],"storeHolidays":[{"date":"Sep 25","description":"Fire Alarm","hours":"09:00 - 22:30 ","comments":"New Product Launch Day Two.","closed":false},{"date":"Sep 26","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":"New Product Launch Day Three.","closed":false},{"date":"Sep 30","description":"Fire Alarm","hours":"10:00 - 22:30 ","comments":"National Holiday","closed":false}],"secureStoreImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R484.png?resize=828:*&output-format=jpg","distance":2.6,"distanceUnit":"km","distanceWithUnit":"2.6 km","timezone":"Asia/Shanghai","storeIsActive":true,"lastUpdated":0.0,"lastFetched":1632533524864,"dateStamp":"24-Sep-2021","distanceSeparator":".","nextAvailableDate":null,"storeHolidayLookAheadWindow":7,"driveDistanceWithUnit":null,"driveDistanceInMeters":null,"dynamicAttributes":{},"storePickupMethodByType":{"INSTORE":{"type":"INSTORE","services":["APU"],"typeCoordinate":{"lat":22.53986,"lon":113.97079},"typeDirection":{"directionByLocale":null},"typeMeetupLocation":{"meetingLocationByLocale":null}}},"storeTimings":null,"availableNow":true},"storelistnumber":1,"storeListNumber":1,"pickupOptionsDetails":{"whatToExpectAtPickup":"<span class=\"as-pickupmethods-intro-header\">取货须知</span><br />当你的订单准备就绪后，我们会向你发送详细的取货说明电子邮件。有关新设备设置，你可以预约免费在线辅导，让 Specialist 专家为你提供指导。","comparePickupOptionsLink":"<a href=\"https://www.apple.com.cn/shop/shipping-pickup\" data-feature-name=\"Astro Link\" data-display-name=\"AOS: shop/shipping-pickup\" target=\"_blank\">进一步了解送货<br />和取货<span class=\"icon icon-after icon-chevronright\"></span><span class=\"visuallyhidden\">(在新窗口中打开)</span></a>","pickupOptions":[{"pickupOptionTitle":"店内","pickupOptionDescription":"提取你的在线订单商品。你可以获得设置帮助，还能选购配件。可能需要测量体温并佩戴口罩。","index":1}]},"rank":1},{"storeEmail":"zhujiangnewtown@apple.com","storeName":"珠江新城","reservationUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","makeReservationUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","state":"广东","storeImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R639.png?resize=828:*&output-format=jpg","country":"CN","city":"广州","storeNumber":"R639","partsAvailability":{"MLDW3CH/A":{"storePickEligible":true,"storeSearchEnabled":true,"storeSelectionEnabled":false,"storePickupQuote":"Apple 珠江新城目前不可取货","storePickupQuote2_0":"<button type=\"button\" class=\"rf-pickup-quote-overlay-trigger as-retailavailabilitytrigger-infobutton retail-availability-search-trigger as-buttonlink\" data-ase-overlay=\"buac-overlay\" data-ase-click=\"show\">Apple 珠江新城</button> 目前无货","pickupSearchQuote":"暂无供应","storePickupLabel":"取货：","partNumber":"MLDW3CH/A","purchaseOption":"","ctoOptions":"","storePickupLinkText":"查看供货情况","storePickupProductTitle":"iPhone 13 128GB 粉色","pickupDisplay":"unavailable","pickupType":"店内取货"}},"phoneNumber":"4006393601","pickupTypeAvailabilityText":"此地点提供店内取货服务。","address":{"address":"Apple 珠江新城","address3":"天汇广场 1 层","address2":"广州珠江新城兴民路 222 号","postalCode":"510623"},"hoursUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","storeHours":{"storeHoursText":"Store Hours","bopisPickupDays":"Days","bopisPickupHours":"Hours","hours":[{"storeTimings":"10:00 - 22:00","storeDays":"周一-周日:"}]},"specialHours":{"specialHoursText":"特别营业时间：","bopisPickupDays":"Days","bopisPickupHours":"Hours","specialHoursData":[{"specialDays":"Sep 25:","specialTimings":"09:00 - 22:00"},{"specialDays":"Sep 26:","specialTimings":"09:00 - 22:00"}],"viewAllSpecialHours":false},"storelatitude":23.11843672960212,"storelongitude":113.32728981971741,"storedistance":89.59,"storeDistanceWithUnit":"89.59 km","storeDistanceVoText":"89.59 km from 510623","retailStore":{"storeNumber":"R639","storeUniqueId":"R639","name":"珠江新城","storeTypeKey":"1","storeSubTypeKey":"0","storeType":"APPLESTORE_DEFAULT","phoneNumber":"4006393601","email":"zhujiangnewtown@apple.com","carrierCode":null,"locationType":null,"latitude":23.11843672960212,"longitude":113.32728981971741,"address":{"city":"广州","companyName":"Apple 珠江新城","countryCode":"CN","county":null,"district":"广州","geoCode":null,"label":null,"languageCode":"zh-cn","mailStop":null,"postalCode":"510623","province":null,"state":"广东","street":"广州珠江新城兴民路 222 号","street2":"天汇广场 1 层","street3":null,"suburb":null,"type":"SHIPPING","addrSourceType":null,"outsideCityFlag":null,"daytimePhoneAreaCode":null,"eveningPhoneAreaCode":null,"daytimePhone":"4006393601","fullPhoneNumber":null,"eveningPhone":null,"emailAddress":null,"firstName":null,"lastName":null,"suffix":null,"lastNamePhonetic":null,"firstNamePhonetic":null,"title":null,"businessAddress":false,"uuid":"3a290981-ca5d-42cd-bfc4-6f7cfbe6b463","mobilePhone":null,"mobilePhoneAreaCode":null,"cityStateZip":null,"daytimePhoneSelected":false,"middleName":null,"residenceStatus":null,"moveInDate":null,"subscriberId":null,"locationType":null,"carrierCode":null,"metadata":{},"verificationState":"UN_VERIFIED","expiration":null,"markForDeletion":false,"primaryAddress":false,"fullEveningPhone":null,"fullDaytimePhone":"4006393601","correctionResult":null,"twoLineAddress":"广州珠江新城兴民路 222 号, 天汇广场 1 层,\n广州, 广东, 510623","addressVerified":false},"urlKey":null,"directionsUrl":null,"storeImageUrl":"http://rtlimages.apple.com/cmc/dieter/store/4_3/R639.png?resize=828:*&output-format=jpg","makeReservationUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","hoursAndInfoUrl":"http://www.apple.com/cn/retail/zhujiangnewtown","storeHours":[{"storeDays":"周一-周日","voStoreDays":null,"storeTimings":"10:00 - 22:00 "}],"storeHolidays":[{"date":"Sep 25","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":"iPhone 13 Launch Day 2","closed":false},{"date":"Sep 26","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":"iPhone 13 Launch Day 3","closed":false}],"secureStoreImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R639.png?resize=828:*&output-format=jpg","distance":89.59,"distanceUnit":"km","distanceWithUnit":"89.59 km","timezone":"Asia/Shanghai","storeIsActive":true,"lastUpdated":0.0,"lastFetched":1632533524864,"dateStamp":"24-Sep-2021","distanceSeparator":".","nextAvailableDate":null,"storeHolidayLookAheadWindow":7,"driveDistanceWithUnit":null,"driveDistanceInMeters":null,"dynamicAttributes":{},"storePickupMethodByType":{"INSTORE":{"type":"INSTORE","services":["APU"],"typeCoordinate":{"lat":23.11843672960212,"lon":113.32728981971741},"typeDirection":{"directionByLocale":null},"typeMeetupLocation":{"meetingLocationByLocale":null}}},"storeTimings":null,"availableNow":true},"storelistnumber":2,"storeListNumber":2,"pickupOptionsDetails":{"whatToExpectAtPickup":"<span class=\"as-pickupmethods-intro-header\">取货须知</span><br />当你的订单准备就绪后，我们会向你发送详细的取货说明电子邮件。有关新设备设置，你可以预约免费在线辅导，让 Specialist 专家为你提供指导。","comparePickupOptionsLink":"<a href=\"https://www.apple.com.cn/shop/shipping-pickup\" data-feature-name=\"Astro Link\" data-display-name=\"AOS: shop/shipping-pickup\" target=\"_blank\">进一步了解送货<br />和取货<span class=\"icon icon-after icon-chevronright\"></span><span class=\"visuallyhidden\">(在新窗口中打开)</span></a>","pickupOptions":[{"pickupOptionTitle":"店内","pickupOptionDescription":"提取你的在线订单商品。你可以获得设置帮助，还能选购配件。可能需要测量体温并佩戴口罩。","index":1}]},"rank":2},{"storeEmail":"parccentral@apple.com","storeName":"天环广场 ","reservationUrl":"http://www.apple.com/cn/retail/parccentral","makeReservationUrl":"http://www.apple.com/cn/retail/parccentral","state":"广东","storeImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R577.png?resize=828:*&output-format=jpg","country":"CN","city":"广州","storeNumber":"R577","partsAvailability":{"MLDW3CH/A":{"storePickEligible":true,"storeSearchEnabled":true,"storeSelectionEnabled":false,"storePickupQuote":"Apple 天环广场目前不可取货","storePickupQuote2_0":"<button type=\"button\" class=\"rf-pickup-quote-overlay-trigger as-retailavailabilitytrigger-infobutton retail-availability-search-trigger as-buttonlink\" data-ase-overlay=\"buac-overlay\" data-ase-click=\"show\">Apple 天环广场</button> 目前无货","pickupSearchQuote":"暂无供应","storePickupLabel":"取货：","partNumber":"MLDW3CH/A","purchaseOption":"","ctoOptions":"","storePickupLinkText":"查看供货情况","storePickupProductTitle":"iPhone 13 128GB 粉色","pickupDisplay":"unavailable","pickupType":"店内取货"}},"phoneNumber":"4006139742","pickupTypeAvailabilityText":"此地点提供店内取货服务。","address":{"address":"Apple 天环广场","address3":null,"address2":"广州市天河区天河路 218 号","postalCode":"510000"},"hoursUrl":"http://www.apple.com/cn/retail/parccentral","storeHours":{"storeHoursText":"Store Hours","bopisPickupDays":"Days","bopisPickupHours":"Hours","hours":[{"storeTimings":"10:00 - 22:00","storeDays":"周一-周日:"}]},"specialHours":{"specialHoursText":"特别营业时间：","bopisPickupDays":"Days","bopisPickupHours":"Hours","specialHoursData":[{"specialDays":"Sep 25:","specialTimings":"09:00 - 22:00"},{"specialDays":"Sep 26:","specialTimings":"09:00 - 22:00"}],"viewAllSpecialHours":false},"storelatitude":23.135,"storelongitude":113.31916,"storedistance":91.47,"storeDistanceWithUnit":"91.47 km","storeDistanceVoText":"91.47 km from 510000","retailStore":{"storeNumber":"R577","storeUniqueId":"R577","name":"天环广场 ","storeTypeKey":"1","storeSubTypeKey":"0","storeType":"APPLESTORE_DEFAULT","phoneNumber":"4006139742","email":"parccentral@apple.com","carrierCode":null,"locationType":null,"latitude":23.135,"longitude":113.31916,"address":{"city":"广州","companyName":"Apple 天环广场","countryCode":"CN","county":null,"district":"广州","geoCode":null,"label":null,"languageCode":"zh-cn","mailStop":null,"postalCode":"510000","province":null,"state":"广东","street":"广州市天河区天河路 218 号","street2":null,"street3":null,"suburb":null,"type":"SHIPPING","addrSourceType":null,"outsideCityFlag":null,"daytimePhoneAreaCode":null,"eveningPhoneAreaCode":null,"daytimePhone":"4006139742","fullPhoneNumber":null,"eveningPhone":null,"emailAddress":null,"firstName":null,"lastName":null,"suffix":null,"lastNamePhonetic":null,"firstNamePhonetic":null,"title":null,"businessAddress":false,"uuid":"b4725efb-202a-4a91-8a36-98c77c267847","mobilePhone":null,"mobilePhoneAreaCode":null,"cityStateZip":null,"daytimePhoneSelected":false,"middleName":null,"residenceStatus":null,"moveInDate":null,"subscriberId":null,"locationType":null,"carrierCode":null,"metadata":{},"verificationState":"UN_VERIFIED","expiration":null,"markForDeletion":false,"primaryAddress":false,"fullEveningPhone":null,"fullDaytimePhone":"4006139742","correctionResult":null,"twoLineAddress":"广州市天河区天河路 218 号,\n广州, 广东, 510000","addressVerified":false},"urlKey":null,"directionsUrl":null,"storeImageUrl":"http://rtlimages.apple.com/cmc/dieter/store/4_3/R577.png?resize=828:*&output-format=jpg","makeReservationUrl":"http://www.apple.com/cn/retail/parccentral","hoursAndInfoUrl":"http://www.apple.com/cn/retail/parccentral","storeHours":[{"storeDays":"周一-周日","voStoreDays":null,"storeTimings":"10:00 - 22:00 "}],"storeHolidays":[{"date":"Sep 25","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":null,"closed":false},{"date":"Sep 26","description":"Fire Alarm","hours":"09:00 - 22:00 ","comments":null,"closed":false}],"secureStoreImageUrl":"https://rtlimages.apple.com/cmc/dieter/store/4_3/R577.png?resize=828:*&output-format=jpg","distance":91.47,"distanceUnit":"km","distanceWithUnit":"91.47 km","timezone":"Asia/Shanghai","storeIsActive":true,"lastUpdated":0.0,"lastFetched":1632533524865,"dateStamp":"24-Sep-2021","distanceSeparator":".","nextAvailableDate":null,"storeHolidayLookAheadWindow":7,"driveDistanceWithUnit":null,"driveDistanceInMeters":null,"dynamicAttributes":{},"storePickupMethodByType":{"INSTORE":{"type":"INSTORE","services":["APU"],"typeCoordinate":{"lat":23.135,"lon":113.31916},"typeDirection":{"directionByLocale":null},"typeMeetupLocation":{"meetingLocationByLocale":null}}},"storeTimings":null,"availableNow":true},"storelistnumber":3,"storeListNumber":3,"pickupOptionsDetails":{"whatToExpectAtPickup":"<span class=\"as-pickupmethods-intro-header\">取货须知</span><br />当你的订单准备就绪后，我们会向你发送详细的取货说明电子邮件。有关新设备设置，你可以预约免费在线辅导，让 Specialist 专家为你提供指导。","comparePickupOptionsLink":"<a href=\"https://www.apple.com.cn/shop/shipping-pickup\" data-feature-name=\"Astro Link\" data-display-name=\"AOS: shop/shipping-pickup\" target=\"_blank\">进一步了解送货<br />和取货<span class=\"icon icon-after icon-chevronright\"></span><span class=\"visuallyhidden\">(在新窗口中打开)</span></a>","pickupOptions":[{"pickupOptionTitle":"店内","pickupOptionDescription":"提取你的在线订单商品。你可以获得设置帮助，还能选购配件。可能需要测量体温并佩戴口罩。","index":1}]},"rank":3}],"overlayInitiatedFromWarmStart":false,"viewMoreHoursLinkText":"查看更多时段","storesCount":"3 stores found near 广东 深圳 南山区","little":false,"location":"广东 深圳 南山区","notAvailableNearby":"距离最近的 [X] 家零售店今日无货。","notAvailableNearOneStore":"距离最近的零售店今天不可取货。","warmDudeWithAPU":false,"viewMoreHoursVoText":"(在新窗口中打开)","availability":{"isComingSoon":false},"viewDetailsText":"查看详情","availabilityStores":"R484,R577,R639"},"deliveryMessage":{"MLDW3CH/A":{"orderByDeliveryBy":"今天订购。","orderByDeliveryBySuffix":"送货至<button class=\"rf-dude-quote-overlay-trigger as-delivery-overlay-trigger as-purchaseinfo-dudetrigger as-buttonlink\" data-autom=\"deliveryDateChecker\" data-ase-overlay=\"dude-overlay\" data-ase-click=\"show\">宝安区†† </button>，预计送达日期：","deliveryOptionMessages":[{"displayName":"2021/10/14 - 2021/10/19 — 免费","inHomeSetup":"false","encodedUpperDateString":"20211019"}],"deliveryOptions":[{"displayName":"标准送货","date":"2021/10/14 - 2021/10/19","shippingCost":"免费","additionalContent":null}],"deliveryOptionsLink":{"text":"宝安区 的可用送货选项†† ","dataVar":{},"newTab":false},"address":{"state":"广东","city":"深圳","district":"宝安区"},"showDeliveryOptionsLink":false,"messageType":"Delivery","basePartNumber":"MLDW3","commitCodeId":"7","dudeAttributes":{"templateID":"DUDE_APU_N","resolvedLabel":"今天订购。","shipMethodsDisplayOrder":["A8"],"fastestShipMethodPriceLabel":null,"leadByPickup":"false","deliveryHeader":null},"inHomeSetup":false,"idl":false,"defaultLocationEnabled":false,"isBuyable":true,"isElectronic":false},"geoLocated":true,"messageType":"regular","geoEnabled":true,"dudeCookieSet":false,"processing":"","contentloaded":"","locationCookieValueFoundForThisCountry":false,"dudeLocated":false,"accessibilityDeliveryOptions":"送货选项"}}}}
