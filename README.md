# CeX Go Client

[CeX](https://uk.webuy.com/) is a UK trade shop for movies, games and electronics.

The CeX website is built with [Vue.js](https://github.com/vuejs/vue) and [Nuxt](https://github.com/nuxt/nuxt.js). It
uses a RESTful API for sourcing products and their categories. It's actually a very pleasant API that's pretty well
designed and easy to use, props to the CeX folks behind it!

This library implements some portion of that API, at least the interesting bits such as listing products. Using this
library, you can pragmatically list all the products available in the store (at the time of writing, this was 369,625!)

The library has some extremely basic unit tests that simply make an API call and print the result. I didn't want to
bother writing a full test suite against a dataset that will change! Run the tests to see an example of the output. The
code is also documented so you should be up and running pretty quickly.

## To CeX

If a CeX employee runs across this: This library is just the result of some mild curiosity, it's in no way intended to
be malicious despite being unofficial and reverse engineered. The data itself is actually quite interesting and being
able to access it pragmatically in this way could result in some interesting things!
