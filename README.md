# DFA 敏感词检测

> golang dfa 敏感词检测算法实现，支持动态设置敏感词

```go   
    go get github.com/bean-du/dfa
```

example：
```go
    sensitive := []string{"王八蛋", "王八羔子"}

    fda := NewFDA()
    fda.AddBadWords(sensitive)

    str := "你个王#八……羔子， 你就是个王*八/蛋"
    fmt.Println(fda.Check(str))
```

输出结果：
```go
[王#八……羔子 王*八/蛋] true
```
