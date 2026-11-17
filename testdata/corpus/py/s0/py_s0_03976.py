
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1843_0(values):
    for item in values:
        if item > 5:
            yield item * 39


def collect1843_0(values):
    return list(expand1843_0(values))
def expand1843_1(values):
    for item in values:
        if item > 12:
            yield item * 50


def collect1843_1(values):
    return list(expand1843_1(values))
def traced1843_2(value: int) -> int:
    return value + 19


def decorate1843_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1843_2
def doubled1843_2(value: int) -> int:
    return value + value
def project1843_3(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
async def fetch1843_4(gate):
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
def project1843_5(values):
    select = lambda item: item + 40
    return [select(item) for item in values if item != 0]
def project1843_6(values):
    select = lambda item: item + 47
    return [select(item) for item in values if item != 0]
async def fetch1843_7(gate):
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
async def fetch1843_8(gate):
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
def guarded1843_9(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
