package api_session

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type APISessionBody struct {
	Meta BodyMeta `json:"meta"`
	Data BodyData `json:"data"`
}
type BodyMeta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type BodySrcIDToMux struct {
	VideoSrcIds []string `json:"video_src_ids"`
	AudioSrcIds []string `json:"audio_src_ids"`
}
type BodyContentSrcIds struct {
	SrcIDToMux BodySrcIDToMux `json:"src_id_to_mux"`
}
type BodyContentSrcIDSets struct {
	ContentSrcIds []BodyContentSrcIds `json:"content_src_ids"`
	AllowSubset   string              `json:"allow_subset"`
}
type BodyHeartbeat struct {
	Lifetime                  int    `json:"lifetime"`
	OnetimeToken              string `json:"onetime_token"`
	DeletionTimeoutOnNoStream int    `json:"deletion_timeout_on_no_stream"`
}
type BodyKeepMethod struct {
	Heartbeat BodyHeartbeat `json:"heartbeat"`
}
type BodyEncryption struct {
	Empty struct{} `json:"empty"`
}
type BodyHlsParameters struct {
	SegmentDuration     int            `json:"segment_duration"`
	TotalDuration       int            `json:"total_duration"`
	TransferPreset      string         `json:"transfer_preset"`
	UseSsl              string         `json:"use_ssl"`
	UseWellKnownPort    string         `json:"use_well_known_port"`
	MediaSegmentFormat  string         `json:"media_segment_format"`
	Encryption          BodyEncryption `json:"encryption"`
	SeparateAudioStream string         `json:"separate_audio_stream"`
}
type BodyHTTPParametersParameters struct {
	HlsParameters BodyHlsParameters `json:"hls_parameters"`
}
type BodyHTTPParameters struct {
	Method     string                       `json:"method"`
	Parameters BodyHTTPParametersParameters `json:"parameters"`
}
type BodyParameters struct {
	HTTPParameters BodyHTTPParameters `json:"http_parameters"`
}
type BodyProtocol struct {
	Name       string         `json:"name"`
	Parameters BodyParameters `json:"parameters"`
}
type BodyPlayControlRange struct {
	MaxPlaySpeed float64 `json:"max_play_speed"`
	MinPlaySpeed float64 `json:"min_play_speed"`
}
type BodySessionOperationAuthBySignature struct {
	CreatedTime int64  `json:"created_time"`
	ExpireTime  int64  `json:"expire_time"`
	Token       string `json:"token"`
	Signature   string `json:"signature"`
}
type BodySessionOperationAuth struct {
	SessionOperationAuthBySignature BodySessionOperationAuthBySignature `json:"session_operation_auth_by_signature"`
}
type BodyContentAuthInfo struct {
	Method string `json:"method"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}
type BodyContentAuth struct {
	AuthType          string              `json:"auth_type"`
	MaxContentCount   int                 `json:"max_content_count"`
	ContentKeyTimeout int                 `json:"content_key_timeout"`
	ServiceID         string              `json:"service_id"`
	ServiceUserID     string              `json:"service_user_id"`
	ContentAuthInfo   BodyContentAuthInfo `json:"content_auth_info"`
}
type BodyRuntimeInfo struct {
	NodeID           string        `json:"node_id"`
	ExecutionHistory []interface{} `json:"execution_history"`
	ThumbnailerState []interface{} `json:"thumbnailer_state"`
}
type BodyClientInfo struct {
	PlayerID     string `json:"player_id"`
	RemoteIP     string `json:"remote_ip"`
	TrackingInfo string `json:"tracking_info"`
}
type BodySession struct {
	ID                   string                   `json:"id"`
	RecipeID             string                   `json:"recipe_id"`
	ContentID            string                   `json:"content_id"`
	ContentSrcIDSets     []BodyContentSrcIDSets   `json:"content_src_id_sets"`
	ContentType          string                   `json:"content_type"`
	TimingConstraint     string                   `json:"timing_constraint"`
	KeepMethod           BodyKeepMethod           `json:"keep_method"`
	Protocol             BodyProtocol             `json:"protocol"`
	PlaySeekTime         int                      `json:"play_seek_time"`
	PlaySpeed            float64                  `json:"play_speed"`
	PlayControlRange     BodyPlayControlRange     `json:"play_control_range"`
	ContentURI           string                   `json:"content_uri"`
	SessionOperationAuth BodySessionOperationAuth `json:"session_operation_auth"`
	ContentAuth          BodyContentAuth          `json:"content_auth"`
	RuntimeInfo          BodyRuntimeInfo          `json:"runtime_info"`
	ClientInfo           BodyClientInfo           `json:"client_info"`
	CreatedTime          int64                    `json:"created_time"`
	ModifiedTime         int64                    `json:"modified_time"`
	Priority             float64                  `json:"priority"`
	ContentRoute         int                      `json:"content_route"`
	Version              string                   `json:"version"`
	ContentStatus        string                   `json:"content_status"`
}
type BodyData struct {
	Session BodySession `json:"session"`
}

func MakeBody(resp *http.Response) APISessionBody {
	var body APISessionBody
	rawJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal([]byte(rawJSON), &body)
	return body
}
