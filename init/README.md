# init

The `init` folder contains system init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs for multiple platforms.

## Conventions

On unix systems agent configurations shall be kept under `/etc/dsn` and data under `/var/lib/dsn/`. These directories are assumed to exist, and the DSN binary is assumed to be located at `/usr/bin/dsn`.

## Agent configuration

The following configuration files are provided as examples:

    * `demo/server.yml`
    * `demo/client.yml`

Place one of these under `/etc/dsn.d` depending on the host's role. You should use `server.yml` to configure a host as a server or `client.yml` to configure a host as a client.

## systemd

On systems using `systemd`, the basic unit file under `systemd/dsn.service` starts and stops the dsn agent. Place it under `/etc/systemd/system/dsn.service`.

You can control DSN with `systemctl start|stop|restart dsn`.