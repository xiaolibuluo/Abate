package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10200 {
    var current: Int by Delegates.observable(73) { _, _, new -> println(new) }

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
data class Pair10201<A, B>(val first: A, val second: B)

fun combine10201(left: Int, right: Int): Pair10201<Int, Int> {
    return Pair10201(left + 80, right + 23)
}
fun <T> List<T>.secondOrNull10202(): T? = if (size > 1) this[1] else null

fun render10202(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome10203 {
    data class Success(val payload: String) : Outcome10203()
    data object Failure : Outcome10203()
}

fun describe10203(outcome: Outcome10203): String = when (outcome) {
    is Outcome10203.Success -> outcome.payload
    Outcome10203.Failure -> "failed"
}
@JvmName("step10204")
fun step10204(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
fun <T, R> List<T>.mapChecked10205(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair10206<A, B>(val first: A, val second: B)

fun combine10206(left: Int, right: Int): Pair10206<Int, Int> {
    return Pair10206(left + 18, right + 17)
}
fun <T, R> List<T>.mapChecked10207(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome10208 {
    data class Success(val payload: String) : Outcome10208()
    data object Failure : Outcome10208()
}

fun describe10208(outcome: Outcome10208): String = when (outcome) {
    is Outcome10208.Success -> outcome.payload
    Outcome10208.Failure -> "failed"
}
