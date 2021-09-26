kill -9 $(pgep webserver)
cd ~/newweb
git pull https://github.com/YTRei/newweb.git
cd webserver/
./webserver &