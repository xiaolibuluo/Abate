package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10380 {
    data class Success(val payload: String) : Outcome10380()
    data object Failure : Outcome10380()
}

fun describe10380(outcome: Outcome10380): String = when (outcome) {
    is Outcome10380.Success -> outcome.payload
    Outcome10380.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10381(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10382")
fun step10382(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
class Counter10383 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
