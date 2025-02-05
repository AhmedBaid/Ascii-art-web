# HTTP Server & Protocol in Go

## 🌍 HTTP Server Overview
A **HTTP server** listens for requests from clients on a specific port, processes them, and sends back responses. Here's how it works:

### 🔄 HTTP Server Lifecycle
1. **Receiving Request**: The server receives a request from a client.
2. **Handling Request**: The server processes the request using a **handler function**. Example: If a client requests `/`, the server asks the handler to return the corresponding HTML page.
3. **Sending Response**: The server sends back a response (HTML, JSON, etc.).

---

## 📡 Understanding HTTP Protocol
HTTP is a protocol that defines how a **server** and a **client** communicate. A request includes:
- **Method** (GET, POST, etc.)
- **URL** (e.g., `/home`)
- **Headers** (metadata like language preferences)

### ✅ Common HTTP Status Codes
| Code | Meaning |
|------|---------|
| **200** | Success ✅ |
| **404** | Not Found 🚫 |
| **500** | Server Error 🔥 |
| **400** | Bad Request ❌ |
| **302** | Found 🔄 |

---

## ⚙️ `ListenAndServe`
This function creates an HTTP server and keeps it running, listening for client requests.

```go
http.ListenAndServe(":8080", nil) // Runs server on port 8080
```

---

## 🛠️ `Handler` in Go
A **handler** is a function that processes HTTP requests. When a request comes in, Go uses `DefaultServeMux` to route the request to the correct handler.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
})
```

---

## 🔌 Understanding Sockets & WebSockets
- **Socket**: A communication link between two programs (e.g., a browser and a server).
- **WebSocket**: A protocol for **real-time** communication (e.g., chat applications, live notifications).
- **TCP & UDP Sockets**: 
  - **TCP**: Reliable, ensures all data arrives.
  - **UDP**: Faster but less reliable.

---

## 🚀 `net/http` Package in Go
This package includes many built-in methods to:
- Start an HTTP server.
- Handle HTTP requests (server-side).
- Send HTTP requests (client-side).
- Manage routes using handlers.

### 🏗️ Important Structures
#### `http.Request`
Represents an **HTTP request**.
- `Method`: GET, POST, etc.
- `URL`: The requested address.
- `Header`: Request headers.
- `Body`: The request payload.

#### `http.Response`
Represents an **HTTP response**.
- `Status`: HTTP status code (200, 404, etc.).
- `Header`: Response headers.
- `Body`: The actual content.

---

## 🛤️ **Default ServeMux** (Built-in Router)
Go provides a built-in router (`ServeMux`) that maps URLs to handlers:

```go
http.HandleFunc("/hello", helloHandler)
```

This ensures requests to `/hello` are handled by `helloHandler`. 🚀


------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------


# 💻 Server HTTP & Protokol f Go

## 🌍 C'est quoi HTTP Server?
Server HTTP kaychd requests mn clients, kay3aljhom, w kayrj3hom responses. Hadi hiya torika dyal service ki ykhdm:

### 🔄 Service HTTP
1. **Receiving Request**: Client kaysift request l server.
2. **Handling Request**: Server kaychof request w kaydir lih traitement m3a **handler function**. Matalan, ila client tlba `/`, server ghay9ol lhandler yjib lih dik l page.
3. **Sending Response**: Server kayrj3 result l client (HTML, JSON, etc.).

---

## 📡 HTTP Protokol
HTTP howa tariqa li kayhadro biha **server** w **client**. Request fih:
- **Method** (GET, POST, etc.)
- **URL** (e.g., `/home`)
- **Headers** (b7al language dyal navigateur)

### ✅ Status Codes M3roofin f HTTP
| Code | M3na |
|------|---------|
| **200** | Success ✅ |
| **404** | Mal9ach 🚫 |
| **500** | Kayan mochkil f server 🔥 |
| **400** | Request ghalat ❌ |
| **302** | L9a chi haja o bddlha 🔄 |

---

## ⚙️ `ListenAndServe`
Hadi method f Go li katkhli server ydemarrer w ybqa kaychd requests.

```go
http.ListenAndServe(":8080", nil) // Server ghadi ykhdm f port 8080
```

---

## 🛠️ `Handler` f Go
**Handler** howa function li kaychd requests w kay3aljhom. Go kayst3ml `DefaultServeMux` bach i3rf handler mn URL.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Salam 3likom!")
})
```

---

## 🔌 C'est quoi Sockets & WebSockets?
- **Socket**: Wasita kaykhdm bin 2 baramij (b7al navigateur w server).
- **WebSocket**: Protokol li kaykhli l'tisaal **direct o mostamir** (b7al chat w live updates).
- **TCP & UDP Sockets**:
  - **TCP**: Kay3tik dawla wach data dkhlt kamla w makynach akhta2.
  - **UDP**: Zriba mais mkaynch verification bzaf.

---

## 🚀 Package `net/http` f Go
Hadi package kay3tik bzaf dyal tools li katkhallik:
- Tdemarrer server HTTP.
- Tsift requestat (côté client).
- Tst3ml handlers bach tdwr les routes.

### 🏗️ Structures Mhimmin
#### `http.Request`
Representi **HTTP request**.
- `Method`: GET, POST, etc.
- `URL`: L’adresse li talabha client.
- `Header`: Les headers dyal request.
- `Body`: Contenu dyal request.

#### `http.Response`
Representi **HTTP response**.
- `Status`: Code dyal réponse (200, 404, etc.).
- `Header`: Metadata dyal response.
- `Body`: Contenu li tsifto server.

---

## 🛤️ **Default ServeMux** (Système de Routing f Go)
Go kay3tik router built-in `ServeMux` li kay9sm requests 3la handlers lmonasibin:

```go
http.HandleFunc("/salam", salamHandler)
```

Hadchi kaykhli requests li jay mn `/salam` i9dro ychdoh `salamHandler`. 🚀

