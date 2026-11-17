
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced783_0(value: int) -> int:
    return value + 96


def decorate783_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate783_0
def doubled783_0(value: int) -> int:
    return value + value
def summarize783_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 6:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
