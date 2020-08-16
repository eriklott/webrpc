// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package embed

var	Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00\xe3\xa9\x10Q\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00client.elm.tmplUT\x05\x00\x01\xda\xa19_<%define \"client\"%>\n<%- if .Services %>\n\n-- CLIENT\n\n<%- range .Services -%>\n<%- $serviceName := .Name -%>\n<%- range .Methods -%>\n\n<% if gt (len .Outputs) 1 %>\ntype alias <% $serviceName %><% .Name.TitleUpcase %>Response =\n    {<% range $index, $element := .Outputs %><% commaAfterFirst $index %> <%if .Optional%>Maybe <% end %><%.Name | safeVarName%> : <% fieldType .Type %> \n    <% end %>}  \n<% end %>\n\n<% $serviceName.TitleDowncase %><%.Name%> :\n    Config\n    -> <% range $index, $element := .Inputs %>\n    <%- if $index | ne 0 %>-> <% end %><%if .Optional%>Maybe <% end %><% .Type | fieldType %>\n    <% end %>\n    \n    <%- if .Inputs%>-> <% end %>\n    \n    <%- if gt (len .Outputs) 1 -%>\n    (Result Error <%$serviceName%><% .Name.TitleUpcase %>Response -> msg)\n    -> Cmd msg\n    <%- else if .Outputs -%>\n    <%- range .Outputs -%>\n    (Result Error <% fieldType .Type %> -> msg)\n    -> Cmd msg\n    <%- end -%>\n    <%- else -%>\n    (Result Error () -> msg)\n    -> Cmd msg\n    <%- end %> \n<% $serviceName.TitleDowncase %><%.Name%> (Config config)<% range .Inputs %> <%.Name | safeVarName%><% end %> toMsg =\n    <%- if or .Inputs .Outputs%>\n    let\n        <%- if .Inputs %>\n        encoder =\n            Encode.object\n                [<% range $index, $element := .Inputs %><% commaAfterFirst $index %> ( \"<%.Name%>\", <% methodArgumentEncoderType . %> <%.Name | safeVarName%> )\n                <% end %>]\n        <% end %>\n        <%- if gt (len .Outputs) 1 %>\n        decoder =\n            Decode.succeed <%$serviceName%><% .Name.TitleUpcase %>Response\n            <%- range .Outputs %>\n                |> andMap (Decode.<%if .Optional%>decodeOptionalField<%else%>field<% end %> \"<%.Name%>\" <%.Type | typeDecoder%>)\n            <%- end %> \n        <% else if eq (len .Outputs) 1 %>\n        <%- range .Outputs %>\n        decoder =\n            Decode.<%if .Optional%>decodeOptionalField<%else%>field<% end %> \"<%.Name%>\" <%.Type | typeDecoder%>            \n        <%- end -%>    \n        <%- end %>        \n    in\n    <%- end%>\n\n    request config\n        { method = \"POST\"\n        , headers = []\n        , url = config.baseUrl ++ \"/rpc/<%$serviceName%>/<%.Name%>\"\n        <%- if .Inputs %>\n        , body = Http.jsonBody encoder\n        <%- else %>\n        , body = Http.jsonBody (Encode.object [])\n        <%- end %>\n        <%- if .Outputs %>\n        , expect = expectJson toMsg decoder\n        <%- else %>\n        , expect = expectWhatever toMsg\n        <%- end %>\n        , timeout = Nothing\n        , tracker = Nothing\n        }\n\n<% end %>     \n<%- end %>   \n<%- end %>    \n<%- end %>\n\n\nPK\x07\x0809\x9f\x9e%\n\x00\x00%\n\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xe3\x18\x86P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00decoders.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^<%define \"decoders\"%>\n-- DECODERS\n<% if .Messages%>\n<%- range .Messages%>\n<%- if .Type | isEnum %>\n<%- $enumName := .Name %>\n<%.Name | messageDecoderName %> : Decoder <%.Name%>\n<%.Name | messageDecoderName %> =\n    let\n        enumDecoder s =\n            case s of\n                <%- range .Fields %>\n                \"<%.Name%>\" ->\n                    Decode.succeed <%$enumName%><%.Name%>\n                <%- end%> \n                _ ->\n                    Decode.fail (\"unknown value for type <%.Name%> : '\" ++ s)\n    in\n    Decode.string |> Decode.andThen enumDecoder     \n<% else if .Type | isStruct %>\n<%.Name | messageDecoderName %> : Decoder <%.Name%>\n<%.Name | messageDecoderName %> =\n    Decode.succeed <%.Name%>\n    <%- range .Fields %>\n        |> andMap (<%if .Optional%>decodeOptionalField<%else%>Decode.field<%end%> \"<%. | exportedJSONField%>\" <%.Type | typeDecoder%>)\n    <%- end %>\n<% end %>\n<%- end %>\n<%- end %>\n<%- end %>PK\x07\x08\xfcYR\xd6\xac\x03\x00\x00\xac\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xe3\x18\x86P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00encoders.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^<%define \"encoders\"%>\n-- ENCODERS\n<% if .Messages%>\n<%- range .Messages%>\n<%- if .Type | isEnum %>\n<%- $enumName := .Name %>\n<%.Name | messageEncoderName %> : <%.Name%> -> Encode.Value\n<%.Name | messageEncoderName %> <%.Name.TitleDowncase%> =\n    case <%.Name.TitleDowncase%> of\n    <%- range .Fields %>\n        <%$enumName%><%.Name%> ->\n            Encode.string \"<%.Name%>\"\n    <% end%> \n<%- else if .Type | isStruct %>\n<%- $messageVarName := .Name.TitleDowncase %>\n<%.Name | messageEncoderName %> : <%.Name%> -> Encode.Value\n<%.Name | messageEncoderName %> <% $messageVarName %> =\n    Encode.object\n        [<% range $index, $element := .Fields %><% commaAfterFirst $index %> ( \"<% . | exportedJSONField %>\", <% . | messageFieldEncoderType %> <% $messageVarName %>.<% . | exportedField %> )\n        <%end%>]\n<% end %>\n<%- end %>\n<%- end %>\n<%- end %>PK\x07\x08B\xb4CLU\x03\x00\x00U\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00)\xa0\x10Q\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00error.elm.tmplUT\x05\x00\x01\x8f\x909_<%define \"error\"%>\n\n-- Error\n\ntype Error\n    = BadUrl String\n    | Timeout\n    | NetworkError\n    | BadStatus Int ErrorResponse\n    | BadBody String\n\ntype alias ErrorResponse =\n    { status : Int\n    , code : String\n    , cause : Maybe String\n    , msg : String\n    , error : String\n    }\n\nerrorResponseDecoder : Decoder ErrorResponse\nerrorResponseDecoder =\n    Decode.succeed ErrorResponse\n        |> andMap (Decode.field \"status\" Decode.int)\n        |> andMap (Decode.field \"code\" Decode.string)\n        |> andMap (decodeOptionalField \"cause\" Decode.string)\n        |> andMap (Decode.field \"msg\" Decode.string)\n        |> andMap (Decode.field \"error\" Decode.string)\n\n\nexpectWhatever : (Result Error () -> msg) -> Http.Expect msg\nexpectWhatever toMsg =\n    Http.expectStringResponse toMsg <|\n        \\response ->\n            case response of\n                Http.BadUrl_ url ->\n                    Err (BadUrl url)\n\n                Http.Timeout_ ->\n                    Err Timeout\n\n                Http.NetworkError_ ->\n                    Err NetworkError\n\n                Http.BadStatus_ metadata body ->\n                    case Decode.decodeString errorResponseDecoder body of\n                        Ok errorResponse ->\n                            Err (BadStatus metadata.statusCode errorResponse)\n\n                        Err err ->\n                            Err (BadBody (Decode.errorToString err))\n\n                Http.GoodStatus_ metadata body ->\n                    Ok ()\n\n\nexpectJson : (Result Error a -> msg) -> Decoder a -> Http.Expect msg\nexpectJson toMsg decoder =\n    Http.expectStringResponse toMsg <|\n        \\response ->\n            case response of\n                Http.BadUrl_ url ->\n                    Err (BadUrl url)\n\n                Http.Timeout_ ->\n                    Err Timeout\n\n                Http.NetworkError_ ->\n                    Err NetworkError\n\n                Http.BadStatus_ metadata body ->\n                    case Decode.decodeString errorResponseDecoder body of\n                        Ok errorResponse ->\n                            Err (BadStatus metadata.statusCode errorResponse)\n\n                        Err err ->\n                            Err (BadBody (Decode.errorToString err))\n\n                Http.GoodStatus_ metadata body ->\n                    case Decode.decodeString decoder body of\n                        Ok value ->\n                            Ok value\n\n                        Err err ->\n                            Err (BadBody (Decode.errorToString err))\n\n\n<%- end %>PK\x07\x08f\x19\xaaZ\xed	\x00\x00\xed	\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x19\xa3\x10Q\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00	\x00helpers.elm.tmplUT\x05\x00\x01\x13\x969_<%define \"helpers\"%>\n-- HELPERS\n\nrequest\n    : ConfigModel\n    -> { method : String\n       , headers : List Http.Header\n       , url : String\n       , body : Http.Body\n       , expect : Http.Expect msg\n       , timeout : Maybe Float\n       , tracker : Maybe String\n       }\n    -> Cmd msg\nrequest config =\n    if config.withCredentials then\n        Http.riskyRequest\n    else\n        Http.request\n\nandMap : Decoder a -> Decoder (a -> b) -> Decoder b\nandMap =\n    Decode.map2 (|>)\n\n\ndecodeOptionalField : String -> Decoder a -> Decoder (Maybe a)\ndecodeOptionalField fieldName decoder =\n    let\n        finishDecoding json =\n            case Decode.decodeValue (Decode.field fieldName Decode.value) json of\n                Ok val ->\n                    Decode.map Just (Decode.field fieldName decoder)\n\n                Err _ ->\n                    Decode.succeed Nothing\n    in\n    Decode.oneOf\n        [ Decode.field fieldName (Decode.null Nothing)\n        , Decode.value |> Decode.andThen finishDecoding\n        ]\n\n\nencodeMaybe : (a -> Encode.Value) -> Maybe a -> Encode.Value\nencodeMaybe encoder =\n    Maybe.map encoder >> Maybe.withDefault Encode.null\n\n\ndecodeTimestamp : Decode.Decoder Time.Posix\ndecodeTimestamp =\n    Decode.string\n        |> Decode.andThen\n            (\\timestampStr ->\n                case Iso8601.toTime timestampStr of\n                    Ok posix ->\n                        Decode.succeed posix\n\n                    Err _ ->\n                        Decode.fail (\"failed to decode iso8601 timestamp : '\" ++ timestampStr)\n            )\n\n\nencodeTimestamp : Time.Posix -> Encode.Value\nencodeTimestamp posix =\n    Encode.string (Iso8601.fromTime posix)\n\n\n<% end %> PK\x07\x083e@E\x96\x06\x00\x00\x96\x06\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xe3\x18\x86P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x12\x00	\x00proto.gen.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^<%- define \"proto\" -%>\nmodule <%.TargetOpts.PkgName%> exposing (Config, configure, defaultConfiguration, Error(..), ErrorResponse, <%.WebRPCSchema | exposingDef%>)\n\n{- <%.Name%> <%.SchemaVersion%>\n  --\n  This file has been generated by https://github.com/webrpc/webrpc using gen/elm\n  Do not edit by hand. Update your webrpc schema and re-generate.\n-}\n\nimport Http\nimport Iso8601\nimport Json.Decode as Decode exposing (Decoder)\nimport Json.Encode as Encode\nimport Time\n\n\n\n-- CONFIG\n\n\ntype Config\n  = Config ConfigModel\n\n\ntype alias ConfigModel =\n  { baseUrl : String\n  , withCredentials : Bool\n  }\n\n\nconfigure :\n  { baseUrl : String\n  , withCredentials : Bool\n  }\n  -> Config\nconfigure impl =\n  Config impl\n\n\ndefaultConfiguration : Config\ndefaultConfiguration =\n  Config\n    { baseUrl = \"\"\n    , withCredentials = False\n    }\n\n<%template \"error\" .%>\n<%template \"types\" .%>\n<%template \"decoders\" .%>\n<%template \"encoders\" .%>\n<%template \"client\" .%>\n<%template \"helpers\" .%>\n<%- end%>PK\x07\x08\xe8\xd3\xc5\x82\xd7\x03\x00\x00\xd7\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xe3\x18\x86P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00types.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^<%define \"types\"%>\n-- TYPES\n<%- if .Messages%>\n<%- range .Messages%>\n<% if .Type | isEnum %>\n<%- $enumName := .Name %>\ntype <%$enumName%>\n<%- range $index, $element := .Fields %>\n    <% if $index | eq 0 %>=<%else%>|<%end%> <%$enumName%><%.Name%>\n<%- end%>\n<%- else if .Type | isStruct %>\ntype alias <%.Name%> =\n    {<% range $index, $element := .Fields %><% commaAfterFirst $index %> <% . | exportedField %> : <% if .Optional %>Maybe <% end %><% . | fieldTypeDef %>\n    <% end -%>}\n<%- end %>\n<%- end %>\n<%- end %>\n<%- end %>PK\x07\x08\x14W\x9c\xea\x0d\x02\x00\x00\x0d\x02\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xe3\xa9\x10Q09\x9f\x9e%\n\x00\x00%\n\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00client.elm.tmplUT\x05\x00\x01\xda\xa19_PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xe3\x18\x86P\xfcYR\xd6\xac\x03\x00\x00\xac\x03\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81k\n\x00\x00decoders.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xe3\x18\x86PB\xb4CLU\x03\x00\x00U\x03\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81_\x0e\x00\x00encoders.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00)\xa0\x10Qf\x19\xaaZ\xed	\x00\x00\xed	\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xfc\x11\x00\x00error.elm.tmplUT\x05\x00\x01\x8f\x909_PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x19\xa3\x10Q3e@E\x96\x06\x00\x00\x96\x06\x00\x00\x10\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81.\x1c\x00\x00helpers.elm.tmplUT\x05\x00\x01\x13\x969_PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xe3\x18\x86P\xe8\xd3\xc5\x82\xd7\x03\x00\x00\xd7\x03\x00\x00\x12\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x0b#\x00\x00proto.gen.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xe3\x18\x86P\x14W\x9c\xea\x0d\x02\x00\x00\x0d\x02\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81+'\x00\x00types.elm.tmplUT\x05\x00\x01\xdb\x9c\x8a^PK\x05\x06\x00\x00\x00\x00\x07\x00\x07\x00\xf0\x01\x00\x00})\x00\x00\x00\x00"
