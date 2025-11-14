# Bomlok

The idea of this library is to reduce the need for reflection when getting values from a struct by utilizing code generation.

The name is heavily inspired by the lombok library that has been written for java to reduce the need for getters and setters,
however this library is more or less the other way around, but to simplify the process for the most common usage of reflection.

## Motivation

There is no getting around, reflection is relatively slow. To speed it up some people have decided to write reflection libraries that utilize the `unsafe` package. However i believe that is the wrong approach to a solution that
can lead to dangerous consequences. Generated code is faster and should have no meaningful overhead compared to other forms of reflection.
