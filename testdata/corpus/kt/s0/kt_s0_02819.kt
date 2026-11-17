package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23240(): T? = if (size > 1) this[1] else null

fun render23240(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull23241(): T? = if (size > 1) this[1] else null

fun render23241(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair23242<A, B>(val first: A, val second: B)

fun combine23242(left: Int, right: Int): Pair23242<Int, Int> {
    return Pair23242(left + 64, right + 12)
}
data class Pair23243<A, B>(val first: A, val second: B)

fun combine23243(left: Int, right: Int): Pair23243<Int, Int> {
    return Pair23243(left + 71, right + 23)
}
@JvmName("step23244")
fun step23244(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
object Registry23245 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23245 = "registry-85"
    }
}
class Counter23246 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
sealed class Outcome23247 {
    data class Success(val payload: String) : Outcome23247()
    data object Failure : Outcome23247()
}

fun describe23247(outcome: Outcome23247): String = when (outcome) {
    is Outcome23247.Success -> outcome.payload
    Outcome23247.Failure -> "failed"
}
