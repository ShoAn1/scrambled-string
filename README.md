# Golang Project For Below Problem Statement 

```
Count how many of the words from a dictionary appear as substrings in a long string of
characters either in their original form or in their scrambled form. The scrambled form of the
dictionary word must adhere to the following rule: the first and last letter must be maintained
while the middle characters can be reorganised.

```

## Building Your Application

Run th Below Command In The Current Directory

```
$ go build -o scrmabled-strings 
```

## Running The Application Using Binary

Run th Below Command In The Current Directory
Update the path for Dictionary and Scrambled String 

``` 
$./scrambled-string --dictionary=[PATH OF DICTIONARY] --input=[PATH FOR SCRAMBLED STRING] 
```
### Details
```
The application requires two filepaths.
1. Path for Dictionary file .
2. Path for Grambled-String file .
```

## Building Docker Image For The Application

Update the dict.txt and strings.txt for Dictionary and Scrambled-String respectively in the current directory
Run the below command in the Current Directory

```
$ docker image build -t <image-name> .
```

## Running Docker Container For The Application 

Build the Docker image using previous step

```
$ docker run <image-name>
```
### Details
```
For Docker container the below two files can be updated in current Directory.
1. strings.txt -> filename for list of grambled strings representing each line as independent string
2. dict.txt    -> dictionary for list of keywords in dictionary representing each line as independent word in the dictionary
```
