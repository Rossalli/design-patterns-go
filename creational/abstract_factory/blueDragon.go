package main

type blueDragon struct {
}

func (rd *blueDragon) makeYoung() iDragon {
	return &blueYoungDragon{
		dragon: dragon{
			color:    "blue",
			category: "young",
			size:     "medium",
			age:      "6- 100 years old",
		},
	}
}

func (rd *blueDragon) makeAncient() iDragon {
	return &blueAncientDragon{
		dragon: dragon{
			color:    "blue",
			category: "ancient",
			size:     "colossal",
			age:      ">= 801 years old",
		},
	}
}
