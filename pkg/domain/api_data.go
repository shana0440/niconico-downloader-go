package domain

import (
	"encoding/json"
	"time"
)

type APIData struct {
	Ads      interface{} `json:"ads"`
	Category interface{} `json:"category"`
	Channel  struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
		IsOfficialAnime   bool   `json:"isOfficialAnime"`
		IsDisplayAdBanner bool   `json:"isDisplayAdBanner"`
		Thumbnail         struct {
			URL      string `json:"url"`
			SmallURL string `json:"smallUrl"`
		} `json:"thumbnail"`
		Viewer struct {
			Follow struct {
				IsFollowed     bool   `json:"isFollowed"`
				IsBookmarked   bool   `json:"isBookmarked"`
				Token          string `json:"token"`
				TokenTimestamp int    `json:"tokenTimestamp"`
			} `json:"follow"`
		} `json:"viewer"`
	} `json:"channel"`
	Client struct {
		Nicosid      string `json:"nicosid"`
		WatchID      string `json:"watchId"`
		WatchTrackID string `json:"watchTrackId"`
	} `json:"client"`
	Comment struct {
		Server struct {
			URL string `json:"url"`
		} `json:"server"`
		Keys struct {
			UserKey string `json:"userKey"`
		} `json:"keys"`
		Layers []struct {
			Index         int  `json:"index"`
			IsTranslucent bool `json:"isTranslucent"`
			ThreadIds     []struct {
				ID   int `json:"id"`
				Fork int `json:"fork"`
			} `json:"threadIds"`
		} `json:"layers"`
		Threads []struct {
			ID                      int         `json:"id"`
			Fork                    int         `json:"fork"`
			IsActive                bool        `json:"isActive"`
			IsDefaultPostTarget     bool        `json:"isDefaultPostTarget"`
			IsEasyCommentPostTarget bool        `json:"isEasyCommentPostTarget"`
			IsLeafRequired          bool        `json:"isLeafRequired"`
			IsOwnerThread           bool        `json:"isOwnerThread"`
			IsThreadkeyRequired     bool        `json:"isThreadkeyRequired"`
			Threadkey               interface{} `json:"threadkey"`
			Is184Forced             bool        `json:"is184Forced"`
			HasNicoscript           bool        `json:"hasNicoscript"`
			Label                   string      `json:"label"`
			PostkeyStatus           int         `json:"postkeyStatus"`
			Server                  string      `json:"server"`
		} `json:"threads"`
		Ng struct {
			NgScore struct {
				IsDisabled bool `json:"isDisabled"`
			} `json:"ngScore"`
			Channel []interface{} `json:"channel"`
			Owner   []interface{} `json:"owner"`
			Viewer  struct {
				Revision int           `json:"revision"`
				Count    int           `json:"count"`
				Items    []interface{} `json:"items"`
			} `json:"viewer"`
		} `json:"ng"`
		IsAttentionRequired bool `json:"isAttentionRequired"`
	} `json:"comment"`
	Community   interface{} `json:"community"`
	EasyComment struct {
		Phrases []struct {
			Text    string `json:"text"`
			Nicodic struct {
				Title     string `json:"title"`
				ViewTitle string `json:"viewTitle"`
				Summary   string `json:"summary"`
				Link      string `json:"link"`
			} `json:"nicodic"`
		} `json:"phrases"`
	} `json:"easyComment"`
	External struct {
		Commons struct {
			HasContentTree bool `json:"hasContentTree"`
		} `json:"commons"`
		Ichiba struct {
			IsEnabled bool `json:"isEnabled"`
		} `json:"ichiba"`
	} `json:"external"`
	Genre struct {
		Key        string `json:"key"`
		Label      string `json:"label"`
		IsImmoral  bool   `json:"isImmoral"`
		IsDisabled bool   `json:"isDisabled"`
		IsNotSet   bool   `json:"isNotSet"`
	} `json:"genre"`
	Marquee struct {
		IsDisabled     bool        `json:"isDisabled"`
		TagRelatedLead interface{} `json:"tagRelatedLead"`
	} `json:"marquee"`
	Media struct {
		Delivery struct {
			RecipeID   string      `json:"recipeId"`
			Encryption interface{} `json:"encryption"`
			Movie      struct {
				ContentID string `json:"contentId"`
				Audios    []struct {
					ID          string `json:"id"`
					IsAvailable bool   `json:"isAvailable"`
					Metadata    struct {
						Bitrate      int `json:"bitrate"`
						SamplingRate int `json:"samplingRate"`
						Loudness     struct {
							IntegratedLoudness float64 `json:"integratedLoudness"`
							TruePeak           float64 `json:"truePeak"`
						} `json:"loudness"`
						LevelIndex         int `json:"levelIndex"`
						LoudnessCollection []struct {
							Type  string `json:"type"`
							Value int    `json:"value"`
						} `json:"loudnessCollection"`
					} `json:"metadata"`
				} `json:"audios"`
				Videos []struct {
					ID          string `json:"id"`
					IsAvailable bool   `json:"isAvailable"`
					Metadata    struct {
						Label      string `json:"label"`
						Bitrate    int    `json:"bitrate"`
						Resolution struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"resolution"`
						LevelIndex                        int `json:"levelIndex"`
						RecommendedHighestAudioLevelIndex int `json:"recommendedHighestAudioLevelIndex"`
					} `json:"metadata"`
				} `json:"videos"`
				Session struct {
					RecipeID  string        `json:"recipeId"`
					PlayerID  string        `json:"playerId"`
					Videos    []string      `json:"videos"`
					Audios    []string      `json:"audios"`
					Movies    []interface{} `json:"movies"`
					Protocols []string      `json:"protocols"`
					AuthTypes struct {
						HTTP string `json:"http"`
						Hls  string `json:"hls"`
					} `json:"authTypes"`
					ServiceUserID     string        `json:"serviceUserId"`
					Token             string        `json:"token"`
					Signature         string        `json:"signature"`
					ContentID         string        `json:"contentId"`
					HeartbeatLifetime int           `json:"heartbeatLifetime"`
					ContentKeyTimeout int           `json:"contentKeyTimeout"`
					Priority          float64       `json:"priority"`
					TransferPresets   []interface{} `json:"transferPresets"`
					Urls              []struct {
						URL             string `json:"url"`
						IsWellKnownPort bool   `json:"isWellKnownPort"`
						IsSsl           bool   `json:"isSsl"`
					} `json:"urls"`
				} `json:"session"`
			} `json:"movie"`
			Storyboard interface{} `json:"storyboard"`
			TrackingID string      `json:"trackingId"`
		} `json:"delivery"`
		DeliveryLegacy interface{} `json:"deliveryLegacy"`
	} `json:"media"`
	OkReason string      `json:"okReason"`
	Owner    interface{} `json:"owner"`
	Payment  struct {
		Video struct {
			IsPpv               bool   `json:"isPpv"`
			IsAdmission         bool   `json:"isAdmission"`
			IsPremium           bool   `json:"isPremium"`
			WatchableUserType   string `json:"watchableUserType"`
			CommentableUserType string `json:"commentableUserType"`
		} `json:"video"`
		Preview struct {
			Ppv struct {
				IsEnabled bool `json:"isEnabled"`
			} `json:"ppv"`
			Admission struct {
				IsEnabled bool `json:"isEnabled"`
			} `json:"admission"`
			Premium struct {
				IsEnabled bool `json:"isEnabled"`
			} `json:"premium"`
		} `json:"preview"`
	} `json:"payment"`
	PcWatchPage struct {
		TagRelatedBanner interface{} `json:"tagRelatedBanner"`
		VideoEnd         struct {
			BannerIn interface{} `json:"bannerIn"`
			Overlay  interface{} `json:"overlay"`
		} `json:"videoEnd"`
		ShowOwnerMenu                bool `json:"showOwnerMenu"`
		ShowOwnerThreadCoEditingLink bool `json:"showOwnerThreadCoEditingLink"`
		ShowMymemoryEditingLink      bool `json:"showMymemoryEditingLink"`
	} `json:"pcWatchPage"`
	Player struct {
		InitialPlayback struct {
			Type        string      `json:"type"`
			PositionSec interface{} `json:"positionSec"`
		} `json:"initialPlayback"`
		Comment struct {
			IsDefaultInvisible bool `json:"isDefaultInvisible"`
		} `json:"comment"`
		LayerMode int `json:"layerMode"`
	} `json:"player"`
	Ppv struct {
		AccessFrom interface{} `json:"accessFrom"`
	} `json:"ppv"`
	Ranking struct {
		Genre struct {
			Rank     int       `json:"rank"`
			Genre    string    `json:"genre"`
			DateTime time.Time `json:"dateTime"`
		} `json:"genre"`
		PopularTag []interface{} `json:"popularTag"`
	} `json:"ranking"`
	Series struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		ThumbnailURL string `json:"thumbnailUrl"`
		Video        struct {
			Prev struct {
				Type         string    `json:"type"`
				ID           string    `json:"id"`
				Title        string    `json:"title"`
				RegisteredAt time.Time `json:"registeredAt"`
				Count        struct {
					View    int `json:"view"`
					Comment int `json:"comment"`
					Mylist  int `json:"mylist"`
					Like    int `json:"like"`
				} `json:"count"`
				Thumbnail struct {
					URL        string `json:"url"`
					MiddleURL  string `json:"middleUrl"`
					LargeURL   string `json:"largeUrl"`
					ListingURL string `json:"listingUrl"`
					NHdURL     string `json:"nHdUrl"`
				} `json:"thumbnail"`
				Duration             int    `json:"duration"`
				ShortDescription     string `json:"shortDescription"`
				LatestCommentSummary string `json:"latestCommentSummary"`
				IsChannelVideo       bool   `json:"isChannelVideo"`
				IsPaymentRequired    bool   `json:"isPaymentRequired"`
				PlaybackPosition     int    `json:"playbackPosition"`
				Owner                struct {
					OwnerType string `json:"ownerType"`
					ID        string `json:"id"`
					Name      string `json:"name"`
					IconURL   string `json:"iconUrl"`
				} `json:"owner"`
				RequireSensitiveMasking bool        `json:"requireSensitiveMasking"`
				VideoLive               interface{} `json:"videoLive"`
				NineD091F87             bool        `json:"9d091f87"`
				Acf68865                bool        `json:"acf68865"`
			} `json:"prev"`
			Next struct {
				Type         string    `json:"type"`
				ID           string    `json:"id"`
				Title        string    `json:"title"`
				RegisteredAt time.Time `json:"registeredAt"`
				Count        struct {
					View    int `json:"view"`
					Comment int `json:"comment"`
					Mylist  int `json:"mylist"`
					Like    int `json:"like"`
				} `json:"count"`
				Thumbnail struct {
					URL        string `json:"url"`
					MiddleURL  string `json:"middleUrl"`
					LargeURL   string `json:"largeUrl"`
					ListingURL string `json:"listingUrl"`
					NHdURL     string `json:"nHdUrl"`
				} `json:"thumbnail"`
				Duration             int    `json:"duration"`
				ShortDescription     string `json:"shortDescription"`
				LatestCommentSummary string `json:"latestCommentSummary"`
				IsChannelVideo       bool   `json:"isChannelVideo"`
				IsPaymentRequired    bool   `json:"isPaymentRequired"`
				PlaybackPosition     int    `json:"playbackPosition"`
				Owner                struct {
					OwnerType string `json:"ownerType"`
					ID        string `json:"id"`
					Name      string `json:"name"`
					IconURL   string `json:"iconUrl"`
				} `json:"owner"`
				RequireSensitiveMasking bool        `json:"requireSensitiveMasking"`
				VideoLive               interface{} `json:"videoLive"`
				NineD091F87             bool        `json:"9d091f87"`
				Acf68865                bool        `json:"acf68865"`
			} `json:"next"`
			First struct {
				Type         string    `json:"type"`
				ID           string    `json:"id"`
				Title        string    `json:"title"`
				RegisteredAt time.Time `json:"registeredAt"`
				Count        struct {
					View    int `json:"view"`
					Comment int `json:"comment"`
					Mylist  int `json:"mylist"`
					Like    int `json:"like"`
				} `json:"count"`
				Thumbnail struct {
					URL        string `json:"url"`
					MiddleURL  string `json:"middleUrl"`
					LargeURL   string `json:"largeUrl"`
					ListingURL string `json:"listingUrl"`
					NHdURL     string `json:"nHdUrl"`
				} `json:"thumbnail"`
				Duration             int    `json:"duration"`
				ShortDescription     string `json:"shortDescription"`
				LatestCommentSummary string `json:"latestCommentSummary"`
				IsChannelVideo       bool   `json:"isChannelVideo"`
				IsPaymentRequired    bool   `json:"isPaymentRequired"`
				PlaybackPosition     int    `json:"playbackPosition"`
				Owner                struct {
					OwnerType string `json:"ownerType"`
					ID        string `json:"id"`
					Name      string `json:"name"`
					IconURL   string `json:"iconUrl"`
				} `json:"owner"`
				RequireSensitiveMasking bool        `json:"requireSensitiveMasking"`
				VideoLive               interface{} `json:"videoLive"`
				NineD091F87             bool        `json:"9d091f87"`
				Acf68865                bool        `json:"acf68865"`
			} `json:"first"`
		} `json:"video"`
	} `json:"series"`
	Smartphone interface{} `json:"smartphone"`
	System     struct {
		ServerTime time.Time `json:"serverTime"`
		IsPeakTime bool      `json:"isPeakTime"`
	} `json:"system"`
	Tag struct {
		Items                 []interface{} `json:"items"`
		HasR18Tag             bool          `json:"hasR18Tag"`
		IsPublishedNicoscript bool          `json:"isPublishedNicoscript"`
		Edit                  struct {
			IsEditable       bool        `json:"isEditable"`
			UneditableReason interface{} `json:"uneditableReason"`
			EditKey          string      `json:"editKey"`
		} `json:"edit"`
		Viewer struct {
			IsEditable       bool        `json:"isEditable"`
			UneditableReason interface{} `json:"uneditableReason"`
			EditKey          string      `json:"editKey"`
		} `json:"viewer"`
	} `json:"tag"`
	Video struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Count       struct {
			View    int `json:"view"`
			Comment int `json:"comment"`
			Mylist  int `json:"mylist"`
			Like    int `json:"like"`
		} `json:"count"`
		Duration  int `json:"duration"`
		Thumbnail struct {
			URL       string `json:"url"`
			MiddleURL string `json:"middleUrl"`
			LargeURL  string `json:"largeUrl"`
			Player    string `json:"player"`
			Ogp       string `json:"ogp"`
		} `json:"thumbnail"`
		Rating struct {
			IsAdult bool `json:"isAdult"`
		} `json:"rating"`
		RegisteredAt             time.Time `json:"registeredAt"`
		IsPrivate                bool      `json:"isPrivate"`
		IsDeleted                bool      `json:"isDeleted"`
		IsNoBanner               bool      `json:"isNoBanner"`
		IsAuthenticationRequired bool      `json:"isAuthenticationRequired"`
		IsEmbedPlayerAllowed     bool      `json:"isEmbedPlayerAllowed"`
		IsGiftAllowed            bool      `json:"isGiftAllowed"`
		Viewer                   struct {
			IsOwner bool `json:"isOwner"`
			Like    struct {
				IsLiked bool        `json:"isLiked"`
				Count   interface{} `json:"count"`
			} `json:"like"`
		} `json:"viewer"`
		WatchableUserTypeForPayment   string `json:"watchableUserTypeForPayment"`
		CommentableUserTypeForPayment string `json:"commentableUserTypeForPayment"`
		NineD091F87                   bool   `json:"9d091f87"`
	} `json:"video"`
	VideoAds struct {
		AdditionalParams struct {
			VideoID                  string `json:"videoId"`
			VideoDuration            int    `json:"videoDuration"`
			IsAdultRatingNG          bool   `json:"isAdultRatingNG"`
			IsAuthenticationRequired bool   `json:"isAuthenticationRequired"`
			IsR18                    bool   `json:"isR18"`
			Nicosid                  string `json:"nicosid"`
			Lang                     string `json:"lang"`
			WatchTrackID             string `json:"watchTrackId"`
			ChannelID                string `json:"channelId"`
			Genre                    string `json:"genre"`
			Gender                   string `json:"gender"`
			Age                      int    `json:"age"`
		} `json:"additionalParams"`
		Items  []interface{} `json:"items"`
		Reason interface{}   `json:"reason"`
	} `json:"videoAds"`
	VideoLive interface{} `json:"videoLive"`
	Viewer    struct {
		ID        int    `json:"id"`
		Nickname  string `json:"nickname"`
		IsPremium bool   `json:"isPremium"`
		Existence struct {
			Age        int    `json:"age"`
			Prefecture string `json:"prefecture"`
			Sex        string `json:"sex"`
		} `json:"existence"`
	} `json:"viewer"`
	Waku struct {
		Information       interface{}   `json:"information"`
		BgImages          []interface{} `json:"bgImages"`
		AddContents       interface{}   `json:"addContents"`
		AddVideo          interface{}   `json:"addVideo"`
		TagRelatedBanner  interface{}   `json:"tagRelatedBanner"`
		TagRelatedMarquee interface{}   `json:"tagRelatedMarquee"`
	} `json:"waku"`
}

func MakeAPIData(rawJSON string) APIData {
	var apiData APIData
	json.Unmarshal([]byte(rawJSON), &apiData)
	return apiData
}
