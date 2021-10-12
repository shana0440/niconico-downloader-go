package heartbeat

import "github.com/shana0440/niconico-downloader-go/pkg/domain/api_session"

type HeartBeatPayload struct {
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
	ContentSrcIds []PayloadContentSrcIDs `json:"content_src_ids"`
	AllowSubset   string                 `json:"allow_subset"`
}
type PayloadHeartbeat struct {
	Lifetime                  int    `json:"lifetime"`
	OnetimeToken              string `json:"onetime_token"`
	DeletionTimeoutOnNoStream int    `json:"deletion_timeout_on_no_stream"`
}
type PayloadKeepMethod struct {
	Heartbeat PayloadHeartbeat `json:"heartbeat"`
}
type PayloadEncryption struct {
	Empty struct{} `json:"empty"`
}
type PayloadHlsParameters struct {
	SegmentDuration     int               `json:"segment_duration"`
	TotalDuration       int               `json:"total_duration"`
	TransferPreset      string            `json:"transfer_preset"`
	UseSsl              string            `json:"use_ssl"`
	UseWellKnownPort    string            `json:"use_well_known_port"`
	MediaSegmentFormat  string            `json:"media_segment_format"`
	Encryption          PayloadEncryption `json:"encryption"`
	SeparateAudioStream string            `json:"separate_audio_stream"`
}
type PayloadHTTPParametersParameters struct {
	HlsParameters PayloadHlsParameters `json:"hls_parameters"`
}
type PayloadHTTPParameters struct {
	Method     string                          `json:"method"`
	Parameters PayloadHTTPParametersParameters `json:"parameters"`
}
type PayloadParameters struct {
	HTTPParameters PayloadHTTPParameters `json:"http_parameters"`
}
type PayloadProtocol struct {
	Name       string            `json:"name"`
	Parameters PayloadParameters `json:"parameters"`
}
type PayloadPlayControlRange struct {
	MaxPlaySpeed float64 `json:"max_play_speed"`
	MinPlaySpeed float64 `json:"min_play_speed"`
}
type PayloadSessionOperationAuthBySignature struct {
	CreatedTime int64  `json:"created_time"`
	ExpireTime  int64  `json:"expire_time"`
	Token       string `json:"token"`
	Signature   string `json:"signature"`
}
type PayloadSessionOperationAuth struct {
	SessionOperationAuthBySignature PayloadSessionOperationAuthBySignature `json:"session_operation_auth_by_signature"`
}
type PayloadContentAuthInfo struct {
	Method string `json:"method"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}
type PayloadContentAuth struct {
	AuthType          string                 `json:"auth_type"`
	MaxContentCount   int                    `json:"max_content_count"`
	ContentKeyTimeout int                    `json:"content_key_timeout"`
	ServiceID         string                 `json:"service_id"`
	ServiceUserID     string                 `json:"service_user_id"`
	ContentAuthInfo   PayloadContentAuthInfo `json:"content_auth_info"`
}
type PayloadRuntimeInfo struct {
	NodeID           string        `json:"node_id"`
	ExecutionHistory []interface{} `json:"execution_history"`
	ThumbnailerState []interface{} `json:"thumbnailer_state"`
}
type PayloadClientInfo struct {
	PlayerID     string `json:"player_id"`
	RemoteIP     string `json:"remote_ip"`
	TrackingInfo string `json:"tracking_info"`
}
type PayloadSession struct {
	ID                   string                      `json:"id"`
	RecipeID             string                      `json:"recipe_id"`
	ContentID            string                      `json:"content_id"`
	ContentSrcIDSets     []PayloadContentSrcIDSets   `json:"content_src_id_sets"`
	ContentType          string                      `json:"content_type"`
	TimingConstraint     string                      `json:"timing_constraint"`
	KeepMethod           PayloadKeepMethod           `json:"keep_method"`
	Protocol             PayloadProtocol             `json:"protocol"`
	PlaySeekTime         int                         `json:"play_seek_time"`
	PlaySpeed            float64                     `json:"play_speed"`
	PlayControlRange     PayloadPlayControlRange     `json:"play_control_range"`
	ContentURI           string                      `json:"content_uri"`
	SessionOperationAuth PayloadSessionOperationAuth `json:"session_operation_auth"`
	ContentAuth          PayloadContentAuth          `json:"content_auth"`
	RuntimeInfo          PayloadRuntimeInfo          `json:"runtime_info"`
	ClientInfo           PayloadClientInfo           `json:"client_info"`
	CreatedTime          int64                       `json:"created_time"`
	ModifiedTime         int64                       `json:"modified_time"`
	Priority             float64                     `json:"priority"`
	ContentRoute         int                         `json:"content_route"`
	Version              string                      `json:"version"`
	ContentStatus        string                      `json:"content_status"`
}

func MakePayload(apiSessionBody api_session.APISessionBody) HeartBeatPayload {
	payload := HeartBeatPayload{
		Session: PayloadSession{
			ID:               apiSessionBody.Data.Session.ID,
			RecipeID:         apiSessionBody.Data.Session.RecipeID,
			ContentID:        apiSessionBody.Data.Session.ContentID,
			ContentSrcIDSets: makePayloadContentSrcIDSets(apiSessionBody.Data.Session.ContentSrcIDSets),
			ContentType:      apiSessionBody.Data.Session.ContentType,
			TimingConstraint: apiSessionBody.Data.Session.TimingConstraint,
			KeepMethod: PayloadKeepMethod{
				Heartbeat: PayloadHeartbeat{
					Lifetime:                  apiSessionBody.Data.Session.KeepMethod.Heartbeat.Lifetime,
					OnetimeToken:              apiSessionBody.Data.Session.KeepMethod.Heartbeat.OnetimeToken,
					DeletionTimeoutOnNoStream: apiSessionBody.Data.Session.KeepMethod.Heartbeat.DeletionTimeoutOnNoStream,
				},
			},
			Protocol: PayloadProtocol{
				Name: apiSessionBody.Data.Session.Protocol.Name,
				Parameters: PayloadParameters{
					HTTPParameters: PayloadHTTPParameters{
						Method: apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Method,
						Parameters: PayloadHTTPParametersParameters{
							HlsParameters: PayloadHlsParameters{
								SegmentDuration:    apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.SegmentDuration,
								TotalDuration:      apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.TotalDuration,
								TransferPreset:     apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.TransferPreset,
								UseSsl:             apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.UseSsl,
								UseWellKnownPort:   apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.UseWellKnownPort,
								MediaSegmentFormat: apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.MediaSegmentFormat,
								Encryption: PayloadEncryption{
									Empty: apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.Encryption.Empty,
								},
								SeparateAudioStream: apiSessionBody.Data.Session.Protocol.Parameters.HTTPParameters.Parameters.HlsParameters.SeparateAudioStream,
							},
						},
					},
				},
			},
			PlaySeekTime: apiSessionBody.Data.Session.PlaySeekTime,
			PlaySpeed:    apiSessionBody.Data.Session.PlaySpeed,
			PlayControlRange: PayloadPlayControlRange{
				MaxPlaySpeed: apiSessionBody.Data.Session.PlayControlRange.MaxPlaySpeed,
				MinPlaySpeed: apiSessionBody.Data.Session.PlayControlRange.MinPlaySpeed,
			},
			ContentURI: apiSessionBody.Data.Session.ContentURI,
			SessionOperationAuth: PayloadSessionOperationAuth{
				SessionOperationAuthBySignature: PayloadSessionOperationAuthBySignature{
					CreatedTime: apiSessionBody.Data.Session.SessionOperationAuth.SessionOperationAuthBySignature.CreatedTime,
					ExpireTime:  apiSessionBody.Data.Session.SessionOperationAuth.SessionOperationAuthBySignature.ExpireTime,
					Token:       apiSessionBody.Data.Session.SessionOperationAuth.SessionOperationAuthBySignature.Token,
					Signature:   apiSessionBody.Data.Session.SessionOperationAuth.SessionOperationAuthBySignature.Signature,
				},
			},
			ContentAuth: PayloadContentAuth{
				AuthType:          apiSessionBody.Data.Session.ContentAuth.AuthType,
				MaxContentCount:   apiSessionBody.Data.Session.ContentAuth.MaxContentCount,
				ContentKeyTimeout: apiSessionBody.Data.Session.ContentAuth.ContentKeyTimeout,
				ServiceID:         apiSessionBody.Data.Session.ContentAuth.ServiceID,
				ServiceUserID:     apiSessionBody.Data.Session.ContentAuth.ServiceUserID,
				ContentAuthInfo: PayloadContentAuthInfo{
					Method: apiSessionBody.Data.Session.ContentAuth.ContentAuthInfo.Method,
					Name:   apiSessionBody.Data.Session.ContentAuth.ContentAuthInfo.Name,
					Value:  apiSessionBody.Data.Session.ContentAuth.ContentAuthInfo.Value,
				},
			},
			RuntimeInfo: PayloadRuntimeInfo{
				NodeID:           apiSessionBody.Data.Session.RuntimeInfo.NodeID,
				ExecutionHistory: apiSessionBody.Data.Session.RuntimeInfo.ExecutionHistory,
				ThumbnailerState: apiSessionBody.Data.Session.RuntimeInfo.ThumbnailerState,
			},
			ClientInfo: PayloadClientInfo{
				PlayerID:     apiSessionBody.Data.Session.ClientInfo.PlayerID,
				RemoteIP:     apiSessionBody.Data.Session.ClientInfo.RemoteIP,
				TrackingInfo: apiSessionBody.Data.Session.ClientInfo.TrackingInfo,
			},
			CreatedTime:   apiSessionBody.Data.Session.CreatedTime,
			ModifiedTime:  apiSessionBody.Data.Session.ModifiedTime,
			Priority:      apiSessionBody.Data.Session.Priority,
			ContentRoute:  apiSessionBody.Data.Session.ContentRoute,
			Version:       apiSessionBody.Data.Session.Version,
			ContentStatus: apiSessionBody.Data.Session.ContentStatus,
		},
	}
	return payload
}

func makePayloadContentSrcIDSets(idSets []api_session.BodyContentSrcIDSets) []PayloadContentSrcIDSets {
	payloadIDSets := make([]PayloadContentSrcIDSets, len(idSets))
	for i, idSet := range idSets {
		payloadIDs := make([]PayloadContentSrcIDs, len(idSet.ContentSrcIds))
		for j, id := range idSet.ContentSrcIds {
			payloadID := PayloadContentSrcIDs{
				SrcIDToMux: PayloadSrcIDToMux{
					VideoSrcIds: id.SrcIDToMux.VideoSrcIds,
					AudioSrcIds: id.SrcIDToMux.AudioSrcIds,
				},
			}
			payloadIDs[j] = payloadID
		}
		payloadIDSets[i] = PayloadContentSrcIDSets{
			ContentSrcIds: payloadIDs,
			AllowSubset:   idSet.AllowSubset,
		}
	}
	return payloadIDSets
}
