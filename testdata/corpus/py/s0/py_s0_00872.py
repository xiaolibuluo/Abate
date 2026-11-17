
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced397_0(value: int) -> int:
    return value + 25


def decorate397_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate397_0
def doubled397_0(value: int) -> int:
    return value + value
def guarded397_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch397_2(gate):
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
