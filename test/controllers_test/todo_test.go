package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

// TestCreateTodo 测试创建Todo接口
func TestCreateTodo(t *testing.T) {
	url := baseURL + "/todos"

	// 构造请求体
	todo := map[string]interface{}{
		"title":    "完成第一阶段功能",
		"completed": false,
		"user_id":  1,
	}
	requestBody, err := json.Marshal(todo)
	if err != nil {
		t.Fatalf("构造请求体失败: %v", err)
	}

	// 发送POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Errorf("创建Todo失败: 状态码 %d, 响应: %s", resp.StatusCode, string(body))
		return
	}

	// 解析响应
	var createdTodo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&createdTodo); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	// 验证响应内容
	if message, ok := createdTodo["message"].(string); !ok || message != "todo created successfully" {
		t.Errorf("响应消息不符合预期: %v", createdTodo)
	}

	t.Logf("创建Todo成功: %+v", createdTodo)
}

// TestUpdateTodo 测试更新Todo接口
func TestUpdateTodo(t *testing.T) {
	todoID := "1" // 假设要更新的Todo ID
	url := baseURL + "/todos/" + todoID

	// 构造请求体
	updatedTodo := map[string]interface{}{
		"title":    "完成第一阶段功能（已更新）",
		"completed": true,
		"user_id":  1,
	}
	requestBody, err := json.Marshal(updatedTodo)
	if err != nil {
		t.Fatalf("构造请求体失败: %v", err)
	}

	// 创建PUT请求
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("创建请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Errorf("更新Todo失败: 状态码 %d, 响应: %s", resp.StatusCode, string(body))
		return
	}

	// 解析响应
	var resultTodo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&resultTodo); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	// 验证响应内容
	if message, ok := resultTodo["message"].(string); !ok || message != "todo updated successfully" {
		t.Errorf("响应消息不符合预期: %v", resultTodo)
	}

	t.Logf("更新Todo成功: %+v", resultTodo)
}