# pdf-annotation-knife
Use go-pdfium to make it easier to handle PDF annotations. But we use an object-oriented approach to PDF annotations.

# Add Attention

We'll show you how to add annotations to a PDF document.

## Text Annotations 

TODO 

## Freetext Annotations
**A free text annotation (PDF 1.3) displays text directly on the page**. Unlike an ordinary text annotation, a free text annotation has no open or closed state; instead of being displayed in a pop-up window, the text is always visible.

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
<img width="485" height="256" alt="image" src="https://github.com/user-attachments/assets/0c3f00c8-34c4-4448-9ae0-dadd51e204f3" style="border: 5px solid #000000" />

## Square Annotations
Square annotations display a rectangle on the page. 

* Create a simple Square

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

Circle annotations display an ellipse on the page. 

* Create a simple Circle

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

Text markup annotations appear as highlights, underlines, strikeouts, or squiggly underlines in the text of a document.

### Highlight Annotations

* How to highlight text with color in a PDF document

But you need to find the position of the text yourself by setting QuadPoints

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

* How to add underlines to the text of a document

But you need to find the position of the text yourself by setting QuadPoints

```go
var underlineAnnot = NewUnderlineAnnotation()
	underlineAnnot.SetRect(Rect{
		Left:   0,
		Top:    410,
		Right:  210,
		Bottom: 0,
	})
underlineAnnot.SetStrikeColor(Color{R: 0, G: 255, B: 0})
underlineAnnot.SetOpacity(120)
underlineAnnot.QuadPoints = []QuadPoint{
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
underlineAnnot.GenerateAppearance()
err = underlineAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="417" height="492" alt="image" src="https://github.com/user-attachments/assets/a1378502-52af-4de5-8b28-d3b4814786b0" />


### Strikeout Annotations

* How to add strikeouts to the text of a document

But you need to find the position of the text yourself by setting QuadPoints

```go
var strikeoutAnnot = NewStrikeoutAnnotation()
	strikeoutAnnot.SetRect(Rect{
		Left:   100,
		Top:    200,
		Right:  200,
		Bottom: 100,
	})
strikeoutAnnot.Width = 2
strikeoutAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 0})
strikeoutAnnot.SetOpacity(120)
strikeoutAnnot.QuadPoints = []QuadPoint{
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
}
strikeoutAnnot.GenerateAppearance()
err = strikeoutAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="439" height="303" alt="image" src="https://github.com/user-attachments/assets/0a0f4696-5acd-4f10-b99d-8bdf564ab72d" />

## Ink Annotations 

An ink annotation represents a freehand “scribble” composed of one or more disjoint paths. 

* Create an simple Ink annotation

```go
var inkAnnot = NewInkAnnotation()
inkAnnot.Points = [][]Point{
	{
		{
			X: 100,
			Y: 100,
		},
		{
			X: 105,
			Y: 105,
		},
		{
			X: 110,
			Y: 100,
		},
		{
			X: 115,
			Y: 95,
		},
		{
			X: 120,
			Y: 100,
		},
	},
}
inkAnnot.SetRect(Rect{
	Left:   0,
	Bottom: 0,
	Top:    200,
	Right:  200,
})
inkAnnot.Width = 4
inkAnnot.SetStrikeColor(Color{R: 255, G: 0, B: 255})
inkAnnot.SetOpacity(120)
inkAnnot.GenerateAppearance()
err = inkAnnot.AddAnnotationToPage(context.Background(), instance, page)
```

<img width="284" height="195" alt="image" src="https://github.com/user-attachments/assets/b303ad98-dd81-460e-87d8-c69fd2030cc5" />


## Stamp Annotations

### stamp with path

* Create a stamp annotation with custom path by setting `PathObject`

```go
var stampAnnot = NewStampAnnotation()
stampAnnot.SetRect(Rect{
	Left:   0,
	Bottom: 0,
	Top:    200,
	Right:  200,
})

stampAnnot.SetPathObject([][]Point{
	{
		{X: 100, Y: 100},
		{X: 102, Y: 98},
		{X: 104, Y: 102},
		{X: 106, Y: 100},
		{X: 108, Y: 98},
		{X: 110, Y: 100},
		{X: 112, Y: 102},
		{X: 114, Y: 100},
		{X: 116, Y: 98},
		{X: 118, Y: 100},
		{X: 120, Y: 102},
	},
	{
		{X: 50, Y: 100},
		{X: 52, Y: 98},
		{X: 54, Y: 102},
		{X: 56, Y: 100},
		{X: 58, Y: 98},
		{X: 60, Y: 100},
		{X: 62, Y: 102},
		{X: 64, Y: 100},
		{X: 66, Y: 98},
		{X: 68, Y: 100},
		{X: 70, Y: 102},
	},
}, 2, Color{R: 255, G: 0, B: 255}, 120)
err = stampAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="471" height="272" alt="image" src="https://github.com/user-attachments/assets/b41d7362-b285-482a-98e3-7ceae3ad3955" />


### stamp with jpeg

```go
var stampAnnot = NewStampAnnotation()
stampAnnot.SetRect(Rect{
	Left:   100,
	Bottom: 100,
	Top:    146,
	Right:  168,
})

outputFile := "data/simple_stamp_img.pdf"
os.Remove(outputFile)

// set image object
stampAnnot.SetImgObject("jpeg", docRes.Document, "simple.jpeg")

// add annotation to page
err = stampAnnot.AddAnnotationToPage(context.Background(), instance, page)
```
<img width="456" height="305" alt="image" src="https://github.com/user-attachments/assets/6847e039-572f-4291-adec-57af73253e2c" />


### stamp with png

TODO

### stamp with text

* Create a stamp annotation with custom text

TODO


# Delete Annotations

TODO
