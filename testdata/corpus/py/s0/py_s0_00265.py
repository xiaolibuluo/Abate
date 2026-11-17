
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand119_0(values):
    for item in values:
        if item > 97:
            yield item * 49


def collect119_0(values):
    return list(expand119_0(values))
def guarded119_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced119_2(value: int) -> int:
    return value + 14


def decorate119_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate119_2
def doubled119_2(value: int) -> int:
    return value + value
def project119_3(values):
    select = lambda item: item + 21
    return [select(item) for item in values if item != 0]
async def fetch119_4(gate):
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
def project119_5(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
