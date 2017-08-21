This little program is my solution to a programming assignment I got as part of the interview process at a company where I 
applied for a freelance position. It does one thing only: show CSV files or fixed-width files in a simple web app.

This is my first program in Go. I put it on Github to be part of my portfolio. Why it deserves street-cred: I wrote it 
in only 4 - 6 hours, without having ever written Go before. It shows my ability to quickly learn a new language and get 
productive with it.

Read on for more info.

## Implementation

You can run this app as follows:

- Make sure you have Go installed and your `$GOPATH` configured and add to your `$PATH`
- Enter the following in your terminal:

```
# might require password
$ go get github.com/DandyDev/data-file-viewer
$ data-file-viewer
```

The application is now serving requests at `http://localhost:8080`

### Features

The application does the following:

- The application will serve your CSV files and fixed-width files in a nice-looking table (Bootstrap thanks!)
    - The path is as follows: `/view/<path/to/your/file>`. The path needs to be relative to the directory you're running 
      the app from
    - The file extension determines how the file is parsed. Supports CSV as `.csv` and fixed-width as `.prn`
- It will *automagically* infer the file encoding, and convert to proper UTF-8. Even the funky encoding of the example 
  files should pose no problems
- For fixed-width:
    - You can supply the column names as a query parameter: `?columns=col1,col2,col3,col4`. Names are case-insensitive. 
      The app will use the supplied column names to calculate where columns begin and end
    - When you don't supply column names, the app will _try_ to infer the column edges by finding common whitespace.
      This will *only work* when there's at whitespace between each column of at least 1 character wide. The supplied 
      `Workbook2.prn` actually doesn't have this because there's no whitespace between the _Phone_ and _Credit Limit_ 
      columns for every row. To be able to demonstrate this feature, I supplied a modified version of this file, 
      `Workbook3.prn` that _does_ have whitespace between these columns on every row.
    
### Caveats

- This is my first Go-code. I gonna assume the code style sucks, and that I missed some obvious best practices, 
  but I'm happy to learn!
- I tried to do error handling wherever it was sensible to do so, but I'm pretty sure I didn't forsee every scenario
- I'm aware of the good practice of writing tests, especially for _critical_ parts of the software, but between learning 
  Go and being clever with extra features, I kinda ran [out of time](http://giphy.com/gifs/aint-nobody-got-time-for-that-gif-hfvihQ6LF9x2o) 
  to write many tests. Try out the `testing` pkg anyways, I wrote one test to test the function that infers column offsets
