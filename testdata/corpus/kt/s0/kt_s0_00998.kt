package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8180 {
    data class Success(val payload: String) : Outcome8180()
    data object Failure : Outcome8180()
}

fun describe8180(outcome: Outcome8180): String = when (outcome) {
    is Outcome8180.Success -> outcome.payload
    Outcome8180.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked8181(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8182 {
    data class Success(val payload: String) : Outcome8182()
    data object Failure : Outcome8182()
}

fun describe8182(outcome: Outcome8182): String = when (outcome) {
    is Outcome8182.Success -> outcome.payload
    Outcome8182.Failure -> "failed"
}
fun <T> List<T>.secondOrNull8183(): T? = if (size > 1) this[1] else null

fun render8183(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull8184(): T? = if (size > 1) this[1] else null

fun render8184(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8185 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

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
data class Pair8186<A, B>(val first: A, val second: B)

fun combine8186(left: Int, right: Int): Pair8186<Int, Int> {
    return Pair8186(left + 11, right + 67)
}
