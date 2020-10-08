# GoKageyoshi
Sometimes subdomains have their own NS servers, so this tool performs an AXFR check for subdomains.
<br>
At this version, Kageyoshi performs AXFR checks only marked NS servers.<br> 
At this version, it is based on *nix dig utility, so it wouldn't work on Windows machines. <br>
<br>
@@ TO DO:
<br>
1)perform NS servers search for each subdomain and add new NS servers to checks;
<br>
2)perform subdomain enumeration automatically (for example, by parsing virustotal, google dorks outputs, dnsdumpster, and etc);
<br>
3)add WindowsOS support;
