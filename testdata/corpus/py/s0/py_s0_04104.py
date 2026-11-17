
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1911_0(value: int) -> int:
    return value + 16


def decorate1911_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1911_0
def doubled1911_0(value: int) -> int:
    return value + value
def traced1911_1(value: int) -> int:
    return value + 23


def decorate1911_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1911_1
def doubled1911_1(value: int) -> int:
    return value + value
def summarize1911_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
