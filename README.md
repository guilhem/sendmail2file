# sendmail2file

A sendmail drop-in replacement that push mail in a file on json format

## Installation

```sh
go install github.com/guilhem/sendmail2file
```

## Usage

```sh
echo -e "Subject:Test Mail using sendmail2file\n\nThis is a Test Message" | sendmail2file
```
