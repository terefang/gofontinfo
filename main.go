package main

import (
	"flag"
	"fmt"
	"log"

	"seehuhn.de/go/sfnt"
)

func main() {
	dumpNames := flag.Bool("names", false, "dump names table")
	dumpMetrics := flag.Bool("metric", false, "dump metrics")
	dumpAttr := flag.Bool("attr", false, "dump attributes")
	fontFile := flag.String("file", "/u/fredo/Fonts/urw/20200910/NimbusSansNarrow-BoldOblique.ttf", "the font file to use")
	flag.Parse()
	font, err := sfnt.ReadFile(*fontFile)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("file = %s\n", *fontFile)
	fmt.Printf("font-family = '%s'\n", font.GetFontInfo().FamilyName)
	fmt.Printf("postscript-name = '%s'\n", font.PostScriptName())
	fmt.Printf("font-fullname = '%s'\n", font.GetFontInfo().FullName)

	if *dumpNames {
		fmt.Printf("font-name = '%s'\n", font.GetFontInfo().FontName)
		fmt.Printf("font-variant = '%s'\n", font.Subfamily())
		fmt.Printf("font-weight = '%s'\n", font.Weight.String())
		fmt.Printf("font-width = '%s'\n", font.Width.String())
	}

	if *dumpAttr {
		fmt.Printf("version = '%s'\n", font.Version.String())
		fmt.Printf("is-bold = %t\n", font.IsBold)
		fmt.Printf("is-cff = %t\n", font.IsCFF())
		fmt.Printf("is-fixed-pitch = %t\n", font.IsFixedPitch())
		fmt.Printf("is-italic = %t\n", font.IsItalic)
		fmt.Printf("is-oblique = %t\n", font.IsOblique)
		fmt.Printf("is-regular = %t\n", font.IsRegular)
		fmt.Printf("is-script = %t\n", font.IsScript)
		fmt.Printf("is-serif = %t\n", font.IsSerif)
	}
	if *dumpMetrics {
		fmt.Printf("num-glyphs = %d\n", font.NumGlyphs())
		fmt.Printf("ascent = %d\n", font.Ascent)
		fmt.Printf("cap-height = %d\n", font.CapHeight)
		fmt.Printf("descent = %d\n", font.Descent)
		fmt.Printf("x-height = %d\n", font.XHeight)
		fmt.Printf("line-gap = %d\n", font.LineGap)
		fmt.Printf("ul-position = %.2f\n", font.UnderlinePosition)
		fmt.Printf("ul-thickness = %.2f\n", font.UnderlineThickness)
		fmt.Printf("units-per-em = %d\n", font.UnitsPerEm)
	}

}
