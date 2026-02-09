package config

import(
	//  "github.com/ram-dev95/go-packages-demo/config"
	 //package github.com/ram-dev95/go-packages-demo
        // imports github.com/ram-dev95/go-packages-demo/config from main.go
        // imports github.com/ram-dev95/go-packages-demo/config from config.go: import cycle not allowed - циклическая ошибка когда пакет ссылаеться сам на себя.
)
var AppName string
var env string
const Version int = 0
const BuildMode = "debug"
