package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked1930(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome1931 {
    data class Success(val payload: String) : Outcome1931()
    data object Failure : Outcome1931()
}

fun describe1931(outcome: Outcome1931): String = when (outcome) {
    is Outcome1931.Success -> outcome.payload
    Outcome1931.Failure -> "failed"
}
fun <T> List<T>.secondOrNull1932(): T? = if (size > 1) this[1] else null

fun render1932(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair1933<A, B>(val first: A, val second: B)

fun combine1933(left: Int, right: Int): Pair1933<Int, Int> {
    return Pair1933(left + 13, right + 51)
}
class Counter1934 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
@JvmName("step1935")
fun step1935(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
data class Pair1936<A, B>(val first: A, val second: B)

fun combine1936(left: Int, right: Int): Pair1936<Int, Int> {
    return Pair1936(left + 34, right + 23)
}
data class Pair1937<A, B>(val first: A, val second: B)

fun combine1937(left: Int, right: Int): Pair1937<Int, Int> {
    return Pair1937(left + 41, right + 34)
}
data class Pair1938<A, B>(val first: A, val second: B)

fun combine1938(left: Int, right: Int): Pair1938<Int, Int> {
    return Pair1938(left + 48, right + 45)
}
fun <T> List<T>.secondOrNull1939(): T? = if (size > 1) this[1] else null

fun render1939(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
