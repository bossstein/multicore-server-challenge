sudo apt update ;
sudo apt install nginx ;

[ ! -f "go1.27.3.linux-amd64.tar.gz" ] && rm -f go1.27.3.linux-amd64.tar.gz && curl -LO https://go.dev/dl/go1.27.3.linux-amd64.tar.gz ;
[ ! -d "/usr/local/go" ] && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.27.3.linux-amd64.tar.gz ;