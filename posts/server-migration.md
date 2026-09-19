---
title: Migrating my Server to Debian
date: 2026-09-18
tags: [it, linux, docker]
---
Unraid, I rebuke thee! I've been stuck with unraid for the past year, and I got sick of it. I don't like paying for security updates, especially with all the fucking exploits that are being discovered in the linux kernel. And manually applying patches is a pain. So I switched to a better os: Debian. 

What I lose is a convenient web server for docker container management, storage array management, and system monitoring. Each of these will have to be replaced by cli native, open source alternatives. I'm talking about Prometheus + Grafana (or maybe just use some syscalls and Go), docker cli, and mdadm. I've worked on two of these so far.

## The First Challenge: Data Integrity
When I first set up my server, I didn't really have to worry about messing up because there was no data to protect. I could wipe entire drives on a whim because there was no data on those drives. Now I have multiple terabytes worth of data, data that needs to be stored during the transition period.

Fortunately, Unraid has each of the drive store its data in an easily understandable xfs format. The plan, then, is simple. First, make a backup of everything on the soon to be boot drive (this is my nvme ssd cache drive). I made two, one on my desktop, one on my laptop, both with a simple `rsync -aHAX`. `rsync` was my best friend throughout this entire process; thank you `--info=progress2`. Backup made, you shutdown the server and physically disconnect the drives. They you boot up, adjusting bios and installing the operating system. The boot drive gets a couple partitions: `/boot`, `/`, an 8GB swap, and the all important `/srv`, which is where server files, like this website, go to be served.

Next, you reconnect the drives and mount their filesystems. Then there's a little dance of files as we try to set up the RAID arrays. Firstly, for each drive array we want, we need to make sure one of the mirror pair drives is free (my arrays are RAID1 and RAID10, setup using mdadm), well, at least backed-up. We then initialize the RAID arrays, taking care to build properly:

```
    sudo mdadm --create /dev/md0 \
        --level=10 \
        --raid-devices=4  \
        /dev/sdb missing /dev/sdc missing
```

The `missing` allows us to build a partial RAID array, which lets us copy our file from the backup to the newly (partially) built array; use `rsync`. We then use `mdadm --add` which adds the backup drives, now that their data is on the array. This process kinda sucked. It took like 4 days for all of the arrays to build properly, and put a lot of stress on the system. But I'm glad it's working now. We finally copy our cache drive contents to the new boot drive, and we are ready to go!

## The Second Challenge: Setup from Scratch

Kinda, data was just one part of the puzzle, the other is getting the internals setup. The system essentially had the basic utils and an open-ssh server. That was it. The first thing I did was use `apt` to install all of my favorite utilities and used git to clone their configuration files (you can check out the dotfiles on my [github](https://github.com/kawiggles/kawiggles-dotfiles). 

Then we go about securing the server. Tailscale is amazing, and it allows us to really restrict traffic immediately. fail2ban, uf2, and nginx handle incoming traffic. We make sure ssh keys are properly transferred and secured, and edit our `ssh_config` to only allow certain users on certain ports only with certain ssh keys to authenticate.

To be honest, still in the middle of setting up docker containers. Got distracted by a talk I gave at OWASP and writing a shell in Zig. More on both soon!
