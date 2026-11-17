package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7520 {
    data class Success(val payload: String) : Outcome7520()
    data object Failure : Outcome7520()
}

fun describe7520(outcome: Outcome7520): String = when (outcome) {
    is Outcome7520.Success -> outcome.payload
    Outcome7520.Failure -> "failed"
}
data class Pair7521<A, B>(val first: A, val second: B)

fun combine7521(left: Int, right: Int): Pair7521<Int, Int> {
    return Pair7521(left + 88, right + 38)
}
class Counter7522 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull7523(): T? = if (size > 1) this[1] else null

fun render7523(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
