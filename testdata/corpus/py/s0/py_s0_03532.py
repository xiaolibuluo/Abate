
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1629_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1629_1(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
def project1629_2(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
def pair1629_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 57}
async def fetch1629_4(gate):
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
def guarded1629_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1629_6(value: int) -> int:
    return value + 78


def decorate1629_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1629_6
def doubled1629_6(value: int) -> int:
    return value + value
