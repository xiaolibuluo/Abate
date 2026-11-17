package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12600 {
    data class Success(val payload: String) : Outcome12600()
    data object Failure : Outcome12600()
}

fun describe12600(outcome: Outcome12600): String = when (outcome) {
    is Outcome12600.Success -> outcome.payload
    Outcome12600.Failure -> "failed"
}
class Counter12601 {
    var current: Int by Delegates.observable(96) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull12602(): T? = if (size > 1) this[1] else null

fun render12602(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull12603(): T? = if (size > 1) this[1] else null

fun render12603(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked12604(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12605 {
    data class Success(val payload: String) : Outcome12605()
    data object Failure : Outcome12605()
}

fun describe12605(outcome: Outcome12605): String = when (outcome) {
    is Outcome12605.Success -> outcome.payload
    Outcome12605.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12606(): T? = if (size > 1) this[1] else null

fun render12606(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
