Step 1. สร้างไฟล์ .env ขึ้นมาใหม่โดยอิงจาก .env.example <br />
API_KEY="นำ API KEY ของคุณมาใส่ในนี้"<br />
API_SECRET="นำ API SECRET ของคุณมาใส่ในนี้"

Step 2. go run main.go

ได้ Output ดังนี้<br />
2024-04-24 03:39:54.227978 +0000 UTC <br />
{Id:302206985 Hash:fwQ6dnQYKnuwtiUrMZp3HbKMpUY Type:market Amount:10 Rate:0 Fee:0 Cred:0 Receive:0 Ts:1713929994}<br />
cron job started...

เมื่อรัน โปรแกรมจะซื้อ Bitcoin แบบ market 10 บาท บน Bitkub<br />
จากนั้นจะซื้อ Bitcoin แบบ market 10 บาท บน Bitkub ทุกๆ 1 ชม.

