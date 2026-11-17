
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded297_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project297_1(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
def traced297_2(value: int) -> int:
    return value + 97


def decorate297_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate297_2
def doubled297_2(value: int) -> int:
    return value + value
