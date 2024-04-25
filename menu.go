package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	//itemNo  itemName   itemPrice
	{1, "Ibihaza", 5000},
	{2, "Isombe", 4000},
	{3, "Brochette", 3000},
	{4, "Ugali", 1500},
	{5, "Ifiriti y'igitoki", 4500},
	{7, "Pilau", 4000},
	{8, "Inkoko", 6000},
	{9, "Agatogo", 3000},
	{10, "Brochettes Mix", 6000},
	{11, "Umutsima", 4000},
	{12, "Ibijumba", 2500},
}

func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}
