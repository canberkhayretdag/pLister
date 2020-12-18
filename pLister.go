package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var topUsernames = []string{"root", "admin", "test", "guest", "administrator", "asdasd", "default", "master"}
var topUsernamesUpper = []string{"Root", "Admin", "Test", "Guest", "Administrator", "Asdasd", "Default", "Master"}
var topPasswords = []string{"root", "admin", "test", "guest", "administrator",
	"password", "pass", "default", "master"}
var topNumPasswords = []string{"1234", "12345", "123456", "1234567",
	"12345678", "123456789", "0123456789", "1234567890", "asdasd"}

var topSpecialChars = []string{"!", "+", "-", "."}
var allSpecialChars = []string{"!", "+", "-", ".", ",", "/", "#", "$", "&", "(", ")", "{", "}", "*", "?", "<", ">", "]", "["}

var outPass []string
var outUname []string

func createUsernameList(s string, size string) {

	if size != "1" && size != "2" && size != "3" {
		fmt.Println("Please type a valid size...")
		usage()
	}

	// Insert the parameter word
	topUsernames = append(topUsernames, s)

	// Get Top usernames from the list
	for _, i := range topUsernames {
		outUname = append(outUname, i)
	}

	// Insert numbers from 0 to 9
	for _, j := range topUsernames {
		for i := 0; i < 10; i++ {
			outUname = append(outUname, j+strconv.Itoa(i))
		}
	}

	if size == "1" {
		// small size
		// top usernames + specific word + numbers from 0 to 9 (for instance: myWord5) + single digit special character (for instance: myWord!)

		// Insert current year
		for _, un := range topUsernames {
			outUname = append(outUname, un+strconv.Itoa(time.Now().Year()))
		}

	} else if size == "2" {
		// medium size list

		for _, v := range topUsernamesUpper {
			outUname = append(outUname, v)
		}

		// concat special word with top usernames
		for _, j := range topUsernames {
			if j != s {
				outUname = append(outUname, j+"-"+s)
				outUname = append(outUname, j+"_"+s)
				outUname = append(outUname, s+"-"+j)
				outUname = append(outUname, s+"_"+j)
			}
		}

		//Insert Last 5 and next 5 years
		year := time.Now().Year()
		for _, un := range topUsernames {
			for i := year - 5; i < year+6; i++ {
				outUname = append(outUname, un+strconv.Itoa(i))
			}
		}

		for _, un := range topUsernamesUpper {
			for i := year - 5; i < year+6; i++ {
				outUname = append(outUname, un+strconv.Itoa(i))
			}
		}

	} else if size == "3" {
		// large list

		for _, v := range topUsernamesUpper {
			outUname = append(outUname, v)
		}

		// concat special word with top usernames
		for _, j := range topUsernames {
			if j != s {
				outUname = append(outUname, j+"-"+s)
				outUname = append(outUname, j+"_"+s)
				outUname = append(outUname, s+"-"+j)
				outUname = append(outUname, s+"_"+j)
			}
		}

		//Insert Last 10 and next 10 years
		year := time.Now().Year()
		for _, un := range topUsernames {
			for i := year - 10; i < year+11; i++ {
				outUname = append(outUname, un+strconv.Itoa(i))
			}
		}

		for _, un := range topUsernamesUpper {
			for i := year - 10; i < year+11; i++ {
				outUname = append(outUname, un+strconv.Itoa(i))
			}
		}

		// Insert numbers from 10 to 99
		for _, j := range topUsernames {
			for i := 10; i < 100; i++ {
				outUname = append(outUname, j+strconv.Itoa(i))
			}
		}
		for _, j := range topUsernamesUpper {
			for i := 10; i < 100; i++ {
				outUname = append(outUname, j+strconv.Itoa(i))
			}
		}
	}
}

func createPasswordList(s string, size string) {

	if size != "1" && size != "2" && size != "3" {
		fmt.Println("Please type a valid size...")
		usage()
	}

	// Insert the parameter word
	topPasswords = append(topPasswords, s)

	// Get Top passwords from the list
	for _, i := range topPasswords {
		outPass = append(outPass, i)
	}

	// Get Top passwords from the list
	for _, i := range topNumPasswords {
		outPass = append(outPass, i)
	}

	// Insert numbers from 0 to 9
	for _, j := range topPasswords {
		for i := 0; i < 10; i++ {
			outPass = append(outPass, j+strconv.Itoa(i))
		}
	}
	if size == "1" {
		// Insert current year
		for _, un := range topPasswords {
			outPass = append(outPass, un+strconv.Itoa(time.Now().Year()))
		}

	} else if size == "2" {

		for _, v := range topPasswords {
			outPass = append(outPass, strings.Title(v))
		}
		// Insert numbers from 0 to 9
		for _, j := range topPasswords {
			for i := 0; i < 10; i++ {
				outPass = append(outPass, strings.Title(j)+strconv.Itoa(i))
			}
		}

		//Insert Last 5 and next 5 years
		year := time.Now().Year()
		for _, un := range topPasswords {
			for i := year - 5; i < year+6; i++ {
				outPass = append(outPass, un+strconv.Itoa(i))
				for _, sc := range topSpecialChars {
					outPass = append(outPass, un+strconv.Itoa(i)+sc)
				}
			}
		}

		// Insert special chars
		for _, un := range topUsernamesUpper {
			for i := year - 5; i < year+6; i++ {
				outPass = append(outPass, un+strconv.Itoa(i))
				for _, sc := range topSpecialChars {
					outPass = append(outPass, un+strconv.Itoa(i)+sc)
				}
			}
		}
	} else if size == "3" {

		for _, v := range topPasswords {
			topPasswords = append(topPasswords, strings.Title(v))
		}
		// Insert numbers from 0 to 99
		for _, j := range topPasswords {
			for i := 0; i < 100; i++ {
				outPass = append(outPass, strings.Title(j)+strconv.Itoa(i))
			}
		}

		//Insert Last 10 and next 10 years
		year := time.Now().Year()
		for _, un := range topPasswords {
			for i := year - 10; i < year+11; i++ {
				outPass = append(outPass, un+strconv.Itoa(i))
				for _, sc := range topSpecialChars {
					outPass = append(outPass, un+strconv.Itoa(i)+sc)
				}
			}
		}

		// Insert special chars
		for _, v := range outPass {
			for _, i := range allSpecialChars {
				outPass = append(outPass, v+i)
			}
		}
	}
}

func createCustomUsernameList(s string, size string) {
	if size != "1" && size != "2" && size != "3" {
		fmt.Println("Please type a valid size...")
		usage()
	}

	outUname = append(outUname, s)
	outUname = append(outUname, strings.Title(s))

	for _, un := range outUname {
		for i := 0; i < 10; i++ {
			outUname = append(outUname, un+strconv.Itoa(i))
		}
	}

	if size == "1" {
		outUname = append(outUname, s+strconv.Itoa(time.Now().Year()))
		outUname = append(outUname, strings.Title(s)+strconv.Itoa(time.Now().Year()))
	} else if size == "2" {
		// concat special word with top usernames
		for _, j := range topUsernames {
			outUname = append(outUname, s+"-"+j)
			outUname = append(outUname, s+"_"+j)
			outUname = append(outUname, j+"-"+s)
			outUname = append(outUname, j+"_"+s)

			outUname = append(outUname, strings.Title(s)+"-"+j)
			outUname = append(outUname, strings.Title(s)+"_"+j)
			outUname = append(outUname, j+"-"+strings.Title(s))
			outUname = append(outUname, j+"_"+strings.Title(s))
		}

		//Insert Last 5 and next 5 years
		year := time.Now().Year()

		for i := year - 5; i < year+6; i++ {
			outUname = append(outUname, s+strconv.Itoa(i))
			outUname = append(outUname, strings.Title(s)+strconv.Itoa(i))
		}

	} else if size == "3" {
		year := time.Now().Year()
		// concat special word with top usernames and year
		for _, j := range topUsernames {
			outUname = append(outUname, s+"-"+j)
			outUname = append(outUname, s+"-"+j+strconv.Itoa(year))
			outUname = append(outUname, s+"_"+j)
			outUname = append(outUname, s+"_"+j+strconv.Itoa(year))
			outUname = append(outUname, j+"-"+s)
			outUname = append(outUname, j+"-"+s+strconv.Itoa(year))
			outUname = append(outUname, j+"_"+s)
			outUname = append(outUname, j+"_"+s+strconv.Itoa(year))

			outUname = append(outUname, strings.Title(s)+"-"+j)
			outUname = append(outUname, strings.Title(s)+"-"+j+strconv.Itoa(year))
			outUname = append(outUname, strings.Title(s)+"_"+j)
			outUname = append(outUname, strings.Title(s)+"_"+j+strconv.Itoa(year))
			outUname = append(outUname, j+"-"+strings.Title(s))
			outUname = append(outUname, j+"-"+strings.Title(s)+strconv.Itoa(year))
			outUname = append(outUname, j+"_"+strings.Title(s))
			outUname = append(outUname, j+"_"+strings.Title(s)+strconv.Itoa(year))
		}

		//Insert Last 10 and next 10 years
		for i := year - 10; i < year+11; i++ {
			outUname = append(outUname, s+strconv.Itoa(i))
			outUname = append(outUname, strings.Title(s)+strconv.Itoa(i))
		}
	}

}

func checkArgs(word string, usernameList string, passwordList string) {
	if word == "-" {
		fmt.Println("Please type a word about target...")
		fmt.Println("----------------------------------")
		usage()
	}

	if usernameList == "no" && passwordList == "no" {
		fmt.Println("Can not be type no for both lists...")
		fmt.Println("----------------------------------")
		usage()
	}

	if usernameList != "no" && usernameList != "yes" {
		fmt.Println("List options must be yes or no...")
		fmt.Println("----------------------------------")
		usage()
	}

	if passwordList != "no" && passwordList != "yes" {
		fmt.Println("List options must be yes or no...")
		fmt.Println("----------------------------------")
		usage()
	}

}

func usage() {
	fmt.Println("C.H. pLister | version: 1.0")
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("Usage of pLister parameters\n")
	fmt.Println("-text       Specific word of your target // for create username list (it's necessary!)")
	fmt.Println("-username   yes or no // for create username list (default yes)")
	fmt.Println("-password   yes or no // for create username list (default yes)")
	fmt.Println("-usize      1,2,3 // size of username list")
	fmt.Println("-psize      1,2,3 // size of password list")
	fmt.Println("-custom     yes or no // create username list with using custom word only (default yes)")
	fmt.Println()
	fmt.Println("Examples\n")
	fmt.Println("pLister -text=mytarget -usize=2 -psize=3")
	fmt.Println("pLister -text=mytarget -username=no")
	fmt.Println("pLister -text=mytarget -custom=no")
	fmt.Println()
	os.Exit(1)
}

func main() {

	usernameList := flag.String("username", "yes", "Create username list")
	usernameListSize := flag.String("usize", "1", "Size of username list")
	passwordList := flag.String("password", "yes", "Create password list")
	passwordListSize := flag.String("psize", "1", "Size of password list")
	word := flag.String("text", "-", "Specific text of your target")
	useOnlyWord := flag.String("custom", "yes", "Create list with using custom word only")
	flag.Parse()

	checkArgs(*word, *usernameList, *passwordList)

	if *usernameList == "yes" {

		_, err1 := os.Create("usernames.txt")

		if err1 != nil {
			log.Fatalf("failed to create usernames file")
			os.Exit(1)
		}

		file, err2 := os.OpenFile("usernames.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err2 != nil {
			log.Fatalf("failed to open usernames file")
			os.Exit(1)
		}

		if *useOnlyWord == "yes" {
			createCustomUsernameList(*word, *usernameListSize)
		} else {
			createUsernameList(*word, *usernameListSize)
		}

		for _, v := range outUname {
			_, err := file.WriteString(v + "\n")

			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
				os.Exit(1)
			}
		}

		file.Close()
	}

	if *passwordList == "yes" {
		_, err1 := os.Create("passwords.txt")

		if err1 != nil {
			log.Fatalf("failed to create passwords file")
			os.Exit(1)
		}

		file, err2 := os.OpenFile("passwords.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err2 != nil {
			log.Fatalf("failed to open passwords file")
			os.Exit(1)
		}

		createPasswordList(*word, *passwordListSize)

		for _, v := range outPass {
			_, err := file.WriteString(v + "\n")

			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
				os.Exit(1)
			}
		}
	}
}
