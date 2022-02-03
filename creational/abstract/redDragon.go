package main

type redDragon struct {
}

func (rd *redDragon) makeYoung() iDragon {
	return &redYoungDragon{
		dragon: dragon{
			color:    "red",
			category: "young",
			size:     "medium",
			age:      "6- 100 years old",
		},
	}
}

func (rd *redDragon) makeAncient() iDragon {
	return &redAncientDragon{
		dragon: dragon{
			color:    "red",
			category: "ancient",
			size:     "colossal",
			age:      ">= 801 years old",
		},
	}
}
