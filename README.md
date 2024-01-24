# Web Brasa
A browser is a medium for viewing web content. Its components include UI, Browser Engine, Rendering Engine, Networking, JavaScript Interpreter, and Data Storage.

Building one helps you understand the concepts much better as a front-end engineer so that you create solutions that work best for the browser. 

So let's build web brasa

## COMPONENTS
**[] UI** responsible for the browser's user interface. It renders the web page, handles user interactions such as URL visits and history navigation, and displays web content.

**[] Browser Engine** loads and renders online pages. Its duties include rendering the resources it retrieves from the network and executing scripts like JavaScript.

**[] Networking** component makes HTTP requests and retrieves resources from the network.

**[] JavaScript Interpreter** 

**[] Data Storage** 

# TECH STACK
- App
    - [Go](https://go.dev/)
- UI
    - [Gotk3](https://github.com/gotk3/gotk3), Go bindings for GTK3
- Browser Engine
- Networking
- JavaScript Interpreter
- Data Storage

# USAGE
- How to run this app locally
    - Install system dependencies
        - Gotk3 dependent on certain 
            - [Linux](https://github.com/gotk3/gotk3/wiki/Installing-on-Linux)
            - [MacOS](https://github.com/gotk3/gotk3/wiki/Installing-on-macOS)
            - [Windows](https://github.com/gotk3/gotk3/wiki/Installing-on-Windows)

    - Clone this repo

    - Install project dependencies


        ```bash
        go mod download
        ```

    - Run the project


        ```bash
        go run main.go
        ```

    - Application GUI window will pop up
- [] Web
- [] How do you integrate Retriedge into VS Code as an extension?
- [] What are the hosting requirements for Retriedge as a SaaS app via Docker?