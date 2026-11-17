
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2595_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2595_1(values):
    select = lambda item: item + 88
    return [select(item) for item in values if item != 0]
def guarded2595_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2595_3(values):
    for item in values:
        if item > 5:
            yield item * 29


def collect2595_3(values):
    return list(expand2595_3(values))
def expand2595_4(values):
    for item in values:
        if item > 12:
            yield item * 40


def collect2595_4(values):
    return list(expand2595_4(values))
