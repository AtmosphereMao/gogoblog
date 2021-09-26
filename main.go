package main
import (
	"myblog/config"
	"myblog/databases"
)

func init(){
	config.Initialize()
}

func main(){
	databases.InitDB()

}