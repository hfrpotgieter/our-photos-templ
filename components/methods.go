package components

import (
	"fmt"
	"time"
)

func StaticBody() Body {
	return Body{Paragraph: "Static"}
}

func DynamicHeader() Header {
	if isValentinesDay() {
		return Header{
			Title:    "Happy Valentines Day!❤️",
			Subtitle: "I loooove you!🤗",
		}
	}
	if isAnniversaryDay() {
		return Header{
			Title:    "Happy Anniversary!❤️",
			Subtitle: fmt.Sprintf("We've been together for %.2f years!🥂", calculateYearsTogether()),
		}
	}
	return Header{
		Title:    "It's not the special day yet!",
		Subtitle: fmt.Sprintf("You'll have to wait for it to come around...⏲️"),
	}
}

var startDate = time.Date(2022, 8, 10, 0, 0, 0, 0, time.UTC)

const (
	anniversaryDay = 10
	valentinesDay  = 14
)

func calculateDaysTogether() float64 {
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	return t.Sub(startDate).Hours() / 24
}

func calculateYearsTogether() float64 {
	return calculateDaysTogether() / 365
}

func calculateMonthsTogether() float64 {
	return calculateDaysTogether() / 30
}

func isValentinesDay() bool {
	t := time.Now()
	return t.Month() == time.February && t.Day() == valentinesDay
}

func isAnniversaryDay() bool {
	t := time.Now()
	return t.Month() == time.August && t.Day() == anniversaryDay
}

func DynamicButton() Button {
	t := time.Now()
	var isValentines, isAnniversary bool
	text := "Not the special day yet!"
	class := btnClassPrimary

	isValentines = isValentinesDay()
	if isValentines {
		text = "Happy Valentine's Day!"
		isValentines = t.Day() == valentinesDay
		class = btnClassRed
	}
	isAnniversary = isAnniversaryDay()
	if isAnniversary {
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
	// TODO (hendrik): set to false once new route has been added.
	return "true"
}

func (b Button) CSSClass() string {
	if b.disabled {
		b.cssClass = fmt.Sprintf("%s disabled", b.cssClass)
	}
	// TODO (hendrik): remove once new route has been added.
	return fmt.Sprintf("%s disabled", b.cssClass)
}
