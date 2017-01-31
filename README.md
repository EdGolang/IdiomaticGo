# Idiomatic Go

This repository explores various different ways of achieving the same end result to various degrees.
> Each manner of doing this could be extended further, and does contain some code not necessary required, but left is so as to mimic what would occur in the real world.

  ---
  
## Problem statement

An employee is either a developer, tester or project manager. Load an employee and print a statement related to them.

### Procedural
  Each employee is a map[string]string
  Uses switch statement to determine what to do
   
### Functional
  Each employee is a map[string]interface{}
  Uses function maps to determine types and related functions
  Uses function parameters to determine final result

### Clojure
  Each employee is a UserType, type defined as a map[string]string
  Very basic use of clojures at top level to show practical uses
  Communication through channels
  > Gets into linking up multiple channels very quickly, so just did this at the top level

### OO
  Employee is a base type, each type of employee derives from there
  Show basic use of interfaces and inheritance as would be done in other languages

### Idiomatic
  Employee is a base type, each type of employee embeds employee
  Use of types occurs solely through interfaces
  Because it is the data different and not the implementation, we return data. If it were implementation, we'd likely pass in an interface and then apply it
 