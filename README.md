Programming Tool Kit - Schema Tool
==========

This projct is part of [my][1] personal toolkit (ptk). Since I'm really just getting
started with this set of tools, this is my first. The `schema` tool is a 
framework-agnostic and db-agnostic way of managing and versioning your schema. There 
are a lot of great schema-management tools out there, but they are either tied to a 
particular framework, written in a set of languages that make it difficult to install
for most people, or are tied to a particular database or set of databases.

The aim of the `schema` tool is to provide one binary (for all platforms) that can be
easily installed and used with a variety of databses (any database that has a schema).
The tool also aims to offer a set of useful features for versioning your schemas, 
tracking changes in your database, and aiding with development within a large team.

## Installation

Currently there is no proper installation script. But you can do a local build
fairly simply and will need `go` installed (>= 1.1.2).

```sh
git checkout git://github.com/JohnMurray/ptk-schema schema
cd ./schema
make

# installed to ./bin/schema
./bin/schema -h
```

For more information see the [wiki][2]



  [1]: http://johnmurray.io
  [2]: http://github.com/johnmurray/ptk-schema
