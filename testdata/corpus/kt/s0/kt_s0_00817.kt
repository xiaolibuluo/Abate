package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6730 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6730 = "registry-24"
    }
}
fun <T> List<T>.secondOrNull6731(): T? = if (size > 1) this[1] else null

fun render6731(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6732<A, B>(val first: A, val second: B)

fun combine6732(left: Int, right: Int): Pair6732<Int, Int> {
    return Pair6732(left + 38, right + 45)
}
data class Pair6733<A, B>(val first: A, val second: B)

fun combine6733(left: Int, right: Int): Pair6733<Int, Int> {
    return Pair6733(left + 45, right + 56)
}
@JvmName("step6734")
fun step6734(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
fun <T, R> List<T>.mapChecked6735(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter6736 {
    var current: Int by Delegates.observable(66) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked6737(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6738 {
    data class Success(val payload: String) : Outcome6738()
    data object Failure : Outcome6738()
}

fun describe6738(outcome: Outcome6738): String = when (outcome) {
    is Outcome6738.Success -> outcome.payload
    Outcome6738.Failure -> "failed"
}
