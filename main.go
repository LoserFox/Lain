package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"lain/bruteforce"
	"lain/lain"
	"net/http"
)

func main() {
	proxy := flag.String("proxy", "", "connect server proxy setting")
	thread := flag.Int("thread", 512, "download thread count")
	version := flag.String("version", GetGameVersion(), "internal game version metadata")
	path := flag.String("path", "Media", "file download path")
	flag.Parse()

	assetsBaseUrl := GetAddressablesCatalogURL(GetAssetMetaData())
	assetsMediaList := GetAssetMediaList(assetsBaseUrl)
	bruteforce.Start(assetsBaseUrl, *thread, assetsMediaList, *path, *version, *proxy)
	fmt.Println("Congratulations. ")
}
func GetAddressablesCatalogURL(in lain.BA_JP_VERSION_METADATA) string {
	return in.ConnectionGroups[0].OverrideConnectionGroups[1].AddressablesCatalogURLRoot
}
func GetAssetMetaData() lain.BA_JP_VERSION_METADATA {
	resp, _ := http.Get(GetAssetDownloadURL())
	defer resp.Body.Close()
	fmt.Println("GetAssetMetaData OK!")
	body, _ := io.ReadAll(resp.Body)
	var data lain.BA_JP_VERSION_METADATA
	json.Unmarshal(body, &data)
	return data
}
func GetAssetMediaList(baseurl string) lain.BA_JP_MEDIA_DATA {
	resp, _ := http.Get(fmt.Sprintf(lain.BA_JP_MEDIA_CATALOG_TEMPLATE, baseurl))
	defer resp.Body.Close()
	fmt.Println("GetAssetMediaList OK!")
	body, _ := io.ReadAll(resp.Body)
	var data lain.BA_JP_MEDIA_DATA
	json.Unmarshal(body, &data)
	return data
}

// create a get asset download url file
func GetAssetDownloadURL() string {
	return fmt.Sprintf(lain.BA_JP_VERSION_METADATA_TEMPLATE, GetGameVersion())
}
func GetGameVersion() string {
	return "r53_29_rfhqxfw36obfz83ei505"
}
