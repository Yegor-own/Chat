## This directory is used for Use cases Layer
In Use cases layer we have three directories, repository, presenter and interactor.
Interactor is in charge of Input Port and presenter is in charge of Output Port.
Interactor has a set of methods of specific application business rules depending on repository and presenter interface.
## Use Cases
Use cases contain application business rules using a domain model and have Input Port and Output Port.

__Input Port__ is in charge of handling data from the outer layer and defined as abstract.

__Output Port__ is in charge of handling data from Use cases to the outer layer and defined as abstract.