package recommendation

import (
	"fmt"
)

// BookRecommendation : the book recommendation model.
type BookRecommendation struct {
	Recommendation
	ShouldRead bool
	BookID     uint
}

// String : the function to convert the BookRecommendation to string.
func (r BookRecommendation) String() string {
	var recommendation = r.Recommendation
	return fmt.Sprintf(
		"BookRecommendation { %d %t %t %s }",
		r.ID,
		recommendation.ShouldBuy,
		r.ShouldRead,
		recommendation.Comment,
	)
}

// SetShouldBuy : the shouldBuy setter method.
func (r *BookRecommendation) SetShouldBuy(shouldBuy bool) {
	r.ShouldBuy = shouldBuy
}

// SetShouldRead : the shouldRead setter method.
func (r *BookRecommendation) SetShouldRead(shouldRead bool) {
	r.ShouldRead = shouldRead
}
