## Blueprint/Boilerplate For Golang Projects

### Using GitHub Registry

Create and Push:

```bash
docker login docker.pkg.github.com -u <USERNAME> -p <GITHUB_TOKEN>
docker build -t  docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest .
# make container
docker push docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest
# make push
```

Pull and Run:

```bash
docker pull docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest
docker run docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest
```
