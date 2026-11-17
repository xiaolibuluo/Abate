
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1870_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1870_1(values):
    for item in values:
        if item > 72:
            yield item * 56


def collect1870_1(values):
    return list(expand1870_1(values))
async def fetch1870_2(gate):
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
def guarded1870_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1870_4(value: int) -> int:
    return value + 93


def decorate1870_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1870_4
def doubled1870_4(value: int) -> int:
    return value + value
