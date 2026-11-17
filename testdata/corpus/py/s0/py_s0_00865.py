
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize393_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 70:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced393_1(value: int) -> int:
    return value + 77


def decorate393_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate393_1
def doubled393_1(value: int) -> int:
    return value + value
