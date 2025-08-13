package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	baseURL    = "http://localhost:8080" // 根据实际情况修改
	username   = "dio"
	password   = "123456"
)

// 测试注册接口
func TestRegisterHandler(t *testing.T) {
	url := baseURL + "/register"
	
	// 构造请求体
	requestBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		t.Fatalf("构造请求体失败: %v", err)
	}

	// 发送POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("读取响应失败: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		t.Errorf("注册失败: 状态码 %d, 响应: %s", resp.StatusCode, string(body))
		return
	}

	// 解析响应JSON
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		t.Errorf("解析响应JSON失败: %v", err)
		return
	}

	// 验证响应内容
	if message, ok := response["message"].(string); !ok || message != "register success" {
		t.Errorf("响应消息不符合预期: %v", response)
	}

	t.Logf("注册成功: %s", string(body))
}

// 测试登录接口
func TestLoginHandler(t *testing.T) {
	url := baseURL + "/login"
	
	// 构造请求体
	requestBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		t.Fatalf("构造请求体失败: %v", err)
	}

	// 发送POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("读取响应失败: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		t.Errorf("登录失败: 状态码 %d, 响应: %s", resp.StatusCode, string(body))
		return
	}

	// 解析响应JSON
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		t.Errorf("解析响应JSON失败: %v", err)
		return
	}

	// 验证响应内容
	if message, ok := response["message"].(string); !ok || message != "login success" {
		t.Errorf("响应消息不符合预期: %v", response)
	}

	t.Logf("登录成功: %s", string(body))
}