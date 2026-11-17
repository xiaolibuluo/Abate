
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand739_0(values):
    for item in values:
        if item > 9:
            yield item * 58


def collect739_0(values):
    return list(expand739_0(values))
def guarded739_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced739_2(value: int) -> int:
    return value + 23


def decorate739_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate739_2
def doubled739_2(value: int) -> int:
    return value + value
