from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Image(_message.Message):
    __slots__ = ["image"]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: bytes
    def __init__(self, image: _Optional[bytes] = ...) -> None: ...

class ImageText(_message.Message):
    __slots__ = ["image", "text"]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    image: bytes
    text: str
    def __init__(self, image: _Optional[bytes] = ..., text: _Optional[str] = ...) -> None: ...
