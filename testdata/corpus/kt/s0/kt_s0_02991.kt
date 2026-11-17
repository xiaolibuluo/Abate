package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24600 {
    data class Success(val payload: String) : Outcome24600()
    data object Failure : Outcome24600()
}

fun describe24600(outcome: Outcome24600): String = when (outcome) {
    is Outcome24600.Success -> outcome.payload
    Outcome24600.Failure -> "failed"
}
@JvmName("step24601")
fun step24601(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
fun <T> List<T>.secondOrNull24602(): T? = if (size > 1) this[1] else null

fun render24602(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24603<A, B>(val first: A, val second: B)

fun combine24603(left: Int, right: Int): Pair24603<Int, Int> {
    return Pair24603(left + 93, right + 60)
}
class Counter24604 {
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
