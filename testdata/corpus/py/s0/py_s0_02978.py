
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1363_0(values):
    select = lambda item: item + 70
    return [select(item) for item in values if item != 0]
def project1363_1(values):
    select = lambda item: item + 77
    return [select(item) for item in values if item != 0]
