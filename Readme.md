1. สร้างไฟล์ .env ขึ้นมาใหม่โดยอิงจาก .env.example
API_KEY="นำ API KEY ของคุณมาใส่ในนี้"
API_SECRET="นำ API SECRET ของคุณมาใส่ในนี้"

2. go run main.go

Output :
2024-04-24 03:39:54.227978 +0000 UTC 
{Id:302206985 Hash:fwQ6dnQYKnuwtiUrMZp3HbKMpUY Type:market Amount:10 Rate:0 Fee:0 Cred:0 Receive:0 Ts:1713929994}
cron job started...

เมื่อรัน โปรแกรมจะซื้อ Bitcoin แบบ market 10 บาท บน Bitkub
จากนั้นจะซื้อ Bitcoin แบบ market 10 บาท บน Bitkub ทุกๆ 1 ชม.

