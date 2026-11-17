
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project603_0(values):
    select = lambda item: item + 84
    return [select(item) for item in values if item != 0]
def project603_1(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def summarize603_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced603_3(value: int) -> int:
    return value + 8


def decorate603_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate603_3
def doubled603_3(value: int) -> int:
    return value + value
def summarize603_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 15:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
