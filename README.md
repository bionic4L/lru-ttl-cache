### LRU (Least Recently Used) Cache
**LRU** — это алгоритм, удаляющий из кэша данные, которые не запрашивались дольше всего. 

>*Представь, что у тебя есть 5 последних вкладок в браузере, а ты открываешь шестую. Если места больше нет, то закроется та вкладка, которой ты пользовался давно, а не та, которую открыл минуту назад.*

Такой подход эффективен в ситуациях, когда данные, использованные недавно, с высокой вероятностью понадобятся снова, например, при кэшировании запросов в базе данных или изображений на сайте.

TTL (Time To Live) модификация позволяет очищать устаревшие элементы по истечении их "срока годности", что позволяет еще более эффективно использовать кеш, так как данные не "тухнут".