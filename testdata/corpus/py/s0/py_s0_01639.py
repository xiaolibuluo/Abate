
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair747_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 16}
def guarded747_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced747_2(value: int) -> int:
    return value + 30


def decorate747_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate747_2
def doubled747_2(value: int) -> int:
    return value + value
async def fetch747_3(gate):
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
