
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1246_0(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
def summarize1246_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 11:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1246_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1246_3(value: int) -> int:
    return value + 25


def decorate1246_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1246_3
def doubled1246_3(value: int) -> int:
    return value + value
def summarize1246_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1246_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 39:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
