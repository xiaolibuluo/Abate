
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2634_0(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def traced2634_1(value: int) -> int:
    return value + 13


def decorate2634_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2634_1
def doubled2634_1(value: int) -> int:
    return value + value
def guarded2634_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
