package domains

import "encoding/json"

type CharacterOnYMMP struct {
	Name                             string           `json:"Name,omitempty"`
	GroupName                        *string          `json:"GroupName,omitempty"`
	Color                            *string          `json:"Color,omitempty"`
	Layer                            *int             `json:"Layer,omitempty"`
	KeyGesture                       *json.RawMessage `json:"KeyGesture,omitempty"`
	Voice                            *json.RawMessage `json:"Voice,omitempty"`
	Volume                           *json.RawMessage `json:"Volume,omitempty"`
	Pan                              *json.RawMessage `json:"Pan,omitempty"`
	PlaybackRate                     *float64         `json:"PlaybackRate,omitempty"`
	VoiceParameter                   *json.RawMessage `json:"VoiceParameter,omitempty"`
	AdditionalTime                   *float64         `json:"AdditionalTime,omitempty"`
	VoiceFadeIn                      *float64         `json:"VoiceFadeIn,omitempty"`
	VoiceFadeOut                     *float64         `json:"VoiceFadeOut,omitempty"`
	EchoIsEnabled                    *bool            `json:"EchoIsEnabled,omitempty"`
	EchoInterval                     *float64         `json:"EchoInterval,omitempty"`
	EchoAttenuation                  *float64         `json:"EchoAttenuation,omitempty"`
	CustomVoiceIsEnabled             *bool            `json:"CustomVoiceIsEnabled,omitempty"`
	CustomVoiceIncludeSubdirectories *bool            `json:"CustomVoiceIncludeSubdirectories,omitempty"`
	CustomVoiceDirectory             *string          `json:"CustomVoiceDirectory,omitempty"`
	CustomVoiceFileName              *string          `json:"CustomVoiceFileName,omitempty"`
	AudioEffects                     *json.RawMessage `json:"AudioEffects,omitempty"`
	IsJimakuVisible                  *bool            `json:"IsJimakuVisible,omitempty"`
	IsJimakuLocked                   *bool            `json:"IsJimakuLocked,omitempty"`
	X                                *json.RawMessage `json:"X,omitempty"`
	Y                                *json.RawMessage `json:"Y,omitempty"`
	Opacity                          *json.RawMessage `json:"Opacity,omitempty"`
	Zoom                             *json.RawMessage `json:"Zoom,omitempty"`
	Rotation                         *json.RawMessage `json:"Rotation,omitempty"`
	JimakuFadeIn                     *float64         `json:"JimakuFadeIn,omitempty"`
	JimakuFadeOut                    *float64         `json:"JimakuFadeOut,omitempty"`
	Blend                            *string          `json:"Blend,omitempty"`
	IsInverted                       *bool            `json:"IsInverted,omitempty"`
	IsAlwaysOnTop                    *bool            `json:"IsAlwaysOnTop,omitempty"`
	IsClippingWithObjectAbove        *bool            `json:"IsClippingWithObjectAbove,omitempty"`
	Font                             *string          `json:"Font,omitempty"`
	FontSize                         *json.RawMessage `json:"FontSize,omitempty"`
	LineHeight2                      *json.RawMessage `json:"LineHeight2,omitempty"`
	LetterSpacing2                   *json.RawMessage `json:"LetterSpacing2,omitempty"`
	DisplayInterval                  *float64         `json:"DisplayInterval,omitempty"`
	WordWrap                         *string          `json:"WordWrap,omitempty"`
	MaxWidth                         *json.RawMessage `json:"MaxWidth,omitempty"`
	BasePoint                        *string          `json:"BasePoint,omitempty"`
	FontColor                        *string          `json:"FontColor,omitempty"`
	Style                            *string          `json:"Style,omitempty"`
	StyleColor                       *string          `json:"StyleColor,omitempty"`
	Bold                             *bool            `json:"Bold,omitempty"`
	Italic                           *bool            `json:"Italic,omitempty"`
	IsDevidedPerCharacter            *bool            `json:"IsDevidedPerCharacter,omitempty"`
	JimakuVideoEffects               *json.RawMessage `json:"JimakuVideoEffects,omitempty"`
	TachieType                       *string          `json:"TachieType,omitempty"`
	TachieCharacterParameter         *json.RawMessage `json:"TachieCharacterParameter,omitempty"`
	IsTachieLocked                   *bool            `json:"IsTachieLocked,omitempty"`
	TachieX                          *json.RawMessage `json:"TachieX,omitempty"`
	TachieY                          *json.RawMessage `json:"TachieY,omitempty"`
	TachieOpacity                    *json.RawMessage `json:"TachieOpacity,omitempty"`
	TachieZoom                       *json.RawMessage `json:"TachieZoom,omitempty"`
	TachieRotation                   *json.RawMessage `json:"TachieRotation,omitempty"`
	TachieFadeIn                     *float64         `json:"TachieFadeIn,omitempty"`
	TachieFadeOut                    *float64         `json:"TachieFadeOut,omitempty"`
	TachieBlend                      *string          `json:"TachieBlend,omitempty"`
	TachieIsInverted                 *bool            `json:"TachieIsInverted,omitempty"`
	TachieIsAlwaysOnTop              *bool            `json:"TachieIsAlwaysOnTop,omitempty"`
	TachieIsClippingWithObjectAbove  *bool            `json:"TachieIsClippingWithObjectAbove,omitempty"`
	TachieDefaultItemParameter       *json.RawMessage `json:"TachieDefaultItemParameter,omitempty"`
	TachieItemVideoEffects           *json.RawMessage `json:"TachieItemVideoEffects,omitempty"`
	TachieDefaultFaceParameter       *json.RawMessage `json:"TachieDefaultFaceParameter,omitempty"`
	TachieDefaultFaceEffects         *json.RawMessage `json:"TachieDefaultFaceEffects,omitempty"`
	AdditionalForegroundTemplateName *json.RawMessage `json:"AdditionalForegroundTemplateName,omitempty"`
	AdditionalBackgroundTemplateName *json.RawMessage `json:"AdditionalBackgroundTemplateName,omitempty"`
	VoiceItemLength                  int              `json:"VoiceItemLength,omitempty"`
	TachieItemLength                 int              `json:"TachieItemLength,omitempty"`
	VoiceItemKeyFrames               *json.RawMessage `json:"VoiceItemKeyFrames,omitempty"`
	TachieItemKeyFrames              *json.RawMessage `json:"TachieItemKeyFrameworks,omitempty"`
}

type ItemOnYMMP struct {
	Type                *string              `json:"$type"`
	Frame               int                  `json:"Frame"`
	Layer               int                  `json:"Layer"`
	Length              int                  `json:"Length"`
	FilePath            string               `json:"FilePath"`
	TachieFaceParameter *TachieFaceParameter `json:"TachieFaceParameter"`
	*UnchangeableFields
}

type TachieFaceParameter struct {
	Type           string          `json:"$type"`
	EyeAnimation   string          `json:"EyeAnimation"`
	MouthAnimation string          `json:"MouthAnimation"`
	Eyebrow        string          `json:"Eyebrow"`
	Eye            string          `json:"Eye"`
	Mouth          string          `json:"Mouth"`
	Hair           string          `json:"Hair"`
	Complexion     string          `json:"Complexion"`
	Body           string          `json:"Body"`
	Back1          string          `json:"Back1"`
	Back2          string          `json:"Back2"`
	Back3          string          `json:"Back3"`
	Etc1           string          `json:"Etc1"`
	Etc2           string          `json:"Etc2"`
	Etc3           string          `json:"Etc3"`
	Face           json.RawMessage `json:"Face"`
	EnableLayers   json.RawMessage `json:"EnableLayers"`
	IsEnabled      bool            `json:"IsEnabled"`
	FilePath       string          `json:"FilePath"`
}

type UnchangeableFields struct {
	IsWaveformEnabled               bool             `json:"IsWaveformEnabled,omitempty"`
	CharacterName                   *string          `json:"CharacterName,omitempty"`
	Serif                           *string          `json:"Serif,omitempty"`
	Decorations                     *json.RawMessage `json:"Decorations,omitempty"`
	Hatsuon                         *json.RawMessage `json:"Hatsuon,omitempty"`
	Pronounce                       *json.RawMessage `json:"Pronounce,omitempty"`
	VoiceLength                     string           `json:"VoiceLength,omitempty"`
	VoiceCache                      *string          `json:"VoiceCache,omitempty"`
	Volume                          *json.RawMessage `json:"Volume,omitempty"`
	Pan                             *json.RawMessage `json:"Pan,omitempty"`
	PlaybackRate                    float64          `json:"PlaybackRate,omitempty"`
	VoiceParameter                  *json.RawMessage `json:"VoiceParameter,omitempty"`
	ContentOffset                   string           `json:"ContentOffset,omitempty"`
	VoiceFadeIn                     float64          `json:"VoiceFadeIn,omitempty"`
	VoiceFadeOut                    float64          `json:"VoiceFadeOut,omitempty"`
	EchoIsEnabled                   bool             `json:"EchoIsEnabled,omitempty"`
	EchoInterval                    float64          `json:"EchoInterval,omitempty"`
	EchoAttenuation                 float64          `json:"EchoAttenuation,omitempty"`
	AudioEffects                    *json.RawMessage `json:"AudioEffects,omitempty"`
	JimakuVisibility                string           `json:"JimakuVisibility,omitempty"`
	Y                               *json.RawMessage `json:"Y,omitempty"`
	X                               *json.RawMessage `json:"X,omitempty"`
	Opacity                         *json.RawMessage `json:"Opacity,omitempty"`
	Zoom                            *json.RawMessage `json:"Zoom,omitempty"`
	Rotation                        *json.RawMessage `json:"Rotation,omitempty"`
	JimakuFadeIn                    float64          `json:"JimakuFadeIn,omitempty"`
	JimakuFadeOut                   float64          `json:"JimakuFadeOut,omitempty"`
	Blend                           string           `json:"Blend,omitempty"`
	IsInverted                      bool             `json:"IsInverted,omitempty"`
	IsAlwaysOnTop                   bool             `json:"IsAlwaysOnTop,omitempty"`
	IsClippingWithObjectAbove       bool             `json:"IsClippingWithObjectAbove,omitempty"`
	Font                            string           `json:"Font,omitempty"`
	FontSize                        *json.RawMessage `json:"FontSize,omitempty"`
	LineHeight2                     *json.RawMessage `json:"LineHeight2,omitempty"`
	LetterSpacing2                  *json.RawMessage `json:"LetterSpacing2,omitempty"`
	DisplayInterval                 float64          `json:"DisplayInterval,omitempty"`
	WordWrap                        string           `json:"WordWrap,omitempty"`
	MaxWidth                        *json.RawMessage `json:"MaxWidth,omitempty"`
	BasePoint                       string           `json:"BasePoint,omitempty"`
	FontColor                       string           `json:"FontColor,omitempty"`
	Style                           string           `json:"Style,omitempty"`
	StyleColor                      string           `json:"StyleColor,omitempty"`
	Bold                            bool             `json:"Bold,omitempty"`
	Italic                          bool             `json:"Italic,omitempty"`
	IsDevidedPerCharacter           bool             `json:"IsDevidedPerCharacter,omitempty"`
	JimakuVideoEffects              *json.RawMessage `json:"JimakuVideoEffects,omitempty"`
	TachieFaceEffects               *json.RawMessage `json:"TachieFaceEffects,omitempty"`
	Group                           float64          `json:"Group,omitempty"`
	IsLocked                        bool             `json:"IsLocked,omitempty"`
	IsHidden                        bool             `json:"IsHidden,omitempty"`
	Text                            *string          `json:"Text,omitempty"`
	FadeIn                          float64          `json:"FadeIn,omitempty"`
	FadeOut                         float64          `json:"FadeOut,omitempty"`
	VideoEffects                    *json.RawMessage `json:"VideoEffects,omitempty"`
	IsLooped                        bool             `json:"IsLooped,omitempty"`
	ShapeType                       string           `json:"ShapeType,omitempty"`
	ShapeParameter                  *json.RawMessage `json:"ShapeParameter,omitempty"`
	TachieItemParameter             *json.RawMessage `json:"TachieItemParameter,omitempty"`
	Blur                            *json.RawMessage `json:"Blur,omitempty"`
	InvertMask                      bool             `json:"InvertMask,omitempty"`
	ClearBackground                 bool             `json:"ClearBackground,omitempty"`
	IsAlphaEnabled                  bool             `json:"IsAlphaEnabled,omitempty"`
	IsFrameOutImageRenderingEnabled bool             `json:"IsFrameOutImageRenderingEnabled,omitempty"`
	GroupRange                      float64          `json:"GroupRange,omitempty"`
	IsGroupOnly                     bool             `json:"IsGroupOnly,omitempty"`
	IsCompressFrame                 bool             `json:"IsCompressFrame,omitempty"`
	KeyFrames                       *json.RawMessage `json:"KeyFrames,omitempty"`
}
type TimelineOnYMMP struct {
	Items         []*ItemOnYMMP    `json:"Items,omitempty"`
	VideoInfo     *json.RawMessage `json:"VideoInfo,omitempty"`
	VerticalLines *json.RawMessage `json:"VerticalLines,omitempty"`
	LayerSettings *json.RawMessage `json:"LayerSettings,omitempty"`
}

type YMMP struct {
	Characters      []*CharacterOnYMMP `json:"Characters,omitempty"`
	Timeline        *TimelineOnYMMP    `json:"Timeline,omitempty"`
	FilePath        string             `json:"FilePath,omitempty"`
	CollapsedGroups *json.RawMessage   `json:"CollapsedGroups,omitempty"`
}

func NewDynamicItem(originalItem *ItemOnYMMP, frame int, length int, layer int, filePath string) *ItemOnYMMP {
	newItem := &ItemOnYMMP{
		Type:                originalItem.Type,
		Frame:               frame,
		Length:              length,
		Layer:               layer,
		FilePath:            filePath,
		TachieFaceParameter: originalItem.TachieFaceParameter,
		UnchangeableFields:  originalItem.UnchangeableFields,
	}
	return newItem
}

func NewSingleItem(originalItem *ItemOnYMMP, frame int, length int, layer int) *ItemOnYMMP {
	newItem := &ItemOnYMMP{
		Type:                originalItem.Type,
		Frame:               frame,
		Length:              length,
		Layer:               layer,
		FilePath:            originalItem.FilePath,
		TachieFaceParameter: originalItem.TachieFaceParameter,
		UnchangeableFields:  originalItem.UnchangeableFields,
	}
	return newItem
}

func NewMultipleItem(originalItem []*ItemOnYMMP, frame int, layer int) []*ItemOnYMMP {
	var newMultipleItems []*ItemOnYMMP
	for _, item := range originalItem {
		newItem := NewSingleItem(item, frame+item.Frame, item.Length, layer+item.Layer)
		newMultipleItems = append(newMultipleItems, newItem)
	}

	return newMultipleItems
}
