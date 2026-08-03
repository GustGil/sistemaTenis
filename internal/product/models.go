package product

type Tenis struct {
	ID       int    `bson:"id"`
	Name     string `bson:"name"`
	Price    string `bson:"price"`
	Category string `bson:"category"`
	Url      string `bson:"url"`
	Status   string `bson:"status"`
}

type CrushionTech struct {
	ID       int
	Name     string
	Category []string
}
