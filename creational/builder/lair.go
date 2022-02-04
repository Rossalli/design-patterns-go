package main

type lair struct {
	dragonBuilder iDragonBuilder
}

func newLair(db iDragonBuilder) *lair {
	return &lair{
		dragonBuilder: db,
	}
}

func (l *lair) setDragonBuilder(db iDragonBuilder) {
	l.dragonBuilder = db
}

func (l *lair) buildDragon() dragon {
	l.dragonBuilder.setColor()
	l.dragonBuilder.setBreath()
	return l.dragonBuilder.getDragon()
}
