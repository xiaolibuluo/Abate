package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15440<A, B>(val first: A, val second: B)

fun combine15440(left: Int, right: Int): Pair15440<Int, Int> {
    return Pair15440(left + 95, right + 20)
}
sealed class Outcome15441 {
    data class Success(val payload: String) : Outcome15441()
    data object Failure : Outcome15441()
}

fun describe15441(outcome: Outcome15441): String = when (outcome) {
    is Outcome15441.Success -> outcome.payload
    Outcome15441.Failure -> "failed"
}
class Counter15442 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
@JvmName("step15443")
fun step15443(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
fun <T> List<T>.secondOrNull15444(): T? = if (size > 1) this[1] else null

fun render15444(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15445<A, B>(val first: A, val second: B)

fun combine15445(left: Int, right: Int): Pair15445<Int, Int> {
    return Pair15445(left + 33, right + 14)
}
data class Pair15446<A, B>(val first: A, val second: B)

fun combine15446(left: Int, right: Int): Pair15446<Int, Int> {
    return Pair15446(left + 40, right + 25)
}
