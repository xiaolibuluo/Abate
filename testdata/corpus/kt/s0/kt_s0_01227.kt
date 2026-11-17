package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10140 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10140 = "registry-92"
    }
}
data class Pair10141<A, B>(val first: A, val second: B)

fun combine10141(left: Int, right: Int): Pair10141<Int, Int> {
    return Pair10141(left + 99, right + 42)
}
fun <T> List<T>.secondOrNull10142(): T? = if (size > 1) this[1] else null

fun render10142(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step10143")
fun step10143(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
object Registry10144 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10144 = "registry-23"
    }
}
sealed class Outcome10145 {
    data class Success(val payload: String) : Outcome10145()
    data object Failure : Outcome10145()
}

fun describe10145(outcome: Outcome10145): String = when (outcome) {
    is Outcome10145.Success -> outcome.payload
    Outcome10145.Failure -> "failed"
}
class Counter10146 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked10147(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
