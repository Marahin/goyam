# GoYAM
#### Golang library & standalone program for YAML comparison
****
**GoYAM** started as a simple program that I’ve used for YAML files comparison (in particular: two swagger API definitions).
Scenario was that more than one person could have changed something in the API definition, and I was working on it as well - so I wanted to see what has changed, and how our definitions diverged.

I have never published a package in Golang, and never I have compared  YAML files before - this was a personal learning curve, comments and reviews are welcome.

## Features
### What does it do?

Imagine a monkey on a tree - a young one, that is still curious of the surrounding world - that is interested in two trees standing next to each other. The monkey wants to see what are the difference between the two: it wants to see how many branches does it have, if the leaves on the branches  are same, if they differ anyhow. **This is the definition of GoYAM.**

**GoYAM explores two YAML definitions and shows every difference it can find, make it missing key, a different value, or a different type for given key.**

GoYAM treats missing keys as *errors*, differing values as *warnings*, and differing types as *warnings*.

### Usage

#### Standalone

GoYAM can be used as standalone program:
```
goyam file_1.yml file_2.yml
```

Using package example:

```
➜  goyam git:(master) ✗ goyam examples/yaml\ files/api-with-examples.yaml examples/yaml\ files/petstore-expanded.yaml
(error 0) root > Key `basePath` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 1) root > Key `definitions` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 2) root > Key `schemes` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 3) root > Key `produces` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 4) root > Key `host` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 5) root > (info) > Key `termsOfService` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 6) root > (info) > Key `contact` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 7) root > (info) > Key `license` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 8) root > (info) > Key `description` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 9) root > (paths) > Key `/` can be found in `examples/yaml files/api-with-examples.yaml` but not in `examples/yaml files/petstore-expanded.yaml`.
(error 10) root > (paths) > Key `/v2` can be found in `examples/yaml files/api-with-examples.yaml` but not in `examples/yaml files/petstore-expanded.yaml`.
(error 11) root > (paths) > Key `/pets` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(error 12) root > (paths) > Key `/pets/{id}` can be found in `examples/yaml files/petstore-expanded.yaml` but not in `examples/yaml files/api-with-examples.yaml`.
(warning 0) root > (info) > Key `title` (type `string`) has a different value (Simple API overview) in `examples/yaml files/api-with-examples.yaml` than `examples/yaml files/petstore-expanded.yaml` (Swagger Petstore). [Simplwagger API etstoverview]
(warning 1) root > (info) > Key `version` (type `string`) has a different value (v2) in `examples/yaml files/api-with-examples.yaml` than `examples/yaml files/petstore-expanded.yaml` (1.0.0). [v21.0.0]
13 errors, 2 warnings found.
```

#### Package

It can also be used as a package in your project.
@TODO - API DOCUMENTATION
#personal #goyam
