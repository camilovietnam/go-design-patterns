# Go design patterns

Repository with examples of design patterns in the Go language. 
This repository was created to put the ideas from this book into code: [Go Design Patterns by Mario Contreras](https://www.amazon.com/Design-Patterns-Mario-Castro-Contreras/dp/1786466201). The book itself has some issues, many of the code seems to have not been tested and some methods change name from one page to another, but I think the ideas are worth exploring. 

01 - Creational Patterns

- Singleton: 
    - Use a shared database connection with a singleton client https://github.com/camilovietnam/go-design-patterns/tree/main/01-creational/01-singleton-pattern/database
    - Share resources with a singleton lock https://github.com/camilovietnam/go-design-patterns/tree/main/01-creational/01-singleton-pattern/lock
- Builder: 
    - Allow multiple logger types by using a builder https://github.com/camilovietnam/go-design-patterns/tree/main/01-creational/02-builder-pattern/logger   
