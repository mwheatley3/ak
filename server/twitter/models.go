package twitter

// ErrorResponse from twitter
type ErrorResponse struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// AuthResponse is the OAuth response from twitter
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// Tweet is the structure of the tweet that is received from twitter
type Tweet struct {
	Contributors         []Contributor          `json:"contributors"` // Not yet generally available to all, so hard to test
	Coordinates          *Coordinates           `json:"coordinates"`
	CreatedAt            string                 `json:"created_at"`
	Entities             Entities               `json:"entities"`
	ExtendedEntities     Entities               `json:"extended_entities"`
	FavoriteCount        int                    `json:"favorite_count"`
	Favorited            bool                   `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
	ID                   int64                  `json:"id"`
	IDStr                string                 `json:"id_str"`
	InReplyToScreenName  string                 `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64                  `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string                 `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64                  `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string                 `json:"in_reply_to_user_id_str"`
	Lang                 string                 `json:"lang"`
	Place                Place                  `json:"place"`
	QuotedStatusID       int64                  `json:"quoted_status_id"`
	QuotedStatusIDStr    string                 `json:"quoted_status_id_str"`
	QuotedStatus         *Tweet                 `json:"quoted_status"`
	PossiblySensitive    bool                   `json:"possibly_sensitive"`
	RetweetCount         int                    `json:"retweet_count"`
	Retweeted            bool                   `json:"retweeted"`
	RetweetedStatus      *Tweet                 `json:"retweeted_status"`
	Source               string                 `json:"source"`
	Scopes               map[string]interface{} `json:"scopes"`
	Text                 string                 `json:"text"`
	Truncated            bool                   `json:"truncated"`
	User                 User                   `json:"user"`
	WithheldCopyright    bool                   `json:"withheld_copyright"`
	WithheldInCountries  []string               `json:"withheld_in_countries"`
	WithheldScope        string                 `json:"withheld_scope"`

	//Geo is deprecated
	//Geo                  interface{} `json:"geo"`
}

// Contributor struct
type Contributor struct {
	ID         int64  `json:"id"`
	IDStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

// Coordinates model
type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"` // Coordinate always has to have exactly 2 values
	Type        string     `json:"type"`
}

// URLEntity model
type URLEntity struct {
	Urls []struct {
		Indices     []int
		URL         string
		DisplayURL  string
		ExpandedURL string
	}
}

// Entities model
type Entities struct {
	Hashtags []struct {
		Indices []int
		Text    string
	}
	Urls []struct {
		Indices     []int
		URL         string
		DisplayURL  string
		ExpandedURL string
	}
	URL          URLEntity
	UserMentions []struct {
		Name       string
		Indices    []int
		ScreenName string
		ID         int64
		IDStr      string
	}
	Media []EntityMedia
}

// EntityMedia model
type EntityMedia struct {
	ID                int64
	IDStr             string
	MediaURL          string
	MediaURLHTTPS     string
	URL               string
	DisplayURL        string
	ExpandedURL       string
	Sizes             MediaSizes
	SourceStatusID    int64
	SourceStatusIDStr string
	Type              string
	Indices           []int
	VideoInfo         VideoInfo `json:"video_info"`
}

// MediaSizes model
type MediaSizes struct {
	Medium MediaSize
	Thumb  MediaSize
	Small  MediaSize
	Large  MediaSize
}

// MediaSize model
type MediaSize struct {
	W      int
	H      int
	Resize string
}

// VideoInfo model
type VideoInfo struct {
	AspectRatio    []int     `json:"aspect_ratio"`
	DurationMillis int64     `json:"duration_millis"`
	Variants       []Variant `json:"variants"`
}

// Variant model
type Variant struct {
	Bitrate     int    `json:"bitrate"`
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
}

// Place model
type Place struct {
	Attributes  map[string]string `json:"attributes"`
	BoundingBox struct {
		Coordinates [][][]float64 `json:"coordinates"`
		Type        string        `json:"type"`
	} `json:"bounding_box"`
	ContainedWithin []struct {
		Attributes  map[string]string `json:"attributes"`
		BoundingBox struct {
			Coordinates [][][]float64 `json:"coordinates"`
			Type        string        `json:"type"`
		} `json:"bounding_box"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		FullName    string `json:"full_name"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		PlaceType   string `json:"place_type"`
		URL         string `json:"url"`
	} `json:"contained_within"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	FullName    string `json:"full_name"`
	Geometry    struct {
		Coordinates [][][]float64 `json:"coordinates"`
		Type        string        `json:"type"`
	} `json:"geometry"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	PlaceType string   `json:"place_type"`
	Polylines []string `json:"polylines"`
	URL       string   `json:"url"`
}

// User model
type User struct {
	ContributorsEnabled            bool     `json:"contributors_enabled"`
	CreatedAt                      string   `json:"created_at"`
	DefaultProfile                 bool     `json:"default_profile"`
	DefaultProfileImage            bool     `json:"default_profile_image"`
	Description                    string   `json:"description"`
	Entities                       Entities `json:"entities"`
	FavouritesCount                int      `json:"favourites_count"`
	FollowRequestSent              bool     `json:"follow_request_sent"`
	FollowersCount                 int      `json:"followers_count"`
	Following                      bool     `json:"following"`
	FriendsCount                   int      `json:"friends_count"`
	GeoEnabled                     bool     `json:"geo_enabled"`
	ID                             int64    `json:"id"`
	IDStr                          string   `json:"id_str"`
	IsTranslator                   bool     `json:"is_translator"`
	Lang                           string   `json:"lang"` // BCP-47 code of user defined language
	ListedCount                    int64    `json:"listed_count"`
	Location                       string   `json:"location"` // User defined location
	Name                           string   `json:"name"`
	Notifications                  bool     `json:"notifications"`
	ProfileBackgroundColor         string   `json:"profile_background_color"`
	ProfileBackgroundImageURL      string   `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string   `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool     `json:"profile_background_tile"`
	ProfileBannerURL               string   `json:"profile_banner_url"`
	ProfileImageURL                string   `json:"profile_image_url"`
	ProfileImageURLHTTPS           string   `json:"profile_image_url_https"`
	ProfileLinkColor               string   `json:"profile_link_color"`
	ProfileSidebarBorderColor      string   `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string   `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string   `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool     `json:"profile_use_background_image"`
	Protected                      bool     `json:"protected"`
	ScreenName                     string   `json:"screen_name"`
	ShowAllInlineMedia             bool     `json:"show_all_inline_media"`
	Status                         *Tweet   `json:"status"` // Only included if the user is a friend
	StatusesCount                  int64    `json:"statuses_count"`
	TimeZone                       string   `json:"time_zone"`
	URL                            string   `json:"url"` // From UTC in seconds
	UtcOffset                      int      `json:"utc_offset"`
	Verified                       bool     `json:"verified"`
	WithheldInCountries            []string `json:"withheld_in_countries"`
	WithheldScope                  string   `json:"withheld_scope"`
}
