package smart

import (
	"fmt"
	"strings"
	"os/exec"
	"io/ioutil"
)

func check_type_sys(disk string) string {
	//You should get 1 for hard disks and 0 for a SSD.
	path := "/sys/block/" + disk + "/queue/rotational"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error", err)
		return "666"
	}
	if string(data) != "1" {
		return "1"
	} else {
		return "0"
	}
}

//func get_list_disks() string {
//	path := "/proc/partitions"
//	data, err := ioutil.ReadFile(path)
//
//	if err != nil {
//		fmt.Println("Error\n", err)
//	}
//
//	//Replace remove empty line between header and table
//	var data_str = s.Replace(string(data), "\n\n", "\n", 1)
//
//	var str_array []string
//	var end_index int
//	var full_lenght = s.Count(data_str, "\n")
//	var i = 0
//
//	//Parse big input string to array of strings
//	for i < full_lenght {
//		end_index = s.IndexByte(data_str, '\n')
//		str_array = append(str_array, data_str[0:end_index])
//		data_str = data_str[end_index+1 : len(data_str)]
//		i = i + 1
//	}
//
//	//Remove first item of array with header text
//	str_array = str_array[1:]
//
//	//Create finall array
//	i = 0
//	for i < full_lenght-1 {
//		str_array[i] = str_array[i][s.LastIndex(str_array[i], " ")+1 : len(str_array[i])]
//		i = i + 1
//	}
//
//	//fmt.Println(str_array)
//	str_out_format := strings.Join(str_array, ", ")
//	return str_out_format
//}

func check_smart_enable(disk string) string {
	//We can have a trouble with this sudo on next line?
	path := "/dev/" + disk
	// sudo removed
	grep := "smartctl -i " + path + " | grep -iPo 'SMART support is: (Disabled|Enabled|Unavailable)'"

	check_smart_enable := exec.Command("bash", "-c", grep)
	execution, err := check_smart_enable.CombinedOutput()

	if err != nil {
		fmt.Println("Error\n", err)
	}

	result_smart := fmt.Sprintf("%s", execution)

	switch result_smart {
	case "SMART support is: Disabled\n":
		return "0"
	case "SMART support is: Enabled\n":
		return "1"
	case "SMART support is: Unavailable\n":
		return "-1"
	default:
		return "666"
	}
}

func check_hw_raid() string {
	path := "/proc/modules"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error", err)
		return "666"
	}

	//Check megaraid
	if strings.Contains(string(data), "megaraid_sas") == true {
		return "1"
	}

	//Check adaptek raid
	if strings.Contains(string(data), "aacraid") == true {
		return "1"
	}

	return "0"
}

//func get_smart_attr(disk string) string {
//
//	path := "/dev/" + disk
//	//sudo removed
//	grep := "smartctl -A " + path
//	get_attr := exec.Command("bash", "-c", grep)
//	data, err := get_attr.CombinedOutput()
//
//	if err != nil {
//		fmt.Println("Error", err)
//	}
//
//	//Remove empty line between header and table
//	var data_str = s.Replace(string(data), "\n\n", "\n", 1)
//	data_str = strings.Replace(data_str, "  ", " ", -1)
//
//	var str_array []string
//	var end_index int
//	var full_lenght = s.Count(data_str, "\n")
//	var i = 0
//
//	//Parse big input string to array of strings
//	for i < full_lenght {
//		end_index = s.IndexByte(data_str, '\n')
//		str_array = append(str_array, data_str[0:end_index])
//		data_str = data_str[end_index+1 : len(data_str)]
//		i = i + 1
//	}
//
//	//Remove first 6 strings with header information and recalculate lenght
//	str_array = str_array[6:]
//	full_lenght = len(str_array) - 1
//
//	//Remove spaces in begining every string like "  5" to "5"
//	i = 0
//	for i < full_lenght {
//		if str_array[i][0] == ' ' {
//			str_array[i] = str_array[i][1:]
//		}
//		i = i + 1
//	}
//
//	//Take first values (ID) before first space + last values (RAW_VALUE) after last space
//	i = 0
//	for i < full_lenght {
//		str_array[i] = str_array[i][:s.IndexAny(str_array[i], " ")] + "," + str_array[i][s.LastIndex(str_array[i], " ")+1:len(str_array[i])]
//		i = i + 1
//	}
//	str_out_text := strings.Join(str_array, ", ")
//	return str_array
//}
