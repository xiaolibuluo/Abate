
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2596_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2596_1(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
def traced2596_2(value: int) -> int:
    return value + 11


def decorate2596_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2596_2
def doubled2596_2(value: int) -> int:
    return value + value
def expand2596_3(values):
    for item in values:
        if item > 18:
            yield item * 36


def collect2596_3(values):
    return list(expand2596_3(values))
def expand2596_4(values):
    for item in values:
        if item > 25:
            yield item * 47


def collect2596_4(values):
    return list(expand2596_4(values))
