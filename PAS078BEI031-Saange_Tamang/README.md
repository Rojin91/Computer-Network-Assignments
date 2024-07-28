# QUESTIONS
- Capture the packet while sending an HTTP request also capture the response.Disect the requrest and response.
- Capture the packet while sending and email.

# HTTP Request and Response Dissection

## HTTP Request Dissection

An HTTP request consists of several components, each serving a specific purpose. It contains:
1. Request Line
2. Headers
3. Blank Line
4. Body (Optional)


### 1. Request Line
- **Method:** `GET` - Requests data from the specified resource.
- **Request URL:** `/` - The root directory of the server.
- **HTTP Version:** `HTTP/1.1` - This version of the HTTP protocol is being used.

### 2. Headers
- **Host:** `r3.i.lencr.org` — The domain name of the server.
- **Connection:** `keep-alive` — Indicates that the client wants to maintain a persistent connection.
- **User-Agent:** `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36` — Identifies the client software making the request.
- **Accept-Encoding:** `gzip, deflate` — Specifies the content encodings that the client can handle.
- **Accept-Language:** `en-US,en;q=0.9` — Specifies the preferred languages for the response.

### 3. Blank Line
This separates the headers from the body. In this request, there is no body.

## HTTP Response Dissection

An HTTP response consists of a status line, headers, and optionally a body.

### Status Line
- **HTTP Version:** `HTTP/1.1`
- **Status Code:** `200` — OK, indicating that the request was successful.
- **Status Message:** `OK`

### Headers
- **Server:** `nginx` — Specifies the server software.
- **Content-Type:** `application/pkix-cert` — The media type of the response body.
- **Last-Modified:** `Fri, 04 Aug 2023 20:57:56 GMT` — The last modification date of the resource.
- **ETag:** `"64dc6654-51a"` — A unique identifier for the version of the resource.
- **Content-Disposition:** `attachment; filename="R3.der"` — Indicates that the content should be downloaded and saved as `R3.der`.
- **Accept-Ranges:** `bytes` — Indicates that the server supports range requests.
- **Vary:** `Accept-Encoding` — Indicates that the response varies based on the `Accept-Encoding` header.
- **Content-Encoding:** `gzip` — The body is compressed using gzip.
- **Content-Length:** `1253` — The length of the response body in bytes.
- **Cache-Control:** `max-age=3600` — Indicates that the response can be cached for 3600 seconds.
- **Expires:** `Sun, 28 Jul 2024 16:40:37 GMT` — The date and time after which the response is considered stale.
- **Date:** `Sun, 28 Jul 2024 15:40:37 GMT` — The date and time at which the response was generated.
- **Connection:** `keep-alive` — Indicates that the connection should be kept open.

### Body
The body of the response contains binary data (likely the certificate file `R3.der`).

### Blank Line
This separates the headers from the body. In this response, the body contains binary data.






