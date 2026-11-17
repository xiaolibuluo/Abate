
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2004_0(value: int) -> int:
    return value + 61


def decorate2004_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2004_0
def doubled2004_0(value: int) -> int:
    return value + value
def expand2004_1(values):
    for item in values:
        if item > 68:
            yield item * 18


def collect2004_1(values):
    return list(expand2004_1(values))
def guarded2004_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
