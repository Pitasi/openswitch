package eshop

import "time"

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

// AmericanGame is the struct returned by the US API for each game.
type AmericanGame struct {
	Type            string        `json:"type"`
	Locale          string        `json:"locale"`
	URL             string        `json:"url"`
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	LastModified    int64         `json:"lastModified"`
	Nsuid           string        `json:"nsuid"`
	Slug            string        `json:"slug"`
	BoxArt          string        `json:"boxArt"`
	Gallery         string        `json:"gallery"`
	Platform        string        `json:"platform"`
	ReleaseDateMask string        `json:"releaseDateMask"`
	Characters      []string      `json:"characters"`
	Categories      []string      `json:"categories"`
	Msrp            int           `json:"msrp"`
	Esrb            string        `json:"esrb"`
	EsrbDescriptors []string      `json:"esrbDescriptors"`
	VirtualConsole  string        `json:"virtualConsole"`
	GeneralFilters  []string      `json:"generalFilters"`
	FilterShops     []interface{} `json:"filterShops"`
	FilterPlayers   []string      `json:"filterPlayers"`
	Publishers      []string      `json:"publishers"`
	Players         string        `json:"players"`
	Featured        bool          `json:"featured"`
	FreeToStart     bool          `json:"freeToStart"`
	PriceRange      string        `json:"priceRange"`
	SalePrice       interface{}   `json:"salePrice"`
	Availability    []string      `json:"availability"`
	ObjectID        string        `json:"objectID"`
	DistinctSeqID   int           `json:"_distinctSeqID"`
	HighlightResult struct {
		Title struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"title"`
		Nsuid struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"nsuid"`
		Publishers []struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"publishers"`
	} `json:"_highlightResult"`
}

// AsianGame is the struct returned by the Japan API.
//
// It's quite difficult to parse since field's values vary a lot.
type AsianGame struct {
	InitialCode      string `json:"InitialCode"`
	TitleName        string `json:"TitleName"`
	MakerName        string `json:"MakerName"`
	MakerKana        string `json:"MakerKana"`
	Price            string `json:"Price"`
	SalesDate        string `json:"SalesDate"`
	SoftType         string `json:"SoftType"`
	PlatformID       int    `json:"PlatformID"`
	DlIconFlg        int    `json:"DlIconFlg"`
	LinkURL          string `json:"LinkURL"`
	ScreenshotImgFlg int    `json:"ScreenshotImgFlg"`
	ScreenshotImgURL string `json:"ScreenshotImgURL"`
}
