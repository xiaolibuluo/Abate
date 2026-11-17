
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2288_0(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
def traced2288_1(value: int) -> int:
    return value + 74


def decorate2288_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2288_1
def doubled2288_1(value: int) -> int:
    return value + value
