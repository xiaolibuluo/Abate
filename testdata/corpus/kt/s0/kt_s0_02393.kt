package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19660 {
    data class Success(val payload: String) : Outcome19660()
    data object Failure : Outcome19660()
}

fun describe19660(outcome: Outcome19660): String = when (outcome) {
    is Outcome19660.Success -> outcome.payload
    Outcome19660.Failure -> "failed"
}
class Counter19661 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

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
@JvmName("step19662")
fun step19662(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
fun <T, R> List<T>.mapChecked19663(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
