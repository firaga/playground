package main

import "fmt"

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

type Animal struct {
	scientificName string
	species        string
	AnimalCategory
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func (a Animal) String() int {
	fmt.Println(a.AnimalCategory.species)
	//return fmt.Sprintf("%s (category: %s)", a.scientificName, a.AnimalCategory)
	return 1
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

func main() {
	category := AnimalCategory{species: "cat", genus: "big cat"}
	//fmt.Printf("The animal category: %s\n",category)
	animal := Animal{
		scientificName: "ccat",
		species:        "new cat",
		AnimalCategory: category,
	}
	//fmt.Printf("The animal : %s\n", animal)
	//fmt.Printf("The animal : %v\n", animal.String())
	fmt.Println(animal.String())
	//fmt.Println(animal.String())
}
