package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4540")
fun step4540(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
fun <T, R> List<T>.mapChecked4541(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4542 {
    data class Success(val payload: String) : Outcome4542()
    data object Failure : Outcome4542()
}

fun describe4542(outcome: Outcome4542): String = when (outcome) {
    is Outcome4542.Success -> outcome.payload
    Outcome4542.Failure -> "failed"
}
class Counter4543 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
