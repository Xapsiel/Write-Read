# Система управления портами

Есть Объект – “Порт”, он бывает 2-ух типов: IN и OUT. При запуске приложения передаются 2 параметра: кол-во IN и OUT портов. Есть API: READ(читает значение IN( random 0\1), параметр: номер порта, ответ значение 0\1) и WRITE(пишет значение в OUT, параметр: номер порта, значение – вывод в консоль).
## Возможности
- Поддержка одновременного чтения и записи значений портов.
- Использование `sync.Mutex` и `sync.WaitGroup` для безопасного доступа к общим ресурсам.
- Инициализация портов как IN или OUT на основе параметров, переданных при запуске.
- Значения входных портов (IN) случайно инициализируются 0 или 1.

## Установка
Клонируйте репозиторий и перейдите в папку проекта:
```sh
git clonehttps://github.com/Xapsiel/Write-Read.git
cd Write-Read
