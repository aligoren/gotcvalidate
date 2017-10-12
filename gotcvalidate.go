/*
Turkish Identification Number.

Wikipedia: http://en.wikipedia.org/wiki/Turkish_Identification_Number
NVI Official Site: http://www.nvi.gov.tr/English/Mernis_EN,Mernis_En.html?pageindex=1
*/

package gotcvalidate

import (
	"strconv"
)

type TCValidate struct {
	ATINUM                                     uint64
	BTINUM                                     uint64
	TiNum                                      uint64
	C1, C2, C3, C4, C5, C6, C7, C8, C9, Q1, Q2 uint64
}

func (tv TCValidate) TiValidate(number string) bool {
	/* bool returnNumber false first time */
	returnNumber := false
	/* Turkish Identification Number length == 11 ? */
	if len(number) == 11 {
		// local test variable
		//var ATINUM, BTINUM, TiNum uint64
		//var C1,C2,C3,C4,C5,C6,C7,C8,C9,Q1,Q2 uint64

		/* tv == TCValidate struct init. TiNum equal == number
		ParseUint for string convert to Uint
		*/
		n, _ := strconv.Atoi(number)
		invalid := [9]int{11111111110, 22222222220, 33333333330, 44444444440, 55555555550, 66666666660, 7777777770, 88888888880, 99999999990}
		for _, i := range invalid {
			if i == n {
				return returnNumber
			}
		}
		if tv.C1 == 0 {
			return returnNumber
		}

		tv.TiNum, _ = strconv.ParseUint(number, 10, 64)
		tv.ATINUM = tv.TiNum / 100
		tv.BTINUM = tv.TiNum / 100
		tv.C1 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C2 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C3 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C4 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C5 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C6 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C7 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C8 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)
		tv.C9 = (tv.ATINUM % 10)
		tv.ATINUM = (tv.ATINUM / 10)

		tv.Q1 = ((10 - ((((tv.C1 + tv.C3 + tv.C5 + tv.C7 + tv.C9) * 3) +
			(tv.C2 + tv.C4 + tv.C6 + tv.C8)) % 10)) % 10)
		tv.Q2 = ((10 - (((((tv.C2 + tv.C4 + tv.C6 + tv.C8) + tv.Q1) * 3) +
			(tv.C1 + tv.C3 + tv.C5 + tv.C7 + tv.C9)) % 10)) % 10)

		returnNumber = ((tv.BTINUM*100)+(tv.Q1*10)+tv.Q2 == tv.TiNum)
	}

	return returnNumber
}
