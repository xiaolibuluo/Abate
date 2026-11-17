
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2587_0(values):
    select = lambda item: item + 74
    return [select(item) for item in values if item != 0]
def traced2587_1(value: int) -> int:
    return value + 81


def decorate2587_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2587_1
def doubled2587_1(value: int) -> int:
    return value + value
def summarize2587_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 88:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
