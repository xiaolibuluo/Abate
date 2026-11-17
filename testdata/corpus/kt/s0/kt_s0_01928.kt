package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15910 {
    data class Success(val payload: String) : Outcome15910()
    data object Failure : Outcome15910()
}

fun describe15910(outcome: Outcome15910): String = when (outcome) {
    is Outcome15910.Success -> outcome.payload
    Outcome15910.Failure -> "failed"
}
object Registry15911 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15911 = "registry-34"
    }
}
class Counter15912 {
    var current: Int by Delegates.observable(41) { _, _, new -> println(new) }

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
@JvmName("step15913")
fun step15913(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
fun <T> List<T>.secondOrNull15914(): T? = if (size > 1) this[1] else null

fun render15914(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15915<A, B>(val first: A, val second: B)

fun combine15915(left: Int, right: Int): Pair15915<Int, Int> {
    return Pair15915(left + 62, right + 38)
}
fun <T> List<T>.secondOrNull15916(): T? = if (size > 1) this[1] else null

fun render15916(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
