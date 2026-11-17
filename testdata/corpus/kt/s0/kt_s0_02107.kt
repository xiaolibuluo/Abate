package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter17400 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull17401(): T? = if (size > 1) this[1] else null

fun render17401(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17402")
fun step17402(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
sealed class Outcome17403 {
    data class Success(val payload: String) : Outcome17403()
    data object Failure : Outcome17403()
}

fun describe17403(outcome: Outcome17403): String = when (outcome) {
    is Outcome17403.Success -> outcome.payload
    Outcome17403.Failure -> "failed"
}
data class Pair17404<A, B>(val first: A, val second: B)

fun combine17404(left: Int, right: Int): Pair17404<Int, Int> {
    return Pair17404(left + 52, right + 33)
}
fun <T> List<T>.secondOrNull17405(): T? = if (size > 1) this[1] else null

fun render17405(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17406 {
    data class Success(val payload: String) : Outcome17406()
    data object Failure : Outcome17406()
}

fun describe17406(outcome: Outcome17406): String = when (outcome) {
    is Outcome17406.Success -> outcome.payload
    Outcome17406.Failure -> "failed"
}
