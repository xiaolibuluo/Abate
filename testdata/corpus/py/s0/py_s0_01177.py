
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize542_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced542_1(value: int) -> int:
    return value + 74


def decorate542_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate542_1
def doubled542_1(value: int) -> int:
    return value + value
