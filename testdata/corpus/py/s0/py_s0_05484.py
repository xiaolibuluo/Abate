
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2532_0(values):
    for item in values:
        if item > 38:
            yield item * 43


def collect2532_0(values):
    return list(expand2532_0(values))
def guarded2532_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
