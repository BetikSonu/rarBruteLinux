package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/raifpy/Go/raiFile"

	"github.com/raifpy/Go/errHandler"
)

func oto() {
	fmt.Println("Unutma , bu araç tamamen unrar 'ın yapısına bağımlı kalarak hazırlanmıştır .\n\033[4mBu sebepten bruteforce 'ın başarılı olup olmadığını tespit edemiyorum ..\033[0m\n\n")
	fmt.Print("rar dosyası : ")
	var args1 string
	fmt.Scanln(&args1)
	os.Args = append(os.Args, args1)
	fmt.Print("wordlist dosyası : ")
	var args2 string
	fmt.Scanln(&args2)
	os.Args = append(os.Args, args2)

	man()

}

func manUsage() {
	fmt.Println("Kullanım : rarbrute <dosya.rar> <wordlist>")
}

func exe(rar, pass string) {
	exec.Command("unrar", "x", "-p"+pass, rar).Run()
	//fmt.Println("unrar", "x", "p"+pass, rar)
}

func dirOku() []string {
	dir := make([]string, 0)
	dirinfo, err := ioutil.ReadDir(".")
	if errHandler.HandlerBool(err) {
		return dir
	}
	for _, eleman := range dirinfo {
		dir = append(dir, eleman.Name())
	}
	return dir
}

func inList(list []string, key string) bool {
	for _, eleman := range list {
		if key == eleman {
			return true
		}
	}
	return false
}

func man() {
	args := os.Args
	//args[1] == fileName/location | args[2] == wordlist.txt
	if len(args) < 3 {
		manUsage()
		os.Exit(1)
	}
	if _, err := os.Stat(args[1]); os.IsNotExist(err) {
		fmt.Println("\033[31m" + args[1] + "\033[0m adlı rar dosyası bulunamadı !")
		os.Exit(1)
	} else if _, err := os.Stat(args[2]); os.IsNotExist(err) {
		fmt.Println("\033[31m" + args[2] + "\033[0m adlı wordlist dosyası bulunamadı !")
		os.Exit(1)
	}
	wordlist, err := raiFile.ReadFile(args[2])
	errHandler.HandlerExit(err)
	//splitList := strings.Split(wordlist, "\n")
	//fmt.Println(splitList[8], len(splitList[8]))
	baslangic := dirOku()
	for index, eleman := range strings.Split(wordlist, "\n") {
		if eleman != "" {
			fmt.Println(index, " : ", eleman)
			exe(args[1], eleman)
		}
	}
	fmt.Println("Sadece birkaç saniye ...")
	time.Sleep(time.Second * 5)
	son := dirOku()

	for _, eleman := range son {
		if !inList(baslangic, eleman) {
			fmt.Println("\033[31mYeni Dosya Bulundu !\033[0m -- " + eleman)
		}
	}
	//fmt.Println("\n---\n")
}

func checkUnRar() bool {
	rarLocation, err := exec.LookPath("unrar")
	errHandler.HandlerExit(err)
	if rarLocation == "" {
		return false
	}
	return true
}

func cikNeden(neden string) {
	fmt.Println(neden)
	os.Exit(1)
}

func main() {
	if !checkUnRar() {
		cikNeden("\033[31munrar bulunamadı !\033[0m\n\nKullandığınız dağıtıma göre yükleyebilirsiniz ;\n\nsudo apt install unrar\nsudo pacman -S urar\nsudo dnf install unrar")
	}

	if len(os.Args) == 1 {
		oto()
	} else {
		man()
	}

}
