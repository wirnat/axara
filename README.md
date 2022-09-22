
# Aksara CLI

Make your life easier, this CLI will autogenerate code base on choosed clean architecture.





[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)


## Overview

- Autogenerate Interface Model base on choosed architecture

- Implement S.O.L.I.D for OOD

- Full integrate with Aksarabase (Struct scanner, Query builder, Database Bussiness Management, ORM)

- Autogenerate basic CRUD code in every layer 

- Autogenerate folder and infrastructure base on choosed architecture

- Easy to use

- Autodetected and generate just for unexecuted module

- Developer friendly
## CLI Command

#### Command 

```http
  unclebob, hexagonal, onion
```

| Command | Description                                | Available Delivery Command
| :-------- | :--------------------------------------- | :------------------------ 
| `unclebob` | https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html| delivery_echo


#### Flag

```http
  help, go-module, path, model
```

| Parameter  | Description                |
| :-------- | :------------------------- |
| `help` | Get detail command and flag for CLI |
| `go-module` | Init your project module |
| `path` | Directory path for compiled code and directory |
| `model` |  Path of existing model directory |
| `delivery` |  Auto generate delivery/ui layer |



## Authors

- [@wira](https://gitlab.com/wirawirw)


## Installation

Install using go module in your project

```bash
  go get -d github.com/wirnat/aksara-cli@latest
```

## Usage/Examples
    1. Create your directory for example: model, and create your model file .
       you must declare ~ModelName exatly same like your model name,
       ~DB in every struct field, and ~EndOfModel at the end of your last curlybracket model.
    
### Model

```go
//~ModelName Company
type Company struct {
	aksarabase.BaseModel        //~DB
	Name                 string `json:"name" adb:"Name"`
	PICName              string `json:"pic_name" adb:"`
	PICPhone             string `json:"pic_phone"`
}

//~EndOfModel
```

###
    2. Run aksara-cli command to execute and auto generate your code
### Terminal

```cmd
example: aksara-cli unclebob --path modules --model model --go-module todo

```
 Note: if aksara-cli command still not found, please reopen your terminal.

###      Done !
If success, it will generate unclebob architecture .
It will create 2 layer(usecase and repository) who contract with generated repository interface and usecase interface.
Now you can just create delivery layer and to struct method in your model.
Congrats you use time wisely :)


