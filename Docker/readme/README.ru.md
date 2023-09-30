# Docker
Docker — инструмент для сборки пакетов, поставки и запуска приложений в «контейнерах»

## Что такое контейнер Docker?

Контейнер — базовая единица программного обеспечения, покрывающая код и все его зависимости
для обеспечения запуска приложения прозрачно, быстро и надежно независимо от окружения. Контейнер
Docker может быть создан с использованием образа Docker. Это исполняемый пакет программного обеспечения,
содержащий все необходимое для запуска приложения, например, системные программы, библиотеки, код, среды исполнения
и настройки.

## Cоставные части архитектуры Docker

Основные составные части архитектуры Docker — это:

- `сервер`, содержит сервис Docker, образы и контейнеры. Сервис связывается с `Registry`,
образы — метаданные приложений, запускаемых в контейнерах Docker.
- `клиент`, применяется для запуска различных действий на сервере Docker.  
- `registry`, используется для хранения образов. Есть публичные, доступные каждому, например, `Docker Hub` и `Docker Cloud`.

## Наиболее важные команды Docker

Наиболее важные команды Docker:
- `build`, сборка образа для Docker
- `create`, создание нового контейнера
- `kill`. принудительная остановка контейнера
- `dockerd`, запуск сервиса Docker
- `commit`, создание нового образа из изменений в контейнере

## Состояние контейнера Docker

Чтобы определить состояние, надо запустить команду:

> docker ps -a

Эта команда выведет список всех доступных контейнеров с их состоянием на сервере.
Из этого списка нужно выбрать требуемый контейнер и узнать его состояние.

## Образ Docker, Docker run

Образ Docker — это набор файлов, соединенный с настройками, с помощью которого можно создать экземпляры,
которые запускаются в отдельных контейнерах в виде изолированных процессов. Образ строится с использованием
инструкций для получения полной исполняемой версии приложения, зависящей от версии ядра сервера.  
Команда `Docker run` используется для создания таких экземпляров, называемых контейнерами, запускаемыми
с использованием образов Docker. При запуске одного образа пользователь может создать несколько контейнеров.

## Различия между виртуализацией и контейнеризацией

`Виртуализация` позволяет запустить несколько операционных систем на одном физическом сервере.  
`Контейнеризация` работает на одной и той же операционной системе, в которой приложения упакованы
в контейнеры и запускаются на одном сервере/виртуальной машине.

## Продвинутые команды Docker

Наиболее важные из них:

- `docker -version`: узнать установленную версию Docker;
- `docker ps`: перечислить все запущенные контейнеры вместе с дополнительной информацией о них;
- `docker ps -a`: перечислить все контейнеры, включая остановленные, вместе с дополнительной информацией о них;
- `docker exec`: войти в контейнер и выполнить в нем команду;
- `docker build`: собрать образ из Dockerfile;
- `docker rm`: удалить контейнер с указанным идентификатором;
- `docker rmi`: удалить образ с указанным идентификатором;
- `docker info`: получить расширенную информацию об установленном Docker, например, сколько запущено контейнеров, образов, версию ядра, доступную оперативную память и т.п.;
- `docker cp`: сохранить файл из контейнера в локальную систему;
- `docker history`: показать историю образа с указанным именем.

## Docker и Docker Compose

Docker применяется для управления отдельными контейнерами (сервисами), из которых состоит приложение.

Docker Compose используется для одновременного управления несколькими контейнерами, входящими в состав приложения.
Этот инструмент предлагает те же возможности, что и Docker, но позволяет работать с более сложными приложениями.

## Дополнительный материал

- [50 вопросов по Docker](https://habr.com/ru/companies/slurm/articles/528206/)
- [20 распространенных вопросов про Docker на собеседовании](https://itsecforu.ru/2020/01/15/%F0%9F%90%B3-20-%D1%80%D0%B0%D1%81%D0%BF%D1%80%D0%BE%D1%81%D1%82%D1%80%D0%B0%D0%BD%D0%B5%D0%BD%D0%BD%D1%8B%D1%85-%D0%B2%D0%BE%D0%BF%D1%80%D0%BE%D1%81%D0%BE%D0%B2-%D0%BF%D1%80%D0%BE-docker-%D0%BD%D0%B0/)
- [Руководство по Docker](https://medium.com/nuances-of-programming/%D1%80%D1%83%D0%BA%D0%BE%D0%B2%D0%BE%D0%B4%D1%81%D1%82%D0%B2%D0%BE-%D0%BF%D0%BE-docker-%D1%87%D0%B0%D1%81%D1%82%D1%8C-1-%D0%BE%D0%B1%D1%80%D0%B0%D0%B7-%D0%BA%D0%BE%D0%BD%D1%82%D0%B5%D0%B9%D0%BD%D0%B5%D1%80-%D1%81%D0%BE%D0%BF%D0%BE%D1%81%D1%82%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5-%D0%BF%D0%BE%D1%80%D1%82%D0%BE%D0%B2-%D0%B8-%D0%BE%D1%81%D0%BD%D0%BE%D0%B2%D0%BD%D1%8B%D0%B5-%D0%BA%D0%BE%D0%BC%D0%B0%D0%BD%D0%B4%D1%8B-4997f2968925#:~:text=%D0%BD%D0%B0%D1%81%D1%82%D1%80%D0%BE%D0%B5%D0%BD%D0%BD%D0%BE%D0%B5%20%D1%81%D0%BE%D0%BF%D0%BE%D1%81%D1%82%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D0%BF%D0%BE%D1%80%D1%82%D0%BE%D0%B2-,%D0%92%D1%8B%D0%B2%D0%BE%D0%B4%D1%8B,%D0%9F%D1%80%D0%B8%D0%BB%D0%BE%D0%B6%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%B0%D0%B5%D1%82%20%D0%B2%D0%BD%D1%83%D1%82%D1%80%D0%B8%20%D0%BA%D0%BE%D0%BD%D1%82%D0%B5%D0%B9%D0%BD%D0%B5%D1%80%D0%B0%20Docker!)

## README.md

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/Docker/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/Docker/readme/README.ru.md)