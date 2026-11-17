
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2248_0(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
def traced2248_1(value: int) -> int:
    return value + 39


def decorate2248_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2248_1
def doubled2248_1(value: int) -> int:
    return value + value
def guarded2248_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
