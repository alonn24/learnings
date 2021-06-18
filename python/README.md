Python learning workspace

## Installation
Install python with pyenv to manage different versions of python:
```
brew update
brew install pyenv

# add to profile file:
export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/shims:$PATH"
```
## Lint
Fix lint issues with
```
pip install --upgrade autopep8
autopep8 --in-place --recursive .
```
Run linter with
```
pip install flake8
flake8
````

## Tests
Run tests with pytest
```
pip install --upgrade pytest
python -m pytest
```