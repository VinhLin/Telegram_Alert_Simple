# Telegram_Alert_Simple
Simple Alert to Telegram Bot

### Câu lệnh khởi tạo:
```
go mod init telegram_alert
```

### Build
```
go build -o telegram_alert *.go
```

-----------------------------------------------------------------------------------------------
### Note
- Mình sẽ cần sử dụng [library notify](https://github.com/nikoksr/notify) cho code của mình.
- Cài đặt libary:
```
go get -u github.com/nikoksr/notify
go get github.com/nikoksr/notify/service/telegram
```
- Có thể kiểm tra lại trong 2 file là **go.mod** và **go.sum**:
```
cat go.mod
cat go.sum
```

-------------------------------------------------------------------------------------------------
### Test Telegram Bot
- Mình đã tạo một bot telegram có token là: 
```
Token: 5698880402:AAHDIojvC9vuuW9Ti0BD2zw-ASSmQmoI-Uo
chat_ID: 5018541524
```



