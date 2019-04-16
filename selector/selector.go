package selector

import (
	"encoding/csv"
	"github.com/annakozyreva1/banner_show/log"
	"os"
	"sort"
	"strconv"
)

var logger = log.Logger

type banner struct {
	URL     string
	counter int
	Shows   int
}

func (b *banner) Rating() int {
	return b.counter * 100 / b.Shows
}

func (b *banner) Showed() {
	b.counter++
}

type Selector struct {
	store map[string][]*banner
}

func (s *Selector) addCategories(categories []string, b *banner) {
	for _, category := range categories {
		s.store[category] = append(s.store[category], b)
	}
}

func (s *Selector) getAllBanners() []*banner {
	var banners []*banner
	for _, ban := range s.store {
		banners = append(banners, ban...)
	}
	return banners
}

func (s *Selector) GetBanner(categories []string) (string, bool) {
	var banners []*banner
	if len(categories) == 0 {
		banners = s.getAllBanners()
	} else {
		for _, category := range categories {
			banners = append(banners, s.store[category]...)
		}
	}
	if len(banners) == 0 {
		return "", false
	}
	sort.Slice(banners, func(i, j int) bool { return banners[i].Rating() < banners[i].Rating() })
	b := banners[0]
	b.Showed()
	return b.URL, true
}

func InitSelector(config string) *Selector {
	f, err := os.Open(config)
	if err != nil {
		logger.Fatalf("failed to open config: %v", err)
	}
	defer f.Close()
	bannersDesc, err := csv.NewReader(f).ReadAll()
	if err != nil {
		logger.Fatalf("failed to read config: %v", err)
	}
	selector := &Selector{
		store: make(map[string] []*banner),
	}
	for _, bannerDesc := range bannersDesc {
		if len(bannerDesc) < 3 {
			logger.Errorf("incorrect banner: %v", bannerDesc)
			continue
		}
		shows, err := strconv.Atoi(bannerDesc[1])
		if err != nil {
			logger.Errorf("incorrect banner shows: %v", bannerDesc)
			continue
		}
		b := banner{
			URL:   bannerDesc[0],
			Shows: shows,
		}
		selector.addCategories(bannerDesc[2:], &b)
	}
	return selector
}
