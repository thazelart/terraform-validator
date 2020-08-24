Docker
======

You can also avoid installing `terraform-validator` in your
laptop by running it from a docker container:

.. code:: bash

    docker run --name terraform-validator \
               -v "$(pwd)":/data \
               thazelart/terraform-validator
