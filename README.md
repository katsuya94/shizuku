# Shizuku

## Bootstrapping

Log in as `root`

```sh
ssh -i ~/.ssh/id_rsa root@ip.address
git clone https://github.com/katsuya94/shizuku.git
shizuku/bootstrap.sh
```

Log in as `kacchan`

```sh
ssh -i ~/.ssh/id_rsa kacchan@ip.address
git clone https://github.com/katsuya94/shizuku.git
cd shizuku
bin/generate_server_key.sh
bin/bootstrap.sh
```
