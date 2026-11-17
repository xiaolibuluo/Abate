
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1084_0(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
def project1084_1(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]
def traced1084_2(value: int) -> int:
    return value + 46


def decorate1084_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1084_2
def doubled1084_2(value: int) -> int:
    return value + value
def summarize1084_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 53:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
