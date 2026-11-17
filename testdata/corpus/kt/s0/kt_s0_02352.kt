package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull19330(): T? = if (size > 1) this[1] else null

fun render19330(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19331(): T? = if (size > 1) this[1] else null

fun render19331(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19332 {
    data class Success(val payload: String) : Outcome19332()
    data object Failure : Outcome19332()
}

fun describe19332(outcome: Outcome19332): String = when (outcome) {
    is Outcome19332.Success -> outcome.payload
    Outcome19332.Failure -> "failed"
}
@JvmName("step19333")
fun step19333(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
data class Pair19334<A, B>(val first: A, val second: B)

fun combine19334(left: Int, right: Int): Pair19334<Int, Int> {
    return Pair19334(left + 39, right + 42)
}
fun <T> List<T>.secondOrNull19335(): T? = if (size > 1) this[1] else null

fun render19335(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter19336 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
