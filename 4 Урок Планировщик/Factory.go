package main

import "log"



type Model interface {
	CreateProduct(model string) Guitar 
}


type Guitar interface {
	Use() string 
}

type ConcreteCreater struct {
}


func Search() Model {
	return &ConcreteCreater{}
}


func (p *ConcreteCreater) CreateProduct(model string) Guitar {
	var material Guitar 

	switch model {
	case "Washburn":
		material = &Product{model}
	case "Dean":
		material = &Product{model}
	case "BcReach":
		material = &Product{model}
	case "Fender":
		material = &Product{model}
	case "Stratocaster":
		material = &Product{model}

	case "Yamaha":
		material = &Product{model}

	default:
		log.Fatalln("Model ", model,  " dont have in storage")
	}

	return material
}

type Product struct {
	model string
}


func (p *Product) Use() string {
	return p.model
}


func main(){

}
