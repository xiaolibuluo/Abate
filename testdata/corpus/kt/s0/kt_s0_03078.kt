package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull25270(): T? = if (size > 1) this[1] else null

fun render25270(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25271(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25272(): T? = if (size > 1) this[1] else null

fun render25272(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25273 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull25274(): T? = if (size > 1) this[1] else null

fun render25274(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25275<A, B>(val first: A, val second: B)

fun combine25275(left: Int, right: Int): Pair25275<Int, Int> {
    return Pair25275(left + 8, right + 63)
}
@JvmName("step25276")
fun step25276(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
sealed class Outcome25277 {
    data class Success(val payload: String) : Outcome25277()
    data object Failure : Outcome25277()
}

fun describe25277(outcome: Outcome25277): String = when (outcome) {
    is Outcome25277.Success -> outcome.payload
    Outcome25277.Failure -> "failed"
}
data class Pair25278<A, B>(val first: A, val second: B)

fun combine25278(left: Int, right: Int): Pair25278<Int, Int> {
    return Pair25278(left + 29, right + 35)
}
