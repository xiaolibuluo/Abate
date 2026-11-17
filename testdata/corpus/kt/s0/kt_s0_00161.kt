package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter1320 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T, R> List<T>.mapChecked1321(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair1322<A, B>(val first: A, val second: B)

fun combine1322(left: Int, right: Int): Pair1322<Int, Int> {
    return Pair1322(left + 86, right + 40)
}
data class Pair1323<A, B>(val first: A, val second: B)

fun combine1323(left: Int, right: Int): Pair1323<Int, Int> {
    return Pair1323(left + 93, right + 51)
}
class Counter1324 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
sealed class Outcome1325 {
    data class Success(val payload: String) : Outcome1325()
    data object Failure : Outcome1325()
}

fun describe1325(outcome: Outcome1325): String = when (outcome) {
    is Outcome1325.Success -> outcome.payload
    Outcome1325.Failure -> "failed"
}
