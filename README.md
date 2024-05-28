# Реализованные функции
* Регистраия
* Создание поста(доступно только для авторизированных пользователей). Пользователь, написавший пост, может запретить оставлять под ним комментарии.
* Получения списка постов(можно задать количество и смещение).
* Получение поста и комментариев под ним(можно задать количество и смещение). Комментарии передаются в иерархическом виде, позволяя вложенность без ограничений, как на Хабре и Редите.
* Написание комментария(доступно только для авторизированных пользователей). Размер комментария ограничен, значение ограничения задаётся в конфиге при запуске.
* Подписаться на пост. В этом случае подписавшемуся человеку будут асинхронно приходить все новые комментарии, появляющиеся под постом.
# Использованные Фреймворки
* В приложении используется авторизация при помощи токенов. Для реализаии был использован фреймворк - https://github.com/golang-jwt/jwt/
* Для работы с GraphQL использовался фреймворк - gqlgen.
* В качестве базы данных использовалась PostgreSQL. При этом также была реализована возможность хранить всю информацию в памяти. Выбрать тип используемого хранилища можно при запуске в конфиге. 0 - PostgreSQL, 1 - в памяти.

# Команда для запуска приложения и бд
docker compose up -d

# Возможные команды
## Регистрация 
### Пример обращения
```
mutation{
  Register(RegisterData:{
    login: "sasha"
  }){
      token
  }
}
```
### Пример ответа
```
{
  "data": {
    "Register": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4ODU5MzgsIkxvZ2luIjoic2FzaGEifQ.2g2t91z9Traoe2RK8_8Qb6MU4tJzW14Nlb7ZpDESmwQ"
    }
  }
}
```
## Создание поста
### Пример обращения
```
mutation{
  CreatePost(IdentificationData:{
    login: "sasha"
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4ODU5MzgsIkxvZ2luIjoic2FzaGEifQ.2g2t91z9Traoe2RK8_8Qb6MU4tJzW14Nlb7ZpDESmwQ"  }
    PostData:"POSTEREL"
    IsCommentEnbale: true){
      result
  }
}
```
### Пример ответа
```
{
  "data": {
    "CreatePost": {
      "result": "Successfully published"
    }
  }
}
```
## Получение списка постов
### Пример обращения
```
query{
  Posts(Limit: 3, Offset: 0){
    ID
    subject
    author
    timeAdd
    isCommentEnable
    
  }
}
```
### Пример ответа
```
{
  "data": {
    "Posts": [
      {
        "ID": "1",
        "subject": "POSTEREL",
        "author": "sasha",
        "timeAdd": "2024-05-28T08:47:15.173342Z",
        "isCommentEnable": false
      },
      {
        "ID": "2",
        "subject": "POSTEREL",
        "author": "sasha",
        "timeAdd": "2024-05-28T08:47:15.78089Z",
        "isCommentEnable": false
      },
      {
        "ID": "3",
        "subject": "POSTEREL",
        "author": "sasha",
        "timeAdd": "2024-05-28T08:47:16.248519Z",
        "isCommentEnable": false
      }
    ]
  }
}
```
## Добавление комментариев
### Пример обращения
```
mutation{
  AddComment(IdentificationData:{
    login: "sasha"
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4MDA5NjAsIkxvZ2luIjoic2FzaGEifQ.J-NPh1rt0sakjTvdkccOzX56HHLIFhpqAiEys746QYU"  }
  Comment:{
    CommentData: "blabla1"
    PostID: "12"
  }), {
      result
  }
}
```
### Пример ответа
```
{
  "data": {
    "AddComment": {
      "result": "Successfully published"
    }
  }
}
```
## Получение комментариев к посту
### Пример обращения
```
query{
  PostAndComment(PostID: 2,Limit: 3, Offset: 0){
    Post{
    ID
    subject
    author
    timeAdd
    isCommentEnable 
    }

    comments{
      CommentData
      ParentID
      PostID
      CommentID
      NestingLevel
    }
  }
}
```
### Пример ответа
{
  "data": {
    "PostAndComment": {
      "Post": {
        "ID": "2",
        "subject": "POSTEREL",
        "author": "sasha",
        "timeAdd": "2024-05-28T08:47:15.78089Z",
        "isCommentEnable": false
      },
      "comments": null
    }
  }
}

## Подписка на пост
### Пример обращения
```
subscription{
  GetCommentsFromPost(IdentificationData:{
    login: "sasha"
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4MDA5NjAsIkxvZ2luIjoic2FzaGEifQ.J-NPh1rt0sakjTvdkccOzX56HHLIFhpqAiEys746QYU"  }
    PostID: "2"
  ), {
      CommentID
      CommentData
  }
}
```
### Пример ответа
```
{
  "data": {
    "GetCommentsFromPost": {
      "CommentID": "6",
      "CommentData": "blabla1"
    }
  }
}
```
