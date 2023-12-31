package main

import "fmt"

const (
	Khan  string = "050000"
	State        = "340000"
	Glmt         = "150000"
	Tdbm         = "040000"
)

var (
	KhanBranchCode  = []string{"5000", "5001", "5003", "5004", "5005", "5006", "5007", "5008", "5009", "5010", "5011", "5012", "5013", "5014", "5015", "5016", "5017", "5018", "5019", "5020", "5021", "5022", "5023", "5024", "5025", "5026", "5027", "5028", "5029", "5030", "5031", "5032", "5033", "5034", "5035", "5036", "5037", "5038", "5039", "5040", "5041", "5042", "5043", "5044", "5045", "5046", "5047", "5048", "5049", "5050", "5051", "5052", "5053", "5054", "5055", "5056", "5057", "5058", "5059", "5060", "5061", "5062", "5063", "5064", "5065", "5066", "5067", "5068", "5069", "5070", "5071", "5072", "5073", "5074", "5075", "5076", "5077", "5078", "5079", "5080", "5081", "5082", "5083", "5084", "5085", "5086", "5087", "5088", "5089", "5090", "5091", "5092", "5093", "5094", "5095", "5096", "5097", "5098", "5099", "5100", "5101", "5102", "5103", "5104", "5105", "5106", "5107", "5108", "5109", "5110", "5111", "5112", "5113", "5114", "5115", "5116", "5117", "5118", "5119", "5120", "5121", "5122", "5123", "5124", "5125", "5126", "5127", "5128", "5129", "5130", "5131", "5132", "5133", "5134", "5135", "5136", "5137", "5138", "5139", "5140", "5141", "5142", "5143", "5144", "5145", "5146", "5147", "5148", "5149", "5150", "5151", "5152", "5153", "5154", "5155", "5156", "5170", "5171", "5172", "5173", "5174", "5175", "5176", "5177", "5178", "5179", "5180", "5181", "5182", "5183", "5184", "5185", "5186", "5187", "5188", "5189", "5190", "5191", "5192", "5193", "5194", "5195", "5196", "5197", "5198", "5201", "5202", "5203", "5204", "5205", "5206", "5207", "5208", "5209", "5210", "5211", "5212", "5213", "5214", "5215", "5216", "5217", "5218", "5219", "5225", "5226", "5227", "5228", "5229", "5230", "5231", "5232", "5233", "5234", "5235", "5236", "5237", "5238", "5239", "5240", "5241", "5242", "5243", "5244", "5245", "5246", "5247", "5248", "5249", "5270", "5271", "5272", "5273", "5274", "5275", "5276", "5277", "5278", "5279", "5280", "5281", "5282", "5283", "5284", "5285", "5286", "5315", "5316", "5317", "5318", "5319", "5320", "5321", "5322", "5323", "5324", "5325", "5326", "5327", "5328", "5329", "5330", "5331", "5332", "5333", "5334", "5335", "5336", "5337", "5360", "5361", "5362", "5363", "5364", "5365", "5366", "5367", "5368", "5369", "5370", "5371", "5372", "5373", "5374", "5375", "5376", "5377", "5378", "5379", "5380", "5381", "5405", "5406", "5407", "5408", "5409", "5411", "5412", "5413", "5414", "5415", "5416", "5417", "5418", "5419", "5420", "5421", "5422", "5423", "5450", "5451", "5452", "5453", "5454", "5455", "5456", "5457", "5458", "5459", "5460", "5461", "5462", "5463", "5464", "5465", "5466", "5495", "5496", "5497", "5498", "5499", "5500", "5501", "5502", "5503", "5504", "5505", "5506", "5507", "5508", "5509", "5510", "5511", "5512", "5513", "5514", "5515", "5516", "5517", "5518", "5519", "5520", "5521", "5522", "5540", "5541", "5542", "5543", "5544", "5545", "5546", "5547", "5548", "5549", "5550", "5551", "5552", "5553", "5554", "5555", "5556", "5557", "5558", "5559", "5560", "5561", "5562", "5563", "5564", "5585", "5586", "5587", "5588", "5589", "5590", "5591", "5592", "5593", "5594", "5595", "5596", "5597", "5598", "5599", "5600", "5601", "5602", "5603", "5604", "5605", "5606", "5607", "5608", "5630", "5631", "5632", "5633", "5634", "5635", "5636", "5637", "5638", "5639", "5640", "5641", "5642", "5643", "5644", "5645", "5646", "5647", "5648", "5675", "5676", "5677", "5678", "5679", "5680", "5681", "5682", "5683", "5684", "5685", "5686", "5687", "5688", "5689", "5690", "5691", "5692", "5700", "5720", "5721", "5722", "5723", "5724", "5725", "5726", "5727", "5728", "5729", "5730", "5731", "5732", "5733", "5734", "5735", "5736", "5737", "5738", "5739", "5740", "5741", "5742", "5743", "5744", "5745", "5746", "5747", "5748", "5749", "5750", "5751", "5752", "5753", "5754", "5755", "5756", "5757", "5758", "5759", "5760", "5761", "5770", "5800", "5801", "5802", "5803", "5804", "5805", "5806", "5807", "5808", "5809", "5810", "5811", "5812", "5813", "5814", "5815", "5816", "5817", "5818", "5819", "5820", "5821", "5822", "5845", "5846", "5847", "5848", "5849", "5850", "5851", "5852", "5853", "5854", "5855", "5856", "5857", "5858", "5859", "5860", "5861", "5862", "5863", "5864", "5865", "5866", "5890", "5891", "5892", "5893", "5894", "5895", "5896", "5897", "5898", "5899", "5900", "5901", "5902", "5903", "5904", "5905", "5906", "5907", "5908", "5909", "5910", "5911", "5912", "5913", "5914", "5915", "5916", "5935", "5936", "5937", "5938", "5939", "5940", "5941", "5942", "5943", "5944", "5945", "5946", "5947", "5948", "5949", "5950", "5951", "5952", "5953", "5954", "5955", "5956", "5957", "5958", "5959", "5990", "5991", "5992", "5993", "5994", "5995", "5996", "5997", "5998", "5166", "5429", "5622", "5619", "5301", "5302", "5304", "5431", "5163", "5926", "5447", "5920", "5664", "5303"}
	StateBranchCode = []string{"1723", "1725", "2004", "2002", "2012", "2008", "2019", "2009", "2015", "2014", "2025", "2022", "2021", "2013", "2024", "2003", "2023", "2000", "2016", "2001", "2005", "2017", "2018", "2805", "2100", "2108", "2107", "2109", "2117", "2112", "3416", "2103", "2818", "2116", "2104", "2106", "3420", "2115", "2111", "2113", "2118", "2110", "2105", "2102", "2114", "2101", "2215", "2201", "2214", "2203", "2210", "2211", "2202", "2206", "2213", "2212", "2208", "2207", "2204", "2200", "2209", "2205", "2307", "2305", "2300", "2324", "2303", "2308", "2318", "2321", "2304", "2309", "2315", "2306", "2316", "2317", "2314", "2313", "2322", "2312", "2311", "2301", "2319", "2323", "2302", "2310", "2400", "2402", "2415", "2419", "2404", "2417", "2401", "2414", "2405", "2411", "2413", "2412", "2420", "2421", "2408", "2410", "2407", "2418", "2416", "2403", "2409", "3200", "2521", "2508", "2514", "2505", "3201", "2517", "2520", "2515", "2512", "2518", "2500", "2504", "2506", "2507", "2511", "2509", "2501", "2503", "2502", "2513", "2618", "2603", "2621", "2620", "2615", "2608", "2619", "2600", "2610", "2604", "2609", "2606", "2617", "2613", "2605", "2614", "2611", "2607", "2602", "2601", "2700", "2705", "2702", "2710", "2706", "2704", "2703", "2708", "2701", "2709", "2714", "2707", "2711", "2712", "2713", "2817", "2807", "2824", "2822", "2823", "2810", "1722", "1709", "1818", "1813", "1814", "1812", "1803", "1808", "1816", "1800", "1801", "1820", "1821", "1819", "1817", "1815", "1804", "1810", "1805", "1811", "1807", "1806", "1906", "1900", "1919", "1917", "1916", "1915", "1920", "1912", "1904", "1911", "1902", "1907", "1918", "1905", "1903", "1914", "1909", "1913", "1922", "1908", "1910", "1921", "1901", "2006", "2010", "2011", "2007", "2020", "1060", "1054", "1074", "1027", "1014", "1075", "3489", "3704", "1094", "1036", "3850", "3418", "1040", "1049", "3401", "3412", "1095", "1084", "1009", "3701", "1033", "1001", "1088", "1022", "3830", "3413", "1023", "3431", "3810", "1052", "3700", "1053", "1063", "1031", "1092", "1038", "1057", "3406", "3702", "1062", "3126", "3469", "1078", "1039", "1044", "3432", "1021", "1026", "1029", "1069", "3800", "3465", "1030", "1050", "3400", "1090", "1067", "1024", "1028", "1041", "3840", "1032", "3460", "1051", "3407", "3851", "1098", "1089", "1047", "1045", "1020", "3130", "1081", "3129", "3131", "3125", "1037", "1011", "1066", "1035", "1068", "3703", "3422", "1048", "1043", "3488", "1059", "3451", "3307", "2715", "2715", "3409", "3409", "3307", "3600", "4000", "1104", "1112", "1108", "1107", "1100", "1116", "1111", "1118", "1109", "1110", "1120", "1101", "1115", "1113", "1106", "1114", "1102", "3414", "1103", "1119", "1105", "1200", "1210", "1201", "1205", "1211", "1203", "1216", "1213", "1207", "1217", "1212", "1215", "1206", "1214", "1209", "1202", "1208", "1218", "1300", "1306", "1308", "1305", "1309", "1301", "1311", "3405", "1304", "1303", "3423", "1307", "1302", "1415", "3402", "1400", "1402", "1416", "1414", "1405", "1403", "1409", "1408", "1407", "1412", "1406", "1512", "1519", "1517", "1501", "1529", "1520", "1515", "1513", "1521", "1510", "1411", "1502", "1526", "1401", "1404", "1504", "1500", "1506", "1522", "1523", "1509", "1507", "1525", "1524", "1505", "1514", "1516", "1508", "1603", "1617", "1614", "1610", "1612", "1615", "1600", "1607", "1609", "1613", "1622", "1606", "1604", "1601", "1605", "1621", "1623", "1602", "1620", "1618", "1616", "1608", "1710", "1724", "1703", "1716", "1719", "1717", "1707", "1720", "1702", "1700", "1714", "1718", "1704", "1712", "1715", "1711", "1713", "1721", "1705", "1708", "3424", "3499", "1706", "1701", "3852", "1064", "1099", "3853", "1093", "1611", "2801", "2800", "2806", "2819", "2802", "2803", "2809", "2811", "2813", "2816", "2804", "2808", "2815", "2812", "2820", "2821", "2911", "2905", "2906", "2912", "2913", "2910", "2915", "2916", "2909", "2903", "2904", "2914", "2907", "2902", "2901", "2900", "2908", "3001", "3003", "3002", "3000", "3102", "3109", "3121", "3114", "3111", "3305", "3112", "3101", "3303", "3127", "3118", "3120", "3304", "3104", "3113", "3110", "3306", "3004", "3105", "3108", "3106", "3119", "3100", "3302", "3107", "3122", "3116", "3117", "3128", "3115", "3300", "3301", "3007", "3005", "1080"}
	GlmtBranchCode  = []string{"110", "111", "112", "113", "114", "116", "117", "118", "120", "121", "122", "123", "124", "130", "132", "140", "141", "150", "160", "162", "163", "170", "171", "172", "173", "174", "180", "181", "182", "190", "191", "192", "193", "195", "196", "197", "200", "201", "202", "203", "210", "211", "212", "213", "220", "221", "240", "241", "250", "251", "270", "271", "272", "273", "280", "281", "290", "300", "301", "302", "303", "305", "310", "315", "320", "321", "322", "323", "325", "330", "335", "340", "341", "345", "346", "350", "351", "352", "360", "361", "362", "363", "364", "365", "366", "367", "368", "400", "401", "402", "403", "430", "431", "432", "433", "434", "435", "436", "437", "460", "470", "471", "472", "490", "491", "520", "521", "522", "523", "524", "525", "526", "527", "528", "529", "540", "541", "560", "561", "562", "563", "564", "565", "580", "581", "582", "590", "600", "601", "610", "611", "612", "620", "625"}
	TdbmBranchCode  = []string{"4"}
)

func main() {
	fmt.Println(bankCode("5180676734", "5"))
	fmt.Println(bankCode("5180676734", ""))
	fmt.Println(bankCode("499259176", "4"))
	fmt.Println(bankCode("499259176", ""))
	fmt.Println(bankCode("3105152644", "15"))
	fmt.Println(bankCode("3105152644", ""))
}

func bankCode(accountNumber, bankCode string) string {
	switch bankCode {
	case "5", "50000":
		return Khan
	case "4":
		return Tdbm
	case "15", "150000":
		return Glmt
	case "34", "340000":
		return State
	}

	if contains(StateBranchCode, accountNumber[:4]) && len(accountNumber) == 12 {
		return State
	} else if contains(KhanBranchCode, accountNumber[:4]) && len(accountNumber) == 10 {
		return Khan
	} else if contains(GlmtBranchCode, accountNumber[:3]) && len(accountNumber) == 10 {
		return Glmt
	} else if contains(TdbmBranchCode, accountNumber[:1]) && len(accountNumber) == 9 {
		return Tdbm
	} else if accountNumber[:1] == "5" {
		// todo end khan banknii corporate gateway ashiglan shalgaj bolno
		// todo /accounts/{accountNumber}/name
	}

	return ""
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
