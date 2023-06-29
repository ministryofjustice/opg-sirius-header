# opg-sirius-header
This is a small template repo imported into other Go repos. <br>
There is custom CSS within the sirius-header.scss that will get pulled through into these other repos.
Use this to update the header template, when the finished changes are merged into main they will pull through into the other repos. <br>

## Development:
If you want to test these changes in the other repos while you are developing you can amend the other repos package.json to see changes made. <br>
(E.g. if developing against workflow amend the package.json dependency in workflow from: <br> `"opg-sirius-header": "ministryofjustice/opg-sirius-header"` <br>
to `"opg-sirius-header": "ministryofjustice/opg-sirius-header#commitId"`. <br> 
Then rebuild the css/ js styling in workflow). <br> It should pull through the latest commit from sirius-header.

### To build locally:
Download dependencies: `yarn install` <br>
Build Sass/ Css: `yarn compile-sass` <br>
Build page: `yarn serve` <br> 
This will then host with http-server, it's usually on 8080 but the console will tell you which port it's been hosted on.

#### Testing:
Run cypress: `yarn serve` then `yarn cypress` in another console window <br>
(NB: Cypress expects the app to be running on 8080 which is the default port, 
if this is taken and the app hosts on another port Cypress will fail)

#### To import into a new app that isn't currently using it:
Add dependency to package.json in repo you want to import it into `"opg-sirius-header": "ministryofjustice/opg-sirius-header"` <br>
Import the SCSS from sirius-header into the repo's main.scss file `@import "node_modules/opg-sirius-header/sass/sirius-header"` <br>
Import module into repo's main.js file `import "opg-sirius-header/sirius-header.js"` <br>
Run `Yarn install` and build the CSS locally
