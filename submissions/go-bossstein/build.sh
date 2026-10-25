mkdir -p /usr/share/nginx ;

rm -f /usr/share/nginx/nginx.conf ;
cp nginx.conf /usr/share/nginx/nginx.conf ;

rm -f main
/usr/local/go/bin/go build main.go