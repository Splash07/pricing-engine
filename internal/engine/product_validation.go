package engine

type ProductSpecification interface {
	IsValid(product Product) bool
}

type AndSpecification struct {
	specifications []ProductSpecification
}

func NewAndSpecification(specifications ...ProductSpecification) ProductSpecification {
	return AndSpecification{
		specifications: specifications,
	}
}

func (s AndSpecification) IsValid(product Product) bool {
	for _, specification := range s.specifications {
		if !specification.IsValid(product) {
			return false
		}
	}
	return true
}

type HasValidMinPrice struct {
	newPrice          Money
	priceType         PriceType
	minimumPercentage float64
}

func newHasValidMinPrice(newPrice Money, priceType PriceType, minimumPercentage float64) ProductSpecification {
	return HasValidMinPrice{
		newPrice, priceType, minimumPercentage,
	}
}

func (h HasValidMinPrice) IsValid(product Product) bool {
	var originalPrice Money
	switch h.priceType {
	case PriceTypeBP:
		originalPrice = product.BlackPrice
	case PriceTypePMD:
		originalPrice = product.PermanentMarkdownPrice
	case PriceTypeTMD:
		originalPrice = product.TemporaryMarkdownPrice
	}

	if originalPrice.LargerThan(h.newPrice) {
		return true
	}

	// logic to compare percentages here

	return false
}
