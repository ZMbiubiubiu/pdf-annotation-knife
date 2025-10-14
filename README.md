# pdf-annotation-knife
Use go-pdfium to make it easier to handle PDF annotations

# Add Attention

## Text Annotations 

TODO 

## Freetext Annotations

```go
var freeTextAnnot = NewFreeTextAnnotation()
	freeTextAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
freeTextAnnot.Width = 2
freeTextAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
freeTextAnnot.SetOpacity(120)
freeTextAnnot.Contents = "Hello, World!"
freeTextAnnot.FontColor = &Color{R: 255, G: 0, B: 0}
freeTextAnnot.GenerateAppearance()
err = freeTextAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="485" height="256" alt="image" src="https://github.com/user-attachments/assets/0c3f00c8-34c4-4448-9ae0-dadd51e204f3" />

## Square Annotations

```go
var squareAnnot = NewSquareAnnotation()
	squareAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
squareAnnot.Width = 10
squareAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
squareAnnot.SetFillColor(Color{R: 0, G: 255, B: 0})
squareAnnot.SetOpacity(120)
squareAnnot.GenerateAppearance()
err = squareAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="538" height="339" alt="image" src="https://github.com/user-attachments/assets/4a831445-2971-41c4-af77-56bb3116603a" />

## Circle Annotations

```go
var circleAnnot = NewCircleAnnotation()
	circleAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
circleAnnot.Width = 2
circleAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
circleAnnot.SetFillColor(Color{R: 0, G: 255, B: 0})
circleAnnot.SetOpacity(120)
circleAnnot.GenerateAppearance()
err = circleAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="433" height="246" alt="image" src="https://github.com/user-attachments/assets/8441f60a-684a-4b0f-8e70-5292a27760eb" />


## Text Markup Annotations 

### Highlight Annotations
```go
var highlightAnnot = NewHighlightAnnotation()
	highlightAnnot.SetRect(Rect{
		Left:   0,
		Top:    400,
		Right:  200,
		Bottom: 0,
	})
highlightAnnot.SetOpacity(120)
highlightAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
// highlightAnnot.SetFillColor(Color{R: 0, G: 255, B: 0})
highlightAnnot.QuadPoints = []QuadPoint{
  {
    LeftTopX:     100,
    LeftTopY:     200,
    RightTopX:    200,
    RightTopY:    200,
    RightBottomX: 200,
    RightBottomY: 100,
    LeftBottomX:  100,
    LeftBottomY:  100,
  },
  {
    LeftTopX:     100,
    LeftTopY:     400,
    RightTopX:    200,
    RightTopY:    400,
    RightBottomX: 200,
    RightBottomY: 300,
    LeftBottomX:  100,
    LeftBottomY:  300,
  },
}
highlightAnnot.GenerateAppearance()
err = highlightAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="385" height="552" alt="image" src="https://github.com/user-attachments/assets/965b8b1c-567f-45eb-a424-25038b473df4" />

### Underline Annotations

TODO

### Strikeout Annotations

TODO

## Ink Annotations 

TODO

## Stamp Annotations

TODO 

# Delete Annotations

TODO
