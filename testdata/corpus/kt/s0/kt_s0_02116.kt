package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17490(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter17491 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked17492(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked17493(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17494 {
    data class Success(val payload: String) : Outcome17494()
    data object Failure : Outcome17494()
}

fun describe17494(outcome: Outcome17494): String = when (outcome) {
    is Outcome17494.Success -> outcome.payload
    Outcome17494.Failure -> "failed"
}
sealed class Outcome17495 {
    data class Success(val payload: String) : Outcome17495()
    data object Failure : Outcome17495()
}

fun describe17495(outcome: Outcome17495): String = when (outcome) {
    is Outcome17495.Success -> outcome.payload
    Outcome17495.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked17496(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17497 {
    data class Success(val payload: String) : Outcome17497()
    data object Failure : Outcome17497()
}

fun describe17497(outcome: Outcome17497): String = when (outcome) {
    is Outcome17497.Success -> outcome.payload
    Outcome17497.Failure -> "failed"
}
