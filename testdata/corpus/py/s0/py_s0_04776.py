
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2211_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2211_1(value: int) -> int:
    return value + 43


def decorate2211_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2211_1
def doubled2211_1(value: int) -> int:
    return value + value
def traced2211_2(value: int) -> int:
    return value + 50


def decorate2211_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2211_2
def doubled2211_2(value: int) -> int:
    return value + value
def summarize2211_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 57:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
