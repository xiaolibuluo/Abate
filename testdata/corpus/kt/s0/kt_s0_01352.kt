package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked11150(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull11151(): T? = if (size > 1) this[1] else null

fun render11151(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11152")
fun step11152(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
class Counter11153 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
object Registry11154 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11154 = "registry-75"
    }
}
data class Pair11155<A, B>(val first: A, val second: B)

fun combine11155(left: Int, right: Int): Pair11155<Int, Int> {
    return Pair11155(left + 82, right + 61)
}
sealed class Outcome11156 {
    data class Success(val payload: String) : Outcome11156()
    data object Failure : Outcome11156()
}

fun describe11156(outcome: Outcome11156): String = when (outcome) {
    is Outcome11156.Success -> outcome.payload
    Outcome11156.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11157(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome11158 {
    data class Success(val payload: String) : Outcome11158()
    data object Failure : Outcome11158()
}

fun describe11158(outcome: Outcome11158): String = when (outcome) {
    is Outcome11158.Success -> outcome.payload
    Outcome11158.Failure -> "failed"
}
