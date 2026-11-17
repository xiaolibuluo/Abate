
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2581_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2581_1(values):
    for item in values:
        if item > 100:
            yield item * 31


def collect2581_1(values):
    return list(expand2581_1(values))
