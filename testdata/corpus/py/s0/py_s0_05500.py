
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2541_0(value: int) -> int:
    return value + 58


def decorate2541_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2541_0
def doubled2541_0(value: int) -> int:
    return value + value
class Slot2541_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2541_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
