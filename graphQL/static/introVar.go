/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package static

var (
	
	intro = `
# In root[queryOp]:
# __schema: __Schema!
# __type(name: String!): __Type

# And anywhere:
# __typename: String!

type __Schema {
	description: String
	types: [__Type!]!
	queryType: __Type!
	mutationType: __Type
	subscriptionType: __Type
	directives: [__Directive!]!
}

type __Type {
	kind: __TypeKind!
	name: String
	description: String

	# must be non-null for OBJECT and INTERFACE, otherwise null.
	fields(includeDeprecated: Boolean = false): [__Field!]

	# must be non-null for OBJECT and INTERFACE, otherwise null.
	interfaces: [__Type!]

	# must be non-null for INTERFACE and UNION, otherwise null.
	possibleTypes: [__Type!]

	# must be non-null for ENUM, otherwise null.
	enumValues(includeDeprecated: Boolean = false): [__EnumValue!]

	# must be non-null for INPUT_OBJECT, otherwise null.
	inputFields: [__InputValue!]

	# must be non-null for NON_NULL and LIST, otherwise null.
	ofType: __Type
	
	# may be non-null for custom SCALAR, otherwise null.
	specifiedByURL: String
}

enum __TypeKind {
	SCALAR
	OBJECT
	INTERFACE
	UNION
	ENUM
	INPUT_OBJECT
	LIST
	NON_NULL
}

type __Field {
	name: String!
	description: String
	args: [__InputValue!]!
	type: __Type!
	isDeprecated: Boolean!
	deprecationReason: String
}

type __InputValue {
	name: String!
	description: String
	type: __Type!
	defaultValue: String
}

type __EnumValue {
	name: String!
	description: String
	isDeprecated: Boolean!
	deprecationReason: String
}

type __Directive {
	name: String!
	description: String
	locations: [__DirectiveLocation!]!
	args: [__InputValue!]!
	isRepeatable: Boolean!
}

enum __DirectiveLocation {
	QUERY
	MUTATION
	SUBSCRIPTION
	FIELD
	FRAGMENT_DEFINITION
	FRAGMENT_SPREAD
	INLINE_FRAGMENT
	VARIABLE_DEFINITION
	SCHEMA
	SCALAR
	OBJECT
	FIELD_DEFINITION
	ARGUMENT_DEFINITION
	INTERFACE
	UNION
	ENUM
	ENUM_VALUE
	INPUT_OBJECT
	INPUT_FIELD_DEFINITION
}
`

)
