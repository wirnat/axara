
# Axara

Make your life easier, this CLI will autogenerate code or architecture base on Template and Model.




[![MIT License](https://img.shields.io/github/license/wirnat/axara?style=flat-square)](https://github.com/wirnat/axara/blob/main/License)
![MIT License](https://img.shields.io/github/go-mod/go-version/wirnat/axara)



## Overview

- Easy to use and learn

- Can be used to CRUD code, generate boilerplate, architectural pattern, layer, small module, and code file

- Generate your repeatable code such as repository, mock, delivery, view, and other some unique repeatable task.

- Speed up your develop process and make you only focus to the business logic

- Customable variable for the template

- Customable Template

- Orchestrate generated file location in jobs, so you can create your own template architecture

- Various template has availabled, such as Clean Architecture, Monorepo Clean Architecture, even repository style code has availabled for speed up your develop.

- Yaml and Json support for orchestrator file

- Get template from github direct from axara command

- Improve Monolith modular development speed until 80%

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
      module: company
    Branch:
      module: branch

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

## Command & Flag

| Command | Description |
|---------|--------------|
| completion  | Generate the autocompletion script for the specified shell |
| generate  |  Auto generate Design Pattern |
| get  |  Get CLI Item from github |
| help  |  Help about any command |
| new  |  New Axara Config file |
|version | Check Version|

| Flag | Short | Description |
|---------|---------|---------|
| --help | -h  | help for axara |
| --models strings |  -m |list of execute models |
| --tags string | -g  |List of execute traits/jobs |


## How its Work?
Unlike typical CRUD generators, Axara CLI generates highly customizable code base on templates, where templates have dynamic variables from each job in the configuration file. The location of the generated files can also be specified in the configuration file, specifically in the job settings.

### Config File
The config file is a configuration file as well as the main executor and orchestrator in this generator, where it determines what jobs are executed, variables that are declared, and other needs of the generator.
| Key | Value Type | Description |
|---------|--------------|-------|
|key | string | this is package key for axara cli, by default just fill with ᬅᬓ᭄ᬱᬭ|
|model_path | string | project model directory|
|lang | string | project language, available: go, typescript |
|models | map[string][map]string | mount your model to config and add dynamic meta variable in the model that's can used in config file / templates|
|meta | map[string]string | add global meta variable that's can used in config file / templates |
|jobs | []Job | List of jobs to be executed during the execution later |
|include_jobs | []string | you can seperate the list of jobs file and add the file path here |

### Jobs
The executor will run the configuration file and read the tasks to be performed. The tasks referred to here include creating files (with directory paths), generating code in files based on templates specified in the job, and assigning tags to jobs to organize which jobs will be executed during generation.
The tasks can be executed either once or per declared model, If executed per model, each model will have its own jobs with their respective meta/variables.

| Key | Value Type | Description | Example |
|---------|--------|---------|---------|
| name | string | name of job | repository |
| dir | string | file directory | app |
| template | string| The template file path that will be compiled into code | templates/repository.text |
| file_name |string| The name of generated file/injected file | repository.go |
| active |bool| job status if false it will passed | false |
| single_execute |bool| If set true the job will not looping per model, and it only executed once | true |
| generate_in |string| replace @Generate [name] with compiled template code   | route |
| tags |[]string| declared which tags will execute the job   | - init |

### Template
The template is the main code generator that is executed by each job. This template is read using Go templates, so users can utilize Go-template features to create/modify the template. Variables used in the template are extracted from meta-data in the config file, whether it is from global meta or model meta.

For more information about text template https://pkg.go.dev/text/template
### Model
When you want to create jobs based on models, such as CRUD repositories or similar entities that have dynamic variables based on the model, the model is the main component that needs to be linked to your generator.

#### Step to Setup Model
1. Add //@Register [model name] after close bracket model object
2. Declare it in config:
	```yaml
	models:
		Company:
		#you can add model meta here
	```
4. When performing the execution, use the "--models Company,Branch" flag to select which model will be executed.

### Meta



