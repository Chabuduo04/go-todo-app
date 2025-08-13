package test

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strconv"
    "testing"
)

const baseURL = "http://localhost:8080" // 改成你的服务地址

// 全局 token & todoID
var token string
var createdTodoID int

func TestTodoAPI(t *testing.T) {
    // 1. 注册 & 登录
    token = registerAndLogin(t)

    // 2. 创建 Todo
    createdTodoID = createTodo(t, token)

    // 3. 获取 Todos
    getTodos(t, token)

    // 4. 更新 Todo
    updateTodo(t, token, createdTodoID)

    // 5. 删除 Todo
    deleteTodo(t, token, createdTodoID)
}

// 注册 & 登录
func registerAndLogin(t *testing.T) string {
    // 注册
    regData := map[string]string{
        "username": "testuser",
        "password": "123456",
    }
    body, _ := json.Marshal(regData)
    resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewBuffer(body))
    if err != nil {
        t.Fatalf("注册失败: %v", err)
    }
    resp.Body.Close()

    // 登录
    loginData := map[string]string{
        "username": "testuser",
        "password": "123456",
    }
    body, _ := json.Marshal(loginData)
    resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(body))
    if err != nil {
        t.Fatalf("登录失败: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        b, _ := ioutil.ReadAll(resp.Body)
        t.Fatalf("登录失败，状态码: %d, 响应: %s", resp.StatusCode, string(b))
    }

    var res map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&res)
    tok, ok := res["token"].(string)
    if !ok || tok == "" {
        t.Fatalf("登录返回 token 无效: %v", res)
    }

    t.Logf("登录成功, token: %s", tok)
    return tok
}

// 创建 Todo
func createTodo(t *testing.T, token string) int {
    todo := map[string]interface{}{
        "title":     "测试Todo",
        "completed": false,
    }
    body, _ := json.Marshal(todo)

    req, _ := http.NewRequest(http.MethodPost, baseURL+"/todos", bytes.NewBuffer(body))
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        t.Fatalf("创建 Todo 请求失败: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        b, _ := ioutil.ReadAll(resp.Body)
        t.Fatalf("创建 Todo 失败，状态码: %d, 响应: %s", resp.StatusCode, string(b))
    }

    var res map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&res)
    todoData := res["todo"].(map[string]interface{})
    id := int(todoData["ID"].(float64))

    t.Logf("创建 Todo 成功, ID=%d", id)
    return id
}

// 获取 Todos
func getTodos(t *testing.T, token string) {
    req, _ := http.NewRequest(http.MethodGet, baseURL+"/todos", nil)
    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        t.Fatalf("获取 Todos 请求失败: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        b, _ := ioutil.ReadAll(resp.Body)
        t.Fatalf("获取 Todos 失败，状态码: %d, 响应: %s", resp.StatusCode, string(b))
    }

    var res map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&res)

    todos := res["todos"].([]interface{})
    if len(todos) == 0 {
        t.Fatalf("Todo 列表为空")
    }

    t.Logf("获取 Todos 成功, 数量=%d", len(todos))
}

// 更新 Todo
func updateTodo(t *testing.T, token string, todoID int) {
    updated := map[string]interface{}{
        "title":     "测试Todo（已更新）",
        "completed": true,
    }
    body, _ := json.Marshal(updated)

    req, _ := http.NewRequest(http.MethodPut, baseURL+"/todos/"+strconv.Itoa(todoID), bytes.NewBuffer(body))
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        t.Fatalf("更新 Todo 请求失败: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        b, _ := ioutil.ReadAll(resp.Body)
        t.Fatalf("更新 Todo 失败，状态码: %d, 响应: %s", resp.StatusCode, string(b))
    }

    t.Logf("更新 Todo 成功, ID=%d", todoID)
}

// 删除 Todo
func deleteTodo(t *testing.T, token string, todoID int) {
    req, _ := http.NewRequest(http.MethodDelete, baseURL+"/todos/"+strconv.Itoa(todoID), nil)
    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        t.Fatalf("删除 Todo 请求失败: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        b, _ := ioutil.ReadAll(resp.Body)
        t.Fatalf("删除 Todo 失败，状态码: %d, 响应: %s", resp.StatusCode, string(b))
    }

    t.Logf("删除 Todo 成功, ID=%d", todoID)
}
