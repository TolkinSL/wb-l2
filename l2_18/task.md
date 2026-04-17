L2.18  
HTTP-сервер «Календарь»  

Пример запуска:  
go run main.go -port 8080

Создание события:
curl -X POST "http://localhost:8080/create_event" \
     -d "user_id=1" \
     -d "date=2026-04-17" \
     -d "title=Семинар Go!"

Получение события:
curl "http://localhost:8080/events_for_day?user_id=1&date=2026-04-17"

Удаление события:
curl -X POST http://localhost:8080/delete_event \
     -d "id=1"

Запуск тестов:
go test ./internal/service/... -v