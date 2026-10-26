mkdir -p /usr/share/nginx ;

rm -f /etc/nginx/nginx.conf ;
cp nginx.conf /etc/nginx/nginx.conf ;

rm -f main
/usr/local/go/bin/go build main.go