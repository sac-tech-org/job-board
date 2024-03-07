# job-board API

This is a demonstration of a HTTP REST API written in Go. It is a recreation of a system in which
multiple interfaces are implemented to serve uniform data from various data sources.

## Project structure

### Code Structure

The `main.go` file sets up dependencies and starts the server, listening on port 8080. The [server](./server/)
package is the central package that handles all the client interactions via REST API. The dependencies are implemented as shown in this graph.

```mermaid
---
title: Server
---
classDiagram
  direction RL


  class Server
  Server: +DataStore dataStore

  class DataStore
  DataStore: +DBClient dbClient
  DataStore: DeleteUser(context.Context, server.DeleteUserInput) error
	DataStore: GetUser(context.Context, server.GetUserInput) (server.User, error)
	DataStore: GetUserList(context.Context, server.GetUserListInput) ([]server.User, error)
	DataStore: PostUser(context.Context, server.PostUserInput) (server.User, error)
	DataStore: PutUser(context.Context, server.PutUserInput) (server.User, error)

  class DBClient
  DBClient: QueryUserByID(uuid.UUID) (DataStore.User, error)
```

### OpenAPI

The OpenAPI spec is in [openapi.yaml](./openapi.yaml), and should be kept up to date with any changes made to the API.
