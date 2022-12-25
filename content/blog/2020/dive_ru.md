---
date: 2020-05-14
image: dive.png
---

# Dive

Если для вас не чуждо слово Докер, то хочу посоветовать небольшую утилиту, которая упростит анализ Докер-образа (как же дико звучит).

Показывает содержимое образа по слоям, подсвечивает изменения. Можно использовать для поиска ненужных файлов чтобы уменьшить размер образа.

https://github.com/wagoodman/dive

<video width="720" height="450" controls>
  <source src="dive.mp4" type="video/mp4">
</video>

Запускается довольно просто:

```bash
dive <your-image-tag>
```

Но можно запустить и как Докер-контейнер, если не хочется устанавливать приложение:

```bash
docker run --rm -it \
    -v /var/run/docker.sock:/var/run/docker.sock \
    wagoodman/dive:latest <dive arguments...>
```

Представляете, можно через Докер запустить контейнер Дайва чтобы проанализировать контейнер Дайва.

#ops