
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1113_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1113_1(values):
    select = lambda item: item + 28
    return [select(item) for item in values if item != 0]
def project1113_2(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def project1113_3(values):
    select = lambda item: item + 42
    return [select(item) for item in values if item != 0]
def expand1113_4(values):
    for item in values:
        if item > 49:
            yield item * 36


def collect1113_4(values):
    return list(expand1113_4(values))
def expand1113_5(values):
    for item in values:
        if item > 56:
            yield item * 47


def collect1113_5(values):
    return list(expand1113_5(values))
def expand1113_6(values):
    for item in values:
        if item > 63:
            yield item * 58


def collect1113_6(values):
    return list(expand1113_6(values))
def project1113_7(values):
    select = lambda item: item + 70
    return [select(item) for item in values if item != 0]
def guarded1113_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1113_9(values):
    select = lambda item: item + 84
    return [select(item) for item in values if item != 0]
def traced1113_10(value: int) -> int:
    return value + 91


def decorate1113_10(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1113_10
def doubled1113_10(value: int) -> int:
    return value + value
def project1113_11(values):
    select = lambda item: item + 98
    return [select(item) for item in values if item != 0]
