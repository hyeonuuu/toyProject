package M3MO

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	errExist    = errors.New("already exists")
	errNotExist = errors.New("doesn't exist")
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CreateMemo
func CreateMemo(title, content string) error {
	f := "./storage/" + title + ".txt"
	// check if title used
	_, err := ioutil.ReadFile(f)
	if err == nil {
		return errExist
	}

	b := []byte(content)
	err = ioutil.WriteFile(f, b, 0644)
	checkErr(err)

	return nil
}

// ChangeTitle
func ChangeTitle(oldTitle, newTitle string) error {
	err := os.Rename("./storage/"+oldTitle+".txt", "./storage/"+newTitle+".txt")
	if err != nil {
		return err
	}
	return nil
}

// ChangeContent
func ChangeContent(title, newContent string) error {
	f := title + ".txt"
	err := os.Rename("./storage/"+f, "./temp/"+f)
	if err != nil {
		return err
	}

	var a string
	fmt.Println("Are you sure to change? [y/n]")

	fmt.Scan(&a)
	switch a {
	case "y":
		err := os.Remove("./temp/" + f)
		if err != nil {
			return err
		}
	case "n":
		err := os.Rename("./temp/"+f, "./storage/"+f)
		if err != nil {
			return err
		}
	}

	err = CreateMemo(title, newContent)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMemo
func DeleteMemo(title string) error {
	f := title + ".txt"
	_, err := ioutil.ReadFile("./storage/" + f)
	if err == nil {
		err := os.Rename("./storage/"+f, "./temp/"+f)
		if err != nil {
			return err
		}

		var a string
		fmt.Println("Are you sure to change? [y/n]")

		fmt.Scan(&a)
		switch a {
		case "y":
			err := os.Remove("./temp/" + f)
			if err != nil {
				return err
			}
		case "n":
			err := os.Rename("./temp/"+f, "./storage/"+f)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return errNotExist
}

// loadMemo
func LoadMemo(title string) (string, error) {
	f := "./storage/" + title + ".txt"
	content, err := ioutil.ReadFile(f)
	if err == nil {
		return fmt.Sprintf("Title: %s\nContents: %s\n", title, content), nil
	}
	return "", errNotExist
}

// M3MO
func M3MO() {
	files, err := ioutil.ReadDir("./storage")
	checkErr(err)
	// show m3mo's storage list
	fmt.Println("............................................")
	fmt.Print("		Your M3MO\n\n")
	for i, file := range files {
		f := strings.Split(file.Name(), ".")
		fmt.Printf("%d. %s\n", i+1, f[0])
	}
	fmt.Println("............................................")
	// show m3mo's util list
	fmt.Println(`		If you wanna
1. create a memo, press 1
2. load memo, press 2
3. change memo's title, press 3
4. change memo's content, press 4
5. delete memo, press 5
6. close prpgram, press 'Ctrl' + C`)
	fmt.Println("............................................")
	// each util's func
	var u string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		u = scanner.Text()
	}
L1:

	for {
		switch u {
		case "1": // CreateMemo()
			var title, content string
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("write your memo's title")
			if scanner.Scan() {
				title = scanner.Text()
			}
			_, err := ioutil.ReadFile("./storage/" + title + ".txt")
			if err == nil {
				fmt.Println("!!! maybe file already exists !!!")
				goto L1
			}
			fmt.Println("write your memo`s content")
			if scanner.Scan() {
				content = scanner.Text()
			}
			for {
				err := CreateMemo(title, content)
				if err == nil {
					fmt.Println("memo is created")
					break
				}
				fmt.Println("!!! maybe file already exists !!!")
				fmt.Println("write your memo's title")
				if scanner.Scan() {
					title = scanner.Text()
				}
				fmt.Println("write your memo`s content")
				if scanner.Scan() {
					content = scanner.Text()
				}
			}

		case "2": // LoadMemo()
			var title string
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("write memo's title wanna load")
			if scanner.Scan() {
				title = scanner.Text()
			}
			for {
				content, err := LoadMemo(title)
				if err == nil {
					fmt.Println(content)
					break
				}
				fmt.Println("!!! maybe file doesn't exist !!!")
				fmt.Println("write memo's title wanna load")
				if scanner.Scan() {
					title = scanner.Text()
				}
			}

		case "3": // ChangeTitle()
			var oldTitle, newTitle string
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("write memo's title wanna changed it's title")
			if scanner.Scan() {
				oldTitle = scanner.Text()
			}
			_, err := ioutil.ReadFile("./storage/" + oldTitle + ".txt")
			if err != nil {
				fmt.Println("!!! maybe file doesn't exists !!!")
				goto L1
			}
			fmt.Println("write memo's title wanna change")
			if scanner.Scan() {
				newTitle = scanner.Text()
			}
			for {
				err := ChangeTitle(oldTitle, newTitle)
				if err == nil {
					fmt.Println("memo's title is changed")
					break
				}
				fmt.Println("!!! maybe file doesn't exist !!!")
				fmt.Println("write memo's title wanna changed it's title")
				if scanner.Scan() {
					oldTitle = scanner.Text()
				}
				fmt.Println("write memo's title wanna change")
				if scanner.Scan() {
					newTitle = scanner.Text()
				}
			}
		case "4": // ChangeContent()
			var title, newContent string
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("write memo's title wanna changed it's content")
			if scanner.Scan() {
				title = scanner.Text()
			}
			_, err := ioutil.ReadFile("./storage/" + title + ".txt")
			if err != nil {
				fmt.Println("!!! maybe file doesn't exists !!!")
				goto L1
			}
			fmt.Println("write memo's content")
			if scanner.Scan() {
				newContent = scanner.Text()
			}
			for {
				err := ChangeContent(title, newContent)
				if err == nil {
					fmt.Println("memo's content is changed")
					break
				}
				fmt.Println("!!! maybe file doesn't exist !!!")
				fmt.Println("write memo's title wanna changed it's content")
				if scanner.Scan() {
					title = scanner.Text()
				}
				fmt.Println("write memo's content")
				if scanner.Scan() {
					newContent = scanner.Text()
				}
			}

		case "5": // DeleteMemo()
			var title string
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("write memo's title wanna delete")
			if scanner.Scan() {
				title = scanner.Text()
			}
			for {
				err := DeleteMemo(title)
				if err == nil {
					fmt.Println("memo is deleted")
					break
				}
				fmt.Println("!!! maybe file doesn't exist !!!")
				fmt.Println("write memo's title wanna delete")
				if scanner.Scan() {
					title = scanner.Text()
				}
			}

		default: // out of 1~5
			fmt.Println("!!! wrong number, press 1 to 5 !!!")
		}
		if u == "1" || u == "2" || u == "3" || u == "4" || u == "5" {
			break
		}
		if scanner.Scan() {
			u = scanner.Text()
		}
	}
}
