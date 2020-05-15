package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

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

func NSUID(g EuropeanGame) (string, error) {
	if len(g.NsuidTxt) == 0 {
		return "", fmt.Errorf("no NSUIDs for the game")
	}
	return g.NsuidTxt[0], nil
}

// EuropeGames calls Nintendo API to fetch a list of all games for the European
// region.
func EuropeGames() ([]EuropeanGame, error) {
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

	res, err := httpClient.Do(req)
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
