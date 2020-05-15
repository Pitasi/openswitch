package eshop

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
