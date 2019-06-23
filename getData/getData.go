package getData

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/**
 *
 * @type {[type]}
 */
func GetAllData(url string) {
	getAttribute(url)
}

func getAttribute(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	eles := doc.Find(".c-job_offer-box")

	eles.Each(func(index int, e2 *goquery.Selection) {
		getSingleRecruitAttribute(e2)
	})

}

func getSingleRecruitAttribute(e2 *goquery.Selection) {

	var recruit Recruit
	var companyName string
	var languages []string
	var frameworks []string

	var address string

	companyName = e2.Find(".c-job_offer-recruiter__name").Text()
	saralyStr := e2.Find(".c-job_offer-detail__salary").Text()

	e2.Find(".c-job_offer-detail tr").Each(func(index int, e3 *goquery.Selection) {
		switch index {
		case 1:
			address = e3.Find(".c-job_offer-detail__description").Text()
		case 2: //language
			e3.Find(".c-job_offer-detail__description .lang_tag").Each(func(index int, e4 *goquery.Selection) {
				languages = append(languages, e4.Text())
			})
		case 3: ///frameworks
			e3.Find(".c-job_offer-detail__description .fw_tag").Each(func(index int, e4 *goquery.Selection) {
				frameworks = append(frameworks, e4.Text())
			})
		}
	})
	salaryArr := strings.Split(saralyStr, "〜")
	fromSalary := strings.Replace(salaryArr[0], "万", "", -1)
	fromSalary = strings.Replace(fromSalary, "円", "", -1)
	fromSalary = strings.Replace(fromSalary, ",", "", -1)

	toSalary := strings.Replace(salaryArr[1], "万", "", -1)
	toSalary = strings.Replace(toSalary, "円", "", -1)
	toSalary = strings.Replace(toSalary, ",", "", -1)

	recruit.companyName = companyName

	recruit.fromSalary = fromSalary
	recruit.toSalary = toSalary
	recruit.address = address

	recruit.languages = languages
	recruit.frameworks = frameworks

	fmt.Println(recruit)
	os.Exit(0)
}
