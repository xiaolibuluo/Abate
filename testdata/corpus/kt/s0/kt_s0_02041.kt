package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16890 {
    data class Success(val payload: String) : Outcome16890()
    data object Failure : Outcome16890()
}

fun describe16890(outcome: Outcome16890): String = when (outcome) {
    is Outcome16890.Success -> outcome.payload
    Outcome16890.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked16891(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter16892 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
sealed class Outcome16893 {
    data class Success(val payload: String) : Outcome16893()
    data object Failure : Outcome16893()
}

fun describe16893(outcome: Outcome16893): String = when (outcome) {
    is Outcome16893.Success -> outcome.payload
    Outcome16893.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked16894(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked16895(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
