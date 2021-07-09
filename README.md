# Prerequisites (applicable for Ubuntu)

## Install go

```bash
wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
```

## Install docker and docker-compose

```bash
sudo apt-get install -y docker-compose docker-ce docker-ce-cli containerd.io
```


# How to use

## Start application

Go to `technical_leboncoin` folder and issue the following command ;
```bash
docker-compose up --build
```

## Make requests

You can make requests to 2 endpoints :

```
http://localhost:8080/api/fizzbuzz/request?int1={int1},int2={int2},limit={limit},str1={str1},str2={str2}
```
Returns a list of numbers from 1 to `limit`, where multiples of `int1` are replaced by `str1`, multiples of `int2` are replaced by `str2`, and multiples of `int1` and `int2` are replaced by `str1str2`


```
http://localhost:8080/api/fizzbuzz/statistics
```
Returns the parameters corresponding to the most used request, as well as the number of hits for this request