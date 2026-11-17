
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize788_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 64:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced788_1(value: int) -> int:
    return value + 71


def decorate788_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate788_1
def doubled788_1(value: int) -> int:
    return value + value
def project788_2(values):
    select = lambda item: item + 78
    return [select(item) for item in values if item != 0]
