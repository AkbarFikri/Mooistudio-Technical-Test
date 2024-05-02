# MooiStudio Technical Test

## 📒 Index

- [📒 Index](#-index)
- [🔰 About](#-about)
- [⚡ Quick Start](#-quick-start)
    - [🔌 Installation](#-installation)
    - [📦 Commands](#-commands)
- [🔧 Development](#-development)
    - [📓 Tech Stack](#-tech-stack)
    - [🔩 API Documentation](#-api-documentation)
- [🔒License](#license)


## 🔰 About

Here's the app that i provided for technical test to get internship on MooiStudio Software House. This apps is fulfil MVP requirement.
This the list of MVP requirement:
- Customer can view product list by product category
- Customer can add product to shopping cart
- Customers can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customers can checkout and make payment transactions (just create order, no need implement payment gateway)
- Login and register customers

**Entity Relational Diagram 📊**
![Alt text](/public/erd.png "a Entity Relational Diagram")

## ⚡ Quick Start

Here's the step for installation and start our app.

_`Note: This is just a backend Apps not include the Frontend Apps.`_

### 🔌 Installation

1. First, make sure that the go language version you have is more than `1.21`
2. Next, you can clone this repository with the command below

```
$ git clone https://github.com/AkbarFikri/mooistudio_technical_test .
```

3. Provide all the `.env.example` file then rename to `.env`
4. Download all packages needed by Go by running the command below

```
$ go mod tidy
```

**❗ YEAYY Installation Finish!!**

### 📦 Commands

- To run the application you can directly open `main.go` in folder `cmd/app` then click the `run without debugging` button in the right corner of vscode or run the command below

```
$ go run cmd/app/main.go
```

### 📓 Tech Stack

List all the Tech Stack we use to build the system in this this project.

| No | Tech    | Details                                                           |
|----|---------| ----------------------------------------------------------------- |
| 1  | Go      | To build a fast and easy Backend App                              |
| 2  | Postman | To build beatiful documentation                                   |

### 🔩 API Documentation

- [Postman](https://documenter.getpostman.com/view/30883191/2sA3JFB4tn)

_Note : If you have question about the documentation feel free to send message to me._

## 🔒License

© Copyright (c) 2024 AkbarFikri
