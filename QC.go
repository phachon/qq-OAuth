package qq_OAuth

import "strings"

const (
	/*----user----*/
	QC_Api_GetUserInfo = "https://graph.qq.com/user/get_user_info"
	QC_Api_GetSimpleUserInfo = "https://graph.qq.com/user/get_simple_userinfo"
	QC_Api_GetWeiboInfo = "https://graph.qq.com/user/get_info"

	/*----vip----*/
	QC_Api_GetVipInfo = "https://graph.qq.com/user/get_vip_info"
	QC_Api_GetVipRichInfo = "https://graph.qq.com/user/get_vip_rich_info"

	/*----pic----*/
	QC_Api_ListAlbum = "https://graph.qq.com/photo/list_album"
	QC_Api_UploadPic = "https://graph.qq.com/photo/upload_pic"
	QC_Api_AddAlbum = "https://graph.qq.com/photo/add_album"
	QC_Api_ListPhoto = "https://graph.qq.com/photo/list_photo"

	/*----fans----*/
	QC_Api_CheckPageFans = "https://graph.qq.com/user/check_page_fans"

	/*----weibo----*/
	QC_Api_AddT = "https://graph.qq.com/t/add_t"
	QC_Api_DeleteT = "https://graph.qq.com/t/del_t"
	QC_Api_AddPicT = "https://graph.qq.com/t/add_pic_t"
	QC_Api_GetRepostList = "https://graph.qq.com/t/get_repost_list"
	QC_Api_GetOtherInfo = "https://graph.qq.com/t/get_other_info"

	/*----relation----*/
	QC_Api_GetFansList = "https://graph.qq.com/relation/get_fanslist"
	QC_Api_GetIdolList = "https://graph.qq.com/relation/get_idollist"
	QC_Api_AddIdol = "https://graph.qq.com/relation/add_idol"
	QC_Api_DelIdol = "https://graph.qq.com/relation/del_idol"

	/*----Tenpay----*/
	QC_Api_GetTenpayAddr = "https://graph.qq.com/cft_info/get_tenpay_addr"
)

type QC struct {
	AppId string
	AccessToken string
	OpenId string
}

func NewQC(appId string, accessToken string, openId string) *QC {
	return &QC{
		AppId: appId,
		AccessToken: accessToken,
		OpenId: openId,
	}
}

// return public params
func (qc *QC) makeParams(privateParams ...map[string]string) map[string]string {
	params := map[string]string{
		"access_token": qc.AccessToken,
		"oauth_consumer_key": qc.AppId,
		"openid": qc.OpenId,
	}
	if len(privateParams) > 0 {
		for key, value := range privateParams[0]  {
			params[key] = value
		}
	}
	return params
}

// get user info
func (qc *QC) GetUserInfo() (body string, code int, err error) {
	params := qc.makeParams()
	body, code, err = NewUtils().HttpGet(QC_Api_GetUserInfo, params, nil)
	return
}

// get user simple info
func (qc *QC) GetSimpleUserInfo() (body string, code int, err error) {
	params := qc.makeParams()
	body, code, err = NewUtils().HttpGet(QC_Api_GetSimpleUserInfo, params, nil)
	return
}

// get vip info
func (qc *QC) GetVipInfo(format string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_GetVipInfo, params, nil)
	return
}

// get vip rich info
func (qc *QC) GetVipRichInfo(format string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_GetVipRichInfo, params, nil)
	return
}

// photo list album
func (qc *QC) ListAlbum(format string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_ListAlbum, params, nil)
	return
}

// add album
func (qc *QC) AddAlbum(albumName string, albumDesc string, priv string, format string) (body string, code int, err error) {
	privateParams := map[string]string{
		"albumname": albumName,
		"albumdesc": albumDesc,
		"priv": priv,
		"format": strings.ToLower(format),
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_AddAlbum, params, nil)
	return
}

// list photo
func (qc *QC) ListPhoto(albumId string, format string) (body string, code int, err error) {
	privateParams := map[string]string{
		"albumid": albumId,
		"format": strings.ToLower(format),
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_ListPhoto, params, nil)
	return
}

// check page fans
func (qc *QC) CheckPageFans (format string, pageId string) (body string, code int, err error) {
	privateParams := map[string]string{
		"page_id": pageId,
		"format": strings.ToLower(format),
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_CheckPageFans, params, nil)
	return
}

// get user weibo info
func (qc *QC) GetInfo (format string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_GetWeiboInfo, params, nil)
	return
}

// add a weibo
func (qc *QC) AddT (format string, content string, clientIp string, longitude string, latitude string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"content": content,
		"clientip": clientIp,
		"longitude": longitude,
		"latitude": latitude,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpPost(QC_Api_AddT, params, nil)
	return
}

// delete a weibo
func (qc *QC) DelT (format string, id string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"id": id,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpPost(QC_Api_DeleteT, params, nil)
	return
}

// add a pic weibo
func (qc *QC) AddPicT (format string, content string, pic string, clientIp string, longitude string, latitude string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"content": content,
		"clientip": clientIp,
		"longitude": longitude,
		"latitude": latitude,
		"pic": pic,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpPost(QC_Api_AddPicT, params, nil)
	return
}

// get repost list
func (qc *QC) GetRepostList (format string, flag string, rootId string, pageFlag string, pageTime string, reqNum string, twitterId string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"flag": flag,
		"rootid": rootId,
		"pageflag": pageFlag,
		"pagetime": pageTime,
		"reqnum": reqNum,
		"twitterid": twitterId,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_GetRepostList, params, nil)
	return
}


// get other info
func (qc *QC) GetOtherInfo (format string, name string, fopenId string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"name": name,
		"fopenid": fopenId,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_GetOtherInfo, params, nil)
	return
}

// get fans list
func (qc *QC) GetFansList  (format string, reqNum string, startIndex string, mode string, install string, sex string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"reqnum": reqNum,
		"startindex": startIndex,
		"mode": mode,
		"install": install,
		"sex": sex,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_GetFansList, params, nil)
	return
}

// get idol list
func (qc *QC) GetIdolList (format string, reqNum string, startIndex string, mode string, install string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"reqnum": reqNum,
		"startindex": startIndex,
		"mode": mode,
		"install": install,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpGet(QC_Api_GetIdolList, params, nil)
	return
}

// add idol
func (qc *QC) AddIdol (format string, name string, fopenIds string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"name": name,
		"fopenids": fopenIds,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpPost(QC_Api_AddIdol, params, nil)
	return
}

// delete a idol
func (qc *QC) DelIdol (format string, name string, fopenId string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"name": name,
		"fopenid": fopenId,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpPost(QC_Api_DelIdol, params, nil)
	return
}

// get Tenpay address
func (qc *QC) GetTenpayAddr (format string, offset string, limit string, ver string) (body string, code int, err error) {
	privateParams := map[string]string{
		"format": strings.ToLower(format),
		"offset": offset,
		"limit": limit,
		"ver": ver,
	}
	params := qc.makeParams(privateParams)
	body, code, err = NewUtils().HttpPost(QC_Api_GetTenpayAddr, params, nil)
	return
}