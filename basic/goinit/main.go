package main

//
//import (
//	"errors"
//	"io"
//	"path"
//	"regexp"
//	"strconv"
//	"sync"
//
//	"fmt"
//	"log"
//	"math/rand"
//	"net/http"
//	"net/url"
//	"strings"
//	"time"
//
//	"crypto/md5"
//	"encoding/json"
//
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/service/s3"
//	"github.com/go-sql-driver/mysql"
//	_ "github.com/go-sql-driver/mysql" // mysql
//
//	"github.com/go-gomail/gomail"
//)
//
//// ossSmith: AKIAUCJPMWS23TXDNMPI/45jnyO5tPakZgcNAA50o/g/yMOevZERgyAB3jY8Y, arn:aws-cn:iam::279809143989:user/ossSmith
//// ossReader:AKIAUCJPMWS2QXN4WVFA/u8lbDu0CUW3AuVCRI12pBF7WPrwzCqNWcu+RVDJy, arn:aws-cn:iam::279809143989:user/ossReader
//// https://wbz-pfile.s3.cn-north-1.amazonaws.com.cn/R-C+(1).jpg
//// https://wbz-pfile.s3.cn-north-1.amazonaws.com.cn/lianpiao/%E9%80%9A%E5%BE%80%E5%B1%B1%E9%A1%B6%E7%9A%84%E8%B7%AF.jpg
//// https://oss-wbz-zone-01-rn4hsnzw11owjcc3qjdxduk9ad4o1cnn1a-s3alias/R-C+(1).jpg
//
//const WBZ_S3_HOST = "https://pfs.lianpiao.net"
//const INTERNAL_APICALL_MAGCODE = "LP_AKcse83js16AHDcCXMVUFs"
//
//const _S3_URL_MAGCODE = "Nxm%&5x.*7N`#B/ c).}"
//const MYSQL_TIME_FORMAT = "2006-01-02 15:04:05"
//const MAX_DOWNLOAD_SIZE = 200 << 20 //200 MB
//
//const DOWNLOAD_URL_EXPIRED = 60 * 15  // seconds
//const DOWNLOAD_URL_EXPIRED_SHORT = 30 // seconds for directly upload, redis key expired
//const AWS_URL_EXPIRED = DOWNLOAD_URL_EXPIRED + 30
//const BLACKLIST_CHECK_MAXCOUNTPERTIME = 10 // 每错一次，跟踪时间重新计时，直到超过规定次数
//const BLACKLIST_CHECK_COUNTTIME = 30
//const BLACKLIST_CHECK_PAUSEDTIME = 15 * 60 // 停止访问30分钟
//
//type UpdoadRuler struct {
//	limitSize int64
//	timeInSec int64
//	alert     bool
//}
//
//type EmailInfo struct {
//	ServerHost string // ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
//	ServerPort int    // ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
//
//	FromEmail  string // FromEmail　发件人邮箱地址
//	FromPasswd string //发件人邮箱密码（注意，这里是明文形式)
//
//	Recipient []string //收件人邮箱
//	CC        []string //抄送
//}
//
///**
// * @Author: dcj
// * @Date: 2020-04-02 15:45:55
// * @Description: 发送邮件
// * @Param : subject[主题]、body[内容]、emailInfo[发邮箱需要的信息(参考EmailInfo)]
// * @Return:
// */
//func SendEmail(subject, body string, emailInfo *EmailInfo) {
//	if len(emailInfo.Recipient) == 0 {
//		log.Print("收件人列表为空")
//		return
//	}
//
//	emailMessage = gomail.NewMessage()
//	//设置收件人
//	emailMessage.SetHeader("To", emailInfo.Recipient...)
//	//设置抄送列表
//	if len(emailInfo.CC) != 0 {
//		emailMessage.SetHeader("Cc", emailInfo.CC...)
//	}
//	// 第三个参数为发件人别名，如"dcj"，可以为空（此时则为邮箱名称）
//	emailMessage.SetAddressHeader("From", emailInfo.FromEmail, "链票系统预警")
//
//	//主题
//	emailMessage.SetHeader("Subject", subject)
//
//	//正文
//	emailMessage.SetBody("text/html", body)
//
//	d := gomail.NewDialer(emailInfo.ServerHost, emailInfo.ServerPort,
//		emailInfo.FromEmail, emailInfo.FromPasswd)
//	err := d.DialAndSend(emailMessage)
//	if err != nil {
//		log.Println("发送邮件失败： ", err)
//	} else {
//		log.Println("已成功发送邮件到指定邮箱")
//	}
//}
//
//func SendAlertEmail(title, body string) {
//	reclist := []string{"richard.li@yodoo.net.cn", "liqt@webaozhang.com"} //收件人邮箱地址
//
//	// tls=yes
//	// ssl=yes
//
//	// from=链票<service@email.lianpiao.net>
//
//	info := &EmailInfo{
//		"smtpdm.aliyun.com",
//		465,
//		"service@email.lianpiao.net", //发件人邮箱地址
//		"jKxUb69VNDaGJh86",
//		reclist,
//		nil,
//	}
//
//	// SendEmail(title, body, info)
//}
//
//var (
//	uploadRulers []UpdoadRuler = []UpdoadRuler{
//		{limitSize: 4 * 500 * 1024 * 1024, timeInSec: 60, alert: true},                   //  500M in 60 sec.
//		{limitSize: 4 * 1024 * 1024 * 1024, timeInSec: 3600, alert: true},                //  1G in 1 hour
//		{limitSize: 4 * 3 * 1024 * 1024 * 1024, timeInSec: 3600 * 24, alert: true},       //  3GB in 1 day.
//		{limitSize: 4 * 10 * 1024 * 1024 * 1024, timeInSec: 3600 * 24 * 30, alert: true}, //  10GB in 30 days.
//	}
//	_remoteAddressReg *regexp.Regexp
//	emailMessage      *gomail.Message
//)
//
//func init() {
//	_remoteAddressReg, _ = regexp.Compile("[^a-zA-Z0-9]+")
//}
//
//func paladingRemoteAddr(hostAddr string) string {
//	return _remoteAddressReg.ReplaceAllString(strings.Split(hostAddr, ":")[0], ".")
//}
//
//type Resp struct {
//	Code int    `json:"code"`
//	Msg  string `json:"msg"`
//	Data string `json:"data,omitempty"`
//}
//
//func reqTrace(next http.HandlerFunc) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//		defer func() { log.Println(r.Method, r.URL.Path, time.Since(start)) }()
//
//		next.ServeHTTP(w, r)
//	})
//}
//
//func cors(f http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		origin := r.Header.Get("Origin")
//		log.Println("CORS: Origin =", origin, ", HTTP ", r.Method)
//
//		switch origin {
//		case "null",
//			"https://dev.lianpiao.net",
//			"https://main.lianpiao.net",
//			"https://main.atpiao.com",
//			"https://uat.lianpiao.net",
//			"https://pfs.lianpiao.net",
//			"https://servicewechat.com",
//			"https://lp-pfiles.s3.cn-north-1.amazonaws.com.cn",
//			"https://wbz-pfile.s3.cn-north-1.amazonaws.com.cn":
//			w.Header().Set("Access-Control-Allow-Origin", origin)
//			// default:
//			//      log.Println("Forbidden! Expect for Origin = http(s)://*.lianpiao.net/")
//			//      w.WriteHeader(http.StatusForbidden)
//			//      return
//		}
//
//		w.Header().Set("Access-Control-Allow-Headers", "Origin,application,language,X-Requested-With,X-Extra-Header,Content-Type,Accept,AccessToken,X-CSRF-Token,Authorization,Token,Cache-Control,Content-Language,errcallback,logintype,notoken,pragma") //header的类型
//		w.Header().Set("Access-Control-Allow-Credentials", "true")                                                                                                                                                                                         //设置为true，允许ajax异步请求带cookie信息
//		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")                                                                                                                                                          //允许请求方法
//		w.Header().Set("Access-Control-Expose-Headers", "Content-Length,application,language,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,errcallback,logintype,notoken,pragma")
//		// w.Header().Set("content-type", "application/json;charset=UTF-8")                                        //返回数据格式是json
//
//		if r.Method == "OPTIONS" {
//			w.WriteHeader(http.StatusNoContent)
//		} else {
//			f(w, r)
//		}
//	}
//}
//
//func main() {
//	defer sqlxdb.Close()
//	defer rdsClient.Close()
//
//	// launch http service
//	var wg sync.WaitGroup
//
//	go func() {
//		log.Println("Start listern on 51000 port...")
//		mux := http.NewServeMux()
//		mux.Handle("/getUploadUrl", reqTrace(http.HandlerFunc(handleS3UploadReq))) // for Java Put
//		mux.Handle("/getPostUrl", reqTrace(http.HandlerFunc(handleS3PostUrlReq)))  // for H5 Post
//		mux.Handle("/deleteS3File", reqTrace(http.HandlerFunc(handleS3DeleteReq)))
//		mux.Handle("/copyS3File", reqTrace(http.HandlerFunc(handleS3CopyReq)))
//		mux.Handle("/upload2S3", reqTrace(http.HandlerFunc(handleS3FileUpload))) // for java Upload immediately
//
//		wg.Add(1)
//		err := http.ListenAndServe(":51000", mux) // in local nerwork
//		if err != nil {
//			log.Println("Failed to listern in local ip, Error -", err)
//		}
//		wg.Done()
//	}()
//
//	func() {
//		log.Println("Start listern on 50000 port...")
//
//		mux := http.NewServeMux()
//		mux.Handle("/s3notify", reqTrace(http.HandlerFunc(handleAwsSNSS3Notification))) // SNS
//		mux.Handle("/awsS3Notify", reqTrace(http.HandlerFunc(handleWbzS3Notification))) // S3 Lambda callback in http
//		mux.Handle("/dl", reqTrace(http.HandlerFunc(handleS3DownloadReq)))
//		mux.Handle("/dl2", reqTrace(http.HandlerFunc(handleS3DownloadReqEx)))
//		mux.Handle("/uploaded", reqTrace(cors(http.HandlerFunc(handleUploadedCallback)))) // 必须有cors
//		s := &http.Server{
//			Addr:           ":50000", // to outer service
//			Handler:        mux,
//			ReadTimeout:    10 * time.Second,
//			WriteTimeout:   10 * time.Second,
//			MaxHeaderBytes: 1 << 20,
//		}
//
//		err := s.ListenAndServe()
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}()
//
//	wg.Wait()
//}
//
//func sendShortResp(w http.ResponseWriter, code int, msg string) {
//	sendResp(w, code, msg, "{}")
//}
//
//func sendResp(w http.ResponseWriter, code int, msg string, data string) {
//	const RESPONE_DATA_TAG = "_$IN_"
//
//	if data == "" {
//		data = RESPONE_DATA_TAG
//	}
//
//	var result Resp
//	result.Code = code
//	result.Msg = msg
//	result.Data = RESPONE_DATA_TAG
//	v, err := json.Marshal(result)
//	if err != nil {
//		log.Println("Send response error -", err)
//		w.WriteHeader(503)
//	} else {
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(200)
//		w.Write([]byte(strings.Replace(string(v), "\""+RESPONE_DATA_TAG+"\"", data, 1)))
//	}
//	// if err := json.NewEncoder(w).Encode(result); err != nil {
//	//      log.Println("Send response error -", err)
//	// }
//}
//
//type S3UploadToken struct {
//	S3Bucket    string         `json:"bucket"` //s3Object.s3.bucket.name
//	S3Key       string         `json:"key"`    //s3Object.s3.object.key
//	RefKey      string         `json:"refkey"`
//	S3Url       string         `json:"url"` // s3Object.s3.object.URLDecodedKey ? Signature
//	DownloadUrl string         `json:"downloadUrl"`
//	FileName    string         `json:"fileName"`
//	FileExt     string         `json:"fileExt"`
//	Length      int64          `json:"length"`
//	ExpireInSec int64          `json:"expired"`
//	RemoteAddr  string         `json:"remoteAddress"`
//	CreateTime  string         `json:"create_time"`
//	OwnerId     int64          `json:"ownerId"`
//	OwnerType   int64          `json:"ownerType"`
//	ProjectName string         `json:"projectName"`
//	StartTime   mysql.NullTime `json:"startTime,omitempty"`
//	EndTime     mysql.NullTime `json:"endTime,omitempty"`
//}
//
//func (i *S3UploadToken) MarshalBinary() ([]byte, error) {
//	return json.Marshal(i)
//}
//
//func handleUploadedCallback(w http.ResponseWriter, r *http.Request) {
//	refKey, project, _, err := checkDownloadUrlSignature("/uploaded", r.URL.Query())
//	if err != nil {
//		log.Println("Check s3 file url failed! -", err)
//		w.WriteHeader(403)
//		return
//	}
//
//	var file S3_Files
//	if err := file.readFromDBByRefkey(refKey, 1); err == nil {
//		log.Printf("Find refKey in database: s3Key =%s, refKey =%s", file.S3key, refKey)
//		sendResp(w, 0, "成功", fmt.Sprintf(`{"url":"%s"}`, file.Url))
//	} else {
//		pb := rdsClient.Subscribe(ctx, AWS_S3FILE_UPLOADED)
//		defer pb.Close()
//		log.Println("阻塞，等待读取 Redis Channel 信息")
//
//		for {
//			select {
//			case mg := <-pb.Channel():
//				// 等待从 channel 中发布 close 关闭服务
//				log.Println("Received POST file, ref key=", mg.Payload)
//				if mg.Payload == refKey {
//					downloadUrl := packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", refKey, project)
//					sendResp(w, 0, "成功", fmt.Sprintf(`{"url":"%s"}`, downloadUrl))
//					return
//				}
//			case <-time.After(45 * time.Second):
//				log.Println("POST action timeout in 45 seconds!! ref key=", refKey)
//				w.WriteHeader(408)
//				return
//			}
//		}
//	}
//}
//
//func getCurrentRemoteAddr(r *http.Request) string {
//	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
//	// 但是要注意的是这两个请求头代表的ip都有可能是伪造的
//	ip := r.Header.Get("X-Real-IP")
//	if ip == "" {
//		// 当请求头不存在即不存在代理时直接获取ip
//		ip = paladingRemoteAddr(r.RemoteAddr)
//
//	}
//	return ip
//}
//
//func checkBlacklist(r *http.Request) bool {
//	ipAddr := getCurrentRemoteAddr(r)
//
//	if strings.HasPrefix(ipAddr, "172.17.") {
//		return true
//	}
//
//	limitKey := fmt.Sprintf("s3:blacklist:%s:int", ipAddr)
//
//	if iCount, err := rdsClient.Get(ctx, limitKey).Int(); err != nil {
//		if err == redis.Nil {
//			return true
//		}
//		log.Println("checkBlacklist: System ERROR! Can not get redis data. key=", limitKey, ", err -", err.Error())
//		return true
//	} else {
//		log.Println("checkBlacklist: Found Blacklist data from redis. key=", limitKey, ", count=", iCount)
//		return iCount <= BLACKLIST_CHECK_MAXCOUNTPERTIME
//	}
//}
//
//func updateBlacklist(r *http.Request) bool {
//	ipAddr := getCurrentRemoteAddr(r)
//
//	if strings.HasPrefix(ipAddr, "172.17.") {
//		return true
//	}
//
//	limitKey := fmt.Sprintf("s3:blacklist:%s:int", ipAddr)
//
//	if iCount, err := rdsClient.Get(ctx, limitKey).Int(); err != nil {
//		if err == redis.Nil {
//			if err := rdsClient.Set(ctx, limitKey, 1, time.Duration(BLACKLIST_CHECK_COUNTTIME)*time.Second).Err(); err != nil {
//				log.Println("updateBlacklist: System ERROR! Can not set redis data. key=", limitKey, ", err -", err.Error())
//				return false
//			}
//			log.Println("updateBlacklist: Catched one mismatched API call. Set redis data into Blacklist, key=", limitKey)
//			return true
//		} else {
//			log.Println("updateBlacklist: System ERROR! Can not get redis data. key=", limitKey, ", err -", err.Error())
//			return false
//		}
//	} else {
//		if iCount < BLACKLIST_CHECK_MAXCOUNTPERTIME {
//			if err := rdsClient.Set(ctx, limitKey, iCount+1, time.Duration(BLACKLIST_CHECK_COUNTTIME)*time.Second).Err(); err != nil {
//				log.Println("updateBlacklist: System ERROR! Can not update redis data. err -", err.Error())
//				return false
//			}
//			log.Println("updateBlacklist: Catched another mismatched API call. Set redis data into Blacklist, key=", limitKey, ", count=", iCount+1)
//		} else {
//			// redis.KeepTTL
//			if err := rdsClient.Set(ctx, limitKey, iCount+1, time.Duration(BLACKLIST_CHECK_PAUSEDTIME)*time.Second).Err(); err != nil {
//				SendAlertEmail("链票告警服务:单位时间内连续发生多次非法下载请求",
//					fmt.Sprintf("<h1>告警类型：暴力破解资源的下载地址</h1><p>发生非法下载请求的次数: %d次/%d秒</p><p>发生时刻:%s</p><p>来源 redis key=%s</p><p>启动拒绝服务的时长: %d 秒",
//						iCount, BLACKLIST_CHECK_COUNTTIME, time.Now().Local().Format(MYSQL_TIME_FORMAT), limitKey, BLACKLIST_CHECK_PAUSEDTIME))
//				log.Println("updateBlacklist: System ERROR! Can not update redis data. err -", err.Error())
//				return false
//			}
//			log.Println("updateBlacklist: Paused remoteaddr from Blacklist! key= ", limitKey, ", count=", iCount+1)
//		}
//		return true
//	}
//}
//
//func handleS3DownloadReq(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "GET" {
//		w.WriteHeader(405)
//		return
//	}
//
//	if !checkBlacklist(r) {
//		log.Println("handleS3DownloadReq: Remote address is in Blacklist!! -", getCurrentRemoteAddr(r))
//		w.WriteHeader(403)
//		return
//	}
//
//	refKey, _, metadata, err := checkDownloadUrlSignature("/dl", r.URL.Query())
//	if err != nil {
//		log.Println("handleS3DownloadReq: Check s3 file url failed! -", err.Error())
//		updateBlacklist(r)
//
//		w.WriteHeader(403)
//		return
//	}
//
//	params := r.URL.Query()
//
//	externalApiCallMark := params.Get("exno_c") != INTERNAL_APICALL_MAGCODE
//	log.Printf("dl headers: %+v\n", r.Header)
//
//	refHead := r.Header.Get("Referer")
//	log.Println("Referer string =", refHead)
//	log.Printf("metadata = %q\n", metadata)
//
//	var url string
//	var mimeType string
//	if metadata != "" {
//		url, err = getRefUrlFromRedis(refKey + ":" + metadata)
//		mimeType, _ = getRefUrlTypeFromRedis(refKey + ":" + metadata)
//	} else {
//		url, err = getRefUrlFromRedis(refKey)
//		mimeType, _ = getRefUrlTypeFromRedis(refKey)
//	}
//	log.Printf("mimeType %q for %q\n", mimeType, url)
//
//	if externalApiCallMark && strings.HasPrefix(mimeType, "image/") {
//		if !strings.HasPrefix(refHead, "https://uat.lianpiao.net/") &&
//			!strings.HasPrefix(refHead, "https://main.lianpiao.net/") &&
//			!strings.HasPrefix(refHead, "https://uat.atpiao.com/") &&
//			!strings.HasPrefix(refHead, "https://main.atpiao.com/") &&
//			!strings.HasPrefix(refHead, "https://dev.lianpiao.net/") &&
//			!strings.HasPrefix(refHead, "https://dev.atpiao.com/") &&
//			!strings.HasPrefix(refHead, "https://dev.webaozhang.com/") &&
//			!strings.HasPrefix(refHead, "https://uat.webaozhang.com/") &&
//			!strings.HasPrefix(refHead, "https://service.webaozhang.com/") &&
//			!strings.HasPrefix(refHead, "https://servicewechat.com/") {
//			log.Printf("Access forbidden to %q from %q outside lianpiao website!!\n", refKey, refHead)
//			w.WriteHeader(403)
//			return
//		}
//	}
//
//	if err == nil && url != "" {
//		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
//	} else {
//		var u S3_Files
//		if err := u.readFromDBByRefkey(refKey, 1); err != nil {
//			w.WriteHeader(404)
//			return
//		}
//
//		log.Printf("mimeType %q from MYSQL for %q\n", u.S3mimeType, url)
//
//		if externalApiCallMark && strings.HasPrefix(u.S3mimeType, "image/") {
//			if !strings.HasPrefix(refHead, "https://uat.lianpiao.net/") &&
//				!strings.HasPrefix(refHead, "https://main.lianpiao.net/") &&
//				!strings.HasPrefix(refHead, "https://uat.atpiao.com/") &&
//				!strings.HasPrefix(refHead, "https://main.atpiao.com/") &&
//				!strings.HasPrefix(refHead, "https://dev.lianpiao.net/") &&
//				!strings.HasPrefix(refHead, "https://dev.atpiao.com/") &&
//				!strings.HasPrefix(refHead, "https://dev.webaozhang.com/") &&
//				!strings.HasPrefix(refHead, "https://uat.webaozhang.com/") &&
//				!strings.HasPrefix(refHead, "https://service.webaozhang.com/") &&
//				!strings.HasPrefix(refHead, "https://servicewechat.com/") {
//				log.Printf("Access forbidden to %q from %q outside lianpiao website!!\n", refKey, refHead)
//				w.WriteHeader(403)
//				return
//			}
//		}
//
//		currentTime := time.Now().UTC()
//		if u.StartTime.Valid {
//			if currentTime.Sub(u.StartTime.Time) < 0 {
//				log.Println("handleS3DownloadReq: request failed! url will open until", u.StartTime.Time)
//				w.WriteHeader(404)
//				return
//			}
//		}
//		if u.EndTime.Valid {
//			if currentTime.Sub(u.EndTime.Time) > 0 {
//				log.Println("handleS3DownloadReq: request failed! url has closed at", u.EndTime.Time)
//				w.WriteHeader(404)
//				return
//			}
//		}
//
//		// url, err := createS3DownloadUrl(u.S3bucket, u.S3key, u.S3mimeType, AWS_URL_EXPIRED)
//		// if err != nil {
//		//      log.Println("handleS3DownloadReq: Cannot create download URL, ERROR -", err.Error())
//		//      w.WriteHeader(500)
//		// } else {
//		//      log.Println("handleS3DownloadReq: Download URL will be expired in 120 seconds - ", url)
//		//      http.Redirect(w, r, url, http.StatusTemporaryRedirect)
//
//		//      if metadata != "" {
//		//              saveRefUrl2Redis(refKey+":"+metadata, url, DOWNLOAD_URL_EXPIRED)
//		//      } else {
//		//      saveRefUrl2Redis(refKey, url, DOWNLOAD_URL_EXPIRED)
//		//      }
//		// }
//		url, err := createS3DownloadUrl(u.S3bucket, u.S3key, u.S3mimeType, AWS_URL_EXPIRED)
//		if err != nil {
//			log.Println("handleS3DownloadReq: Cannot create download URL, ERROR -", err.Error())
//			w.WriteHeader(500)
//		} else {
//			log.Println("handleS3DownloadReq: Download URL will be expired in 120 seconds - ", url)
//			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
//
//			if metadata != "" {
//				saveRefUrl2Redis(refKey+":"+metadata, url, DOWNLOAD_URL_EXPIRED)
//				saveRefUrlType2Redis(refKey+":"+metadata, u.S3mimeType, DOWNLOAD_URL_EXPIRED)
//			} else {
//				saveRefUrl2Redis(refKey, url, DOWNLOAD_URL_EXPIRED)
//				saveRefUrlType2Redis(refKey, u.S3mimeType, DOWNLOAD_URL_EXPIRED)
//			}
//		}
//	}
//}
//
//func CopyHttpData(url string, w http.ResponseWriter, maxLength int64) error {
//	// Get the data
//	resp, err := http.Get(url)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// 然后将响应流和文件流对接起来
//	// _, err = io.Copy(out, resp.Body)
//	size, err := io.Copy(w, io.LimitReader(resp.Body, maxLength))
//	if err != nil {
//		return err
//	}
//	if size >= maxLength {
//		log.Println("downloading file size =", size, "exceed", maxLength)
//		return fmt.Errorf("要下载的文件超过了允许的大小(%d)", maxLength)
//	}
//
//	return nil
//}
//
//func handleS3DownloadReqEx(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "GET" {
//		w.WriteHeader(405)
//		return
//	}
//
//	if !checkBlacklist(r) {
//		log.Println("handleS3DownloadReq: Remote address is in Blacklist!! -", getCurrentRemoteAddr(r))
//		w.WriteHeader(403)
//		return
//	}
//
//	refKey, _, metadata, err := checkDownloadUrlSignature("/dl", r.URL.Query())
//	if err != nil {
//		log.Println("handleS3DownloadReq: Check s3 file url failed! -", err.Error())
//		updateBlacklist(r)
//
//		w.WriteHeader(403)
//		return
//	}
//
//	params := r.URL.Query()
//	externalApiCallMark := params.Get("exno_c") != INTERNAL_APICALL_MAGCODE
//
//	log.Printf("dl headers: %+v\n", r.Header)
//	refHead := r.Header.Get("Referer")
//	log.Println("Referer string =", refHead)
//	log.Printf("metadata = %q\n", metadata)
//
//	var url string
//	var mimeType string
//	if metadata != "" {
//		url, err = getRefUrlFromRedis(refKey + ":" + metadata)
//		mimeType, _ = getRefUrlTypeFromRedis(refKey + ":" + metadata)
//	} else {
//		url, err = getRefUrlFromRedis(refKey)
//		mimeType, _ = getRefUrlTypeFromRedis(refKey)
//	}
//	log.Printf("mimeType %q for %q\n", mimeType, url)
//
//	if externalApiCallMark && strings.HasPrefix(mimeType, "image/") {
//		if !strings.HasPrefix(refHead, "https://uat.lianpiao.net/") &&
//			!strings.HasPrefix(refHead, "https://main.lianpiao.net/") &&
//			!strings.HasPrefix(refHead, "https://uat.atpiao.com/") &&
//			!strings.HasPrefix(refHead, "https://main.atpiao.com/") &&
//			!strings.HasPrefix(refHead, "https://dev.lianpiao.net/") &&
//			!strings.HasPrefix(refHead, "https://dev.atpiao.com/") &&
//			!strings.HasPrefix(refHead, "https://dev.webaozhang.com/") &&
//			!strings.HasPrefix(refHead, "https://uat.webaozhang.com/") &&
//			!strings.HasPrefix(refHead, "https://service.webaozhang.com/") &&
//			!strings.HasPrefix(refHead, "https://servicewechat.com/") {
//			log.Printf("Access forbidden to %q from %q outside lianpiao website!!\n", refKey, refHead)
//			w.WriteHeader(403)
//			return
//		}
//	}
//
//	if err != nil || url == "" {
//		var u S3_Files
//		if err := u.readFromDBByRefkey(refKey, 1); err != nil {
//			w.WriteHeader(404)
//			return
//		}
//
//		if externalApiCallMark && strings.HasPrefix(u.S3mimeType, "image/") {
//			if !strings.HasPrefix(refHead, "https://uat.lianpiao.net/") &&
//				!strings.HasPrefix(refHead, "https://main.lianpiao.net/") &&
//				!strings.HasPrefix(refHead, "https://uat.atpiao.com/") &&
//				!strings.HasPrefix(refHead, "https://main.atpiao.com/") &&
//				!strings.HasPrefix(refHead, "https://dev.lianpiao.net/") &&
//				!strings.HasPrefix(refHead, "https://dev.atpiao.com/") &&
//				!strings.HasPrefix(refHead, "https://dev.webaozhang.com/") &&
//				!strings.HasPrefix(refHead, "https://uat.webaozhang.com/") &&
//				!strings.HasPrefix(refHead, "https://service.webaozhang.com/") &&
//				!strings.HasPrefix(refHead, "https://servicewechat.com/") {
//				log.Printf("Access forbidden to %q from %q outside lianpiao website!!\n", refKey, refHead)
//				w.WriteHeader(403)
//				return
//			}
//		}
//
//		currentTime := time.Now().UTC()
//		if u.StartTime.Valid {
//			if currentTime.Sub(u.StartTime.Time) < 0 {
//				log.Println("handleS3DownloadReq: request failed! url will open until", u.StartTime.Time)
//				w.WriteHeader(404)
//				return
//			}
//		}
//		if u.EndTime.Valid {
//			if currentTime.Sub(u.EndTime.Time) > 0 {
//				log.Println("handleS3DownloadReq: request failed! url has closed at", u.EndTime.Time)
//				w.WriteHeader(404)
//				return
//			}
//		}
//
//		url, err = createS3DownloadUrl(u.S3bucket, u.S3key, u.S3mimeType, AWS_URL_EXPIRED)
//		if err != nil || url == "" {
//			log.Println("handleS3DownloadReq: Cannot create download URL, ERROR -", err.Error())
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		log.Println("handleS3DownloadReq: Download URL will be expired in 120 seconds - ", url)
//		// http.Redirect(w, r, url, http.StatusTemporaryRedirect)
//
//		if metadata != "" {
//			saveRefUrl2Redis(refKey+":"+metadata, url, DOWNLOAD_URL_EXPIRED)
//			saveRefUrlType2Redis(refKey+":"+metadata, u.S3mimeType, DOWNLOAD_URL_EXPIRED)
//		} else {
//			saveRefUrl2Redis(refKey, url, DOWNLOAD_URL_EXPIRED)
//			saveRefUrlType2Redis(refKey, u.S3mimeType, DOWNLOAD_URL_EXPIRED)
//		}
//	}
//
//	if err = CopyHttpData(url, w, MAX_DOWNLOAD_SIZE); err != nil {
//		log.Println("dl failed! Error:", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//	} else {
//		w.WriteHeader(http.StatusOK)
//	}
//}
//
//func handleS3FileUpload(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		w.WriteHeader(405)
//		return
//	}
//
//	params := r.URL.Query()
//	log.Println("Query = ", params)
//
//	expired, err := strconv.ParseInt(params.Get("expired"), 10, 64)
//	if err != nil {
//		// io.WriteString(w, `{"code": 1000, "message":"未指定接口过期时间"}`)
//		sendShortResp(w, -1001, "未指定接口过期时间")
//		return
//	}
//	if expired <= 0 {
//		expired = 60
//	}
//
//	project := strings.ToUpper(strings.Trim(params.Get("proj"), " "))
//	// userId := params.Get("ownerid")
//	ownerId, err := strconv.ParseInt(params.Get("ownerid"), 10, 64)
//	if ownerId < 0 || err != nil || project != "LP" && project != "WBZ" && project != "ATP" {
//		// io.WriteString(w, `{"code": 1000, "message":"缺少用户ID, 文件名, 或项目名"}`)
//		sendShortResp(w, -1003, "无效用户ID, 或项目名")
//		return
//	}
//
//	ownerType, err := strconv.ParseInt(params.Get("ownerType"), 10, 64)
//	if ownerType < 0 || err != nil {
//		ownerType = 0
//	}
//
//	var startTime, endTime time.Time
//	startTimeString := strings.Trim(params.Get("startTime"), " ")
//	endTimeString := strings.Trim(params.Get("endTime"), " ")
//	// loc, _ := time.LoadLocation("Asia/Shanghai")
//	loc, err := time.LoadLocation("Asia/Shanghai") //设置时区
//	if err != nil {
//		loc = time.FixedZone("CST", 8*3600)
//	}
//	if startTimeString != "" {
//		if t, err := time.ParseInLocation(MYSQL_TIME_FORMAT, startTimeString, loc); err != nil {
//			sendShortResp(w, -1003, "错误的起始时间格式")
//			return
//		} else {
//			startTime = t
//		}
//	}
//	if endTimeString != "" {
//		if t, err := time.ParseInLocation(MYSQL_TIME_FORMAT, endTimeString, loc); err != nil {
//			sendShortResp(w, -1003, "错误的停止服务时间格式")
//			return
//		} else {
//			endTime = t
//		}
//	}
//
//	const maxUploadSize = 100 << 20 // 100 MB
//	// const uploadPath = "./tmp"
//	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
//		log.Printf("Could not parse multipart form: %v\n", err)
//		sendShortResp(w, -1001, "CANT_PARSE_FORM") //http.StatusInternalServerError
//		return
//	}
//	fileType := r.PostFormValue("type")
//	file, fileHeader, err := r.FormFile("uploadFile")
//	if err != nil {
//		log.Printf("Could not parse multipart form: %v\n", err)
//		sendShortResp(w, -1001, "INVALID_FILE") //, http.StatusBadRequest)
//		return
//	}
//	defer file.Close()
//
//	fileSize := fileHeader.Size
//	fileName := fileHeader.Filename
//	log.Printf("File size (bytes): %v, type: %s, file: %s\n", fileSize, fileType, fileName)
//	if fileSize > maxUploadSize {
//		log.Printf("File size exceeds max length: %d\n", fileSize)
//		sendShortResp(w, -1001, "FILE_TOO_BIG") // http.StatusBadRequest
//		return
//	}
//	fileName = strings.ReplaceAll(fileName, "+", "_")
//
//	if !checkUploadRulers(ownerId, ownerType, project, fileSize, fileName) {
//		sendShortResp(w, -1005, "申请上传的文件大小超出用户当前可用存储容量的上限")
//		return
//	}
//
//	_, mimeType := getContentType(fileName)
//
//	// s3Key := fmt.Sprintf("%X/%d/%s", md5.Sum([]byte(project+fileName+strconv.FormatInt(time.Now().Unix(), 16))), ownerId, path.Base(fileName))
//	// refKey := fmt.Sprintf("%X/%s", md5.Sum([]byte(s3Key)), path.Base(fileName))
//	s3Key, refKey := generateFileKeys(project+fileName, fileName, ownerId, ownerType)
//
//	var s3Bucket, downloadUrl string // := "s3-lp-zone-01-eno9s1x41q96cr94wmzh7jsabio1rcnn1a-s3alias"
//
//	switch project {
//	case "LP", "ATP":
//		s3Bucket = "lp-pfiles" //"s3-lp-zone-01-eno9s1x41q96cr94wmzh7jsabio1rcnn1a-s3alias"
//		downloadUrl = packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", refKey, project)
//	case "WBZ":
//		s3Bucket = "wbz-pfile" //"oss-wbz-zone-01-rn4hsnzw11owjcc3qjdxduk9ad4o1cnn1a-s3alias"
//		downloadUrl = packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", refKey, project)
//	default:
//		// io.WriteString(w, `{"code": 1000, "message":"非法的项目名称"}`)
//		sendShortResp(w, -1004, "项目不存在")
//		return
//	}
//
//	var s3Token = &S3UploadToken{
//		S3Bucket: s3Bucket,
//		S3Key:    s3Key,
//		RefKey:   refKey,
//		// S3Url:       "", // upload url
//		DownloadUrl: downloadUrl,
//		Length:      fileSize,
//		ExpireInSec: expired,
//		RemoteAddr:  r.RemoteAddr,
//		CreateTime:  time.Now().UTC().Format(MYSQL_TIME_FORMAT),
//		OwnerId:     ownerId,
//		OwnerType:   ownerType,
//		ProjectName: project,
//		FileName:    fileName,
//		FileExt:     mimeType,
//	}
//
//	if startTime != endTime {
//		s3Token.StartTime.Scan(startTime)
//		s3Token.EndTime.Scan(endTime)
//	}
//
//	if err := rdsClient.Set(ctx, s3Key, s3Token, time.Duration(expired)*time.Second).Err(); err != nil {
//		log.Printf("handleS3FileUpload: redis Error :%v", err)
//		sendShortResp(w, -1005, "系统错误(redis)")
//		return
//	}
//
//	// pb := rdsClient.Subscribe(ctx, AWS_S3FILE_UPLOADED)
//	// defer pb.Close()
//	// log.Println("handleS3FileUpload: 阻塞，等待读取 Redis Channel 信息")
//
//	if err := uploadS3FileTarget(s3Bucket, s3Key, file); err != nil {
//		sendShortResp(w, -1005, "失败(s3)")
//		return
//	}
//
//	currentTime := time.Now().UTC()
//	if s3Token.StartTime.Valid {
//		if currentTime.Sub(s3Token.StartTime.Time) < 0 {
//			data := fmt.Sprintf(`{"url":"%s", "length":%d, "file":"%s", "expired":%d}`, downloadUrl, fileSize, fileName, expired)
//			log.Printf("handleS3FileUpload: Redis = %+v\nS3 file's info:%s\n", s3Token, data)
//			sendResp(w, 0, "成功", data)
//			return
//		}
//	}
//	if s3Token.EndTime.Valid {
//		if currentTime.Sub(s3Token.EndTime.Time) > 0 {
//			data := fmt.Sprintf(`{"url":"%s", "length":%d, "file":"%s", "expired":%d}`, downloadUrl, fileSize, fileName, expired)
//			log.Printf("handleS3FileUpload: Redis = %+v\nS3 file's info:%s\n", s3Token, data)
//			sendResp(w, 0, "成功", data)
//			return
//		}
//	}
//
//	// url, err := createS3DownloadUrl(s3Bucket, s3Key, mimeType, AWS_URL_EXPIRED)
//	// if err != nil {
//	//      log.Println("handleS3FileUpload: Cannot create download URL, ERROR -", err.Error())
//	//      sendShortResp(w, -1005, "系统错误(s3)")
//	// } else {
//	//      log.Println("handleS3FileUpload: Download URL will be expired in 120 seconds - ", url)
//	//      if err := saveRefUrl2Redis(refKey, url, DOWNLOAD_URL_EXPIRED_SHORT); err != nil {
//	//              log.Println("handleS3FileUpload: POST action timeout in 45 seconds!! ref key=", refKey)
//	//              sendShortResp(w, -1005, "系统错误(redis)")
//	//      } else {
//	//              data := fmt.Sprintf(`{"url":"%s", "length":%d, "file":"%s", "expired":%d}`, downloadUrl, fileSize, fileName, expired)
//	//              log.Printf("handleS3FileUpload: Redis = %+v\nS3 file's info:%s\n", s3Token, data)
//	//              sendResp(w, 0, "成功", data)
//	//      }
//	// }
//
//	url, err := createS3DownloadUrl(s3Bucket, s3Key, mimeType, AWS_URL_EXPIRED)
//	if err != nil {
//		log.Println("handleS3FileUpload: Cannot create download URL, ERROR -", err.Error())
//		sendShortResp(w, -1005, "系统错误(s3)")
//	} else {
//		log.Println("handleS3FileUpload: Download URL will be expired in 120 seconds - ", url)
//		if err := saveRefUrl2Redis(refKey, url, DOWNLOAD_URL_EXPIRED_SHORT); err != nil {
//			log.Println("handleS3FileUpload: POST action timeout in 45 seconds!! ref key=", refKey)
//			sendShortResp(w, -1005, "系统错误(redis)")
//		} else {
//			saveRefUrlType2Redis(refKey, mimeType, DOWNLOAD_URL_EXPIRED_SHORT)
//
//			data := fmt.Sprintf(`{"url":"%s", "length":%d, "file":"%s", "expired":%d}`, downloadUrl, fileSize, fileName, expired)
//			log.Printf("handleS3FileUpload: Redis = %+v\nS3 file's info:%s\n", s3Token, data)
//			sendResp(w, 0, "成功", data)
//		}
//	}
//
//	// for {
//	//      select {
//	//      case mg := <-pb.Channel():
//	//              // 等待从 channel 中发布 close 关闭服务
//	//              log.Println("handleS3FileUpload: Received POST file, ref key=", mg.Payload)
//	//              if mg.Payload == refKey {
//	//                      data := fmt.Sprintf(`{"url":"%s", "length":%d, "file":"%s", "expired":%d}`, downloadUrl, fileSize, fileName, expired)
//	//                      log.Printf("handleS3FileUpload: Redis = %+v\nS3 file's info:%s\n", s3Token, data)
//	//                      sendResp(w, 0, "成功", data)
//	//                      return
//	//              }
//	//      case <-time.After(60 * time.Second):
//	//              log.Println("handleS3FileUpload: POST action timeout in 45 seconds!! ref key=", refKey)
//	//              sendShortResp(w, -1005, "系统错误(S3服务超时)")
//	//              return
//	//      }
//	// }
//
//	// tm := time.Now().UTC()
//	// fileData := &S3_Files{
//	//      CreateTime: tm,
//	//      UpdateTime: tm,
//
//	//      OwnerId:     ownerId,
//	//      ProjectName: project, // 1:lp, 2:wbz
//	//      FileName:    fileName,
//	//      Url:         downloadUrl,
//
//	//      S3bucket:     s3Bucket,
//	//      S3key:        s3Key,
//	//      S3mimeType:   mimeType,
//	//      S3size:       fileSize,
//	//      S3Action:     "ObjectCreated:Put",
//	//      S3time:       tm,
//	//      S3remoteAddr: r.RemoteAddr,
//	//      Status:       1, //
//	//      RefKey:       refKey,
//	// }
//	// if startTime != endTime {
//	//      fileData.StartTime.Scan(startTime)
//	//      fileData.EndTime.Scan(endTime)
//	// }
//
//	// if err := fileData.saveDB(); err != nil {
//	//      sendShortResp(w, -1005, "系统错误(mysql)")
//	// } else {
//	//      // io.WriteString(w, fmt.Sprintf(`{"code": 0, "message":"成功", "data":{"url":"%s", "length":%d, "file":"%s", "expired":%d}}`, str, s3Length, fileName, expired))
//	//      data := fmt.Sprintf(`{"url":"%s", "length":%d, "file":"%s", "expired":%d}`, downloadUrl, fileSize, fileName, expired)
//	//      log.Printf("Redis = %+v\nS3 file's info:%s\n", s3Token, data)
//	//      sendResp(w, 0, "成功", data)
//	// }
//}
//
//func handleS3CopyReq(w http.ResponseWriter, r *http.Request) {
//	// if r.Method != "GET" {
//	//      w.WriteHeader(405)
//	//      return
//	// }
//	var err error
//
//	params := r.URL.Query()
//	log.Println("Query = ", params)
//
//	var ownerId, targetId int64
//	if ownerId, err = strconv.ParseInt(params.Get("ownerid"), 10, 64); err != nil || ownerId < 0 {
//		// io.WriteString(w, `{"code": 1000, "message":"未指定接口过期时间"}`)
//		sendShortResp(w, -1001, "非法用户")
//		return
//	}
//	ownerType, err := strconv.ParseInt(params.Get("ownerType"), 10, 64)
//	if ownerType < 0 || err != nil {
//		ownerType = 0
//	}
//
//	if targetId, err = strconv.ParseInt(params.Get("targetid"), 10, 64); err != nil || targetId < 0 {
//		sendShortResp(w, -1001, "缺少参数或参数不合法")
//		return
//	}
//	targetType, err := strconv.ParseInt(params.Get("targetType"), 10, 64)
//	if targetType < 0 || err != nil {
//		targetType = 0
//	}
//
//	project := strings.ToUpper(params.Get("proj"))
//	refKey := params.Get("key")
//	if project == "" || refKey == "" {
//		sendShortResp(w, -1001, "项目名或文件key值为空")
//		return
//	}
//
//	var ref_file S3_Files
//
//	if err = ref_file.readFromDBByRefkey(refKey, 1); err != nil {
//		sendShortResp(w, -1002, "无效的文件")
//	} else if ref_file.OwnerId != ownerId && ref_file.OwnerType != ownerType { //&& ref_file.ProjectName == project {
//		sendShortResp(w, -1003, "不能拷贝其他人的文件")
//	} else if ref_file.OwnerId == targetId && ref_file.OwnerType == targetType && ref_file.ProjectName == project {
//		// sendShortResp(w, -1004, "目标用户与当前用户是同一个")
//		sendResp(w, 0, "成功", fmt.Sprintf(`{"url":"%s"}`, ref_file.Url))
//	} else {
//		// if s3_file, err := readAllFromDBByS3key(ref_file.S3key, 1); err != nil {
//		//      sendShortResp(w, -1001, "系统错误")
//		// } else {
//		//      for _, value := range *s3_file {
//		//              if value.OwnerId == targetId && value.OwnerType == targetType && value.ProjectName == project {
//		//                      // sendShortResp(w, -1004, "目标用户已经拥有过该文件")
//		//                      sendResp(w, 0, "成功", fmt.Sprintf(`{"url":"%s"}`, value.Url))
//		//                      return
//		//              }
//		//      }
//
//		// s3Key := fmt.Sprintf("%X/%d/%s", md5.Sum([]byte(project+ref_file.S3key+strconv.FormatInt(time.Now().Unix(), 16))), targetId, path.Base(ref_file.FileName))
//		// ref_file.RefKey = fmt.Sprintf("%X/%s", md5.Sum([]byte(s3Key)), path.Base(ref_file.FileName))
//		_, refKey := generateFileKeys(project+ref_file.S3key, ref_file.FileName, targetId, targetType)
//
//		ref_file.RefKey = refKey
//		ref_file.OriginId = ref_file.ID
//		ref_file.OwnerId = targetId
//		ref_file.OwnerType = targetType
//		ref_file.ProjectName = project
//		ref_file.Url = packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", ref_file.RefKey, project)
//		if err := ref_file.saveDB(); err == nil {
//			sendResp(w, 0, "成功", fmt.Sprintf(`{"url":"%s"}`, ref_file.Url))
//		} else {
//			sendShortResp(w, -1005, "失败")
//		}
//		// }
//	}
//}
//
//func handleS3DeleteReq(w http.ResponseWriter, r *http.Request) {
//	// if r.Method != "GET" {
//	//      w.WriteHeader(405)
//	//      return
//	// }
//
//	params := r.URL.Query()
//	log.Println("Query = ", params)
//
//	project := strings.ToUpper(params.Get("proj"))
//	refKey := params.Get("key")
//
//	ownerId, err := strconv.ParseInt(params.Get("ownerid"), 10, 64)
//	if err != nil || ownerId < 0 {
//		// io.WriteString(w, `{"code": 1000, "message":"未指定接口过期时间"}`)
//		sendShortResp(w, -1001, "非法用户")
//		return
//	}
//	ownerType, err := strconv.ParseInt(params.Get("ownerType"), 10, 64)
//	if ownerType < 0 || err != nil {
//		ownerType = 0
//	}
//
//	if project == "" || refKey == "" {
//		sendShortResp(w, -1002, "项目名, 或文件key值为空")
//		return
//	}
//
//	var file S3_Files
//	if err := file.readFromDBByRefkeyAndProj(refKey, project, ownerId, ownerType, 1); err != nil {
//		sendShortResp(w, -1001, "文件不存在")
//		return
//	}
//
//	removeRefUrlFromRedis(refKey)
//
//	result, err := sqlxdb.Exec("UPDATE s3_files SET status=3 WHERE refKey=? AND projectName=? AND ownerId=? AND ownerType=? AND status=?", refKey, project, ownerId, ownerType, 1)
//	if err != nil {
//		log.Println("exec failed:", err)
//		return
//	}
//
//	idAff, err := result.RowsAffected()
//	if err != nil {
//		log.Println("RowsAffected failed:", err)
//		sendShortResp(w, -1001, "数据库系统错误")
//		return
//	}
//	log.Println("Success: deleted rows", idAff)
//	sendShortResp(w, 0, "成功删除")
//}
//
//func checkUploadRulers(ownerId, ownerType int64, project string, s3Length int64, fileName string) bool {
//	if ownerId < 0 || ownerType < 0 || s3Length <= 0 {
//		log.Printf("checkUploadRulers: Invalid userid-%d and file size %d\n", ownerId, s3Length)
//		return false
//	}
//
//	for i, v := range uploadRulers {
//		limitKey := fmt.Sprintf("s3:limit:%du%d.%s.%d:int", ownerType, ownerId, project, i)
//
//		if leftSize, err := rdsClient.Get(ctx, limitKey).Int64(); err == nil {
//			log.Printf("checkUploadRulers: Check limitation of [%d] from redis, Key= %s\n", leftSize, limitKey)
//
//			if leftSize-s3Length <= 0 {
//				log.Printf("checkUploadRulers: User (%d)'s S3 space left %d in %d seconds, but want to upload %d for file %s\n", ownerId, leftSize, v.timeInSec, s3Length, fileName)
//				if v.alert {
//					SendAlertEmail("链票告警服务:用户上传文件超限制",
//						fmt.Sprintf("<h1>错误信息：</h1><p>违反规则: 限制规则-%d, %d字节/%d秒</p><p>发生时刻:%s</p><p>用户ID:%d</p><p>redis key=%s</p><p>待上传文件:%s</p><p>文件大小:%d 字节</p><p>剩余额度: %d 字节",
//							i, v.limitSize, v.timeInSec, time.Now().Local().Format(MYSQL_TIME_FORMAT), ownerId, limitKey, fileName, s3Length, leftSize))
//				}
//				return false
//			}
//			if err := rdsClient.Set(ctx, limitKey, leftSize-s3Length, redis.KeepTTL).Err(); err != nil {
//				log.Printf("checkUploadRulers: System ERROR! Can not update redis data. err - %+v\n", err)
//				// SendAlertEmail("链票告警服务:Redis服务异常", "<h1>错误信息：</h1><p>System error:"+err.Error()+"</p>"+"<p>写入出错, redis key="+limitKey+"</p>")
//				return false
//			}
//		} else if err == redis.Nil {
//			if err := rdsClient.Set(ctx, limitKey, v.limitSize-s3Length, time.Duration(v.timeInSec)*time.Second).Err(); err != nil {
//				log.Printf("checkUploadRulers: System ERROR! Can not update redis data. err - %+v\n", err)
//				// SendAlertEmail("链票告警服务:Redis服务异常", "<h1>错误信息：</h1><p>System error:"+err.Error()+"</p>"+"<p>写入出错, redis key="+limitKey+"</p>")
//				return false
//			}
//		} else {
//			log.Printf("checkUploadRulers: Redis system ERROR! err:%+v\n", err)
//			// SendAlertEmail("链票告警服务:Redis服务异常", "<h1>错误信息：</h1><p>System error:"+err.Error()+"</p>"+"<p>读取出错, redis key="+limitKey+"</p>")
//			return false
//		}
//	}
//
//	return true
//}
//
//func generateFileKeys(root, filename string, ownerid, ownerType int64) (s3key, refkey string) {
//	s3key = fmt.Sprintf("%X/%d-%d/%s", md5.Sum([]byte(root+strconv.FormatInt(time.Now().UnixNano(), 16))), ownerType, ownerid, path.Base(filename))
//	refkey = fmt.Sprintf("%X/%s", md5.Sum([]byte(s3key)), path.Base(filename))
//	return s3key, refkey
//}
//
//func handleS3PostUrlReq(w http.ResponseWriter, r *http.Request) {
//	// if r.Method != "GET" {
//	//      w.WriteHeader(405)
//	//      return
//	// }
//
//	params := r.URL.Query()
//	log.Println("Query = ", params)
//
//	expired, err := strconv.ParseInt(params.Get("expired"), 10, 64)
//	if err != nil {
//		// io.WriteString(w, `{"code": 1000, "message":"未指定接口过期时间"}`)
//		sendShortResp(w, -1001, "未指定接口过期时间")
//		return
//	}
//	if expired <= 0 {
//		expired = 60
//	}
//
//	s3Length, err := strconv.ParseInt(params.Get("length"), 10, 64)
//	if err != nil || s3Length > 30*1024*1024 || s3Length <= 0 { // 30MB
//		// io.WriteString(w, `{"code": 1000, "message":"未指定待上传文件的大小,或者文件超过30MB"}`)
//		sendShortResp(w, -1002, "未指定待上传文件的大小,或者文件超过30MB")
//		return
//	}
//
//	project := strings.ToUpper(strings.Trim(params.Get("proj"), " "))
//
//	ownerId, err := strconv.ParseInt(params.Get("ownerid"), 10, 64)
//	_, fileName := path.Split(params.Get("file"))
//	if ownerId < 0 || err != nil || project != "LP" && project != "WBZ" && project != "ATP" || fileName == "" {
//		// io.WriteString(w, `{"code": 1000, "message":"缺少用户ID, 文件名, 或项目名"}`)
//		sendShortResp(w, -1003, "无效用户ID, 文件名, 或项目名")
//		return
//	}
//	fileName = strings.ReplaceAll(fileName, "+", "_")
//
//	ownerType, err := strconv.ParseInt(params.Get("ownerType"), 10, 64)
//	if ownerType < 0 || err != nil {
//		ownerType = 0
//	}
//
//	if !checkUploadRulers(ownerId, ownerType, project, s3Length, fileName) {
//		sendShortResp(w, -1005, "申请上传的文件大小超出用户当前可用存储容量的上限")
//		return
//	}
//
//	_, mimeType := getContentType(fileName)
//
//	// s3Key := fmt.Sprintf("%X/%d/%s", md5.Sum([]byte(project+fileName+strconv.FormatInt(time.Now().Unix(), 16))), ownerId, path.Base(fileName))
//	// refkey := fmt.Sprintf("%X/%s", md5.Sum([]byte(s3Key)), path.Base(fileName))
//	s3Key, refKey := generateFileKeys(project+fileName, fileName, ownerId, ownerType)
//
//	var s3Bucket, downloadUrl string // := "s3-lp-zone-01-eno9s1x41q96cr94wmzh7jsabio1rcnn1a-s3alias"
//
//	switch project {
//	case "LP", "ATP":
//		s3Bucket = "lp-pfiles" //"s3-lp-zone-01-eno9s1x41q96cr94wmzh7jsabio1rcnn1a-s3alias"
//		downloadUrl = packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", refKey, project)
//	case "WBZ":
//		s3Bucket = "wbz-pfile" //"oss-wbz-zone-01-rn4hsnzw11owjcc3qjdxduk9ad4o1cnn1a-s3alias"
//		downloadUrl = packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", refKey, project)
//	default:
//		// io.WriteString(w, `{"code": 1000, "message":"非法的项目名称"}`)
//		sendShortResp(w, -1004, "项目不存在")
//		return
//	}
//
//	//
//	c := &Credentials{
//		Region:          AWS_REGION,
//		Bucket:          s3Bucket,
//		AccessKeyID:     AWS_ACCESSKEY,
//		SecretAccessKey: AWS_SECRETKEY,
//	}
//	o := &PolicyOptions{
//		ExpirySeconds: int(expired),
//		MinFileSize:   int(s3Length - 1),
//		MaxFileSize:   int(s3Length),
//	}
//	refUrl := packUploadAndDownloadUrl(WBZ_S3_HOST, "/uploaded", refKey, project)
//
//	post, err := NewPresignedPOST(s3Key, refUrl, c, o)
//	// log.Printf("The POST policy is: %+v, err: %v", post, err)
//	if err != nil {
//		log.Printf("Create POST policy error: %v\n", err)
//		sendShortResp(w, -1, "系统错误!")
//		return
//	}
//
//	// data := fmt.Sprintf(`{"uploadUrl":"%s", "url":"%s", "length":%d, "file":"%s", "expired":%d, "userId":"%d"}`, post.Policy, downloadUrl, s3Length, fileName, expired, ownerId)
//	// sendResp(w, 0, "成功", data)
//
//	var s3Token = &S3UploadToken{
//		S3Bucket:    s3Bucket,
//		S3Key:       s3Key,
//		RefKey:      refKey,
//		S3Url:       post.Signature, // upload url
//		DownloadUrl: downloadUrl,
//		Length:      s3Length,
//		ExpireInSec: expired,
//		RemoteAddr:  r.RemoteAddr,
//		CreateTime:  time.Now().UTC().Format(MYSQL_TIME_FORMAT),
//		OwnerId:     ownerId,
//		OwnerType:   ownerType,
//		ProjectName: project,
//		FileName:    fileName,
//		FileExt:     mimeType,
//	}
//
//	if err := rdsClient.Set(ctx, s3Key, s3Token, time.Duration(expired)*time.Second).Err(); err == nil {
//		// io.WriteString(w, fmt.Sprintf(`{"code": 0, "message":"成功", "data":{"url":"%s", "length":%d, "file":"%s", "expired":%d}}`, str, s3Length, fileName, expired))
//		data := fmt.Sprintf(`{"Key": "%s","Policy": "%s","Signature": "%s","RedirUrl": "%s","uploadUrl": "%s","Credential": "%s","Date": "%s","Acl": "private","length":%d, "file":"%s", "expired":%d}`,
//			post.Key, post.Policy, post.Signature, post.RedirUrl, post.Action, post.Credential, post.Date,
//			s3Length, fileName, expired)
//		log.Printf("Redis = %+v\nResponse: %s\n", s3Token, data)
//		sendResp(w, 0, "成功", data)
//		return
//	} else {
//		log.Printf("Error :%v", err)
//		sendShortResp(w, -2, "系统错误!")
//	}
//}
//
//func handleS3UploadReq(w http.ResponseWriter, r *http.Request) {
//	// if r.Method != "GET" {
//	//      w.WriteHeader(405)
//	//      return
//	// }
//
//	params := r.URL.Query()
//	log.Println("Query = ", params)
//
//	expired, err := strconv.ParseInt(params.Get("expired"), 10, 64)
//	if err != nil {
//		// io.WriteString(w, `{"code": 1000, "message":"未指定接口过期时间"}`)
//		sendShortResp(w, -1001, "未指定接口过期时间")
//		return
//	}
//	if expired <= 0 {
//		expired = 60
//	}
//
//	s3Length, err := strconv.ParseInt(params.Get("length"), 10, 64)
//	if err != nil || s3Length > 30*1024*1024 { // 30MB
//		// io.WriteString(w, `{"code": 1000, "message":"未指定待上传文件的大小,或者文件超过30MB"}`)
//		sendShortResp(w, -1002, "未指定待上传文件的大小,或者文件超过30MB")
//		return
//	}
//
//	project := strings.ToUpper(strings.Trim(params.Get("proj"), " "))
//
//	_, fileName := path.Split(params.Get("file"))
//	ownerId, err := strconv.ParseInt(params.Get("ownerid"), 10, 64)
//	if ownerId < 0 || err != nil || project != "LP" && project != "WBZ" && project != "ATP" || fileName == "" {
//		// io.WriteString(w, `{"code": 1000, "message":"缺少用户ID, 文件名, 或项目名"}`)
//		sendShortResp(w, -1003, "无效用户ID, 文件名, 或项目名")
//		return
//	}
//	fileName = strings.ReplaceAll(fileName, "+", "_")
//
//	ownerType, err := strconv.ParseInt(params.Get("ownerType"), 10, 64)
//	if ownerType < 0 || err != nil {
//		ownerType = 0
//	}
//
//	if !checkUploadRulers(ownerId, ownerType, project, s3Length, fileName) {
//		sendShortResp(w, -1005, "申请上传的文件大小超出用户当前可用存储容量的上限")
//		return
//	}
//
//	_, mimeType := getContentType(fileName)
//
//	// s3Key := fmt.Sprintf("%X/%d/%s", md5.Sum([]byte(project+fileName+strconv.FormatInt(time.Now().Unix(), 16))), ownerId, path.Base(fileName))
//	// refKey := fmt.Sprintf("%X/%s", md5.Sum([]byte(s3Key)), path.Base(fileName))
//	s3Key, refKey := generateFileKeys(project+fileName, fileName, ownerId, ownerType)
//
//	var s3Bucket, downloadUrl string // := "s3-lp-zone-01-eno9s1x41q96cr94wmzh7jsabio1rcnn1a-s3alias"
//
//	switch project {
//	case "LP", "ATP":
//		s3Bucket = "lp-pfiles" //"s3-lp-zone-01-eno9s1x41q96cr94wmzh7jsabio1rcnn1a-s3alias"
//		downloadUrl = packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", refKey, project)
//	case "WBZ":
//		s3Bucket = "wbz-pfile" //"oss-wbz-zone-01-rn4hsnzw11owjcc3qjdxduk9ad4o1cnn1a-s3alias"
//		downloadUrl = packUploadAndDownloadUrl(WBZ_S3_HOST, "/dl", refKey, project)
//	default:
//		// io.WriteString(w, `{"code": 1000, "message":"非法的项目名称"}`)
//		sendShortResp(w, -1004, "项目不存在")
//		return
//	}
//
//	svc := s3.New(s3Sess)
//
//	option := s3.PutObjectInput{
//		Bucket:        aws.String(s3Bucket),
//		Key:           aws.String(s3Key),
//		ContentLength: aws.Int64(s3Length),
//		//      // Body:   strings.NewReader("EXPECTED CONTENTS"),
//	}
//
//	req, _ := svc.PutObjectRequest(&option)
//	url, err := req.Presign(time.Duration(expired) * time.Second)
//
//	log.Println("The URL is:", url, " err:", err)
//	if err != nil {
//		sendShortResp(w, -1, "系统错误!")
//		return
//	}
//
//	var s3Token = &S3UploadToken{
//		S3Bucket:    s3Bucket,
//		S3Key:       s3Key,
//		RefKey:      refKey,
//		S3Url:       url, // upload url
//		DownloadUrl: downloadUrl,
//		Length:      s3Length,
//		ExpireInSec: expired,
//		RemoteAddr:  r.RemoteAddr,
//		CreateTime:  time.Now().UTC().Format(MYSQL_TIME_FORMAT),
//		OwnerId:     ownerId,
//		OwnerType:   ownerType,
//		ProjectName: project,
//		FileName:    fileName,
//		FileExt:     mimeType,
//	}
//
//	if err := rdsClient.Set(ctx, s3Key, s3Token, time.Duration(expired)*time.Second).Err(); err == nil {
//		// io.WriteString(w, fmt.Sprintf(`{"code": 0, "message":"成功", "data":{"url":"%s", "length":%d, "file":"%s", "expired":%d}}`, str, s3Length, fileName, expired))
//		data := fmt.Sprintf(`{"uploadUrl":"%s", "url":"%s", "length":%d, "file":"%s", "expired":%d}`, url, downloadUrl, s3Length, fileName, expired)
//		log.Printf("Redis = %+v\nS3 file's info:%s\n", s3Token, data)
//		sendResp(w, 0, "成功", data)
//		return
//	} else {
//		log.Printf("Error :%v", err)
//		sendShortResp(w, -2, "系统错误!")
//	}
//}
//
//func packUploadAndDownloadUrl(host, path, key, project string) string {
//	timeStamp := time.Now().Unix()
//
//	signature := fmt.Sprintf("%s4%s-%s+%s.v2@%s", path, project, _S3_URL_MAGCODE, key, strconv.FormatInt(timeStamp, 8))
//	// log.Println("Pack signature =", signature)
//	signatueUrl := fmt.Sprintf("%s%s?a=%s&a1=%s&b=%X&c=%03d%d&d=v2", host, path, url.QueryEscape(key), project, md5.Sum([]byte(signature)), rand.Intn(1000), timeStamp)
//	log.Printf("signatueUrl = [%s] for %s", signatueUrl, key)
//	return signatueUrl
//}
//
//func checkDownloadUrlSignature(path string, params url.Values) (targetKey, project, meta string, err error) {
//	log.Println("Query = ", params)
//
//	urlKey := params.Get("a")
//	proj := params.Get("a1")
//	cipher := params.Get("b")
//	timeS := params.Get("c")
//	version := params.Get("d")
//	metadata := strings.Trim(params.Get("signature"), " ") // fake signature string
//
//	if metadata != "" { //len(metadata) > 32 {
//		metadata = fmt.Sprintf("%X", md5.Sum([]byte(metadata)))
//	}
//
//	if urlKey == "" || proj == "" || cipher == "" || timeS == "" || len(timeS) < 10 || version == "" {
//		return urlKey, proj, metadata, errors.New("无效请求(参数不合法)")
//	}
//
//	timeStamp, err := strconv.ParseInt(timeS[3:], 10, 64)
//	if err != nil {
//		log.Println("Parse time stamp failed! -", timeS)
//		return urlKey, proj, metadata, errors.New("时间戳字符串信息不正确")
//	}
//	log.Println("Parsed time stamp =", timeStamp)
//
//	key, err := url.QueryUnescape(urlKey)
//	if err != nil {
//		log.Println("Decode url key value failed! -", urlKey)
//		return urlKey, proj, metadata, errors.New("请求参数信息错误")
//	}
//
//	var signature string
//	if version == "v1" {
//		signature = fmt.Sprintf("%s+%s.v1@%s", _S3_URL_MAGCODE, key, strconv.FormatInt(timeStamp, 8))
//	} else {
//		signature = fmt.Sprintf("%s4%s-%s+%s.v2@%s", path, proj, _S3_URL_MAGCODE, key, strconv.FormatInt(timeStamp, 8))
//	}
//
//	// log.Println("Check signature =", signature)
//	if fmt.Sprintf("%X", md5.Sum([]byte(signature))) == cipher {
//		return key, proj, metadata, nil
//	} else {
//		return key, proj, metadata, errors.New("签名信息不正确")
//	}
//}
