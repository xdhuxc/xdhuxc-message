### xdhuxc-message

用于发送邮件和钉钉消息。


### 功能
1、支持定时发送消息

2、支持重试和重发：已经支持

3、metrics 接口：已经支持

4、全字段文本查找：已经支持，但是有点问题

5、钉钉消息发送到个人工作通知好群组：目前可以发送到个人消息通知

6、重新设计错误码格式

7、增加操作审计：已经支持

8、支持邮件发送至多人及抄送：已经支持


### 错误码设计

错误码 | 类型 | 含义
- | - | - 
00001 | 

### 接口说明

1）请求方法：POST

2）请求地址：<http://127.0.0.1:8080/api/v1/message>

3）请求头：

请求头 | 值
--- | ---
Content-Type | application/json
Accept | application/json

4）请求体

请求体数据格式为 `JSON`，字段包括：

 字段  | 类型 | 是否必需 | 含义
:------:| :-----:|:------:|------|
sender | 字符串 | 是 | 消息发送者
messageType | 字符串 | 是 | 消息类型，目前支持：email，dingtalk
content | 字符串 | 是 | 消息内容
receivers | 字符串数组 | 是 | 对于钉钉，是消息接收者在钉钉系统中的 user_id 列表；对于邮件，是邮件接收者的邮箱数组
cc | 字符串数组 | 否 | 邮件的抄送者邮箱数组
emailType | 字符串 | 否，对于邮件消息，必填 | 邮件类型，可选项为：text/plain 或 text/html
subject | 字符串 | 否，对于邮件消息，必填 | 邮件的主题
description | 字符串 | 否 | 描述信息

5）响应

响应状态码 | 含义
--- | ---
200 | 处理程序接收到了请求，并成功处理
非 200 | 出现了错误，处理程序没有接收到请求或其他情况

### 示例

对于钉钉消息

请求体为：
```markdown
{
    "sender": "xdhuxc",
    "messageType": "dingtalk",
    "content": "xdhuxc-dingtalk-1",
    "receivers": ["8", "88"],
    "description": "it is only a test"
}
```
响应体为：
```markdown
{
    "result": {
        "id": 1,
        "sender": "xdhuxc",
        "messageType": "dingtalk",
        "isSent": true,
        "content": "xdhuxc-dingtalk-1",
        "description": "it is only a test",
        "createTime": "2019-11-25T16:13:12.596565+08:00",
        "updateTime": "2019-11-25T16:13:12.596565+08:00",
        "receivers": [
            "8", "88"
        ]
    },
    "code": 0
}
```

对于邮件消息

请求体为：
```markdown
{
    "sender": "xdhuxc"
    "messageType": "email",
    "content": "xdhuxc-email-test-1",
    "receivers": ["xdhuxc@163.com", "xdhuxc@qq.com"],
    "cc": ["xdhuxc@163.com", "xdhuxc@qq.com"],
    "emailType": "text/plain",
    "subject": "xdhuxc test",
    "description": "it is only a test"
}
```
响应体为：
```markdown
{
    "result": {
        "id": 1,
        "sender": "xdhuxc"
        "messageType": "email",
        "isSent": true,
        "content": "xdhuxc-email-test-1",
        "description": "it is only a test",
        "createTime": "2019-11-25T17:19:00.898447+08:00",
        "updateTime": "2019-11-25T17:19:00.898448+08:00",
        "receivers": [
            "xdhuxc@163.com",
            "xdhuxc@qq.com"
        ],
        "cc": [
            "xdhuxc@163.com",
            "xdhuxc@qq.com"
        ],
        "emailType": "text/plain",
        "subject": "xdhuxc test"
    },
    "code": 0
}
```
### 指标接口

http://localhost:8080/metrics

### 常见问题及解决
1、发送邮件消息时，报错如下：
```
{
    "code": 100007,
    "type": "",
    "result": "gomail: could not send email 1: 501 mail from address must be same as authorization user"
}
```
原因：邮件的实际发送人和邮件服务器的 user 不一致导致的。


解决：修改邮件的实际发送人为邮件服务器的 user。


### 参考资料

MySQL 全文索引

https://segmentfault.com/a/1190000020470079?utm_source=tag-newest
