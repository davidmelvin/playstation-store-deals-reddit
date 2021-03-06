package main

type AutoGenerated struct {
	Props         Props         `json:"props"`
	Page          string        `json:"page"`
	Query         Query         `json:"query"`
	BuildID       string        `json:"buildId"`
	RuntimeConfig RuntimeConfig `json:"runtimeConfig"`
	IsFallback    bool          `json:"isFallback"`
	CustomServer  bool          `json:"customServer"`
	Gip           bool          `json:"gip"`
	AppGip        bool          `json:"appGip"`
}
type PageProps struct {
	Batarangs      interface{} `json:"batarangs"`
	CategoryID     string      `json:"categoryId"`
	IsInvalidRoute bool        `json:"isInvalidRoute"`
	Locale         string      `json:"locale"`
	Page           int         `json:"page"`
	StatusCode     int         `json:"statusCode"`
}
type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}
type RealName struct {
	Name Name `json:"name"`
}
type Phones struct {
	ID          string `json:"id"`
	IsVerified  bool   `json:"isVerified"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
	Qualifier   string `json:"qualifier"`
	IsMain      bool   `json:"isMain"`
}
type EmsCriteria struct {
	PsPlus   string `json:"psPlus"`
	PsNow    string `json:"psNow"`
	EaAccess string `json:"eaAccess"`
}
type UserData struct {
	Age                  int         `json:"age"`
	Country              string      `json:"country"`
	Language             string      `json:"language"`
	Locale               string      `json:"locale"`
	AccountID            string      `json:"accountId"`
	AccountUUID          string      `json:"accountUuid"`
	RealName             RealName    `json:"realName"`
	IsBanned             bool        `json:"isBanned"`
	IsSuspended          bool        `json:"isSuspended"`
	OnlineID             string      `json:"onlineId"`
	SecurityResetVersion int         `json:"securityResetVersion"`
	ToSUAVersion         int         `json:"toSUAVersion"`
	HasPin               bool        `json:"hasPin"`
	CustomerSince        string      `json:"customerSince"`
	Phones               []Phones    `json:"phones"`
	VerificationStatus   string      `json:"verificationStatus"`
	LastUpdateDate       string      `json:"lastUpdateDate"`
	Region               string      `json:"region"`
	NpStatusValue        string      `json:"npStatusValue"`
	IsPsPlusMember       bool        `json:"isPsPlusMember"`
	EmsCriteria          EmsCriteria `json:"emsCriteria"`
	HashedAccountID      string      `json:"hashedAccountId"`
	IsLoggedIn           bool        `json:"isLoggedIn"`
	Name                 string      `json:"name"`
	Relationship         string      `json:"relationship"`
}
type Session struct {
	UserData    UserData `json:"userData"`
	RedirectURI string   `json:"redirectUri"`
	ReferrerURL string   `json:"referrerUrl"`
}
type AppProps struct {
	Session Session `json:"session"`
}
type NavTree struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}
type Views struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}
type Items struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}
type Components struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}

type Link struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}
type Price struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}
type Platforms struct {
	Type string        `json:"type"`
	JSON []interface{} `json:"json"`
}
type Media struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}
type Skus struct {
	Type      string `json:"type"`
	Generated bool   `json:"generated"`
	ID        string `json:"id"`
	Typename  string `json:"typename"`
}

type ServiceBranding struct {
	Type string        `json:"type"`
	JSON []interface{} `json:"json"`
}

type Props struct {
	PageProps   PageProps              `json:"pageProps"`
	AppProps    AppProps               `json:"appProps"`
	ApolloState map[string]interface{} `json:"apolloState"`
}
type Query struct {
	Locale     string `json:"locale"`
	CategoryID string `json:"categoryId"`
	Page       string `json:"page"`
}
type GqlSsr struct {
	Host string `json:"host"`
}
type GqlBrowser struct {
	Host string `json:"host"`
}
type ServiceSession struct {
	Host string `json:"host"`
}
type Checkout struct {
	Host string `json:"host"`
}
type TelemetryURL struct {
	Host string `json:"host"`
}
type Service struct {
	GqlSsr       GqlSsr         `json:"gqlSsr"`
	GqlBrowser   GqlBrowser     `json:"gqlBrowser"`
	Session      ServiceSession `json:"session"`
	Checkout     Checkout       `json:"checkout"`
	TelemetryURL TelemetryURL   `json:"telemetryUrl"`
}
type Pdc struct {
	Host string `json:"host"`
}
type StaticAsset struct {
	Host           string `json:"host"`
	FontsMain      string `json:"fontsMain"`
	BaseStylesMain string `json:"baseStylesMain"`
}
type Feature struct {
	QueryWhitelist                  bool `json:"queryWhitelist"`
	PdpEnabled                      bool `json:"pdpEnabled"`
	TelemetrySamplingThreshold      int  `json:"telemetrySamplingThreshold"`
	TelemetryFilterLoadEventMetrics bool `json:"telemetryFilterLoadEventMetrics"`
}
type RuntimeConfig struct {
	Env         string      `json:"env"`
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	GitBranch   string      `json:"gitBranch"`
	GitSha      string      `json:"gitSha"`
	Isoenv      string      `json:"isoenv"`
	Service     Service     `json:"service"`
	Pdc         Pdc         `json:"pdc"`
	EmsClientID string      `json:"emsClientId"`
	StaticAsset StaticAsset `json:"staticAsset"`
	Feature     Feature     `json:"feature"`
}

type Product struct {
	ID                                  string      `json:"id"`
	Name                                string      `json:"name"`
	Price                               Price       `json:"price"`
	Platforms                           Platforms   `json:"platforms"`
	LocalizedStoreDisplayClassification interface{} `json:"localizedStoreDisplayClassification"`
	Media                               []Media     `json:"media"`
	Skus                                []Skus      `json:"skus"`
	Typename                            string      `json:"__typename"`
}

type SkuPrice struct {
	ProductID             string
	BasePrice             string          `json:"basePrice"`
	DiscountedPrice       string          `json:"discountedPrice"`
	DiscountText          interface{}     `json:"discountText"`
	IsFree                bool            `json:"isFree"`
	ServiceBranding       ServiceBranding `json:"serviceBranding"`
	UpsellServiceBranding interface{}     `json:"upsellServiceBranding"`
	UpsellText            interface{}     `json:"upsellText"`
	Typename              string          `json:"__typename"`
}
