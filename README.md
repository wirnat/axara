
# Axara

Make your life easier, this CLI will autogenerate code and architecture base on Template and Model.




[![MIT License](https://img.shields.io/github/license/wirnat/axara?style=flat-square)](https://github.com/wirnat/axara/blob/main/License)
![MIT License](https://img.shields.io/github/go-mod/go-version/wirnat/axara)



## Overview

- Easy to use and learn

- Can be used to generate boilerplate, architectural pattern, layer, small module, and code file

- Generate your repeatable code such as repository, mock, delivery, view, and other some unique repeatable task.

- Speed up your develop process and make you only focus to the business logic

- Customable variable for the template

- Powerfull meta/variable feature

- Can be used to build your favorite boilerplate

- Various template has availabled, such as Clean Architecture, Monorepo Clean Architecture, even repository style code has availabled for speed up your develop.

- Yaml and Json support for orchestrator file

- Get template from github direct from axara command

- Improve Monolith modular development speed until 80%

- Fully Customable Template

- Can inject code to existing file

- You can choose which job will run when generate

## Instalation
```bash
  go install github.com/wirnat/axara@latest
```
## Grab Template
```bash
  axara get [Template Git URL] [Result Directory]
```
## Generate 
```bash
  axara generate [template config yaml/json] -g [list of job tags]
```
## Quick Example 
#### 1. Get Template
```bash
  axara get https://github.com/wirnat/axara-template-go-clean-architecture templates/clean-architecture
```
#### 2. Setup Model
Lets create example model that's will generate into code and run in every jobs:

model/base_model.go
```go

package model

import (
	"time"
)

type BaseModel struct {
	ID        int64      `json:"id"`
	UUID      string     `json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"` //~ignore
	DeletedAt *time.Time `json:"deleted_at"` //~ignore
}
```

model/company.go

```go
package model

type Company struct {
  BaseModel
  Name string `json:"name"` //@meta validate:true
}

//@Register Company
```

model/branch.go

```go
package model

type Branch struct {
  BaseModel
  Name string `json:"name"` //@meta validate:true
  CompanyUUID `json:"company_uuid"` //@meta validate:true
}

//@Register Branch
```

Add //@Register [ModelName] comment in end of model struct, it will mark the model and add to config executor

#### 3. Setup Template Config File
If you success get the template , you can modify template config in templates/clean-architecture or your result get directory.
Now open index.yaml file, and adjust the following configurations:

```yaml
model_path: model #or your model directory
```

Listing the model and add meta, just adjust into which models will executed: 
  ```yaml
  models:
    Company:
      module: company
    Branch:
      module: branch
  ```
  
By default, the template creator will guide you to which need to adjust in config file especially in meta section:
```yaml
  main_dir: .
  app_dir: .
  app:  gitlab.com/aksaratech/example-quick #your gomod module name
  import:  gitlab.com/aksaratech/example-quick #your gomod module name or import root
  this: ~main_dir~/templates/clean-architecture  #this config template directory
  template: ~this~/templates #templates directory
  import_model: ~import~/model #import meta + adjust to your model directory
```

The final result of the configuration file will be as follows.
```yaml
key: ᬅᬓ᭄ᬱᬭ #init aksara key
model_path: model  #init model dir
lang: golang #available for now: golang, typescript, js, dart
models:
    Company:
      company: company
    Branch:
      branch: branch

meta:
  main_dir: .
  app_dir: .
  app:  gitlab.com/aksaratech/example-quick #your gomod module name
  import:  gitlab.com/aksaratech/example-quick #your gomod module name or import root
  this: ~main_dir~/templates/clean-architecture  #this config template directory
  template: ~this~/templates #templates directory
  import_model: ~import~/model #import meta + adjust to your model directory

  infrastructure_dir: ~app_dir~/infrastructure
  import_infrastructure: ~import~/infrastructure
  import_pagination: ~import_infrastructure~/paginator
  import_trx: ~import_infrastructure~/db_transaction
  import_trx_gorm: ~import_infrastructure~/db_transaction/gorm_transaction
  import_request: ~import~/domain/request/~model_snake~_request
  import_response: ~import~/domain/response/~model_snake~_response
  import_conf: ~import_infrastructure~/env/conf
  import_env: ~import_infrastructure~/env
  import_repository: ~import~/repository/~model_snake~_repository
  import_usecase: ~import~/usecase/~model_snake~_usecase


  #you can fill other static meta here

#jobs is queue of generator file
include_jobs:
  - ~this~/jobs/structure.yaml
  - ~this~/jobs/repository.yaml
  - ~this~/jobs/param.yaml
  - ~this~/jobs/infrastructure.yaml
  - ~this~/jobs/response.yaml
  - ~this~/jobs/usecase.yaml
  - ~this~/jobs/delivery.yaml
```

#### 4. Execute
```bash
  axara generate templates/clean-architecture/index.yaml -g init --models Company,Branch
```
The command will execute tasks that are tagged as 'init' and loop the job in every Company and Branch Model

## Command
## How its Work?
### Config File
### Jobs
### Template
### Model
### Meta



