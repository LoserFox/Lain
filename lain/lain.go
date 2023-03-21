package lain

const BA_JP_VERSION_METADATA_TEMPLATE = "https://yostar-serverinfo.bluearchiveyostar.com/%s.json"
const BA_JP_ANDROID_BUNDLE_DOWNLOAD_INFO_TEMPLATE = "%s/Android/bundleDownloadInfo.json"
const BA_JP_ANDROID_BUNDLE_BASEURL_TEMPLATE = "%s/Android/"
const BA_JP_IOS_BUNDLE_DOWNLOAD_INFO_TEMPLATE = "%s/iOS/bundleDownloadInfo.json"
const BA_JP_IOS_BUNDLE_BASEURL_TEMPLATE = "%s/iOS/"
const BA_JP_MEDIA_CATALOG_TEMPLATE = "%s/MediaResources/MediaCatalog.json"
const BA_JP_MEDIA_BASEURL_TEMPLATE = "%s/MediaResources/%s"
const BA_JP_TABLE_BUNDLES_CATALOG_TEMPLATE = "%s/TableBundles/TableCatalog.json"
const BA_JP_TABLE_BUNDLES_BASEURL_TEMPLATE = "%s/TableBundles/"

type BA_JP_VERSION_METADATA struct {
	ConnectionGroups []struct {
		Name                       string `json:"Name"`
		ManagementDataURL          string `json:"ManagementDataUrl"`
		IsProductionAddressables   bool   `json:"IsProductionAddressables"`
		APIURL                     string `json:"ApiUrl"`
		GatewayURL                 string `json:"GatewayUrl"`
		KibanaLogURL               string `json:"KibanaLogUrl"`
		ProhibitedWordBlackListURI string `json:"ProhibitedWordBlackListUri"`
		ProhibitedWordWhiteListURI string `json:"ProhibitedWordWhiteListUri"`
		CustomerServiceURL         string `json:"CustomerServiceUrl"`
		OverrideConnectionGroups   []struct {
			Name                       string `json:"Name"`
			AddressablesCatalogURLRoot string `json:"AddressablesCatalogUrlRoot"`
		} `json:"OverrideConnectionGroups"`
		BundleVersion string `json:"BundleVersion"`
	} `json:"ConnectionGroups"`
}

type BA_JP_MEDIA_DATA struct {
	Table     map[string]BA_JP_MEDIA_DATA_TABLE `json:"Table"`
	MediaList string                            `json:"MediaList"`
}
type BA_JP_MEDIA_DATA_TABLE struct {
	IsChanged bool   `json:"isChanged"`
	MediaType int    `json:"mediaType"`
	Path      string `json:"path"`
	FileName  string `json:"fileName"`
	Bytes     int    `json:"bytes"`
	Crc       int64  `json:"Crc"`
	IsInbuild bool   `json:"isInbuild"`
}
