# Ascii-Art-Web (add stylize and dockerise and export)

## Description
Ascii-Art-Web is a web-based application that allows users to generate ASCII art from text using different banner styles. The project extends the functionality of the previous ascii-art project by providing a graphical user interface (GUI) via a web server. Users can input text, select a banner style (shadow, standard, or thinkertoy), and view the ASCII art output on a webpage.

## Objectives
- Create and run an HTTP server.
- Implement a web GUI for generating ASCII art.
- Support the following banner styles:
  - `shadow`
  - `standard`
  - `thinkertoy`
- Provide the following HTTP endpoints:
  - `GET /`: Displays the main page with the input form.
  - `POST /ascii-art`: Processes user input and returns the generated ASCII art.

## HTTP Status Codes
- **200 OK**: Request was successful.
- **404 Not Found**: Templates or banners not found.
- **400 Bad Request**: Invalid requests.
- **500 Internal Server Error**: Unhandled server errors.

## Main Page Features
The main page includes:
1. **Text Input**: A field to input text for ASCII art generation.
2. **Banner Selection**: Radio buttons or a dropdown menu to select one of the banner styles.
3. **Submit Button**: A button to send the input and banner style to the server.
4. **Output Display**: Displays the generated ASCII art on the same page or on a separate results page.

## Authors
- Ahmed Baid (abaid)
- Reda Anniz (ranniz)

## Usage
### Running the Project
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ascii-art-web
   ```
2. Build and run the server:
   ```bash
   go run main.go
   ```
3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

### Input Format
- Enter any text in the input field.
- Select a banner style (shadow, standard, or thinkertoy).
- Click the "Generate" button to display the ASCII art.

## Implementation Details
### Algorithm
1. **Text Input Parsing**: The server receives text input and banner style via a POST request.
2. **ASCII Art Generation**:
   - The backend processes the text input using the selected banner style.
   - The ASCII art generation logic is reused from the ascii-art project.
3. **Response Handling**:
   - If successful, the ASCII art is displayed on the webpage.
   - If an error occurs (e.g., invalid input), an appropriate error message is shown.

### Project Structure
```
ASCII-ART-WEB/
├── cmd/         // entry point of the project
│   ├── main.go
├── files/           // Contains ASCII art files (e.g., shadow.txt, standard.txt, thinkertoy.txt)
│   ├── shadow.txt
│   ├── standard.txt
│   ├── thinkertoy.txt
├── functions/       // Go logic for ASCII art generation
│   ├── Ascii.go
│   ├── PrintAscii.go
├── handler/          // Handler functions (StyleHandler,ResultHandler,FormHandler)
│   ├── Handlers.go
├── styles/          // CSS styles for the application
│   ├── index.css
│   ├── result.css
│   ├── stausPage.css
├── template/        // HTML templates
│   ├── index.html
│   ├── result.html
│   ├── stausPage.html
├── go.mod           // Go module file
└── README.md        // Documentation

```

### Instructions
1. **Server Implementation**:
   - Write the HTTP server in Go.
   - Use `net/http` for routing and handling requests.
2. **Templates**:
   - Store all HTML templates in the `templates/` directory.
   - Use Go's `html/template` package to parse and render templates.
3. **Best Practices**:
   - Ensure clean and modular code.
   - Handle errors gracefully.
4. **Allowed Packages**:
   - Only standard Go packages are permitted.

## Example
1. **Input**:
   - Text: `Hello`
   - Banner: `shadow`

2. **Output**:
```                              
_|                _| _|          
_|_|_|     _|_|   _| _|   _|_|   
_|    _| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
                                 
```

Visit the application at [http://localhost:8080](http://localhost:8080) to test it yourself!


## Partie docker :
## docker integration 
- **Docker Integration**:  
  The project will include at least one Dockerfile, one image, and one container. Proper metadata will be applied to Docker objects, and unused objects (garbage collection) will be handled carefully.
  
- **Web Server in Go**:  
  The web server will be written in Go and must adhere to Go best practices.

- **Docker Practices**:  
  The project will follow Docker best practices, including the use of a Dockerfile to build and configure the container.

- **Standard Go Packages**:  
  Only the standard Go packages are allowed to be used in this project.

## Technologies Used

- **Go**: For the web server and ASCII art generation logic.
- **Docker**: To containerize the application, create images, and manage containers.
- **HTML & HTTP**: For the web interface and handling HTTP requests.

## Learning Objectives

This project will help you learn the following:

- Client utilities for handling web requests.
- Basics of web servers in Go.
- HTML and HTTP protocols.
- Methods for receiving and outputting data via web requests.
- Introduction to Docker concepts and commands.
- Using Docker to set up services and dependencies.
- Containerizing applications and creating Docker images.
- Managing application dependencies and ensuring compatibility.

## Project Requirements

1. **Create a Dockerfile**: The Dockerfile should respect good Docker practices for building and containerizing the Go web server.
2. **Create a Docker Image**: Build the Docker image from the Dockerfile.
3. **Create a Docker Container**: Run the application inside a Docker container.
4. **Metadata Application**: Apply metadata to Docker objects to ensure proper organization and management.
5. **Garbage Collection**: Handle unused Docker objects properly to avoid clutter and unnecessary resource consumption.


## Partie export
### Description
ASCII-Art-Web-Export est une extension du projet ASCII-Art-Web qui vise à permettre l'exportation des résultats ASCII générés par l'application web. L'objectif est de fournir une fonctionnalité permettant à l'utilisateur de télécharger le résultat sous un format exportable.

### Objectifs
- Assurer qu'il est possible d'exporter le résultat du projet [ASCII-Art](../../ascii-art/README.md) implémenté sur le site web.
- Le fichier exporté doit avoir les bonnes permissions (**lecture et écriture**) pour l'utilisateur.
- Comprendre et utiliser les en-têtes HTTP appropriées pour le transfert de fichiers :
  - [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type)
  - [Content-Length](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length)
  - [Content-Disposition](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition)
- Implémenter une interface utilisateur incluant un bouton ou un lien pour télécharger/exporter le fichier.

### Instructions
- Créer un nouvel **endpoint HTTP** pour permettre le transfert du fichier vers le client.
- Le serveur web doit **supporter au moins un format d'exportation**.
- Le serveur web doit être implémenté en **Go**.
- Assurer la gestion des erreurs du site web.
- Le code doit respecter les [bonnes pratiques](../../good-practices/README.md).

### Formats d'exportation possibles
- **Fichier texte (TXT)**
- [Autres formats](https://en.wikipedia.org/wiki/Document_file_format) peuvent être envisagés si pertinents.

### Packages autorisés
- Uniquement les packages standards de [Go](https://golang.org/pkg/).

### Compétences acquises
Ce projet vous aidera à apprendre :
- Les bases des formats d'exportation de fichiers.
- L'utilisation des [en-têtes HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) pour le transfert de fichiers.
- Les différentes manières de recevoir et d'envoyer des données en HTTP.
- L'amélioration de l'interface utilisateur en intégrant une option de téléchargement/exportation.
