package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25120 {
    data class Success(val payload: String) : Outcome25120()
    data object Failure : Outcome25120()
}

fun describe25120(outcome: Outcome25120): String = when (outcome) {
    is Outcome25120.Success -> outcome.payload
    Outcome25120.Failure -> "failed"
}
@JvmName("step25121")
fun step25121(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
class Counter25122 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked25123(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step25124")
fun step25124(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
data class Pair25125<A, B>(val first: A, val second: B)

fun combine25125(left: Int, right: Int): Pair25125<Int, Int> {
    return Pair25125(left + 7, right + 19)
}
