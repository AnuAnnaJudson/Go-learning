Local setup for Gqlgen(steps)

1. delete scehema.resolvers.go
2. Go to models_gen.go remove all types
3. schema.graphqls clear all schemas
4. Copy and paste your schema
5. resolver.go include the line first which is go generate command
	//go:generate go run github.com/99designs/gqlgen
6. cd to graph folder
7. run command go generate(when you change schema do steps 6,7)
8. Resolver are there to inject dependencies in your mutation 
   and queries.
