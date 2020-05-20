# NucleServer 
[![Build Status](https://travis-ci.org/nucleome/nucleserver.svg?branch=master)](https://travis-ci.org/nucleome/nucleserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/nucleome/nucleserver)](https://goreportcard.com/report/github.com/nucleome/nucleserver)
[![Licenses](https://img.shields.io/badge/license-bsd-orange.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub Repository](https://img.shields.io/badge/GitHub-Repository-blue.svg)](https://github.com/nucleome/nucleserver)

[*NucleServer*](http://doc.nucleome.org/data/server) is a simple standalone tool to host an additional data server for [*Nucleome Browser*](https://vis.nucleome.org). A typical usage is to host a set of cumtomized genome data files that is not on the default server, such as additional genome tracks in [bigWig](https://genome.ucsc.edu/goldenpath/help/bigWig.html), [bigBed](https://genome.ucsc.edu/goldenpath/help/bigBed.html) and [.hic](https://github.com/aidenlab/Juicebox/blob/master/HiC_format_v8.docx) formats. To facilitate users with limited commandline experience, we also implemented an simple GUI called [*NucleData*](https://github.com/nucleome/nucledata). However, it's only working for setting up a **local server** on a personal PC for now.


## Quick Start
This is a quick demo on setting up a typical server with sample data. To start, you can download the pre-compiled excutables from the main server (Current Build Version: 01-07-2020  v0.1.5). If they are not compatible to your machine, you can try to compile it from source.
- [Linux](https://vis.nucleome.org/static/nucleserver/current/linux/nucleserver)
- [MacOS](https://vis.nucleome.org/static/nucleserver/current/mac/nucleserver)
- [Windows](https://vis.nucleome.org/static/nucleserver/current/win64/nucleserver.exe)

> If you are using Windows and not familiar with runnning command line tool in Windows, please read [this article](https://www.computerhope.com/issues/chusedos.htm) first. Then,you can run `nucleserver` as a command line tool in terminal.

> As a side note, please note that you'd have to grant the excutable the correct permission. In linux/Mac, this can be done with the command ``` chmod +x nucleserver ```.

The next step is to create an excel table for configurations. You can download a simple template [Here](https://docs.google.com/spreadsheets/d/1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM/edit#gid=1744383077). 
Please note this google sheet can be saved in .xlsx format (now called nucle.xlsx in this demo). 
This excel table will point to a bigBed file host remotely by ENCODE through the Internet. 
If you'd like to host the data in your own server, you can download this [bigBed file](https://www.encodeproject.org/files/ENCFF845IDA/@@download/ENCFF845IDA.bigBed) manually. 
And then, you can point to this local file by modifying the followings in the nucle.xlsx file.
The following command will start the data server.
```
nucleserver start -i nucle.xlsx
```
If you would like to add local files.

- In the Config sheet, define the root variable as a PATH to the data folder, such as `/home/yourusername/trackdata`.
- In the "ENCODE_ChIPSeq" sheet, you can use the URL directly or change the URL to a relative PATH pointing to the bigBed file, such as `./ENCFF845IDA.bigBed`, if you download it to your local drive. If you use the URL directly, NucleServer will only fetch the index and store it in "$HOME/.nucle/index".

You many want to put the process in background using **screen** or nohup. The simple command using nohup is provided below.
```
nohup nucleserver start -i nucle.xlsx &
```
If everthing goes fine, you should then be able to add this additional server to your browser. The URL can be the following if you are running the server at a local machine with the default 8611 port.
```
http://127.0.0.1:8611
```
> Please note that you don't have add this particular URL. The localhost http://127.0.0.1:8611 is one of default servers in Nucleome Browser. If user starts a data server in localhost and the port is the default 8611, you can just reload server content or add new genome browser panel after the local server start, the custom data will show up in this genome browser config panel.

If the data server location is differnt from the URL mentioned above, you'd have to add it manually to [Nucleome Browser](https://vis.nucleome.org).  
> If you don't have a genome browser panel to start with, please add one at first. The add button is in submenu of panels in the menu bar. After clicking it, please follow this little guide: "Click Config tracks → Click Config Servers → Input Server URI and any Id you'd like into table → Click Refresh Button to reload". 
![screenshot](https://nucleome.github.io/image/configServers.png) 

> If you open a new genome browser panel, it will automatically copy the previous configurations. 

## Install by compiling source code
Users can download the compiled binaries for Linux, Mac and Windows OS as described in the quick start. 
However, if the binary is not working or you are trying to install the most recent experimental version, you can alway compile from the source code. NucleServer is implemented in [GoLang](https://golang.org) ( version > 1.11 ) and hosted on Github. With the Golang environment installed, the source code can be cloned simply by the following command.
```
go get -u github.com/nucleome/nucleserver
```

##  Start a data service
The command to start a data service in Mac OS or Linux is the following.
```shell
./nucleserver start -i [google sheet id or excel file] -p [port default:8611]
```
The command to start a data service in Windows is the following.
```shell
nucleserver.exe start -i [google sheet id or excel file] -p [port default:8611]
```

> The **Google Sheet ID** can be found as part of the url in the google sheet webpage. It is indicated by a blue background in the following screenshot.
![demo image](https://nucleome.github.io/image/google_sheet_id_demo.png).

> If this is the **first time** you are using `nucleserver` with google sheet, it will firstly print a web link in terminal, asking for permissions. Please visit this link in a browser and grant the permissions. Google should provide you a token in respond. Please then enter this token in the terminal. As the result,  a credential token will be stored in `[Your Home Dir]/.nucle/credentials/gsheet.json`. 

## Config input file

A config input file can either be an Excel file or a Google Sheets. The file must contain two sheets, namely "Config" and "Index".  
- The “Config” sheet stores the configuration variable values. Currently, `root` variable is the only variable needed for NucleServer. It is the root path for you store all track data files. (As a result, user can easily migrating data between servers.) All the URI/PATH in other sheets will be relative to this `root`. The only exception is for URIs starting with `http` or `https`.
![screenshot](https://nucleome.github.io/image/sheetConfig.png).
- The “Index” sheet stores the configuration information for organizing the track groups, each with a unique sheet title. The sheet titles not present in Index sheet will be ignored by the browser. The Name and Value columns define the corresponding columns in the track group sheet. 
![screenshot](https://nucleome.github.io/image/sheetIndex.png).
- The track group sheets provide information such as file location(uri), short label(shortLabel), long label(longLabel) and weblink(metaLink) for the tracks. As mentioned, these data files can be files in a local personal PC/server or an web link pointing to a remote server. 
- If the track group sheet contains four columns, the columns name should be "shortLabel", "uri", "metaLink" and "longLabel”. The corresponding column header in the "Index" sheet should be "A" and "B,C,D", so that they are defined accordlingly. 
![screenshot demo](https://nucleome.github.io/image/sheetData4.png) 

- If using two columns, the column name can be any string user defined. Please just filled the "Index" sheet accordingly.
![screenshot demo](https://nucleome.github.io/image/sheetSimpleData.png)
> In sheet "Index", those entries which Id starts with “#” will be ignored when loading. Column "Type" is designed for future data type. Currently, please just use "track" in this column. It support bigWig, bigBed and .hic format files.


## Other hints
### Host private and public data for community in "HTTPS"
We highly recommend the host servers to support "HTTPS", as it promote the browser's functionality in progressive web application, google based permission management and session storage. If the data is sensitive, you can also host it locally. It is then not accessible by other users or web application administrator. In addition, we also provides a simple password protection option (currently experimental) for user access data in internet. As demostrated below, user can add a password when starting the server.
```
nucleserver start -i nucle.xlsx -c password
```
As an result, only users login with the password through the following webpage can access the hosted data.
```
http://yourwebsite:8611/main.html
```
### Converting local data server to the public using reverse Proxy
A Reverse Proxy implemented in GoLang [Traefik](https://traefik.io/) is recommended for convert local data service to https global data service.  [Nginx](https://www.nginx.com/) is also working here. 


### Using Reverse Proxy to host more data services in same domain
Nucleome Browser default data services are "/d/portal" and "https://127.0.0.1:8611".
Nucleome Browser supports URL like "https://[youdomain]/path/to/dataservice". 


### Build an Entry to A Nucleome Browser with customized data services. 
Easiest way is configure your panel and save as a session to your google sheet.
Copy this saved session to a Google Sheet with shareable view link.
The entry will be on the following link.
```
https://vis.nucleome.org/entry/?sheetid=[your public google sheet id]&title=[Sheet1]
```

### Local index for remote data
If acessing data from other servers such as ENCODE, NucleServer will fetch index from the web link and store them locally, which is on average 1% of the original data file in size. It is stored in `[Your Home Dir]/.nucle/index`. As a result, while browsing the genome, NucleServer will fetch the corresponding data from ENCODE each time based on the index. 

