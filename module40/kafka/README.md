# Учим kafka

## Подготовка. DEPRICATED

Т.к. будем использовать библиотеку [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go),
необходимо установить библиотеку [librdkafka](https://github.com/confluentinc/librdkafka).

Для ubuntu или wsl ubuntu: `sudo apt install librdkafka-dev`.

Внимание! Если работа из системы windows и устанавливаете из wsl ubuntu, необходимо запускать код так же из 
wsl терминала

Для Windows:
1. Скачиваем nuget [тут](https://www.nuget.org/packages/librdkafka.redist/)
2. Открываем архиватором, например, Winrar
3. Переходим в /runtimes/<платформа x64 или x86>/native
4. Копируем все либо в C:/Windows/system32 или в любую другую папку
5. Если в другую папку, нужно добавить ее в переменные окружения PATH