package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull7780(): T? = if (size > 1) this[1] else null

fun render7780(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked7781(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7782 {
    data class Success(val payload: String) : Outcome7782()
    data object Failure : Outcome7782()
}

fun describe7782(outcome: Outcome7782): String = when (outcome) {
    is Outcome7782.Success -> outcome.payload
    Outcome7782.Failure -> "failed"
}
fun <T> List<T>.secondOrNull7783(): T? = if (size > 1) this[1] else null

fun render7783(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter7784 {
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
fun <T> List<T>.secondOrNull7785(): T? = if (size > 1) this[1] else null

fun render7785(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7786<A, B>(val first: A, val second: B)

fun combine7786(left: Int, right: Int): Pair7786<Int, Int> {
    return Pair7786(left + 73, right + 31)
}
sealed class Outcome7787 {
    data class Success(val payload: String) : Outcome7787()
    data object Failure : Outcome7787()
}

fun describe7787(outcome: Outcome7787): String = when (outcome) {
    is Outcome7787.Success -> outcome.payload
    Outcome7787.Failure -> "failed"
}
class Counter7788 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
