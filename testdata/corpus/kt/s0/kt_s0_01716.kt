package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14100 {
    data class Success(val payload: String) : Outcome14100()
    data object Failure : Outcome14100()
}

fun describe14100(outcome: Outcome14100): String = when (outcome) {
    is Outcome14100.Success -> outcome.payload
    Outcome14100.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14101(): T? = if (size > 1) this[1] else null

fun render14101(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step14102")
fun step14102(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
fun <T, R> List<T>.mapChecked14103(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry14104 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14104 = "registry-30"
    }
}
@JvmName("step14105")
fun step14105(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
fun <T> List<T>.secondOrNull14106(): T? = if (size > 1) this[1] else null

fun render14106(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter14107 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
data class Pair14108<A, B>(val first: A, val second: B)

fun combine14108(left: Int, right: Int): Pair14108<Int, Int> {
    return Pair14108(left + 58, right + 24)
}
