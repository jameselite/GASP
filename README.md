<br>

<h1> GASP: Golang CLI Assistant for backend projects ! </h1>

![License](https://img.shields.io/badge/license-MIT-cyan)  ![License](https://img.shields.io/badge/Version-1.0.0-black)   ![License](https://img.shields.io/badge/Maintainer-Soroush_GH-blue)  ![License](https://img.shields.io/badge/status-active-purple)


<br>

#### Forgot the setup of your backend project in Go !

<br>
GASP help you by generating boilerplate, making folder structure based on the architect of your project,config files, generating backend components such as controllers,routers, middlewares etc.

<br>
<br>

> [!NOTE]
> It installs and setup <a href="https://github.com/sqlc-dev/sqlc">sqlc</a> for generating type safe code from raw SQL
<br>


## Installation
To get started with this project, follow these steps:

  <br>
  
  1. Make sure you have go installed and it is in your PATH.
  
  <br>
  
  2. Run this command :
  ```
  go install github.com/jameselite/gasp@latest
  ```   

  <br>

## Usage

<br>

1. Initialize a project with:
```
gasp init example_project
```

2. Move to the created directory

3. Run start command to start the proccess of setup:
```
gasp start
```

4. if anything missed or you want to add it later you can use add commands, more information in:
```
gasp --help
```

5. generating backend components are also easy, just a few notes:

<h1> Tips </h1>
1. in making a controller with "gasp generate controller", you MUST have a router, and you can make a router with:

```
gasp generate router example_group
```

2. group path means: the path that every controller in that router have it is called Grouping routes

<br>

## Contributing
We’d love to have your contributions! Here’s how you can help:

1. Fork this repository.
2. Create a new branch for your changes:
  ```
   git checkout -b feature/your-feature-name
  ```
3. Make your changes and commit them:
  ```
   git commit -m "Add your message here"
  ```
4. Push to your fork:
  ```
   git push origin feature/your-feature-name
  ```
5. Submit a pull request to the main repository.

Please follow the project’s coding guidelines and include tests where applicable.

## License
This project is licensed under the MIT License See the [LICENSE](LICENSE) file for more details.
