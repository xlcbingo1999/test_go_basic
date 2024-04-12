package Construct

import "fmt"

// 工厂模式的核心就是为了生成某个实例的时候, 传入字符串用于让工厂根据对应的情况生成

// 1. 简单工厂模式: 就是整个程序中只有一种工厂，无法实现扩展
// 2. 工厂方法模式: 工厂也是一种接口, 支持重新扩展新的工厂 [依赖倒置原则, 需要依赖于抽象而不是具体类方法]
// 3. 抽象工厂模式: 每个工厂所生成的产品不是同种类型的, 而是具有多种类型的产品 [本程序就是抽象工厂, 每个厂不只是生成Shoe, 还会生成Tshirt等]

// 核心: 4个要素[抽象工厂、具体工厂、抽象产品、具体产品]

// 产品 =================================================
type ShoeBase interface {
	setLogo(logo string)
	getLogo() string
}

type Shoe struct {
	logo string
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

type NikeShoe struct {
	Shoe
}

// end 产品 =================================================

// 工厂 =================================================
type Factory interface {
	makeShoe() ShoeBase
}

type Niki struct{}

func (n *Niki) makeShoe() ShoeBase {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Niki OK",
		},
	}
}

func GetFactory(name string) Factory {
	if name == "Niki" {
		return &Niki{}
	}
	panic("not defined")
}

func RunAbstractFactory() {
	factory := GetFactory("Niki")

	nikiShoe := factory.makeShoe()
	fmt.Println(nikiShoe.getLogo())
}
