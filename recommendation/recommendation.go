package recommendation

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Recommendation : the model representing a recommendation to buy.
type Recommendation struct {
	gorm.Model
	ShouldBuy bool
	Comment   string `gorm:"type:text"`
}

// String : function to get the string representation of the model.
func (r Recommendation) String() string {
	return fmt.Sprintf(
		"Recommendation { %d %t %s }",
		r.ID,
		r.ShouldBuy,
		r.Comment,
	)
}

// SetShouldBuy : shouldBuy setter method.
func (r *Recommendation) SetShouldBuy(shouldBuy bool) {
	r.ShouldBuy = shouldBuy
}

// SetComment : comment setter method.
func (r *Recommendation) SetComment(comment string) {
	r.Comment = comment
}
