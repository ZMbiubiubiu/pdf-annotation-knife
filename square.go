// 矩形
package annotation

type SquareAnnotation struct {
	BaseAnnotation
}

func NewSquareAnnotation() *SquareAnnotation {
	
	return &SquareAnnotation{}
}

func (s *SquareAnnotation) GenerateAppearance() error {
	return nil
}
