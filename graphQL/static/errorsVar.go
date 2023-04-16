/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package static

var (
	
	errEn = `STRINGS

ArgNotRequiredByInterface	The argument ^0 is not required by the interface
AbsTypeResolverAlrdyDefined	The abstract type resolver is already defined
AbsTypeNotResolved	The abstract type ^0 has not been resolved
AbsTypeResolverNotDefined	The abstract type resolver is not defined
CoercerNotDefined	The coercer for ^0 is not defined
CircularInputs	^0: Circular References of Input Objects
CycleThrough	Cycle through ^0
DefValIncoercibleToTypeIn	In ^0, the default value is incoercible to its type
DirectiveInWrongLocation	The directive ^0 is in a wrong location
DupArgName	^0 is a duplicated argument name
DupDirName	^0 is a duplicated directive name
DupEnumValue	^0 is a duplicated enum value
DupFieldName	^0 is a duplicated field name
DupInterface	^0 is a duplicated interface name
DupName	^0 is a duplicated name
DupTypeName	^0 is a duplicated type name
DupVarName	^0 is a duplicated variable name
EnumExtWoDef	^0: enum extension without definition
FieldNotAvailableFor	^0: this field is not available for the type ^1
FieldsCantMerge	Fields ^0 and ^1 can't merge
FragmentNotUsed	Fragment ^0 is not used
FragSpreadImpossible	Type ^0 is impossible here
IncorrectArgsNb	Incorrect number of arguments in ^0
IncorrectArgumentIn	^0: incorrect argument of ^1
IncorrectFragmentType	^0: incorrect fragment type
IncorrectReturnTypeOf	^0: incorrect return type of ^1
IncorrectValueType	^0: incorrect value type
IncorrectVarValue	^0: incorrect variable value
InitialValueNotFixed	The initial value was never fixed
InputObjExtWoDef	^0: input object extension without definition
IntOutOfRange	Int out of range
IntroName	^0 begins with __
IsNotFieldFor	^0 is not a field of ^1
IsNotInputFieldFor	^0 is not an input argument of ^1
MisusedName	^0: misused predeclared name
NotAnInputObject	^0 is not an input object
NoArgForNonNullType	No argument for non-null ^0 argument in ^1
NoArgumentIn	No argument in ^0
NonBoolValueForBoolType	Non Boolean value for Boolean type
NonEnumValueForEnumType	Non Enum value for ^0 Enum type
NonFloatOrIntValueForFloatType	Non Float or Int value for Float type
NonIntValueForIntType	Non Int value for Int type
NonListValueForListType	Non list value for ^0 list type
NotObjectTypeInUnion	In union ^1, ^0 is not an object type
NonObjValForInputObjType	Non Object value for Input Object type
NonObjValForObjType	Non Object value for Object type
NonScalarValueForScalarType	Non scalar value for ^0 scalar type
NonStringValueForIDType	Non String value for ID type
NonStringValueForStringType	Non String value for String type
NoSubscriptionRoot	^0: the subscription root is not defined
NoSubselOfObjInterOrUnion	Missing subselection of ^0 (OBJECT, INTERFACE or UNION)
NotABoolean	Not a Boolean
NotAFloat	Not a Float
NotAnID	Not an ID
NotAnInteger	Not an Integer
NotAString	Not a String
NotDefinedInScope	^0 is not defined in scope
NotDefinedOp	The operation to be executed is not defined
NotImplementedIn	The field ^0 is not implemented in the interface ^1
NotInputField	^0 is not an input field
NotInputType	^0 has not an input type
NotLoneAnonOp	Anonymous operation is not alone
NotOutputType	^0 has not an output type
NotSingleSubscriptRoot	The subscription root ^0 is not single
NotSubtype	The type of ^0 is not a subtype of ^1
NotUniqueOperation	The operation ^0 is not unique
NoValueForVar	^0: variable with no value
NullArgForNonNullType	Null argument ^0 for non-null type in ^1
NullValueWithNonNullType	Null value with non-null type
NullVarWithNonNullType	^0: Null variable with non-null type
ObjectExtWoDef	^0: object extension without definition
QueryRootNotDefined	The query root is not defined
ResolverAlrdyDefined	The resolver for ^1 in ^0 is already defined
ResolverNotDefined	The resolver for ^1 in ^0 is not defined
RootAlrdyDefined	The root ^0 is already defined
RootNotDefinedFor	The root for ^0 operation is not defined
ScalarCoercerAlrdyDefined	The scalar coercer for ^0 is already defined
ScalarExtWoDef	^0: scalar extension without definition
SchemaExtWoDefinition	Schema extension without definition
SelfInterface	The interface ^0 can't implement itself
StreamResolverAlrdyDefined	The stream resolver for ^0 is already defined
StreamResolverNotDefined	The stream resolver for ^0 is not defined
SubselectionOfScalarOrEnum	Incorrect subselection of ^0 (SCALAR or ENUM)
UnableToCoerce	Unable to coerce ^0 to its type
UnionExtWoDef	^0: union extension without definition
UnknownArgument	^0: Unknown argument
UnknownDirective	^0: Unknown directive
UnknownFragment	^0: Unknown fragment
UnknownObject	^0: Unknown object type
UnknownOperation	^0: Unknown operation
UnknownScalarType	^0: Unknown scalar type
UnknownType	^0: Unknown type
UnknownVariable	$^0: Unknown variable in ^1
VariableNotUsed	$^0: this variable is not used
VarUsageNotAllowed	Incorrect usage of $^0
WrongArgumentType	^0 has a wrong argument type
WrongEnumValueForEnumType	Wrong Enum value ^0 for ^1 Enum type
WrongRootType	The root ^0 is not an object type
`

)
