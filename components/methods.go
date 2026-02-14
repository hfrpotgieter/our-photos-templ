package components

import (
	"fmt"
	"time"
)

func StaticBody() Body {
	return Body{Paragraph: "Static"}
}

func StaticHeader() Header {
	return Header{Title: "Static", Subtitle: "Header"}
}

const (
	anniversaryDay = 10
	valentinesDay  = 14
)

var startDate = time.Date(2022, 8, 10, 0, 0, 0, 0, time.UTC)

func DynamicButton() Button {
	t := time.Now()
	var isValentines, isAnniversary bool
	text := "Not the special day yet!"
	class := btnClassPrimary
	switch t.Month() {
	case time.February:
		text = "Happy Valentine's Day!"
		isValentines = t.Day() == valentinesDay
		class = btnClassRed
	default:
		text = "Happy Anniversary!"
		isAnniversary = t.Day() == anniversaryDay
	}
	return Button{
		Text:     text,
		cssClass: class,
		Href:     "/TODO",
		disabled: !(isValentines || isAnniversary),
	}
}

func (b Button) IsDisabled() string {
	if b.disabled {
		return "true"
	}
	return "false"
}

func (b Button) CSSClass() string {
	if b.disabled {
		b.cssClass = fmt.Sprintf("%s disabled", b.cssClass)
	}
	return b.cssClass
}
