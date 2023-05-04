package main

import (
	"demo/utils"
	"demo/utils/authv3"
	"fmt"
)

// 您的应用ID
var appKey = ""

// 您的应用密钥
var appSecret = ""

func main() {
	// 添加请求参数
	paramsMap := createRequestParams()
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/v2/correct_writing_text", header, paramsMap, "application/json")
	// 打印返回结果
	if result != nil {
		fmt.Print(string(result))
	}
}

func createRequestParams() map[string][]string {

	/*
		note: 将下列变量替换为需要请求的参数
		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/%E4%BD%9C%E6%96%87%E6%89%B9%E6%94%B9/API%E6%96%87%E6%A1%A3/%E8%8B%B1%E8%AF%AD%E4%BD%9C%E6%96%87%E6%89%B9%E6%94%B9%EF%BC%88%E6%96%87%E6%9C%AC%E8%BE%93%E5%85%A5%EF%BC%89/%E8%8B%B1%E8%AF%AD%E4%BD%9C%E6%96%87%E6%89%B9%E6%94%B9%EF%BC%88%E6%96%87%E6%9C%AC%E8%BE%93%E5%85%A5%EF%BC%89-API%E6%96%87%E6%A1%A3.html
	*/
	q := "正文文本"
	grade := "作文等级"
	title := "作文标题"
	modelContent := "作文参考范文"
	isNeedSynonyms := "是否查询同义词"
	correctVersion := "作文批改版本：基础，高级"
	isNeedEssayReport := "是否返回写作报告"

	return map[string][]string{
		"q":                 {q},
		"grade":             {grade},
		"title":             {title},
		"modelContent":      {modelContent},
		"isNeedSynonyms":    {isNeedSynonyms},
		"correctVersion":    {correctVersion},
		"isNeedEssayReport": {isNeedEssayReport},
	}
}
