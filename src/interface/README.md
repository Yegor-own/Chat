## This directory is used for Interface Adapter Layer
# Interface Adapter
Interface Adapter handles the communication with the inner and the outer layer. It has only concerns about technological logic, not business logic.

__Controllers__ are a set of a specific implementation of Input Port in Use Cases.

ex.) Convert form data before saving it in database

__Presenter__ is a set of a specific implementation of Output Port in Use Cases.

ex.) Convert data from the database before passing to View