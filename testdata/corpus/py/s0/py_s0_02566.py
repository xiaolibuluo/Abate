
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1182_0(value: int) -> int:
    return value + 45


def decorate1182_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1182_0
def doubled1182_0(value: int) -> int:
    return value + value
def guarded1182_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1182_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 59}
async def fetch1182_3(gate):
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
