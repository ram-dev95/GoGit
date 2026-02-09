package main

import
(
	"fmt"
	c "github.com/ram-dev95/go-packages-demo/config"
	"github.com/ram-dev95/go-packages-demo/utils/text"// 
	"github.com/ram-dev95/go-packages-demo/store"
	
  )      //main.go:7:2: no required module provides package github.com/ram-dev95/go-packages-demo/utils/tetx; to add it:
        // go get github.com/ram-dev95/go-packages-demo/utils/tetx - ошибка показывает что в данном модуле нет зависимостей с таким именем.

func main() {
	fmt.Println("Модуль работает")
	c.AppName = "GoGit"
	fmt.Printf("Название приложения: %s\n", c.AppName)
	const version = c.Version
	fmt.Printf("Версия приложения: %d\n", version)
	// config.env = "production"
	// fmt.Printf("Окружение: %s\n", config.env)	
	text.Greeting()
	fmt.Println(text.Hello)
	// fmt.Println(text.hellosuffix) обращение не к глобальной переменной
	Name := "Shadowed"
	fmt.Println(store.Name, Name)

  	fmt.Println(c.AppName)// обращение к переменной из другового пакета с помощью псевдонима удобного для чтения
	fmt.Println(c.Version)
	fmt.Printf("App %s, build=%s | Store=%s | Greeting=\"%s\"\n", c.AppName, c.BuildMode, store.AddressPublic, text.Greeting())


}