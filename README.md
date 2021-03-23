# NucleServer 
[![Build Status](https://travis-ci.org/nucleome/nucleserver.svg?branch=master)](https://travis-ci.org/nucleome/nucleserver)
[![Releases](https://img.shields.io/github/release/nucleome/nucleserver.svg)](https://github.com/nucleome/nucleserver/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nucleome/nucleserver)](https://goreportcard.com/report/github.com/nucleome/nucleserver)
[![Licenses](https://img.shields.io/badge/license-gpl3-orange.svg)](https://opensource.org/licenses/GPL-3.0)
[![GitHub Repository](https://img.shields.io/badge/GitHub-Repository-blue.svg)](https://github.com/nucleome/nucleserver)

[*NucleServer*](http://doc.nucleome.org/data/server) is a standalone tool to host a data service for [*Nucleome Browser*](https://vis.nucleome.org). You can use this tool to host either genomic data or 3D structural modeling data on a personal computer or data server. For genomic data, it supports common data types such as genome tracks in [bigWig](https://genome.ucsc.edu/goldenpath/help/bigWig.html), [bigBed](https://genome.ucsc.edu/goldenpath/help/bigBed.html), [.hic](https://github.com/aidenlab/Juicebox/blob/master/HiC_format_v8.docx), and [tabix](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC3042176/) formats. For 3D structural modeling data, it supports customized [Nucle3D](https://github.com/nucleome/nucle3d) format. You also view the [documentation](https://nb-docs.readthedocs.io/en/latest/data_service.html) for more examples. 

## Quick Start
Below is a quick demo showing you how to set up a genomic data service with sample data. 

First, you need to download the pre-compiled executable files from [the latest release](https://github.com/nucleome/nucleserver/releases) in this GitHub repository. 
We provide pre-compiled files in Linux, Windows, and Mac OS. 
You can directly download the pre-compiled program based on your operating system. 
If the program is not compatible with your machine, you can also try to compile it by cloning this repository.

> TIPS: If you are using Windows and not familiar with running a command-line tool in Windows, please refer to [this article](https://www.computerhope.com/issues/chusedos.htm). If everything goes well, you should be able to run `nucleserver` as a command-line tool in the terminal.

> TIPS: Please also note that you need to grant the correct permission to the program. In Linux/Mac, this can be done with the command ``` chmod +x nucleserver ```.

Second, you need to prepare an excel file for configuration. 
For this quick demo, we have prepared this [template file](https://docs.google.com/spreadsheets/d/1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM/edit#gid=1744383077).
You can download it and save it with .xlsx as the suffix of the file (let's name in nucle.xlsx in this demo).
The content of this excel shows metatable pointing to a bigBed file host remotely by ENCODE on the Internet.
You can also download the data in your own machine and modify the excel to reflect the correct file location in the datasheet (see below)

Finally, you can start the data server using the following command: 

```
nucleserver start -i nucle.xlsx
```

If everthing goes well, you should see messages showing you datasheet has been added to a local data server. The URL of the data server in this case (i.e., local machine with a default 8611 port) is:

```
http://127.0.0.1:8611
```

> TIPS: Please note that you don't have to add this particular URL (http://127.0.0.1:8611) to Nucelome Browser. This localhost URL is one of the default servers in Nucleome Browser. If you start a data server under this URL, you can just reload server content or add a new genome browser panel after the local server starts. Your custom data will automatically show up in this genome browser config panel. If the data server URL is differnt from the URL mentioned above, you would have to add it manually to [Nucleome Browser](https://vis.nucleome.org). Please see [the document](https://nb-docs.readthedocs.io/en/latest/data_service.html#genomic-data) for details. 
> TIPS: If you do not know how to add a new genome browser panel, you can watch animation [here](https://nb-docs.readthedocs.io/en/latest/animation.html#panel-oraganization). Basically, in [Nucleome Browser](https://vis.nucleome.org), there is a plus button on the top menu bar. You can click it and select the genome browser panel. In a genome browser panel, you can then config data service using the configuration interface. 

![screenshot](https://nucleome.github.io/image/configServers.png) 

## Install by compiling source code
If the binary is not working for you or you want to install the latest experimental version, you can also compile from the source code. 
NucleServer is implemented in [GoLang](https://golang.org) ( version > 1.11 ) and hosted on Github. 
With the Golang environment installed, the source code can be cloned simply by the following command.
```
go get -u github.com/nucleome/nucleserver
```

##  Start a data service
The following command is used to start a data service in Mac OS or Linux:
```shell
./nucleserver start -i [google sheet id or excel file] -p [port default:8611]
```
Thhis command to start a data service in Windows:
```shell
nucleserver.exe start -i [google sheet id or excel file] -p [port default:8611]
```
> TIPS: The **Google Sheet ID** can be found as part of the URL in the Google Sheets webpage. It is shown with a blue background in the following screenshot.
![demo image](https://nucleome.github.io/image/google_sheet_id_demo.png).
> TIPS: If this is the **first time** you run `nucleserver` with a google sheet, the program will firstly print a web link in the terminal, asking you for permissions. Please visit the link by copying it into a web browser and grant the permissions following the instructions from Google. Google should provide you a token. Please then paste this token into the terminal. After the authorization step is done, a credential token will be stored in `[Your Home Dir]/.nucle/credentials/gsheet.json`. 

## Configuration excel file

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

If you would like to add local files.

- In the Config sheet, define the root variable as a PATH to the data folder, such as `/home/yourusername/trackdata`.
- In the "ENCODE_ChIPSeq" sheet, you can use the URL directly or change the URL to a relative PATH pointing to the bigBed file, such as `./ENCFF845IDA.bigBed`, if you download it to your local drive. If you use the URL directly, NucleServer will only fetch the index and store it in "$HOME/.nucle/index".


## Other hints

### Start a data service in background.
You many want to put a data servic in background, using command line such as nohup. The simple command using nohup is provided below.
```
nohup nucleserver start -i [nucle.xlsx] &
```

### Host private and public data for community in "HTTPS"
We highly recommend the host servers to support "HTTPS", as it promote the browser's functionality in progressive web application, google based permission management and session storage. If the data is sensitive, you can also host it locally. It is then not accessible by other users or web application administrator. In addition, we also provides a simple password protection option (currently experimental) for user access data in internet. As demonstrated below, user can add a password when starting the server.
```
nucleserver start -i [nucle.xlsx] -c password
```
As an result, only users login with the password through the following webpage can access the hosted data.
```
http://yourwebsite:8611/main.html
```
### Converting local data server to the public using reverse Proxy
A Reverse Proxy implemented in GoLang [Traefik](https://traefik.io/) is recommended for convert local data services to a https global data service.  [Nginx](https://www.nginx.com/) is also working here. 


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

### Host your data in servers without https
If you host your data in servers without ssl certification.
Either you can use "http://vis.nucleome.org" which supports fetch data from non-ssl web sites.
Or you can make a tunnel between your server with localhost using the command below.
```
ssh -N -L 8611:localhost:8611 server
```

# Depedencies
- fetch data from bigwig and bigbed files https://github.com/pbenner/gonetics
- fetch data from tabix files github.com/brentp/bix
