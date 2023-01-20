# Creating a slave server

**This has been tested on ubuntu 20.04**

## Setup
Run the `/scripts/setup_ubuntu_slave.sh` script on the VPS.

That should be it. The script install dependencies and reboot the server and start the slave.

## Debugging 
### Startup issues

Start by inspecting the rc.local service

```bash
systemctl status rc-local.service
```
If anything failed here, try the commands manually:
```bash
git checkout rework # can be skipped
git -C /root/block pull
```
```bash
make -C /root/block compile
```
```bash
/root/block/bin/slave
```

If there are no issues there, check the syslog

```bash
journalctl -u rc-local.service
```