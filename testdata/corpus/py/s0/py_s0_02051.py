
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand934_0(values):
    for item in values:
        if item > 22:
            yield item * 20


def collect934_0(values):
    return list(expand934_0(values))
def project934_1(values):
    select = lambda item: item + 29
    return [select(item) for item in values if item != 0]
def traced934_2(value: int) -> int:
    return value + 36


def decorate934_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate934_2
def doubled934_2(value: int) -> int:
    return value + value
