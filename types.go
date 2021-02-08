package twitterscraper

import "time"

type (
	// Video type.
	Video struct {
		ID      string
		Preview string
		URL     string
	}

	// Tweet type.
	Tweet struct {
		Hashtags     []string
		HTML         string
		ID           string
		IsQuoted     bool
		IsPin        bool
		IsReply      bool
		IsRetweet    bool
		Likes        int
		PermanentURL string
		Photos       []string
		Replies      int
		Retweets     int
		Text         string
		TimeParsed   time.Time
		Timestamp    int64
		URLs         []string
		UserID       string
		Username     string
		Videos       []Video
	}

	// Result of scrapping.
	Result struct {
		Tweet
		Error error
	}

	// User type.
	User struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		Entities    struct {
			URL struct {
				Urls []struct {
					ExpandedURL string `json:"expanded_url"`
				} `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount      int      `json:"favourites_count"`
		FollowersCount       int      `json:"followers_count"`
		FriendsCount         int      `json:"friends_count"`
		IDStr                string   `json:"id_str"`
		ListedCount          int      `json:"listed_count"`
		Name                 string   `json:"name"`
		Location             string   `json:"location"`
		PinnedTweetIdsStr    []string `json:"pinned_tweet_ids_str"`
		ProfileBannerURL     string   `json:"profile_banner_url"`
		ProfileImageURLHTTPS string   `json:"profile_image_url_https"`
		Protected            bool     `json:"protected"`
		ScreenName           string   `json:"screen_name"`
		StatusesCount        int      `json:"statuses_count"`
		Verified             bool     `json:"verified"`
	}

	// timeline JSON
	timeline struct {
		GlobalObjects struct {
			Tweets map[string]struct {
				ConversationIDStr string `json:"conversation_id_str"`
				CreatedAt         string `json:"created_at"`
				FavoriteCount     int    `json:"favorite_count"`
				FullText          string `json:"full_text"`
				Entities          struct {
					Hashtags []struct {
						Text string `json:"text"`
					} `json:"hashtags"`
					Media []struct {
						MediaURLHttps string `json:"media_url_https"`
						Type          string `json:"type"`
						URL           string `json:"url"`
					} `json:"media"`
					URLs []struct {
						ExpandedURL string `json:"expanded_url"`
						URL         string `json:"url"`
					} `json:"urls"`
				} `json:"entities"`
				ExtendedEntities struct {
					Media []struct {
						IDStr         string `json:"id_str"`
						MediaURLHttps string `json:"media_url_https"`
						Type          string `json:"type"`
						VideoInfo     struct {
							Variants []struct {
								Bitrate int    `json:"bitrate,omitempty"`
								URL     string `json:"url"`
							} `json:"variants"`
						} `json:"video_info"`
					} `json:"media"`
				} `json:"extended_entities"`
				InReplyToStatusIDStr string    `json:"in_reply_to_status_id_str"`
				ReplyCount           int       `json:"reply_count"`
				RetweetCount         int       `json:"retweet_count"`
				RetweetedStatusIDStr string    `json:"retweeted_status_id_str"`
				QuotedStatusIDStr    string    `json:"quoted_status_id_str"`
				Time                 time.Time `json:"time"`
				UserIDStr            string    `json:"user_id_str"`
			} `json:"tweets"`
			Users map[string]struct {
				CreatedAt   string `json:"created_at"`
				Description string `json:"description"`
				Entities    struct {
					URL struct {
						Urls []struct {
							ExpandedURL string `json:"expanded_url"`
						} `json:"urls"`
					} `json:"url"`
				} `json:"entities"`
				FavouritesCount      int      `json:"favourites_count"`
				FollowersCount       int      `json:"followers_count"`
				FriendsCount         int      `json:"friends_count"`
				IDStr                string   `json:"id_str"`
				ListedCount          int      `json:"listed_count"`
				Name                 string   `json:"name"`
				Location             string   `json:"location"`
				PinnedTweetIdsStr    []string `json:"pinned_tweet_ids_str"`
				ProfileBannerURL     string   `json:"profile_banner_url"`
				ProfileImageURLHTTPS string   `json:"profile_image_url_https"`
				Protected            bool     `json:"protected"`
				ScreenName           string   `json:"screen_name"`
				StatusesCount        int      `json:"statuses_count"`
				Verified             bool     `json:"verified"`
			} `json:"users"`
		} `json:"globalObjects"`
		Timeline struct {
			Instructions []struct {
				AddEntries struct {
					Entries []struct {
						Content struct {
							Item struct {
								Content struct {
									Tweet struct {
										ID string `json:"id"`
									} `json:"tweet"`
								} `json:"content"`
							} `json:"item"`
							Operation struct {
								Cursor struct {
									Value      string `json:"value"`
									CursorType string `json:"cursorType"`
								} `json:"cursor"`
							} `json:"operation"`
							TimelineModule struct {
								Items []struct {
									Item struct {
										ClientEventInfo struct {
											Details struct {
												GuideDetails struct {
													TransparentGuideDetails struct {
														TrendMetadata struct {
															TrendName string `json:"trendName"`
														} `json:"trendMetadata"`
													} `json:"transparentGuideDetails"`
												} `json:"guideDetails"`
											} `json:"details"`
										} `json:"clientEventInfo"`
									} `json:"item"`
								} `json:"items"`
							} `json:"timelineModule"`
						} `json:"content,omitempty"`
					} `json:"entries"`
				} `json:"addEntries"`
				PinEntry struct {
					Entry struct {
						Content struct {
							Item struct {
								Content struct {
									Tweet struct {
										ID string `json:"id"`
									} `json:"tweet"`
								} `json:"content"`
							} `json:"item"`
						} `json:"content"`
					} `json:"entry"`
				} `json:"pinEntry,omitempty"`
			} `json:"instructions"`
		} `json:"timeline"`
	}

	fetchFunc func(user string, maxTweetsNbr int, cursor string) ([]*Tweet, string, error)
)

type TwitterGlobal struct {
	GlobalObjects GlobalObjects `json:"globalObjects"`
}

type GlobalObjects struct {
	Users map[string]TwAccount `json:"users"`
}

type TwAccount struct {
	ID                                      int64                  `json:"id"`
	IDStr                                   string                 `json:"id_str"`
	Name                                    string                 `json:"name"`
	ScreenName                              string                 `json:"screen_name"`
	Location                                string                 `json:"location"`
	Description                             string                 `json:"description"`
	URL                                     string                 `json:"url"`
	Entities                                Entities               `json:"entities"`
	Protected                               bool                   `json:"protected"`
	FollowersCount                          int64                  `json:"followers_count"`
	FastFollowersCount                      int64                  `json:"fast_followers_count"`
	NormalFollowersCount                    int64                  `json:"normal_followers_count"`
	FriendsCount                            int64                  `json:"friends_count"`
	ListedCount                             int64                  `json:"listed_count"`
	CreatedAt                               string                 `json:"created_at"`
	FavouritesCount                         int64                  `json:"favourites_count"`
	UTCOffset                               interface{}            `json:"utc_offset"`
	TimeZone                                interface{}            `json:"time_zone"`
	GeoEnabled                              bool                   `json:"geo_enabled"`
	Verified                                bool                   `json:"verified"`
	StatusesCount                           int64                  `json:"statuses_count"`
	MediaCount                              int64                  `json:"media_count"`
	Lang                                    interface{}            `json:"lang"`
	ContributorsEnabled                     bool                   `json:"contributors_enabled"`
	IsTranslator                            bool                   `json:"is_translator"`
	IsTranslationEnabled                    bool                   `json:"is_translation_enabled"`
	ProfileBackgroundColor                  string                 `json:"profile_background_color"`
	ProfileBackgroundImageURL               string                 `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS          string                 `json:"profile_background_image_url_https"`
	ProfileBackgroundTile                   bool                   `json:"profile_background_tile"`
	ProfileImageURL                         string                 `json:"profile_image_url"`
	ProfileImageURLHTTPS                    string                 `json:"profile_image_url_https"`
	ProfileBannerURL                        string                 `json:"profile_banner_url"`
	ProfileImageExtensionsAltText           interface{}            `json:"profile_image_extensions_alt_text"`
	ProfileImageExtensionsMediaAvailability interface{}            `json:"profile_image_extensions_media_availability"`
	ProfileImageExtensionsMediaColor        interface{}            `json:"profile_image_extensions_media_color"`
	ProfileImageExtensions                  ProfileImageExtensions `json:"profile_image_extensions"`
	ProfileLinkColor                        string                 `json:"profile_link_color"`
	ProfileSidebarBorderColor               string                 `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor                 string                 `json:"profile_sidebar_fill_color"`
	ProfileTextColor                        string                 `json:"profile_text_color"`
	ProfileUseBackgroundImage               bool                   `json:"profile_use_background_image"`
	HasExtendedProfile                      bool                   `json:"has_extended_profile"`
	DefaultProfile                          bool                   `json:"default_profile"`
	DefaultProfileImage                     bool                   `json:"default_profile_image"`
	PinnedTweetIDS                          []interface{}          `json:"pinned_tweet_ids"`
	PinnedTweetIDSStr                       []interface{}          `json:"pinned_tweet_ids_str"`
	HasCustomTimelines                      bool                   `json:"has_custom_timelines"`
	CanDm                                   interface{}            `json:"can_dm"`
	Following                               interface{}            `json:"following"`
	FollowRequestSent                       interface{}            `json:"follow_request_sent"`
	Notifications                           interface{}            `json:"notifications"`
	Muting                                  interface{}            `json:"muting"`
	Blocking                                interface{}            `json:"blocking"`
	BlockedBy                               interface{}            `json:"blocked_by"`
	WantRetweets                            interface{}            `json:"want_retweets"`
	AdvertiserAccountType                   string                 `json:"advertiser_account_type"`
	AdvertiserAccountServiceLevels          []interface{}          `json:"advertiser_account_service_levels"`
	ProfileInterstitialType                 string                 `json:"profile_interstitial_type"`
	BusinessProfileState                    string                 `json:"business_profile_state"`
	TranslatorType                          string                 `json:"translator_type"`
	FollowedBy                              interface{}            `json:"followed_by"`
	EXT                                     EXT                    `json:"ext"`
	RequireSomeConsent                      bool                   `json:"require_some_consent"`
}

type EXT struct {
	HighlightedLabel HighlightedLabel `json:"highlightedLabel"`
}

type HighlightedLabel struct {
	R   HighlightedLabelR `json:"r"`
	TTL int64             `json:"ttl"`
}

type HighlightedLabelR struct {
	Ok Ok `json:"ok"`
}

type Ok struct {
}

type Entities struct {
	URL         Description `json:"url"`
	Description Description `json:"description"`
}

type Description struct {
	Urls []URL `json:"urls"`
}

type URL struct {
	URL         string  `json:"url"`
	ExpandedURL string  `json:"expanded_url"`
	DisplayURL  string  `json:"display_url"`
	Indices     []int64 `json:"indices"`
}

type ProfileImageExtensions struct {
	MediaStats MediaStats `json:"mediaStats"`
}

type MediaStats struct {
	R   MediaStatsR `json:"r"`
	TTL int64       `json:"ttl"`
}

type MediaStatsR struct {
	Missing interface{} `json:"missing"`
}
