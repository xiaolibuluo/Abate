
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1523_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 16:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1523_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1523_2(value: int) -> int:
    return value + 30


def decorate1523_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1523_2
def doubled1523_2(value: int) -> int:
    return value + value
