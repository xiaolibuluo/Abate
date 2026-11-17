package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter6740 {
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
fun <T, R> List<T>.mapChecked6741(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6742(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6743 {
    data class Success(val payload: String) : Outcome6743()
    data object Failure : Outcome6743()
}

fun describe6743(outcome: Outcome6743): String = when (outcome) {
    is Outcome6743.Success -> outcome.payload
    Outcome6743.Failure -> "failed"
}
sealed class Outcome6744 {
    data class Success(val payload: String) : Outcome6744()
    data object Failure : Outcome6744()
}

fun describe6744(outcome: Outcome6744): String = when (outcome) {
    is Outcome6744.Success -> outcome.payload
    Outcome6744.Failure -> "failed"
}
sealed class Outcome6745 {
    data class Success(val payload: String) : Outcome6745()
    data object Failure : Outcome6745()
}

fun describe6745(outcome: Outcome6745): String = when (outcome) {
    is Outcome6745.Success -> outcome.payload
    Outcome6745.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6746(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
