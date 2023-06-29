from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class DatasetPath(_message.Message):
    __slots__ = ["file_path", "json_path"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    JSON_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    json_path: str
    def __init__(self, file_path: _Optional[str] = ..., json_path: _Optional[str] = ...) -> None: ...

class ResultJson(_message.Message):
    __slots__ = ["data"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: str
    def __init__(self, data: _Optional[str] = ...) -> None: ...
