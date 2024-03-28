# OOP in go

### 3 oop principles?
Abstraction, Encapsulation, Polymorphism

### SOLID principles?
SOLID stands for: Single Responsibility, Open-Closed, Liskov Substitution, Interface Segregation, and Dependency Inversion

### Single responsibility in go? 
Struct (object) should have 1 reason to change.

### Open-closed in go? 

### Liskov Substitution in go?
This principle states that subtypes should be substitutable for their base types. In go it is related to composition

### Interface Segregation Principle
no code should be forced to depend on methods it does not use. That is why we have `Reader`, `Closer` so on

### Dependency Inversion Principle
- High-level modules should not import anything from low-level modules
- Abstractions should not depend on details
