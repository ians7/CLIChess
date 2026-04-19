# Project purpose

I created this project because I wanted to learn a new language over the summer, and the language I chose was golang. Although there are a million chess games out there, I wanted to choose something that would keep me interested in the project. I've been playing chess for years and am very familiar with the domain, so I thought creating a CLI chess app could be a lot of fun, and it was!

# How to play

Run the application by typing `./app` into the command line when in the `src/offlineApp` directory. Then, the user will be prompted to input proper chess notation. However, the user is not required to specify check or checkmate. For example, Rh3, Ngf3, Qxb7 would be proper input, but Rh3+, Ngf3+, Qxg7# would be improper input. 

# Future of the project

 - My main priority is to create a TCP server to run the game and be able to play with friends remotely
 - I would like to implement a rematch function where the users will receive the opposite colored pieces that they received the game prior. The idea of this is to make it easier for two people to play consecutive matches against each other.
 - I would like to log all of the moves and create functionality to import the game to a text file with all of the moves properly notated

# Reflection

Ultimately I had a lot of fun creating this project. I think that I would need to pursue more project in Go to really get the hang of it, but the project was a lot of fun.

Some of my struggles in this project include: 
 - Detecting checkmate. I didn't want to have to find every single possible move for each piece to try to be more efficient, but I ultimately decided to in order to see if a piece could block a check/checkmate
 - Piece movement, especially pawns. There were so many different considerations in the movement of a pawn that it took a lot of thought to design a system for handling piece movement.

In the end, I decided to add a Square struct that only contains the row/file of a square so that I wouldn't have to return so many rows/files from functions when I was returning two pieces. I could almost certainly have used this struct in functions created prior to the creation of Square, but I didn't go back through to refactor. This is something that I would possibly like to refactor later on.
