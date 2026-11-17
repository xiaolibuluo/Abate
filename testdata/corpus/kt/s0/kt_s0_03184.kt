package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26160 {
    data class Success(val payload: String) : Outcome26160()
    data object Failure : Outcome26160()
}

fun describe26160(outcome: Outcome26160): String = when (outcome) {
    is Outcome26160.Success -> outcome.payload
    Outcome26160.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked26161(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull26162(): T? = if (size > 1) this[1] else null

fun render26162(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair26163<A, B>(val first: A, val second: B)

fun combine26163(left: Int, right: Int): Pair26163<Int, Int> {
    return Pair26163(left + 84, right + 54)
}
@JvmName("step26164")
fun step26164(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
fun <T> List<T>.secondOrNull26165(): T? = if (size > 1) this[1] else null

fun render26165(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter26166 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
