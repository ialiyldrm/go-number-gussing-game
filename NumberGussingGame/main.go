package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	gameInfo()
	
	number := numRand(1,100)
	isWrong := false
	isWin := false

	//reader := bufio.NewReader(os.Stdin)
	attemps :=0;
	attempsLimit := 5

	for  attemps < attempsLimit  {
		/* fmt.Print("Tahmininizi giriniz: ")

		input,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		guess,err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		} */

		guess := input()

		if guess < 0 || guess > 99 {
			fmt.Println("Lütfen 0-100 arasında bir değer giriniz.")
			if isWrong {
				attemps++
				fmt.Println("Çok fazla hatalı giriş yaptınız. Kalan hak :",attempsLimit - attemps )
			} else {
				isWrong = true
				fmt.Println("Bir daha hatalı girişinizde hakkınızdan düşülecektir.")
			}
			continue
		}

		if guess == number {
			fmt.Println("Tebrikler, doğru tahmin ! Tahmin ettiğiniz sayı : ",number)
			isWin = true
			break
		} else {
			fmt.Println("Yanlış tahmin!")
			if guess > number {
				if (guess - number) < 10 {
					fmt.Println(guess," sayısı, gizli sayıdan büyüktür. Sıcak tahmin!")
				}else {
					fmt.Println(guess," sayısı, gizli sayıdan büyüktür. Soğuk tahmin!")
				}		
			} else {
				if (number - guess) < 10 {
					fmt.Println(guess," sayısı, gizli sayıdan küçüktür. Sıcak tahmin!")
				}else {
					fmt.Println(guess," sayısı, gizli sayıdan küçüktür. Soğuk tahmin!")
				}	
			}
			attemps++
			fmt.Println(attempsLimit - attemps,"hakkınız kaldı!")
		}
	}

	if !isWin {
		fmt.Println("Kaybettiniz!"+"\nGizli sayı :",number)
	}



	
	
}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func gameInfo(){
	fmt.Println("Sayı tahmin oyununa hoşgeldiniz!\n"+
				" 1 ile 100 arasındaki sayıyı bulmaya çalışın!")
	fmt.Println("Tahmininiz gizli sayıya yakın ise 'Sıcak Tahmin',değilse 'Soğuk Tahmin'dir.")
	fmt.Println("Haydi oynayalım!")
}

func input()int{
	fmt.Print("Tahmininizi giriniz: ")
	reader := bufio.NewReader(os.Stdin)

	input,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	guess,err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return guess
}
