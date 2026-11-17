
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand323_0(values):
    for item in values:
        if item > 33:
            yield item * 13


def collect323_0(values):
    return list(expand323_0(values))
def traced323_1(value: int) -> int:
    return value + 40


def decorate323_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate323_1
def doubled323_1(value: int) -> int:
    return value + value
def expand323_2(values):
    for item in values:
        if item > 47:
            yield item * 35


def collect323_2(values):
    return list(expand323_2(values))
def expand323_3(values):
    for item in values:
        if item > 54:
            yield item * 46


def collect323_3(values):
    return list(expand323_3(values))
def traced323_4(value: int) -> int:
    return value + 61


def decorate323_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate323_4
def doubled323_4(value: int) -> int:
    return value + value
def guarded323_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded323_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
