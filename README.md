HTTP-Traceroute
=======

This project has been inspired by Digininja's HTTP traceroute available [here]() and because I wanted to give it a shot with Go. 

#### How does it work? 

Install Go and compile: 

```
go build http-traceroute.go
```

Then, launch it: 

```
$ ./http-traceroute "http://www.yahoo.com"
Initial request to: http://www.yahoo.com 

#### Redirected ####

Request for: http://www.yahoo.com
Redirected to: https://www.yahoo.com/
Status code: 301


#### Headers ####

Date :  Tue, 07 Jul 2015 12:38:52 GMT
Location :  https://www.yahoo.com/
Content-Type :  text/html
Connection :  keep-alive
Via :  http/1.1 ir14.fp.ir2.yahoo.com (ApacheTrafficServer)
Server :  ATS
Content-Language :  en
Cache-Control :  no-store, no-cache
Content-Length :  375


#### Cookies ####

No cookies created.


#### Redirected ####

Request for: https://www.yahoo.com/
Redirected to: https://uk.yahoo.com/?p=us
Status code: 302


#### Headers ####

Server :  ATS
Date :  Tue, 07 Jul 2015 12:38:52 GMT
X-Frame-Options :  DENY
Strict-Transport-Security :  max-age=2592000
Vary :  Accept-Encoding
Content-Type :  text/html; charset=utf-8
Age :  0
Connection :  keep-alive
P3p :  policyref="http://info.yahoo.com/w3c/p3p.xml", CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE LOC GOV"
Cache-Control :  private
Set-Cookie :  DNR=deleted; expires=Mon, 07-Jul-2014 12:38:51 GMT; path=/; domain=.www.yahoo.com
Via :  http/1.1 usproxy2.fp.bf1.yahoo.com (ApacheTrafficServer), http/1.1 ir1.fp.ir2.yahoo.com (ApacheTrafficServer)
Location :  https://uk.yahoo.com/?p=us


#### Cookies ####

DNR=deleted; expires=Mon, 07-Jul-2014 12:38:51 GMT; path=/; domain=.www.yahoo.com
DNR=deleted; expires=Mon, 07-Jul-2014 12:38:51 GMT; path=/; domain=.www.yahoo.com
hpc=d=s7s1cD8m54yOZTEyvcWpFnm1DjHePFelGJ1IrykZsEDgYQv_0k31CI.3TxxEMryhiutdsEksWwGcVuGgC9ZHbDzOx2FETDjY_A6j1Sb8A3aq.9egk4X0fgu7LFrwZBIKsKKnoJn8vY_U9iowuufFuFune57kF6.i9s7ABAGIJLbffXBtItlmP4XJusUenPyN4ZtyZ4k-&v=2; expires=Wed, 06-Jul-2016 12:38:52 GMT; path=/; domain=www.yahoo.com


Final URL is https://uk.yahoo.com/?p=us  with status code:  200


#### Headers ####

Age :  0
Connection :  keep-alive
Via :  http/1.1 ir1.fp.ir2.yahoo.com (ApacheTrafficServer)
Server :  ATS
Cache-Control :  no-store, no-cache, private, max-age=0
Expires :  -1
Date :  Tue, 07 Jul 2015 12:38:53 GMT
Vary :  Accept-Encoding
Content-Type :  text/html; charset=utf-8
P3p :  policyref="http://info.yahoo.com/w3c/p3p.xml", CP="CAO DSP COR CUR ADM DEV TAI PSA PSD IVAi IVDi CONi TELo OTPi OUR DELi SAMi OTRi UNRi PUBi IND PHY ONL UNI PUR FIN COM NAV INT DEM CNT STA POL HEA PRE LOC GOV"
Strict-Transport-Security :  max-age=2592000
X-Frame-Options :  DENY


#### Cookies ####

No cookies created.

```

#### Contributing

Feel free to report any bug, or if you want to do any Pull Request.
Code is a bit messy and was mostly to get an overview of the language. 