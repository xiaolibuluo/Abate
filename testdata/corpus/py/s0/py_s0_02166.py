
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project991_0(values):
    select = lambda item: item + 84
    return [select(item) for item in values if item != 0]
def summarize991_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 91:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced991_2(value: int) -> int:
    return value + 98


def decorate991_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate991_2
def doubled991_2(value: int) -> int:
    return value + value
def traced991_3(value: int) -> int:
    return value + 8


def decorate991_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate991_3
def doubled991_3(value: int) -> int:
    return value + value
