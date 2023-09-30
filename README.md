# Challenges


### CAPSTONE: TICKET TO MARS

#### Welcome to the first challenge. It’s time to take everything covered in unit 1 and write a program on your own. Your challenge is to write a ticket generator in the Go Playground that makes use of variables, constants, switch, if, and for. It should also draw on the fmt and math/rand packages to display and align text and to generate random numbers.
#### When planning a trip to Mars, it would be handy to have ticket pricing from multiple spacelines in one place. Websites exist that aggregate ticket prices for airlines, but so far nothing exists for spacelines. That’s not a problem for you, though. You can use Go to teach your computer to solve problems like this.Start by building a prototype that generates 10 random tickets and displays them in a tabular format with a nice header, as follows:

```
SpaceLine         Days    Trip-Type       Price
======================================================
SpaceX                   9       Round-trip      $ 58
Space Adventures         21      Round-trip      $ 21
SpaceX                   24      Round-trip      $ 77
SpaceX                   40      Round-trip      $ 80
SpaceX                   49      Round-trip      $ 62
Space Adventures         33      Round-trip      $ 31
Virgin Galactic          22      Round-trip      $ 85
Virgin Galactic          21      Round-trip      $ 71
SpaceX                   20      Round-trip      $ 36
```

#### The table should have four columns:
- The spaceline company providing the service
- The duration in days for the trip to Mars (one-way)
- Whether the price covers a return trip
- The price in millions of dollars 
#### For each ticket, randomly select one of the following spacelines: Space Adventures, SpaceX, or Virgin Galactic.Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km away from Earth at the time.Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine the duration for the trip to Mars and also the ticket price.

#### Make faster ships more expensive, ranging in price from $36 million to $50 million. Double the price for round trips.When you’re done, post your solution to the Get Programming with Go forums at forums.manning.com/forums/get-programming-with-go. If you get stuck, feel free to ask questions on the forums, or take a peek at the appendix for our solution.
