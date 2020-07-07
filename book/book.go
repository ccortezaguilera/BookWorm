package book

import (
	"fmt"

	"github.com/ccortezaguilera/BookWorm/recommendation"
	"github.com/jinzhu/gorm"
)

// Book : The book model to store the book information.
type Book struct {
	gorm.Model
	Title          string `gorm:"varchar(100);index;"`
	Isbn           string `gorm:"size:13;index;"`
	PurchaseURL    string `gorm:"type:text"`
	HasRead        bool
	DoesOwn        bool
	Recommendation recommendation.BookRecommendation
}

// String : return the Book struct as a string.
func (b Book) String() string {
	return fmt.Sprintf(
		"Book { %d %s Isbn: %q Url: %s }",
		b.ID,
		b.Title,
		b.Isbn,
		b.PurchaseURL,
	)
}

// HasReadStr : convert the bool into yes or no string.
func (b Book) HasReadStr() string {
	return GetAnswer(b.HasRead)
}

// DoesOwnStr : convert the bool into yes or no string.
func (b Book) DoesOwnStr() string {
	return GetAnswer(b.DoesOwn)
}

// ShouldBuyStr : convert the bool into yes or no string.
func (b Book) ShouldBuyStr() string {
	var recommendation = b.Recommendation
	var rec = recommendation.Recommendation
	return GetAnswer(rec.ShouldBuy)
}

// ShouldReadStr : convert the bool into yes or no string.
func (b Book) ShouldReadStr() string {
	var recommendation = b.Recommendation
	return GetAnswer(recommendation.ShouldRead)
}

// GetAnswer : convert a bool into yes or no string.
func GetAnswer(predicate bool) string {
	if predicate {
		return "yes"
	}
	return "no"
}

// SetDoesOwn : doesOwn setter method.
func (b *Book) SetDoesOwn(doesOwn bool) {
	b.DoesOwn = doesOwn
}

// SetHasRead : hasRead setter method.
func (b *Book) SetHasRead(hasRead bool) {
	b.HasRead = hasRead
}

// SetRecommendation : recommendation setter method.
func (b *Book) SetRecommendation(rec *recommendation.BookRecommendation) {
	b.Recommendation = *rec
}

// GetBookWithID : Query selector to find record with id.
func GetBookWithID(db *gorm.DB, recordID uint) *gorm.DB {
	return db.Where("id = ?", recordID)
}

// GetFirstTwentyRecords : Limit the Query to 20 records.
func GetFirstTwentyRecords(db *gorm.DB) *gorm.DB {
	return db.Limit(20)
}

// ListBooks : A scope function to address queries.
func ListBooks() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(GetFirstTwentyRecords)
	}
}
