# Goal
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Eventually the goal is to inform and provide an equal sound of the positive side. Not to bash, hate or attack "killedbygoogle". Everyone should have the right to express their thoughts, opinion, whatnot.

# Contribute entries

## How this website works

- We use data.json as data-generator
- We iterate over data.json with a Go "script" (`build-html.go`)
- The script will place an `index-middle.tmpl` in the `blocks` directory. 
- The script will combine the `index-top.tmpl` + `index-middle.tmpl` + `index-bottom.tmpl` into a `index.html` and places it into the `public` folder.
- Everything in the `public` folder gets uploaded to Google Cloud Storage
- Contents get served via a loadbalancer & CDN

### Why that script!?

- Because I wanted to keep it simple.
- I don't want to use JS
- A static html generator such as Hugo is then an overkill for 1 single HTML page with some dynamic content.
- Just working in 1 single HTML file is annoying.

Keep it simple. 


## Adding entries yourself

- Create a feature branch and make a merge request.

In the `data.json` you can add entries. 

The base for an entry is the following:

```
    {
        "title": "",
        "description": "",
        "value": "",
        "img": "img/",
        "notables": [
            "",
            ""
        ]
    }
```

- title: Used for the title
- description: A summery of what the product/service is
- value: What did it mean to you/others. What is the value
- img: A reference to the image of the product/service
- notables: This is an array of "facts". I.e.: Launched in X, has X users, etc.

The image can be placed into `public/img` - The entry in the data.json should be `img/your-image.ext`.

## Create an issue

If you are unfamiliar with Git or you are getting stuck somewhere: Create an issue with either:

- An entry that you would like to get added. Please include all the information for that entry
- Provide feedback on how to improve this repository to contributions.


# Contribute on all other parts

If you want to be certain that a possible PR gets merged, perhaps make an issue first explaining on what you want to achieve. This way we can have a discussion on the viability of this feature.

Personally I have a few "rules" I would like to maintain. 

- No ads
- No Javascript
- Keep it simple (more vanilla than..exploding libraries.)

Other than that I'm always open to completely change the entire setup aslong as:

- It should be easy usable in some way. Currently everyone could alter the `data.json`. I see this as something with little barriers. 
- It should not give that much overhead (seriously, it's just a single HTML page!)

# General things

- Let's keep it chill and releaxed. No heated arguments. After all we just want a positive vibe.
- Respect eachother. 
- I assume people can be mature, otherwise I'll have to incorporate some code of conduct..

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/dickynovanto1103"><img src="https://avatars1.githubusercontent.com/u/22437392?v=4" width="100px;" alt=""/><br /><sub><b>Dicky Novanto</b></sub></a><br /><a href="https://github.com/wiardvanrij/Alive-By-Google/commits?author=dickynovanto1103" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/fabianmet"><img src="https://avatars2.githubusercontent.com/u/4489146?v=4" width="100px;" alt=""/><br /><sub><b>fabianmet</b></sub></a><br /><a href="#content-fabianmet" title="Content">🖋</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
