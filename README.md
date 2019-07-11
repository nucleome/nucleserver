# NucleServer 
[![Build Status](https://travis-ci.org/nucleome/nucleserver.svg?branch=master)](https://travis-ci.org/nucleome/nucleserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/nucleome/nucleserver)](https://goreportcard.com/report/github.com/nucleome/nucleserver)
[![Licenses](https://img.shields.io/badge/license-bsd-orange.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub Repository](https://img.shields.io/badge/GitHub-Repository-blue.svg)](https://github.com/nucleome/nucleserver)

[*NucleServer*](http://doc.nucleome.org/data/server) is a command line tool for users to start a [*Nucleome Browser*](https://vis.nucleome.org) data service in local or remote servers to host their multi-resolution genome related data files, such as [bigWig](https://genome.ucsc.edu/goldenpath/help/bigWig.html), [bigBed](https://genome.ucsc.edu/goldenpath/help/bigBed.html) and [.hic](https://github.com/aidenlab/Juicebox/blob/master/HiC_format_v8.docx). 

If you are looking a GUI tool to host these files in local computer, please use this tool [*NucleData*](https://github.com/nucleome/nucledata). 


## Install

This software is implemented in [GoLang](https://golang.org/).
User can either download the binary exectuable files we have compiled, or compile from source code.

### Download Binary Executable Files
 
Download Binary Exectuable Files in Linux, Windows and Mac OS without installation.

Current Build Version: 07-10-2019  v0.1.2

- [Linux](https://vis.nucleome.org/static/nucleserver/current/linux/nucleserver)
- [Windows](https://vis.nucleome.org/static/nucleserver/current/win64/nucleserver.exe)
- [MacOS](https://vis.nucleome.org/static/nucleserver/current/mac/nucleserver)

Then change the mode of this file into executable. In Linux or Mac OS, this can be done in a terminal, using command `chmod`.

```shell
chmod +x nucleserver
```

### Alternativley, Compile From Source Code
NucleServer is implemented in [GoLang](https://golang.org) ( version > 1.11 ).  
```
go get -u github.com/nucleome/nucleserver
```

## Get Started with Examples

### Quick Start

[Example Input: Google sheet](https://docs.google.com/spreadsheets/d/1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM/edit#gid=1744383077)

For a quick start, please download the Google Sheet above as an excel file and named it as `nucle.xlsx`.
Then run the command below in your local computer.

`nucleserver start -i nucle.xlsx`

Or skip downloading and use **Google Sheet ID** directly like this.

`nucleserver start -i 1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM`

> The **Google Sheet ID** is part of the url in the google sheet webpage. It is in blue background in the following demostration image.
> ![Google Sheet ID Demo](https://nucleome.github.io/image/google_sheet_id_demo.png)

> When **first time** use `nucleserver` with google sheet, it will prompt a link in terminal to ask for permission to access user's Google Sheets, copy this link to browser and get back a token, then copy and paste the token to the command terminal, a credential token will be stored in `[Your Home Dir]/.nucle/credentials/gsheet.json`. 

After the data service is ready. Open [Nucleome Browser](https://vis.nucleome.org) in your web browser. You should be able to browsing MTA1 ChIPSeq narrow peaks from ENCODE project. However, the bigBed data is not downloaded to your computer yet. NucleServer fetch index from ENCODE http web link and store the index, which is average one percent data file size in `[Your Home Dir]/.nucle/index`. When you browsing genome, NucleServer will fetch the corresponding data from ENCODE each time. 

### Local Files

We would like to demonstrate how to start a data service with local files.

- Download this example file [MTA1 ChIPSeq narrow peaks bigBed](https://www.encodeproject.org/files/ENCFF845IDA/@@download/ENCFF845IDA.bigBed) from ENCODE to a directory, for example `~/Downloads`.
- Open the [Example Input Template](https://docs.google.com/spreadsheets/d/1gdK9L2DuJ7hln1ouLy8pQcvX6Fbrm6EUv28Al7ivmKw/edit?usp=sharing) sheet. Download it 
as an excel file and named it as `nucle.xlsx`.
- Change root variable to your home directory such as `/home/yourusername` in Config sheet.
- Change the uri of MTA1 entry in "ENCODE_ChIPSeq" sheet to the file relative path to root.
Start nucleserver.

`nucleserver start -i nucle.xlsx`

This time you should be browsing MTA1 ChIPSeq narrow peaks and the file are stored in local drive.



## User Manual 

###  Start a data Service

in Mac OS or Linux
```shell
./nucleserver start -i [google sheet id or excel file] -p [port default:8611]
```
in Windows 
```shell
nucleserver.exe start -i [google sheet id or excel file] -p [port default:8611]
```

If you are using Windows and not familiar with runnning command line tool in Windows, please read [this article](https://www.computerhope.com/issues/chusedos.htm) first. Then,you can run `nucleserver` as a command line tool in terminal.

The track configuration input for nucleserver could be an Excel file or Google Sheet ID.

### Input file format

User's private data are not accessible by other users or web application administrator if his/her data server is in localhost. *NucleServer*　also provides a simple password protection option for user access data in internet.

The input is an Excel file or a Google Sheets which has the basic information such as file location(uri), short label(shortLabel), long label(longLabel) and weblink(metaLink) of further track description. These data files can be either located in local drive or just an http web adress link.

Two sheets "Config" and "Index" are required for start this data server.  

“Config” sheet stores the configuration variable values. Currently, `root` variable is the only variable needed for NucleServer. It is the root path for you store all track data files. It is designed for user conveniently migrating data between servers. All the URI in other sheets will be the relative path to `root` if their URI are not start with `http` or `https`.

![Sheet Config Example](https://nucleome.github.io/image/sheetConfig.png)

The “Index” sheet stores the configuration information of all other sheets which are needed to use in NucleServer. The sheet titles which are not in Index sheet will be ignored.

![Sheet Index Example](https://nucleome.github.io/image/sheetIndex.png)

For track format data sheet, if using four columns, the columns name should be “shortLabel” , “uri,metaLink,longLabel”, and the corresponding column header such as A,B et al. should put into the 4th and 5th column.

If using two columns, the column name could be any string user defined. Just filled in the column index into the fourth and the fifth column accordingly. In sheet "Index", those entries which Id starts with “#” will be ignored when loading.
Column "Type" is a reserve entry for future data server. Currently, just use "track" in this column. It support bigWig, bigBed and .hic format files.
#### Simple Name and URI
![Sheet Data Example](https://nucleome.github.io/image/sheetSimpleData.png)

#### With Long Label and Meta Link
![Sheet Data Example](https://nucleome.github.io/image/sheetData4.png)


The localhost http://127.0.0.1:8611 is one of default servers in Nucleome Browser. If user starts a data server in localhost and the port is default 8611, user doesn’t need to configure the server list. Just reload server content or add new genome browser panel after the local server start, the custom data will show in this genome browser config panel.

If Data server is in other port or other web servers instead of localhost, user need to add the server into server lists. Open the [Nucleome Browser](https://vis.nucleome.org) in your chrome browser. 

If user don't have a genome browser panel, please add a genome browser panel, the add button is in submenu of panels in the menu bar. Then, in this genome browser, then Click Config tracks → Click Config Servers → Input Server URI and any Id into table → Click Refresh Button to reload.


![Config Servers](https://nucleome.github.io/image/configServers.png)

If user open a new genome browser panel , it will loading servers as last configuration. Servers configuration is stored as settings for this panel, if user duplicate this panel, the servers setting will be automatically copied too.


## Host public data for community in "HTTPS"

### Why we need https
*Nucleome Browser* in "HTTPS" would provide more functions than "HTTP", such as Progressive Web Application, using private Google Sheets or store sessions in Google Sheet. However, it only can fetch data service from "HTTPS" or localhost due to web security reason.

### Solution: Reverse Proxy
A Reverse Proxy implemented in GoLang [Traefik](https://traefik.io/) is recommended for convert local data service to https global data service. 

[Nginx](https://www.nginx.com/) is also working. 


### Using Reverse Proxy to host more data services in same domain
Nucleome Browser default data services are "/d/portal" and "https://127.0.0.1:8611".
Nucleome Browser supports URL like "https://youdomain.com/path/to/dataservice". 


### Build an Entry to A Nucleome Browser with customized data services. 
Easiest way is configure your panel and save as a session to your google sheet.
Copy this saved session to a Google Sheet with shareable view link.
The entry will be on the following link http(s)://vis.nucleome.org/v1/pub.html?sheetid=[your public google sheet id]

## Alternative Way to provide public data 
Provide a Google Sheet with public data web links. User can start a local service with this google sheet.  It would be even better if data hosters can provide tar file of pre build index files to download.

## Host private data in internet with password protection (Experimental)
`nucleserver start -i nucle.xlsx -c password`

http://yourwebsite:8611/main.html to sign in with `password`

## TODOs
- Supporting Large Set Data Host
