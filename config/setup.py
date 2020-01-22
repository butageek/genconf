# see details: https://setuptools.readthedocs.io/en/latest/setuptools.html
# additional details: https://docs.python.org/3/distutils/setupscript.html

from setuptools import setup

# read README.md as long description
with open('README.md', 'r') as f:
    long_description = f.read()

setup(
    # project name
    # the distributed package name starts as 'name'
    # eg. 'helloworld'
    name='',
    # the version number
    # shows in the distributed package name following 'name'
    # eg. '0.0.1'
    version='',
    description='',
    # directory to be searched for package
    # '' refers to root folder, eg. {'': 'src'} means project_root_folder/src/
    package_dir={'': 'src'},
    # modules to be searched
    # eg. 'helloworld'
    py_modules=[],
    classifiers=[
        # classifier configuration
        # see details: https://pypi.org/classifiers/
        'Programming Language :: Python :: 3.7',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
    ],
    long_description=long_description,
    long_description_content_type='text/markdown',
    # url of project website
    # eg. 'https://github.com/butageek/helloworld'
    url='',
    author='',
    author_email='',
)
