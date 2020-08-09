# Pathscan
Web path scanner
## Installing
Install and update using pip:  
``
root@kali# go build
root@kali# ./pathscan -h
``

## Help Section
``
Usage of ./pathscan:
  -t int
    	the num of threads (default 30)
  -u string
    	website url
``

## Example Section
``
root@kali# ./pathscan -t 40 -u "http://test/"
``
