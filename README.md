# Go design patterns

Repository with examples of design patterns in the Go language. 
This repository was created to put the ideas from this book into code: [Go Design Patterns by Mario Contreras](https://www.amazon.com/Design-Patterns-Mario-Castro-Contreras/dp/1786466201). The book itself has some issues, sometimes the methods declared in one page are being called with a different name in the next page, or return a different type of object...but it's not hard to overlook these issues and gain some knowledge from it. 

![Image](https://raw.githubusercontent.com/camilovietnam/go-design-patterns/main/go-logo.png)

## 01 - Creational Patterns

- Singleton: 
    - Use a shared database connection with a singleton client https://github.com/camilovietnam/go-design-patterns/tree/main/01-creational/01-singleton-pattern/database
    - Share resources with a singleton lock https://github.com/camilovietnam/go-design-patterns/tree/main/01-creational/01-singleton-pattern/lock
- Builder: 
    - Allow multiple logger types by using a builder https://github.com/camilovietnam/go-design-patterns/tree/main/01-creational/02-builder-pattern/logger   
