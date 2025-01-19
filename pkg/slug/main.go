package slug

import "github.com/gosimple/slug"

func Slugify(s string) string {
	slug.Lowercase = true
	return slug.MakeLang(s, "en")
}
