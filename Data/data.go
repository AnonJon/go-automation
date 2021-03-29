package data

import "fmt"

var (
	auPrograms = map[string]int{
		"clinical medical assistant":                                       1405,
		"clinical medical assistant with ob/gyn":                           1405,
		"clinical medical assistant with pediatric specialist":             1405,
		"dental assisting":                                                 1825,
		"ekg technician cert program":                                      1250,
		"medical billing and coding":                                       1215,
		"medical billing and coding with medical administrative assistant": 1370,
		"medical billing and coding with medical admin":                    1370,
		"medical billing and coding with medical administration":           1370,
		"medical billing & coding w/ medical administrative assistant certificate program includes cmaa and cpc national certification exams": 1370,
		"pharmacy technician":                             1200,
		"pharmacy technician with medical administration": 1400,
		"pharmacy technician with medical admin online inc. national cert. and clinical ext.": 1400,
		"phlebotomy technician":                 1575,
		"physicians office assistant with ehrm": 1370,
		"veterinary assistant specialist":       1013}

	metPrograms = map[string]float64{
		"accounting professional":                                         2999.25,
		"administrative assistant with quickbooks":                        2999.25,
		"administrative assistant with bookkeeping and quickbooks":        2999.25,
		"associate professional in human resources":                       2999.25,
		"bookeeping with quickbooks":                                      2849.25,
		"business management professional":                                2999.25,
		"childcare specialist":                                            2999.25,
		"child day care management cert program":                          2962.50,
		"computer security technician (comp tia security+ and network +)": 2999.25,
		"drug and alcohol counselor":                                      2962.50,
		"event planning entrepreneur":                                     2962.50,
		"event panning and banquet management":                            2962.50,
		"full stack web developer with mean stack":                        2999.25,
		"full stack web developer with ruby on rails":                     2999.25,
		"front end web developer":                                         2999.25,
		"health & fitness industry professional":                          2962.50,
		"homeland security specialist":                                    2849.25,
		"human resources professional":                                    2999.25,
		"interior decorating and design entrepreneur":                     2962.50,
		"it cyber security professional with comp tia security+":          2999.25,
		"it network professional with comptia network+":                   2999.25,
		"life skills coach":                                               2962.50,
		"massage practitioner program (500 hr)":                           3000,
		"massage practitioner program (620 hr)":                           3000,
		"massage practitioner program (650 hr)":                           3000,
		"massage practitioner program (700 hr)":                           3000,
		"massage practitioner program (750 hr)":                           3000,
		"marketing professional":                                          2849.25,
		"mental health technician specialist cert":                        2962.50,
		"nutrition and fitness professional":                              2962.50,
		"ophthalmic assistant specialist":                                 2962.50,
		"paralegal certificate program":                                   2999.25,
		"patient advocate specialist":                                     2962.50,
		"personal fitness trainer specialist":                             2999.25,
		"photography entrepreneur with adobe certificate":                 2962.50,
		"photography entrepreneur with adobe":                             2962.50,
		"physical therapy aide":                                           2962.50,
		"professional cooking and catering":                               2962.50,
		"project management professional (pmp)":                           2999.25,
		"real estate law professional":                                    2849.25,
		"stress management coach":                                         2962.50,
		"teachers aide":                                                   2999.25,
		"technical writing":                                               1649.25,
		"travel agent specialist":                                         2962.50,
		"veterinary office assistant specialist":                          2962.50,
		"wedding consultant entrepreneur":                                 2962.50,
		"wellness coaching":                                               2437.50,
		"wellness and life skills coaching specialist":                    3000}
)

//CourseAuPrograms is being used in mycaa_to_main.go
func CourseAuPrograms(course string) bool {

	for k := range auPrograms {
		if k == course {
			return true
		}
	}
	return false
}

//SetPricingColumn is being used in mycaa main
func SetPricingColumn(school string) string {
	schoolList := map[string]string{
		"AU ED4": "L",
		"AU":     "M",
		"MET":    "L",
		"WKU":    "Q",
		"LSU":    "S",
		"UTEP":   "T",
		"CLEM":   "U",
		"UWLAX":  "V",
		"CSU":    "Y",
		"TAMU":   "Z",
		"MSU":    "AB",
		"UNH":    "AE",
		"DESU":   "AF",
		"TAMIU":  "AG",
	}
	for k, v := range schoolList {
		if k == school {
			return v
		}
	}
	fmt.Println("School not there")
	return "School not there"
}

//SetPricingAU is being used in mycaa main
func SetPricingAU(course string) int {
	for k, v := range auPrograms {
		if course == k {
			return v
		}
	}

	fmt.Println(
		"\033[1;31mNo class update pricing for au, smart update started... \033[0;0m")
	return -0

}

//SetPricingMET is being used in mycaa main
func SetPricingMET(course string) float64 {
	for k, v := range metPrograms {
		if course == k {
			return v
		}
	}

	fmt.Println(
		"\033[1;31mNo class update pricing for met, smart update started... \033[0;0m")
	return -0.0

}
