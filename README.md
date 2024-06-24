# Project purpose

I created this project because I wanted to learn a new language over the summer, and the language I chose was golang. Although there are a million chess games out there, I wanted to choose something that would keep me interested in the project. I've been playing chess for years and am very familiar with the domain, so I thought creating a CLI chess app could be a lot of fun, and it was! 

# How to play

Run the application with the 'app' object file by typing './app' into the command line. Then, the user will be prompted to input proper chess notation. However, the user is not required to specify check or checkmate. For example, Rh3, Ngf3, Qxb7 would be proper input, but Rh3+, Ngf3+, Qxg7# would be improper input. 

# Future of the project

 - My main priority is to connect the game to a TCP server that I can use to play with my friends
 - Sometime this summer, I would like to implement a rematch function where the users will receive the opposite colored pieces that they received the game prior
 - I would like to log all of the moves and create functionality to import the game to a text file with all of the moves properly notated

# Reflection

Ultimately I had a lot of fun creating this project. I definitely still don't know how to write 'effective' Go, but I hope to continue writing it in my free time as I quite enjoyed working with it.

Some of my struggles in this project include: 
 - Detecting checkmate. I didn't want to have to find every single possible move for each piece to try to be more efficient, but I ultimately decided to in order to see if a piece could block a check/checkmate
 - Piece movement, especially pawns. I have to do math on the 2D array (the board) to try to determine where pieces can move, and I have to iterate through the board to see if pieces are in the way, see if pieces are proteceted (for king captures), I have to account for en passant, pawns can only move diagonally on captures, determine when a pawn can move 1 or 2 squares, and the list goes on. The amount of double nested for loops and ridiculous number of if statments felt really stupid, and I feel like it shouldn't be like that, but I have no idea.

In the end, I decided to add a Square struct that only contains the row/file of a square so that I wouldn't have to return so many rows/files from functions when I was returning two pieces. I could almost certainly have used this struct in functions created prior to the creation of Square, but I didn't go back through to refactor. This is something that I would possibly like to refactor later on.
