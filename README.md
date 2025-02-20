Описание проекта:

Проект GO_FINAL_PROJECT — реализация To-Do списка для локального использования. Включает пакеты:

database — создание БД sheduler.db с таблицей sheduler.

handlers — обработчики для добавления, удаления, обновления, завершения задач, получения списка задач и поиска.

models — структура TASK.

nextDate — функция NextDate для вычисления следующей даты выполнения задачи.

response — вспомогательные функции Error и Success.


Запуск:

Локально: go run main.go.

Docker:

docker pull lenarasp/todo-list-backend

docker run -d -p 7540:7540 lenarasp/todo-list-backend

Сайт: http://localhost:7540/


Тестирование:

В tests/settings.go измените параметр Search на true.

Запустите тесты: go test ./tests.