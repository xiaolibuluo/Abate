
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1330_0(value: int) -> int:
    return value + 29


def decorate1330_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1330_0
def doubled1330_0(value: int) -> int:
    return value + value
def expand1330_1(values):
    for item in values:
        if item > 36:
            yield item * 58


def collect1330_1(values):
    return list(expand1330_1(values))
async def fetch1330_2(gate):
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
def guarded1330_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
