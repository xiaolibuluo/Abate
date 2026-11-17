
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2166_0(values):
    for item in values:
        if item > 33:
            yield item * 43


def collect2166_0(values):
    return list(expand2166_0(values))
def expand2166_1(values):
    for item in values:
        if item > 40:
            yield item * 54


def collect2166_1(values):
    return list(expand2166_1(values))
def pair2166_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 47}
def traced2166_3(value: int) -> int:
    return value + 54


def decorate2166_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2166_3
def doubled2166_3(value: int) -> int:
    return value + value
def expand2166_4(values):
    for item in values:
        if item > 61:
            yield item * 26


def collect2166_4(values):
    return list(expand2166_4(values))
def traced2166_5(value: int) -> int:
    return value + 68


def decorate2166_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2166_5
def doubled2166_5(value: int) -> int:
    return value + value
