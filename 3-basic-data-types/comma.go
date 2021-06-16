package __basic_data_types

func comma(num string) string {
	if len(num) <= 3 {
		return num
	}

	return comma(num[:len(num)-3]) + "," + comma(num[len(num)-3:])
}