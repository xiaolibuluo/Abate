package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10460 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
sealed class Outcome10461 {
    data class Success(val payload: String) : Outcome10461()
    data object Failure : Outcome10461()
}

fun describe10461(outcome: Outcome10461): String = when (outcome) {
    is Outcome10461.Success -> outcome.payload
    Outcome10461.Failure -> "failed"
}
sealed class Outcome10462 {
    data class Success(val payload: String) : Outcome10462()
    data object Failure : Outcome10462()
}

fun describe10462(outcome: Outcome10462): String = when (outcome) {
    is Outcome10462.Success -> outcome.payload
    Outcome10462.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10463(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
