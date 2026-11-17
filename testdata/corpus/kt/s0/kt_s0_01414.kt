package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11640<A, B>(val first: A, val second: B)

fun combine11640(left: Int, right: Int): Pair11640<Int, Int> {
    return Pair11640(left + 5, right + 44)
}
fun <T> List<T>.secondOrNull11641(): T? = if (size > 1) this[1] else null

fun render11641(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11642")
fun step11642(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
@JvmName("step11643")
fun step11643(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
@JvmName("step11644")
fun step11644(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
class Counter11645 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull11646(): T? = if (size > 1) this[1] else null

fun render11646(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome11647 {
    data class Success(val payload: String) : Outcome11647()
    data object Failure : Outcome11647()
}

fun describe11647(outcome: Outcome11647): String = when (outcome) {
    is Outcome11647.Success -> outcome.payload
    Outcome11647.Failure -> "failed"
}
fun <T> List<T>.secondOrNull11648(): T? = if (size > 1) this[1] else null

fun render11648(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
