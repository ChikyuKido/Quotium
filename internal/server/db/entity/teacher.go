package entity

type Teacher struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255;not null"`
	Title      string `gorm:"size:255;not null"`
	ShortName  string `gorm:"size:5;not null"`
	QuoteCount int64  `gorm:"column:quote_count"`
}

type ByQuoteCount []Teacher

func (a ByQuoteCount) Len() int           { return len(a) }
func (a ByQuoteCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByQuoteCount) Less(i, j int) bool { return a[i].QuoteCount < a[j].QuoteCount }

type ByName []Teacher

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByShortName []Teacher

func (a ByShortName) Len() int           { return len(a) }
func (a ByShortName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByShortName) Less(i, j int) bool { return a[i].ShortName < a[j].ShortName }
