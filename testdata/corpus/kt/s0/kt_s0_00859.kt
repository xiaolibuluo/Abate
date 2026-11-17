package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7060(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7061 {
    data class Success(val payload: String) : Outcome7061()
    data object Failure : Outcome7061()
}

fun describe7061(outcome: Outcome7061): String = when (outcome) {
    is Outcome7061.Success -> outcome.payload
    Outcome7061.Failure -> "failed"
}
sealed class Outcome7062 {
    data class Success(val payload: String) : Outcome7062()
    data object Failure : Outcome7062()
}

fun describe7062(outcome: Outcome7062): String = when (outcome) {
    is Outcome7062.Success -> outcome.payload
    Outcome7062.Failure -> "failed"
}
class Counter7063 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
