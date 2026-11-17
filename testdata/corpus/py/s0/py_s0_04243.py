
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1980_0(gate):
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
def expand1980_1(values):
    for item in values:
        if item > 47:
            yield item * 33


def collect1980_1(values):
    return list(expand1980_1(values))
def expand1980_2(values):
    for item in values:
        if item > 54:
            yield item * 44


def collect1980_2(values):
    return list(expand1980_2(values))
def guarded1980_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1980_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 68}
def expand1980_5(values):
    for item in values:
        if item > 75:
            yield item * 16


def collect1980_5(values):
    return list(expand1980_5(values))
def guarded1980_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1980_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 89}
