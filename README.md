# screenshot
A Docker container with a REST API that takes screenshots of web pages


## Usage

### Pull Docker image

```sh
docker pull ghcr.io/preetam/screenshot:main
```

### Run Docker container

```sh
docker run -p 8000:8000 ghcr.io/preetam/screenshot:main
```

### Request a screenshot

http://localhost:8000/?url=https://google.com/

![image](https://user-images.githubusercontent.com/379404/148623283-61d7eb03-87e5-442d-a29b-914cdd38bd53.png)

Optionally set window size: http://localhost:8000/?url=https://google.com/&window_size=800,600

![image](https://user-images.githubusercontent.com/379404/148623306-149bf87c-1a12-41f7-93a9-21dabcc1ac4d.png)


## License

Apache 2.0 (see LICENSE)
