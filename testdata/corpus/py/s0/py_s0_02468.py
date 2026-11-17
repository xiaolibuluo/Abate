
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1143_0(value: int) -> int:
    return value + 23


def decorate1143_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1143_0
def doubled1143_0(value: int) -> int:
    return value + value
async def fetch1143_1(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def project1143_2(values):
    select = lambda item: item + 37
    return [select(item) for item in values if item != 0]
def expand1143_3(values):
    for item in values:
        if item > 44:
            yield item * 52


def collect1143_3(values):
    return list(expand1143_3(values))
def project1143_4(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def traced1143_5(value: int) -> int:
    return value + 58


def decorate1143_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1143_5
def doubled1143_5(value: int) -> int:
    return value + value
