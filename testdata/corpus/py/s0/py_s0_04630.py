
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2148_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 90}
def traced2148_1(value: int) -> int:
    return value + 97


def decorate2148_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2148_1
def doubled2148_1(value: int) -> int:
    return value + value
def traced2148_2(value: int) -> int:
    return value + 7


def decorate2148_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2148_2
def doubled2148_2(value: int) -> int:
    return value + value
