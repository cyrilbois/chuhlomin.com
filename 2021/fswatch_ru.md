---
date: 2021-10-22
comments_enabled: true
---

# fswatch

Иногда бывает нужно регулярно запускать программу в терминале при изменении файлов в определённой директории. Этакий "локальный CI", с которым знакомы фронтэндеры: поменял JavaScript -- и тебе страница сразу сама обновилась.

Написанная Enrico Maria Crisostomo программа `fswatch` выводит пути измененных файлов в указанных директориях, что можно дальше передать "по цепочке" (pipe) в другую программу.

https://emcrisostomo.github.io/fswatch/

В примере ниже `fswatch` запускает программу `genblog` при изменении любых файлов в текущей директории (кроме файлов в директориях `.git` и `output`):

```bash
fswatch --one-per-batch --recursive --exclude="output" --exclude=".git" . | xargs -n1 sh -c "genblog"
```

#cli