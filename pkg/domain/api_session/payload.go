package api_session

import "github.com/shana0440/niconico-downloader-go/pkg/domain"

type APISessionPayload struct {
	Session PayloadSession `json:"session"`
}
type PayloadSrcIDToMux struct {
	VideoSrcIds []string `json:"video_src_ids"`
	AudioSrcIds []string `json:"audio_src_ids"`
}
type PayloadContentSrcIDs struct {
	SrcIDToMux PayloadSrcIDToMux `json:"src_id_to_mux"`
}
type PayloadContentSrcIDSets struct {
	ContentSrcIDs []PayloadContentSrcIDs `json:"content_src_ids"`
}
type PayloadHeartbeat struct {
	Lifetime int `json:"lifetime"`
}
type PayloadKeepMethod struct {
	Heartbeat PayloadHeartbeat `json:"heartbeat"`
}
type PayloadHlsParameters struct {
	UseWellKnownPort string `json:"use_well_known_port"`
	UseSsl           string `json:"use_ssl"`
	TransferPreset   string `json:"transfer_preset"`
	SegmentDuration  int64  `json:"segment_duration"`
}
type PayloadHTTPParametersParameters struct {
	HlsParameters PayloadHlsParameters `json:"hls_parameters"`
}
type PayloadHTTPParameters struct {
	Parameters PayloadHTTPParametersParameters `json:"parameters"`
}
type PayloadParameters struct {
	HTTPParameters PayloadHTTPParameters `json:"http_parameters"`
}
type PayloadProtocol struct {
	Name       string            `json:"name"`
	Parameters PayloadParameters `json:"parameters"`
}
type PayloadSessionOperationAuthBySignature struct {
	Token     string `json:"token"`
	Signature string `json:"signature"`
}
type PayloadSessionOperationAuth struct {
	SessionOperationAuthBySignature PayloadSessionOperationAuthBySignature `json:"session_operation_auth_by_signature"`
}
type PayloadContentAuth struct {
	AuthType          string `json:"auth_type"`
	ContentKeyTimeout int    `json:"content_key_timeout"`
	ServiceID         string `json:"service_id"`
	ServiceUserID     string `json:"service_user_id"`
}
type PayloadClientInfo struct {
	PlayerID string `json:"player_id"`
}
type PayloadSession struct {
	RecipeID             string                      `json:"recipe_id"`
	ContentID            string                      `json:"content_id"`
	ContentType          string                      `json:"content_type"`
	ContentSrcIDSets     []PayloadContentSrcIDSets   `json:"content_src_id_sets"`
	TimingConstraint     string                      `json:"timing_constraint"`
	KeepMethod           PayloadKeepMethod           `json:"keep_method"`
	Protocol             PayloadProtocol             `json:"protocol"`
	ContentURI           string                      `json:"content_uri"`
	SessionOperationAuth PayloadSessionOperationAuth `json:"session_operation_auth"`
	ContentAuth          PayloadContentAuth          `json:"content_auth"`
	ClientInfo           PayloadClientInfo           `json:"client_info"`
	Priority             float64                     `json:"priority"`
}

func MakePayload(apiData domain.APIData) APISessionPayload {
	payload := APISessionPayload{
		Session: PayloadSession{
			RecipeID:    apiData.Media.Delivery.Movie.Session.RecipeID,
			ContentID:   apiData.Media.Delivery.Movie.Session.ContentID,
			ContentType: "movie",
			ContentSrcIDSets: []PayloadContentSrcIDSets{PayloadContentSrcIDSets{
				ContentSrcIDs: []PayloadContentSrcIDs{PayloadContentSrcIDs{
					SrcIDToMux: PayloadSrcIDToMux{
						VideoSrcIds: []string{apiData.Media.Delivery.Movie.Videos[0].ID},
						AudioSrcIds: []string{apiData.Media.Delivery.Movie.Audios[0].ID},
					},
				}},
			}},
			TimingConstraint: "unlimited",
			KeepMethod: PayloadKeepMethod{
				Heartbeat: PayloadHeartbeat{
					Lifetime: apiData.Media.Delivery.Movie.Session.HeartbeatLifetime,
				},
			},
			Protocol: PayloadProtocol{
				Name: "http",
				Parameters: PayloadParameters{
					HTTPParameters: PayloadHTTPParameters{
						Parameters: PayloadHTTPParametersParameters{
							HlsParameters: PayloadHlsParameters{
								UseWellKnownPort: "yes",
								UseSsl:           "yes",
								TransferPreset:   "",
								SegmentDuration:  6000,
							},
						},
					},
				},
			},
			ContentURI: "",
			SessionOperationAuth: PayloadSessionOperationAuth{
				SessionOperationAuthBySignature: PayloadSessionOperationAuthBySignature{
					Token:     apiData.Media.Delivery.Movie.Session.Token,
					Signature: apiData.Media.Delivery.Movie.Session.Signature,
				},
			},
			ContentAuth: PayloadContentAuth{
				AuthType:          apiData.Media.Delivery.Movie.Session.AuthTypes.HTTP,
				ContentKeyTimeout: apiData.Media.Delivery.Movie.Session.ContentKeyTimeout,
				ServiceID:         "nicovideo",
				ServiceUserID:     apiData.Media.Delivery.Movie.Session.ServiceUserID,
			},
			ClientInfo: PayloadClientInfo{
				PlayerID: apiData.Media.Delivery.Movie.Session.PlayerID,
			},
			Priority: apiData.Media.Delivery.Movie.Session.Priority,
		},
	}
	return payload
}
