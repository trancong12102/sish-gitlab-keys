# Deploy

Make `.env` file from `.env.example` and fill in the values.

```shell
cp .env.example .env
```

Edit [config.yml](config.yml) and [le-config.yml](le-config.yml) to match your domain and DNS auth info.

Then, start the Let's Encrypt container to get the certificates:

```shell
docker compose up -d
```

Then, create a symlink that points to your domain's Let's Encrypt certificates like:

```shell
ln -s /etc/letsencrypt/live/<your domain>/fullchain.pem ssl/<your domain>.crt
ln -s /etc/letsencrypt/live/<your domain>/privkey.pem ssl/<your domain>.key
```

**Careful**: the symlinks need to point to `/etc/letsencrypt`, not a relative path. The symlinks will not resolve on the
host filesystem, but they will resolve inside the sish container because it mounts the letsencrypt files
in `/etc/letsencrypt`, not `./letsencrypt`.

Finally, you can restart the sish container:

```shell
docker compose restart
```
