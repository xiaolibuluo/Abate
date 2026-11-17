
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1993_0(gate):
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
def traced1993_1(value: int) -> int:
    return value + 22


def decorate1993_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1993_1
def doubled1993_1(value: int) -> int:
    return value + value
async def fetch1993_2(gate):
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
async def fetch1993_3(gate):
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
def guarded1993_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1993_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1993_6(values):
    for item in values:
        if item > 57:
            yield item * 57


def collect1993_6(values):
    return list(expand1993_6(values))
def traced1993_7(value: int) -> int:
    return value + 64


def decorate1993_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1993_7
def doubled1993_7(value: int) -> int:
    return value + value
def pair1993_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 71}
