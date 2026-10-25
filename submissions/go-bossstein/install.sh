sudo apt update ;
sudo apt install nginx ;

rm -rf /usr/local/go
rm -f go1.27.3.linux-amd64.tar.gz
curl -LO https://go.dev/dl/go1.27.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.27.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin