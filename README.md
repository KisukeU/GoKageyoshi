# GoKageyoshi
Sometimes subdomains have their own NS servers, so this tool performs an AXFR check for subdomains.

At this version, Kageyoshi performs AXFR checks only marked NS servers. 
At this version, it is based on *nix dig utility, so it wouldn't work on Windows machines. 

@@ TO DO:
1)perform NS servers search for each subdomain and add new NS servers to checks;
2)perform subdomain enumeration automatically (for example, by parsing virustotal, google dorks outputs, dnsdumpster, and etc);
3)add WindowsOS support;
