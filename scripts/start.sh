cd /vagrant

bin/caddy -pidfile='/vagrant/.caddypid' &
sleep 2
echo "Finished. Server is running in process "
cat /vagrant/.caddypid