package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Pitasi/openswitch/internal/eshop"
	"github.com/Pitasi/openswitch/internal/httpclient"
	"github.com/sirupsen/logrus"
)

type EuropeanEshop struct {
	countries []string
	log       *logrus.Entry
}

func NewEuropeanEshop(countries []string) *EuropeanEshop {
	return &EuropeanEshop{
		countries: countries,
		log:       logrus.New().WithField("ID", "eshop-eu"),
	}
}

func (es *EuropeanEshop) ID() string {
	return fmt.Sprintf("eshop-eu")
}

func (es *EuropeanEshop) Provide(ctx context.Context) ([]Game, error) {
	es.log.Info("Provider starts")

	games, err := europeFetchGames(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching games: %w", err)
	}

	nsuids := europeNsuids(games)

	// map id->games
	m := make(map[string]*Game)
	for _, eg := range games {
		g := es.adaptGame(eg)
		m[g.ProviderGameID] = g
	}

	// fetch price for country
	for i, country := range es.countries {
		es.log.Infof("Fetch for %s (%d/%d)", country, i+1, len(es.countries))

		prices, err := eshop.Prices(ctx, country, nsuids)
		if err != nil {
			return nil, fmt.Errorf("fetching prices for %s: %w", country, err)
		}

		// add country offer to games
		for id, g := range m {
			p, found := prices[id]
			if !found {
				continue
			}
			offer, err := es.adaptPrice(country, p)
			if err != nil {
				es.log.WithError(err).Warnf("invalid offer")
				continue
			}
			if offer == nil {
				continue
			}

			g.AddOffer(offer)
		}
	}

	// m to slice
	gamesWithOffers := make([]Game, 0, len(m))
	for _, g := range m {
		gamesWithOffers = append(gamesWithOffers, *g)
	}

	es.log.Infof("Done. Fetched %d games.", len(gamesWithOffers))

	return gamesWithOffers, nil
}

func (es *EuropeanEshop) adaptGame(g EuropeanGame) *Game {
	id, err := g.NSUID()
	if err != nil {
		id = g.Title
	}
	return &Game{
		ProviderID:     es.ID(),
		ProviderGameID: id,

		Title:       g.Title,
		Description: g.Excerpt,
		ImageURL:    g.ImageURL,
	}
}

func (es *EuropeanEshop) adaptPrice(country string, p *eshop.APIPrice) (*Offer, error) {
	if !p.IsOnSale() {
		return nil, nil
	}

	regular, err := strconv.ParseFloat(p.RegularPrice.RawValue, 32)
	if err != nil {
		return nil, err
	}

	discounted := 0.
	var (
		discountStart *time.Time
		discountEnd   *time.Time
	)
	if p.IsDiscounted() {
		discounted, err = strconv.ParseFloat(p.DiscountPrice.RawValue, 32)
		if err != nil {
			return nil, err
		}

		discountStart = &p.DiscountPrice.StartDatetime
		discountEnd = &p.DiscountPrice.EndDatetime
	}

	// TODO: convert regular and discounted to EUR

	return &Offer{
		ProviderID:    fmt.Sprintf("%s-%s", es.ID(), country),
		BuyLink:       p.BuyLink,
		RegularPrice:  float32(regular),
		DiscountStart: discountStart,
		DiscountEnd:   discountEnd,
		DiscountPrice: float32(discounted),
	}, nil
}

func europeNsuids(games []EuropeanGame) []string {
	nsuids := make([]string, 0, len(games))
	for _, g := range games {
		nsuid, err := g.NSUID()
		if err != nil {
			log.Printf("no nsuid for %s, skipping price fetch\n", g.Title)
		}
		nsuids = append(nsuids, nsuid)
	}
	return nsuids
}

// EuropeanGame is the struct returned by the EU API for each game.
//
// It's the most complete of the three regions.
type EuropeanGame struct {
	FsID                             string      `json:"fs_id"`
	ChangeDate                       time.Time   `json:"change_date"`
	URL                              string      `json:"url"`
	Type                             string      `json:"type"`
	DatesReleasedDts                 []time.Time `json:"dates_released_dts"`
	ClubNintendo                     bool        `json:"club_nintendo"`
	PrettyDateS                      string      `json:"pretty_date_s"`
	PlayModeTvModeB                  bool        `json:"play_mode_tv_mode_b"`
	PlayModeHandheldModeB            bool        `json:"play_mode_handheld_mode_b"`
	ProductCodeTxt                   []string    `json:"product_code_txt"`
	ImageURLSqS                      string      `json:"image_url_sq_s"`
	DeprioritiseB                    bool        `json:"deprioritise_b"`
	PgS                              string      `json:"pg_s"`
	GiftFinderDetailPageImageURLS    string      `json:"gift_finder_detail_page_image_url_s"`
	CompatibleController             []string    `json:"compatible_controller"`
	ImageURL                         string      `json:"image_url"`
	OriginallyForT                   string      `json:"originally_for_t"`
	PaidSubscriptionRequiredB        bool        `json:"paid_subscription_required_b"`
	CloudSavesB                      bool        `json:"cloud_saves_b"`
	DigitalVersionB                  bool        `json:"digital_version_b"`
	TitleExtrasTxt                   []string    `json:"title_extras_txt"`
	ImageURLH2X1S                    string      `json:"image_url_h2x1_s"`
	SystemType                       []string    `json:"system_type"`
	AgeRatingSortingI                int         `json:"age_rating_sorting_i"`
	GameCategoriesTxt                []string    `json:"game_categories_txt"`
	PlayModeTabletopModeB            bool        `json:"play_mode_tabletop_mode_b"`
	Publisher                        string      `json:"publisher"`
	ProductCodeSs                    []string    `json:"product_code_ss"`
	Excerpt                          string      `json:"excerpt"`
	NsuidTxt                         []string    `json:"nsuid_txt"`
	DateFrom                         time.Time   `json:"date_from"`
	LanguageAvailability             []string    `json:"language_availability"`
	PriceHasDiscountB                bool        `json:"price_has_discount_b"`
	PriceDiscountPercentageF         float64     `json:"price_discount_percentage_f"`
	Title                            string      `json:"title"`
	SortingTitle                     string      `json:"sorting_title"`
	CopyrightS                       string      `json:"copyright_s"`
	GiftFinderCarouselImageURLS      string      `json:"gift_finder_carousel_image_url_s"`
	WishlistEmailSquareImageURLS     string      `json:"wishlist_email_square_image_url_s"`
	PlayersTo                        int         `json:"players_to"`
	WishlistEmailBanner640WImageURLS string      `json:"wishlist_email_banner640w_image_url_s"`
	VoiceChatB                       bool        `json:"voice_chat_b"`
	PlayableOnTxt                    []string    `json:"playable_on_txt"`
	HitsI                            int         `json:"hits_i"`
	PrettyGameCategoriesTxt          []string    `json:"pretty_game_categories_txt"`
	GiftFinderWishlistImageURLS      string      `json:"gift_finder_wishlist_image_url_s"`
	SwitchGameVoucherB               bool        `json:"switch_game_voucher_b"`
	GameCategory                     []string    `json:"game_category"`
	SystemNamesTxt                   []string    `json:"system_names_txt"`
	PrettyAgeratingS                 string      `json:"pretty_agerating_s"`
	PriceRegularF                    float64     `json:"price_regular_f"`
	EshopRemovedB                    bool        `json:"eshop_removed_b"`
	PlayersFrom                      int         `json:"players_from"`
	AgeRatingType                    string      `json:"age_rating_type"`
	PriceSortingF                    float64     `json:"price_sorting_f"`
	PriceLowestF                     float64     `json:"price_lowest_f"`
	AgeRatingValue                   string      `json:"age_rating_value"`
	PhysicalVersionB                 bool        `json:"physical_version_b"`
	WishlistEmailBanner460WImageURLS string      `json:"wishlist_email_banner460w_image_url_s"`
	Version                          int         `json:"_version_"`
	Popularity                       int         `json:"popularity"`
}

func (g *EuropeanGame) NSUID() (string, error) {
	if len(g.NsuidTxt) == 0 {
		return "", fmt.Errorf("no NSUIDs for the game")
	}
	return g.NsuidTxt[0], nil
}

func europeFetchGames(ctx context.Context) ([]EuropeanGame, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "search.nintendo-europe.com",
		Path:   "en/select", // en is one of the available locales
		RawQuery: url.Values{
			"rows":  []string{"9999"}, // max number of games returned
			"fq":    []string{"type:GAME AND system_type:nintendoswitch* AND product_code_txt:*"},
			"q":     []string{"*"},
			"sort":  []string{"sorting_title asc"},
			"start": []string{"0"},
			"wt":    []string{"json"},
		}.Encode(),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	res, err := httpclient.New(10*time.Second, 5).Do(req)
	if err != nil {
		return nil, err
	}

	var apiRes struct {
		Response struct {
			Docs []EuropeanGame `json:"docs"`
		} `json:"response"`
	}

	err = json.NewDecoder(res.Body).Decode(&apiRes)
	if err != nil {
		return nil, err
	}

	return apiRes.Response.Docs, nil
}
