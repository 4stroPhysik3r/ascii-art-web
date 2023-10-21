# _**Ascii-Art-Web - Stylize**_

### Description

ascii-art-web is a web application that changes the users input into a "Ascii-Art" representation. It is possible to choose from three different fonts: standard, shadow or thinkertoy from a dropdown menu.

### Usage: how to run

Type `go run .` into the console to start the server at port 3030.
You can now visit the URL: http://localhost:3030/ and enter your text there or click directly on the link in terminal.

### Implementation

The GUI (Graphical User Interface) is implemented with the help of the "html/template" package. The user can enter their text into the input field and choose a font. The input is then sent to the server via a POST request. The server then calls the "ascii-art" function to convert the input into Ascii-Art.
The result will be printed onto the ".../result"-page underneath the submit button.

### Author

- [Alexander Embacher](https://01.kood.tech/git/4stroPhysik3r)
