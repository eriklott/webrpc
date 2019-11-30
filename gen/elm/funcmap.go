package elm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

var fieldTypeMap = map[schema.DataType]string{
	schema.T_Uint:      "Int",
	schema.T_Uint8:     "Int",
	schema.T_Uint16:    "Int",
	schema.T_Uint32:    "Int",
	schema.T_Uint64:    "Int",
	schema.T_Int:       "Int",
	schema.T_Int8:      "Int",
	schema.T_Int16:     "Int",
	schema.T_Int32:     "Int",
	schema.T_Int64:     "Int",
	schema.T_Float32:   "Float",
	schema.T_Float64:   "Float",
	schema.T_String:    "String",
	schema.T_Timestamp: "String",
	schema.T_Bool:      "Bool",
}

var decoderTypeMap = map[schema.DataType]string{
	schema.T_Uint:      "int",
	schema.T_Uint8:     "int",
	schema.T_Uint16:    "int",
	schema.T_Uint32:    "int",
	schema.T_Uint64:    "int",
	schema.T_Int:       "int",
	schema.T_Int8:      "int",
	schema.T_Int16:     "int",
	schema.T_Int32:     "int",
	schema.T_Int64:     "int",
	schema.T_Float32:   "float",
	schema.T_Float64:   "float",
	schema.T_String:    "string",
	schema.T_Timestamp: "string",
	schema.T_Bool:      "bool",
}

var encoderTypeMap = map[schema.DataType]string{
	schema.T_Uint:      "int",
	schema.T_Uint8:     "int",
	schema.T_Uint16:    "int",
	schema.T_Uint32:    "int",
	schema.T_Uint64:    "int",
	schema.T_Int:       "int",
	schema.T_Int8:      "int",
	schema.T_Int16:     "int",
	schema.T_Int32:     "int",
	schema.T_Int64:     "int",
	schema.T_Float32:   "float",
	schema.T_Float64:   "float",
	schema.T_String:    "string",
	schema.T_Timestamp: "string",
	schema.T_Bool:      "bool",
}

func messageDecoderName(v schema.VarName) string {
	name := string(v)
	if name == "" {
		return ""
	}
	return strings.ToLower(name[0:1]) + name[1:] + "Decoder"
}

func fieldType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_List:
		z, err := fieldType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "(List " + z + ")", nil
	case schema.T_Struct:
		return in.Struct.Name, nil
	default:
		if fieldTypeMap[in.Type] != "" {
			return fieldTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent type: %#v", in)
}

func fieldTypeDef(in *schema.MessageField) (string, error) {
	return fieldType(in.Type)
}

func isStruct(t schema.MessageType) bool {
	return t == "struct"
}

func isEnum(t schema.MessageType) bool {
	return t == "enum"
}

func capitalizeName(v interface{}) (string, error) {
	capitalizeFn := func(s string) string {
		if s == "" {
			return ""
		}
		return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
	}
	switch t := v.(type) {
	case schema.VarName:
		return capitalizeFn(string(t)), nil
	case string:
		return capitalizeFn(t), nil
	default:
		return "", errors.New("capitalizeFieldName, unknown arg type")
	}
}

func exportedJSONField(in schema.MessageField) (string, error) {
	for i := range in.Meta {
		for k := range in.Meta[i] {
			if k == "json" {
				s := strings.Split(fmt.Sprintf("%v", in.Meta[i][k]), ",")
				if len(s) > 0 {
					return s[0], nil
				}
			}
		}
	}
	return string(in.Name), nil
}

func messageEncoderName(v schema.VarName) string {
	if string(v) == "" {
		return ""
	}
	return v.TitleDowncase() + "Encoder"
}

func methodArgumentEncoderType(in schema.MethodArgument) (string, error) {
	encoderStr, err := encoderTypeRec(in.Type)
	if err != nil {
		return encoderStr, err
	}
	if in.Optional {
		encoderStr = "encodeMaybe (" + encoderStr + ")"
	}
	return encoderStr, nil
}

func messageFieldEncoderType(in schema.MessageField) (string, error) {
	encoderStr, err := encoderTypeRec(in.Type)
	if err != nil {
		return encoderStr, err
	}
	if in.Optional {
		encoderStr = "encodeMaybe (" + encoderStr + ")"
	}
	return encoderStr, nil
}

func encoderTypeRec(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_List:
		subEncoder, err := encoderTypeRec(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "Encode.list (" + subEncoder + ")", nil
	case schema.T_Struct:
		return messageEncoderName(in.Struct.Message.Name), nil
	default:
		if decoderTypeMap[in.Type] != "" {
			return "Encode." + decoderTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent encoder: %#v", in)
}

func decoderType(in *schema.VarType) (string, error) {
	switch in.Type {
	case schema.T_List:
		z, err := decoderType(in.List.Elem)
		if err != nil {
			return "", err
		}
		return "(Decode.list " + z + ")", nil
	case schema.T_Struct:
		return messageDecoderName(in.Struct.Message.Name), nil
	default:
		if decoderTypeMap[in.Type] != "" {
			return "Decode." + decoderTypeMap[in.Type], nil
		}
	}
	return "", fmt.Errorf("could not represent decoder: %#v", in)
}

func exposingDef(in *schema.WebRPCSchema) string {
	exposedNames := []string{}
	for _, message := range in.Messages {
		exposedNames = append(exposedNames, message.Name.TitleUpcase())
	}
	for _, service := range in.Services {
		for _, method := range service.Methods {
			if len(method.Outputs) > 0 {
				exposedNames = append(exposedNames, service.Name.TitleUpcase()+method.Name.TitleUpcase()+"Data")
			}
			exposedNames = append(exposedNames, service.Name.TitleDowncase()+method.Name.TitleUpcase())
		}
	}
	return strings.Join(exposedNames, ", ")
}

func safeVarName(in schema.VarName) string {
	s := string(in)
	if s == "type" {
		return "type_"
	}
	return s
}

func exportedField(in *schema.MessageField) string {
	s := safeVarName(in.Name)

	nameTag := "elm.field.name"
	for k := range in.Meta {
		for k, v := range in.Meta[k] {
			if k == nameTag {
				s = fmt.Sprintf("%v", v)
			}
		}
	}

	return s
}

func commaAfterFirst(index int) string {
	if index == 0 {
		return ""
	}
	return ","
}

var templateFuncMap = map[string]interface{}{
	"fieldType":                 fieldType,
	"fieldTypeDef":              fieldTypeDef,
	"isStruct":                  isStruct,
	"isEnum":                    isEnum,
	"capitalizeName":            capitalizeName,
	"exportedJSONField":         exportedJSONField,
	"decoderType":               decoderType,
	"methodArgumentEncoderType": methodArgumentEncoderType,
	"messageFieldEncoderType":   messageFieldEncoderType,
	"messageDecoderName":        messageDecoderName,
	"messageEncoderName":        messageEncoderName,
	"exposingDef":               exposingDef,
	"exportedField":             exportedField,
	"safeVarName":               safeVarName,
	"commaAfterFirst":           commaAfterFirst,
}
