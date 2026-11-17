package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11060 {
    data class Success(val payload: String) : Outcome11060()
    data object Failure : Outcome11060()
}

fun describe11060(outcome: Outcome11060): String = when (outcome) {
    is Outcome11060.Success -> outcome.payload
    Outcome11060.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11061(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome11062 {
    data class Success(val payload: String) : Outcome11062()
    data object Failure : Outcome11062()
}

fun describe11062(outcome: Outcome11062): String = when (outcome) {
    is Outcome11062.Success -> outcome.payload
    Outcome11062.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11063(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked11064(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter11065 {
    var current: Int by Delegates.observable(62) { _, _, new -> println(new) }

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
