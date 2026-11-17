
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2649_0(value: int) -> int:
    return value + 7


def decorate2649_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2649_0
def doubled2649_0(value: int) -> int:
    return value + value
def traced2649_1(value: int) -> int:
    return value + 14


def decorate2649_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2649_1
def doubled2649_1(value: int) -> int:
    return value + value
def summarize2649_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2649_3(value: int) -> int:
    return value + 28


def decorate2649_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2649_3
def doubled2649_3(value: int) -> int:
    return value + value
