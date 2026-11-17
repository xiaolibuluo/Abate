
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded105_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand105_1(values):
    for item in values:
        if item > 19:
            yield item * 23


def collect105_1(values):
    return list(expand105_1(values))
def project105_2(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
