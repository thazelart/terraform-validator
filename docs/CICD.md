# CI/CD

This docker file run automatically terraform in the root directory. If you want to run it in your build pipeline, you can configure if very easily.

Let’s see how we can get it working on popular CI software.                  

## Google cloud build
```yaml
# .cloudbuild.yaml
steps:
- name: 'thazelart/terraform-validator'
```

## Travis
To be defined

## Circleci
To be defined

## Gitlab CI
To be defined

## Github Action
To be defined
