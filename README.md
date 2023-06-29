# opg-sirius-header
<p> 
This is a small template repo imported into other Go repos. <br>
There is custom CSS within the sirius-header.scss that will get pulled through into these other repos.
Use this to update the header template, when the finished changes are merged into main they will pull through into the other repos. <br>
</p>

## Development:
If you want to test these changes in the other repos while you are developing you can amend the other repos package.json to see changes made. <br>
(E.g. if developing against workflow amend the package.json dependency in workflow from: <br> `"opg-sirius-header": "ministryofjustice/opg-sirius-header"` <br>
to `"opg-sirius-header": "ministryofjustice/opg-sirius-header#commitId"`. <br> 
Then rebuild the css/ js styling in workflow). <br> It should pull through the latest commit from sirius-header.

### Useful commands:

Download dependencies: `yarn install` <br>
Build Sass/ Css: `yarn compile-sass` <br>
Build page: `yarn serve` <br> 
Run cypress: `yarn serve` then `yarn cypress` <br>
(NB: Cypress expects the app to be running on 8080 which is the default port, 
if this is taken and the app hosts on another app Cypress will fail)