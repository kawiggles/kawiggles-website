---
title: Test Post
date: 2026-08-30
tags: [coding, retrospective]
---

This post is partially a literal test of my new website's blog functionality, and partially a retrospective on why I thought it necessary to rewrite my website in the first place. There's basically three primary motivators: maturation of skill, expression of identity through aesthetics, and practicality.

# I got better at coding

When I first built my website, I was extremely new to programming as a whole. I had just set up my github account only a few weeks prior for the purpose of tracking my now defunct game project that I was writing in C++. I had no idea how to make a functioning website, and new little about websites beyond what I had learned about HTML and CSS in middle school. My old website was a product of this skillset; it was essentially a few HTML files alongside a stylesheet being served by nginx. I was proud of this website. It showed that I knew how to deploy nginx in a docker container running on my server, wire up network rules properly, and serve it out to the world. 

It's been six months since then, and in that period of time I've put serious effort into learning more about how computers work. Firstly, I've learned both Go and Rust to an extent which I never achieved in either C or C++. I've figured out how to build actual things, including web servers. Having built a solid understanding of network protocols and programming in general, it seemed silly to have my website be restricted to a set of static files served by an external application when I can build a server myself. It's not particularly difficult with Go. And most importantly, writing my website myself gives me a chance to build a real application that is used every day, even if I'm the only user. Building this website is therefore simultaneously a form of skill expression and a chance to challenge myself to build something practical and usable.

# The old website looked super ugly

It really did. I did not know how to design a UI at all. The entire thing had to be reworked. I still don't know how to design a UI, but I know CSS enough to copy from other people. It's much easier once you have a baseline to work from.

The biggest issue with the design, however, was that it didn't feel like anything. I'm no visual design person, but I know that a visual design is supposed to evoke a feeling. My previous website evoked no such feeling. A personal website should give you a sense of the person behind the website. As I add more to it over time, I hope to achieve this goal

# The website was practically nonfunctional

This ties into the previous problem of aesthetics. It was so ugly that I wasn't showing it to anyone when networking. My point of pride became a point of embarrassment, especially as I met people in the tech world with much nicer looking websites than mine. And if no-one was viewing the website, what was the point of it?

Another thing, it was an absolute pain in the ass to edit. Editing raw html files is not fun. That's a lot of what motivated the current system of markdown parsing: ease of updating content. As my life changes, I need to be updating my resume, my contact information, and my blog info constantly. Discrete markdown files make this much easier.

# Conclusion

Recently, I've been thinking a lot about how visceral the Dunning-Kruger effect is in tech. Unlike some other domains, there's no one point of expertise you can reach in tech. I constantly am learning about how deep this world is, and how I really don't know anything about anything. You can't expect to master a single subdomain in tech, much less the entirety of it. It makes me feel overwhelmed sometimes. 

Projects like this counteract that feeling. While I may never be a master, I know that every day I improve at what I do. This website is the perfect example. 6 months ago, I wouldn't even know where to begin with building a web server from scratch. Today, I'm standing up this website. And the nice thing about tech is that some years down the line, I'll probably look back at this website and criticize present day me for being so naive. There's something comforting in the knowledge that you have room to grow.
