
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1108_0(gate):
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
def traced1108_1(value: int) -> int:
    return value + 60


def decorate1108_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1108_1
def doubled1108_1(value: int) -> int:
    return value + value
def project1108_2(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
def expand1108_3(values):
    for item in values:
        if item > 74:
            yield item * 51


def collect1108_3(values):
    return list(expand1108_3(values))
def project1108_4(values):
    select = lambda item: item + 81
    return [select(item) for item in values if item != 0]
def expand1108_5(values):
    for item in values:
        if item > 88:
            yield item * 12


def collect1108_5(values):
    return list(expand1108_5(values))
def guarded1108_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1108_7(value: int) -> int:
    return value + 5


def decorate1108_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1108_7
def doubled1108_7(value: int) -> int:
    return value + value
def guarded1108_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
