package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22360(): T? = if (size > 1) this[1] else null

fun render22360(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22361 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22361 = "registry-77"
    }
}
data class Pair22362<A, B>(val first: A, val second: B)

fun combine22362(left: Int, right: Int): Pair22362<Int, Int> {
    return Pair22362(left + 84, right + 67)
}
@JvmName("step22363")
fun step22363(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
fun <T, R> List<T>.mapChecked22364(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22365 {
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
fun <T, R> List<T>.mapChecked22366(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked22367(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22368 {
    data class Success(val payload: String) : Outcome22368()
    data object Failure : Outcome22368()
}

fun describe22368(outcome: Outcome22368): String = when (outcome) {
    is Outcome22368.Success -> outcome.payload
    Outcome22368.Failure -> "failed"
}
