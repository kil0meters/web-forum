# web-forum


An experimental web forum to be used as hub for my non-existent community.

<<<<<<< Updated upstream:README.md
<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**
=======

* Building
>>>>>>> Stashed changes:README.org

- [web-forum](#web-forum)
- [Building](#building)
- [Project Philosophy](#project-philosophy)
    - [No Echo Chambers](#no-echo-chambers)
        - [Vote Counts are Invisible](#vote-counts-are-invisible)
    - [Maximize Community Engagement](#maximize-community-engagement)
    - [Reward High-Quality Contributions](#reward-high-quality-contributions)

<!-- markdown-toc end -->

# Building

Dependencies: sqlite3

First, install [go](https://go-la), then:
```
go get https://github.com/kil0meters/web-forum
```

# Project Philosophy

Why not just make a subreddit? The answer is simple. When Reddit does something
 like redesigning their website to appeal to younger users, [rolling out
 auto-playing video ads](https://www.digitaltrends.com/social-media/reddit-video-ads-announced/), or
 [turning a long standing Reddit meme into some cash
 crab](https://www.reddit.com/coins), their motivators are not to improve their
 community, but to line the pockets of [their
 investors](https://techcrunch.com/2019/02/11/reddit-300-million/). 
 
At small scales, you can combat these fundamental troubles with the platform by
strictly moderating, but once your community grows to a to a certain size, you
have to start making compromises such as banning memes, banning users based on
the communities they interact with. 

Therefore, in order to foster a good community at large scales, the only option
is to roll out an independent forum.

There are three main goals this project seeks to attain: 

 - Minimize potential for echo chambers
 - Maximize community engagement
 - Reward high-quality contributions to the community


## No Echo Chambers

This is something that has plagued communities, especially political ones, since
the start of time, and there's probably no perfect solution to this, but there
are a few things which should help combat this on this forum:

### Vote Counts are Invisible

It's a well-documented phenomenon in social media: a user is more likely to
"like" a post if they see that a lot of their friends have liked it as well.
Simply going away with visible score counters is the best option, in my opinion. 

It's worth exploring removing voting entirely and have a ranking system based
purely on engagement (e.g. views and comments), however.


<!-- ### {{ .OTHER_THINGS }} -->

## Maximize Community Engagement

Dedicated fans are worth their weight in gold. They are the types of people that
will go out and promote your content any chance they get, and they contribute
the most to the community, so it's understandable why one would want to maximize
the frequency of these such individuals.

Building out a custom site does something which is impossible to attain on
reddit: be its own island.

```
this forum <=[!]=> reddit.com
```

When you're on a subreddit, you are inside the Reddit community, people will
reference subreddits, and travel around the site as they please, but with a
custom forum; however, they are there to do one thing: interact with your
community. While there will be less casual passers by, the people that *do* come
are going to be much more dedicated on average.


## Reward High-Quality Contributions

When someone posts an essay they spent several hours on, they should get
"rewarded" more than someone who reposed a le epic dank meme. 

> I wrote a manifesto for a site I only have a few hundred lines of code in :)
