Getting started with terraform-validator
========================================

.. toctree::
   :maxdepth: 2
   :hidden:
   :titlesonly:
   :caption: Getting started

   /getting-started/install
   /getting-started/first-steps
   /getting-started/docker


.. toctree::
   :maxdepth: 2
   :hidden:
   :caption: Configuration

   /configuration/introduction
   /configuration/layers
   /configuration/current-layer
   /configuration/recursivity

.. toctree::
   :maxdepth: 2
   :hidden:
   :titlesonly:
   :caption: Contributing

   /contributing/how-to
   /contributing/code-of-conduct
   /contributing/thanks

| |Mentioned in Awesome Go| |GoDoc| |License|
| |Build Status| |CodeCov| |Go Report Card|
| |Docker Cloud Build Status| |Docker Pulls|

This tool will help you ensure that a terraform folder answer to your
norms and conventions rules. This can be really useful in several cases :

- You're a team that want to have a clean and maintainable code.
- You're a lonely developer that develop a lot of modules and you want to have a certain consistency between them.

.. |check| raw:: html

    <input checked=""  disabled="" type="checkbox">

.. |uncheck| raw:: html

    <input disabled="" type="checkbox">

| **Features:**
|  |check| make sure that the block names match a certain pattern.
|  |check| make sure that the code is properly dispatched. To do this you can decide what type of block can contain each file (for example output blocks must be in ``outputs.tf``).
|  |check| ensure that mandatory ``.tf`` files are present.
|  |check| ensure that the terraform version has been defined.
|  |check| ensure that the providers' version has been defined.
|  |check| make sure that the variables and/or outputs blocks have the description argument filled in.
|  |check| layered terraform folders (test recursively).


.. warning::
   Terraform 0.12+ is supported only by the versions 2.0.0 and higher.

.. |Mentioned in Awesome Go| image:: https://awesome.re/mentioned-badge.svg
   :target: https://github.com/avelino/awesome-go#validation
.. |GoDoc| image:: https://godoc.org/github.com/thazelart/terraform-validator?status.svg
   :target: https://godoc.org/github.com/thazelart/terraform-validator
.. |License| image:: https://img.shields.io/badge/License-Apache%202.0-blue.svg
   :target: https://github.com/gojp/goreportcard/blob/master/LICENSE
.. |Build Status| image:: https://travis-ci.com/thazelart/terraform-validator.svg?branch=master
   :target: https://travis-ci.com/thazelart/terraform-validator
.. |CodeCov| image:: https://codecov.io/gh/thazelart/terraform-validator/branch/master/graph/badge.svg
   :target: https://codecov.io/gh/thazelart/terraform-validator
.. |Go Report Card| image:: https://goreportcard.com/badge/github.com/thazelart/terraform-validator
   :target: https://goreportcard.com/report/github.com/thazelart/terraform-validator
.. |Docker Cloud Build Status| image:: https://img.shields.io/docker/cloud/build/thazelart/terraform-validator.svg
   :target: https://hub.docker.com/r/thazelart/terraform-validator
.. |Docker Pulls| image:: https://img.shields.io/docker/pulls/thazelart/terraform-validator
   :target: https://hub.docker.com/r/thazelart/terraform-validator
