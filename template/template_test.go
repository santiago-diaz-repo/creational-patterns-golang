package template

import (
	"math/rand"
	"testing"
)

func TestSocialMedia_Post(t *testing.T) {
	i := rand.Intn(2)
	if i == 1 {
		socialFe := Facebook{}
		social := socialMedia{
			socialMedia: socialFe,
		}
		social.Post("hello")
	} else {
		socialTu := Twitter{}
		social := socialMedia{
			socialMedia: socialTu,
		}
		social.Post("hello")
	}

}
