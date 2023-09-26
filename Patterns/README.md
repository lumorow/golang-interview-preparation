# Design Patterns

<h1 align="center"><img class="goldT" src="../../img/patterns.webp" ></h1>

You can go straight to the explanation and implementation of patterns:
[Design patterns with examples in Golang](https://github.com/AlexanderGrom/go-patterns)


## Generators.
***
Such templates are needed to optimize the creation of a particular object.
Generative patterns help create objects so that they communicate effectively with others and control their operation.

### Here are some examples:
“Factory” - a separate class is invented to create new objects. It creates objects as copies of some standard;
“Prototype” - the object itself creates copies of itself;
Builder - similar to a factory, but new objects can be modified. They are created according to complex logic, and do not copy the reference one;
“Singleton” - implies the presence of one large object that has global access to everything;
“Lazy Initialization” is a method in which an object is not initialized immediately, but as it progresses.

There are other patterns of varying complexity. For each task, one or another pattern is more optimal. The exact solution depends on the application, but the result should be an efficient and optimized system.

## Structural.
***
If generative patterns are responsible for the creation and interaction of objects, then structural patterns are responsible for how these objects are structured in code.
They describe how simple classes and objects are “assembled” into more complex ones.

### Here are examples:

“Decorator” is a template for attaching additional behavior to an object;
“Composite” is a pattern that combines several objects into a tree structure;
“Bridge” is the principle of dividing an entity into abstraction and implementation, so that the theoretical structure and the concrete object can change independently;
“Facade” is a method for reducing external calls to a single object;
“Proxy” is a pattern similar to “Facade”, but with a special proxy object that controls access to the main one.
These are just some examples. There are many more real patterns.

## Behavioral.
***
These are design patterns that describe how objects behave and interact with others.
They are used, for example, to divide responsibilities between different entities or to react
for changes without errors.

### Examples of behavioral patterns:

“Iterator” - one object sequentially gives access to various others, without using their complex descriptions;
"Observer" - a pattern in which objects become aware of changes in others;
“Keeper” (Memento) - helps to preserve an object in some state with the ability to return to it in the future;
“Chain of Responsibility” - distributes responsibility for certain tasks to different objects;
“Mediator” - organizes weak connections between objects to reduce their dependence on each other.
How to understand which pattern to apply
Each pattern has its own area of use. Experienced developers understand where to use what, based on the very specifics of the task,
but at the beginning of the journey it can be difficult.

### Selecting a pattern into smaller steps:

- highlight the entities that are used in the process;
- think through the connections between them;
- abstract the resulting system from a specific task;
- see if the problem matches the meaning of something for which there is a pattern;
- select several patterns from the desired group and see which one suits best;
- think over a specific implementation of this pattern, taking into account the specifics of the task.

It looks difficult, but over time the habit will come. Experienced developers have already gotten the hang of it,
so they have much less problems with choosing a pattern. Once understanding is formed and practice is gained,
it will be easier.

### Benefits of Design Patterns
- Speed up and make writing code easier.
- They allow you not to “reinvent the wheel”, but to use a ready-made, proven principle.
- When used correctly, they make the code more readable and efficient.
- Simplify mutual understanding between developers.
- Helps get rid of common mistakes.
- Do not depend on the programming language and its features.
- Allows you to complete complex tasks faster and easier.

### Disadvantages of Design Patterns
- Using patterns for the sake of patterns, on the contrary, complicates the code and confuses developers.
- Incorrect use of one or another template can make the program less effective.
- Patterns are not universal: a specific pattern will work for one task, but not another.
- In the early stages of learning, it can be difficult to choose the appropriate pattern for a specific problem.
- Due to its strong association with object-oriented programming, the use of patterns in other paradigms is limited. Although, for example, in functional programming they can be used - they are simply implemented differently.

### Is it worth using patterns?
The presence of shortcomings does not make the idea of patterns bad. It's just a tool that needs to be used wisely.
You should not use templates where you can do without them, just for the sake of “beauty”. If you use them in places
where they are really needed - they will become good

### Resources
- [Go Design patterns (Introduction)](https://medium.com/@mr_apr/go-design-patterns-introduction-9c5e57a3cb13)
- [Паттерны проектирования с примерами на Golang](https://github.com/AlexanderGrom/go-patterns)
- [Golang | Паттерны проектирования - Агентство цифровых технологий](https://www.youtube.com/playlist?list=PLxj7Nz8YYkVW5KHnsb9qWUDP2eD1TXl1N)

## README.md
***

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/Patterns/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/Patterns/readme/README.ru.md)