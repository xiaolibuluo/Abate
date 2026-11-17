package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6650(): T? = if (size > 1) this[1] else null

fun render6650(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6651<A, B>(val first: A, val second: B)

fun combine6651(left: Int, right: Int): Pair6651<Int, Int> {
    return Pair6651(left + 24, right + 39)
}
class Counter6652 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
@JvmName("step6653")
fun step6653(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
fun <T, R> List<T>.mapChecked6654(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6655<A, B>(val first: A, val second: B)

fun combine6655(left: Int, right: Int): Pair6655<Int, Int> {
    return Pair6655(left + 52, right + 22)
}
sealed class Outcome6656 {
    data class Success(val payload: String) : Outcome6656()
    data object Failure : Outcome6656()
}

fun describe6656(outcome: Outcome6656): String = when (outcome) {
    is Outcome6656.Success -> outcome.payload
    Outcome6656.Failure -> "failed"
}
