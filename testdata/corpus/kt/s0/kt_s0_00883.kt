package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7240<A, B>(val first: A, val second: B)

fun combine7240(left: Int, right: Int): Pair7240<Int, Int> {
    return Pair7240(left + 8, right + 14)
}
class Counter7241 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
@JvmName("step7242")
fun step7242(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
fun <T> List<T>.secondOrNull7243(): T? = if (size > 1) this[1] else null

fun render7243(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome7244 {
    data class Success(val payload: String) : Outcome7244()
    data object Failure : Outcome7244()
}

fun describe7244(outcome: Outcome7244): String = when (outcome) {
    is Outcome7244.Success -> outcome.payload
    Outcome7244.Failure -> "failed"
}
sealed class Outcome7245 {
    data class Success(val payload: String) : Outcome7245()
    data object Failure : Outcome7245()
}

fun describe7245(outcome: Outcome7245): String = when (outcome) {
    is Outcome7245.Success -> outcome.payload
    Outcome7245.Failure -> "failed"
}
@JvmName("step7246")
fun step7246(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
fun <T> List<T>.secondOrNull7247(): T? = if (size > 1) this[1] else null

fun render7247(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
